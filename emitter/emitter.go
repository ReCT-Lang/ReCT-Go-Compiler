package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/symbols"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Emitter struct {
	Program         binder.BoundProgram
	Module          *ir.Module
	UseFingerprints bool

	// referenced functions
	CFuncs   map[string]*ir.Func
	ArcFuncs map[string]*ir.Func

	// referenced classes
	Classes map[string]Class

	// global variables
	Globals          map[string]Global
	Functions        map[string]Function
	FunctionWrappers map[string]*ir.Func
	FunctionLocals   map[string]map[string]Local
	StrConstants     map[string]value.Value
	StrNameCounter   int

	// local things for this current class
	Class    Class
	ClassSym symbols.ClassSymbol

	// local things for this current function
	Function    *ir.Func
	FunctionSym symbols.FunctionSymbol
	Locals      map[string]Local
	Temps       []string
	Labels      map[string]*ir.Block

	// flags
	IsInClass bool
}

const verboseARC = false

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:          program,
		Module:           ir.NewModule(),
		UseFingerprints:  useFingerprints,
		Globals:          make(map[string]Global),
		Functions:        make(map[string]Function),
		CFuncs:           make(map[string]*ir.Func),
		ArcFuncs:         make(map[string]*ir.Func),
		FunctionLocals:   make(map[string]map[string]Local),
		StrConstants:     make(map[string]value.Value),
		Classes:          make(map[string]Class),
		FunctionWrappers: make(map[string]*ir.Func),
		Temps:            make([]string, 0),
	}

	emitter.EmitBuiltInFunctions()

	// declare all class structs
	for _, cls := range emitter.Program.Classes {
		emitter.EmitClass(cls)

		// declare all function names inside the class
		for _, fnc := range cls.Functions {
			if !fnc.Symbol.BuiltIn { // ignore system functions
				// ignore constructor and Die as they are declared automatically
				if fnc.Symbol.Name != "Constructor" && fnc.Symbol.Name != "Die" {
					function := emitter.EmitClassFunction(cls.Symbol, fnc.Symbol, fnc.Body)
					functionName := emitter.Id(fnc.Symbol)
					emitter.Classes[emitter.Id(cls.Symbol.Type)].Functions[functionName] = function
				}
			}
		}
	}

	// declare all function names
	for _, fnc := range emitter.Program.Functions {
		if !fnc.Symbol.BuiltIn {
			function := Function{IRFunction: emitter.EmitFunction(fnc.Symbol, fnc.Body), BoundFunction: fnc}
			functionName := emitter.Id(fnc.Symbol)
			emitter.Functions[functionName] = function
		}
	}

	// emit main function first
	mainName := emitter.Id(program.MainFunction)
	emitter.FunctionSym = emitter.Functions[mainName].BoundFunction.Symbol
	emitter.EmitBlockStatement(emitter.Functions[mainName].BoundFunction.Symbol, emitter.Functions[mainName].IRFunction, emitter.Functions[mainName].BoundFunction.Body)

	// emit function bodies
	for _, fnc := range emitter.Functions {
		if !fnc.BoundFunction.Symbol.BuiltIn && fnc.BoundFunction.Symbol.Fingerprint() != program.MainFunction.Fingerprint() {
			emitter.FunctionSym = fnc.BoundFunction.Symbol
			emitter.EmitBlockStatement(fnc.BoundFunction.Symbol, fnc.IRFunction, fnc.BoundFunction.Body)
		}
	}

	// emit class function bodies
	for _, cls := range emitter.Program.Classes {
		for _, fnc := range cls.Functions {
			if fnc.Symbol.BuiltIn {
				continue
			}

			emitter.FunctionSym = fnc.Symbol
			emitter.Class = emitter.Classes[emitter.Id(cls.Symbol.Type)]
			emitter.ClassSym = cls.Symbol
			emitter.IsInClass = true

			// find out if this is the constructor
			if fnc.Symbol.Name == "Constructor" {
				// if it is, hand it the already prepared constructor function
				emitter.EmitBlockStatement(fnc.Symbol, emitter.Classes[emitter.Id(cls.Symbol.Type)].Constructor, fnc.Body)

			} else if fnc.Symbol.Name == "Die" {
				// if its the destructor, clear all auto-generated code inside it
				emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor.Blocks = make([]*ir.Block, 0)
				emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor.NewBlock("")

				// hand it the already prepared destructor function
				emitter.EmitBlockStatement(fnc.Symbol, emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor, fnc.Body)
			} else {
				// if not, emit the function like normal
				emitter.EmitBlockStatement(fnc.Symbol, emitter.Classes[emitter.Id(cls.Symbol.Type)].Functions[emitter.Id(fnc.Symbol)], fnc.Body)
			}

		}
	}

	return emitter.Module
}

// <CLASSES>-------------------------------------------------------------------
func (emt *Emitter) EmitClass(cls binder.BoundClass) {
	// create the class' vTable type
	clsvTable := emt.Module.NewTypeDef("struct."+cls.Symbol.Name+"_vTable",
		types.NewStruct(
			types.NewPointer(emt.Classes[emt.Id(builtins.Any)].vTable),
			types.I8Ptr,
			types.NewPointer(types.NewFunc(types.Void, types.I8Ptr)),
		),
	)

	// =====================================================================
	// create the llvm type
	// ---------------------------------------------------------------------
	// general struct format:
	// ----------------------
	// [0] Class vTable, contains the destructor and any virtual methods
	// [1] Object reference counter for the ARC
	// ... Class data
	// =====================================================================

	// list of our type's fields
	clsFields := make([]types.Type, 0)
	clsFieldMap := make(map[string]int)

	clsFields = append(clsFields, types.NewPointer(clsvTable)) // vTable
	clsFields = append(clsFields, types.I32)                   // ARC counter

	// add our types one by one
	for i, field := range cls.Symbol.Fields {
		clsFieldMap[emt.Id(field)] = i + 2
		clsFields = append(clsFields, emt.IRTypes(field.VarType()))
	}

	// create the class' struct
	clsType := emt.Module.NewTypeDef("struct.class_"+cls.Symbol.Name,
		types.NewStruct(clsFields...),
	)

	// ---------------------------------------------------------------------
	// create the destructor
	// ---------------------------------------------------------------------

	// create an IR function definition for the destructor
	destructor := emt.Module.NewFunc(cls.Symbol.Name+"_public_Die", types.Void, ir.NewParam("obj", types.I8Ptr))

	// create a new root block
	root := destructor.NewBlock("")

	// bitcast the given void* into a class pointer
	clsPtr := root.NewBitCast(destructor.Params[0], types.NewPointer(clsType))

	// generate cleanup for all object members
	root.Insts = append(root.Insts, NewComment("<DieARC>"))

	for _, field := range cls.Symbol.Fields {
		if field.VarType().IsObject {
			root.Insts = append(root.Insts, NewComment(fmt.Sprintf("-> destroying reference to '%s [Field %d]'", emt.Id(field), clsFieldMap[emt.Id(field)])))

			// get the field's pointer
			ptr := root.NewGetElementPtr(clsType, clsPtr, CI32(0), CI32(int32(clsFieldMap[emt.Id(field)])))

			// load the pointer
			obj := root.NewLoad(emt.IRTypes(field.VarType()), ptr)

			// decrement its ARC reference counter
			emt.DestroyReference(&root, obj, "")
		}
	}

	root.Insts = append(root.Insts, NewComment("</DieARC>"))

	root.NewRet(nil)

	// create the vTable constant
	vtc := constant.NewStruct(
		clsvTable.(*types.StructType),
		emt.Classes[emt.Id(builtins.Any)].vConstant,
		emt.GetConstantStringConstant(cls.Symbol.Name),
		destructor,
	)
	clsvConstant := emt.Module.NewGlobalDef(cls.Symbol.Name+"_vTable_Const", vtc)

	// ---------------------------------------------------------------------
	// create the constructor
	// ---------------------------------------------------------------------

	// create an IR function definition for the constructor
	clsParams := make([]*ir.Param, 0)
	clsParams = append(clsParams, ir.NewParam("me", types.NewPointer(clsType)))

	// look for an explicit constructor
	var constructorFunction *binder.BoundFunction
	for _, fnc := range cls.Functions {
		if fnc.Symbol.Name == "Constructor" {
			constructorFunction = &fnc
			break
		}
	}

	// constructor found
	if constructorFunction != nil {
		// add the constructors parameters onto our template
		for _, param := range constructorFunction.Symbol.Parameters {
			clsParams = append(clsParams, ir.NewParam(emt.Id(param), emt.IRTypes(param.Type)))
		}
	}

	constructor := emt.Module.NewFunc(cls.Symbol.Name+"_public_Constructor", types.Void, clsParams...)

	// create a basic constructor in IR
	// --------------------------------

	// get the objects own reference
	croot := constructor.NewBlock("")
	clsMePtr := croot.NewAlloca(types.NewPointer(clsType))
	croot.NewStore(clsParams[0], clsMePtr)
	clsMe := croot.NewLoad(types.NewPointer(clsType), clsMePtr)

	// set the vTable to the vTable constant
	clsMyVTable := croot.NewGetElementPtr(clsType, clsMe, CI32(0), CI32(0))
	croot.NewStore(clsvConstant, clsMyVTable)

	// set the reference count to 0
	clsMyRefCount := croot.NewGetElementPtr(clsType, clsMe, CI32(0), CI32(1))
	croot.NewStore(CI32(0), clsMyRefCount)

	// store a NULL in any object fields, to tell the ARC that they are empty
	for _, field := range cls.Symbol.Fields {
		if field.VarType().IsObject {
			ptr := croot.NewGetElementPtr(clsType, constructor.Params[0], CI32(0), CI32(int32(clsFieldMap[emt.Id(field)])))
			croot.NewStore(constant.NewNull(emt.IRTypes(field.VarType()).(*types.PointerType)), ptr)
		}
	}

	// don
	croot.NewRet(nil)

	// create locals array for constructor and destructor
	if constructorFunction != nil {
		// create locals array
		locals := make(map[string]Local)

		// create all needed locals in the root block to GC can trash them anywhere
		for _, stmt := range constructorFunction.Body.Statements {
			if stmt.NodeType() == boundnodes.BoundVariableDeclaration {
				declStatement := stmt.(boundnodes.BoundVariableDeclarationStatementNode)

				if declStatement.Variable.IsGlobal() {
					continue
				}

				varName := emt.Id(declStatement.Variable)

				// create local variable
				local := croot.NewAlloca(emt.IRTypes(declStatement.Variable.VarType()))
				local.SetName(varName)

				// save it for referencing later
				locals[varName] = Local{IRLocal: local, IRBlock: croot, Type: declStatement.Variable.VarType()}
			}
		}

		// store this for later
		emt.FunctionLocals[cls.Symbol.Name+"_public_"+emt.Id(constructorFunction.Symbol)] = locals
	}

	// create the class object to keep track of things
	emt.Classes[emt.Id(cls.Symbol.Type)] = Class{
		Name: cls.Symbol.Name,
		Type: clsType,

		vTable:    clsvTable,
		vConstant: clsvConstant,

		Constructor: constructor,
		Destructor:  destructor,

		Functions: make(map[string]*ir.Func),
		Fields:    clsFieldMap,
	}
}

// </CLASSES>------------------------------------------------------------------
// <FUNCTIONS>-----------------------------------------------------------------

func (emt *Emitter) EmitFunction(sym symbols.FunctionSymbol, body boundnodes.BoundBlockStatementNode) *ir.Func {
	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := emt.Id(param)

		// create it
		params = append(params, ir.NewParam(paramName, emt.IRTypes(param.Type)))
	}

	// figure out the return type
	returnType := emt.IRTypes(sym.Type)

	// the function name
	functionName := emt.Id(sym)
	irName := tern(sym.Fingerprint() == emt.Program.MainFunction.Fingerprint(), "main", functionName)

	// create an IR function definition
	function := emt.Module.NewFunc(irName, returnType, params...)

	// create a root block
	root := function.NewBlock("")

	// create locals array
	locals := make(map[string]Local)

	// create all needed locals in the root block to GC can trash them anywhere
	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundVariableDeclaration {
			declStatement := stmt.(boundnodes.BoundVariableDeclarationStatementNode)

			if declStatement.Variable.IsGlobal() {
				continue
			}

			varName := emt.Id(declStatement.Variable)

			// create local variable
			local := root.NewAlloca(emt.IRTypes(declStatement.Variable.VarType()))
			local.SetName(varName)

			// save it for referencing later
			locals[varName] = Local{IRLocal: local, IRBlock: root, Type: declStatement.Variable.VarType()}
		}
	}

	// store this for later
	emt.FunctionLocals[functionName] = locals

	return function
}

func (emt *Emitter) EmitClassFunction(cls symbols.ClassSymbol, sym symbols.FunctionSymbol, body boundnodes.BoundBlockStatementNode) *ir.Func {
	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	params = append(params, ir.NewParam("$me", types.NewPointer(emt.Classes[emt.Id(cls.Type)].Type)))

	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := emt.Id(param)

		// create it
		params = append(params, ir.NewParam(paramName, emt.IRTypes(param.Type)))
	}

	// figure out the return type
	returnType := emt.IRTypes(sym.Type)

	// the function name
	functionName := emt.Id(sym)
	irName := cls.Name + "_" + tern(sym.Public, "public", "private") + "_" + functionName

	// create an IR function definition
	function := emt.Module.NewFunc(irName, returnType, params...)

	// create a root block
	root := function.NewBlock("")

	// create locals array
	locals := make(map[string]Local)

	// create all needed locals in the root block to GC can trash them anywhere
	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundVariableDeclaration {
			declStatement := stmt.(boundnodes.BoundVariableDeclarationStatementNode)

			if declStatement.Variable.IsGlobal() {
				continue
			}

			varName := emt.Id(declStatement.Variable)

			// create local variable
			local := root.NewAlloca(emt.IRTypes(declStatement.Variable.VarType()))
			local.SetName(varName)

			// save it for referencing later
			locals[varName] = Local{IRLocal: local, IRBlock: root, Type: declStatement.Variable.VarType()}
		}
	}

	// store this for later
	emt.FunctionLocals[irName] = locals

	return function
}

func (emt *Emitter) EmitBlockStatement(sym symbols.FunctionSymbol, fnc *ir.Func, body boundnodes.BoundBlockStatementNode) {
	// set up our environment
	functionName := emt.Id(sym)
	irName := functionName

	if emt.IsInClass {
		irName = emt.Class.Name + "_" + tern(sym.Public, "public", "private") + "_" + functionName
	}

	emt.Function = fnc
	emt.Locals = emt.FunctionLocals[irName]
	emt.Labels = make(map[string]*ir.Block)

	// create a semi-root block
	// each label statement will create a new block and store it in this variable
	currentBlock := fnc.Blocks[0]
	semiroot := fnc.NewBlock("semiroot")

	currentBlock.NewBr(semiroot)
	currentBlock = semiroot

	// go through the body and register all label blocks
	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// create a new block when we encounter a label
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			emt.Labels[string(labelStatement.Label)] = fnc.NewBlock(string(labelStatement.Label))
		}
	}

	skipToNextBlock := false

	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// if we encounter a label statement -> switch to the label's block
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			currentBlock = emt.Labels[string(labelStatement.Label)]
			skipToNextBlock = false

		} else {
			if skipToNextBlock {
				continue
			}

			// emit a statement to the current block
			switch stmt.NodeType() {
			case boundnodes.BoundVariableDeclaration:
				emt.EmitVariableDeclarationStatement(&currentBlock, stmt.(boundnodes.BoundVariableDeclarationStatementNode))

			case boundnodes.BoundGotoStatement:
				emt.EmitGotoStatement(&currentBlock, stmt.(boundnodes.BoundGotoStatementNode))
				skipToNextBlock = true

			case boundnodes.BoundConditionalGotoStatement:
				emt.EmitConditionalGotoStatement(&currentBlock, stmt.(boundnodes.BoundConditionalGotoStatementNode))

			case boundnodes.BoundExpressionStatement:
				emt.EmitBoundExpressionStatement(&currentBlock, stmt.(boundnodes.BoundExpressionStatementNode))

			case boundnodes.BoundReturnStatement:
				emt.EmitReturnStatement(&currentBlock, stmt.(boundnodes.BoundReturnStatementNode))
				// skip forward until we either hit a new block or the end of the function
				skipToNextBlock = true

			case boundnodes.BoundGarbageCollectionStatement:
				emt.EmitGarbageCollectionStatement(&currentBlock, stmt.(boundnodes.BoundGarbageCollectionStatementNode))
			}
		}
	}
}

// </FUNCTIONS>----------------------------------------------------------------
// <STATEMENTS>----------------------------------------------------------------

func (emt *Emitter) EmitVariableDeclarationStatement(blk **ir.Block, stmt boundnodes.BoundVariableDeclarationStatementNode) {
	varName := emt.Id(stmt.Variable)
	var expression value.Value

	// if there's an initializer given
	if stmt.Initializer != nil {
		expression = emt.EmitExpression(blk, stmt.Initializer)

		// if the expression is a variable and contains an object -> increase reference counter
		if stmt.Initializer.IsPersistent() && stmt.Initializer.Type().IsObject {
			emt.CreateReference(blk, expression, "variable declaration ["+varName+"]")
		}
	} else {
		// if there's no initializer -> use the type's default
		expression = emt.DefaultConstant(blk, stmt.Variable.VarType())
	}

	if stmt.Variable.IsGlobal() {
		// create a new global
		global := emt.Module.NewGlobalDef(varName, emt.DefaultConstant(blk, stmt.Variable.VarType()))

		// save it for referencing later
		emt.Globals[varName] = Global{IRGlobal: global, Type: stmt.Variable.VarType()}

		// emit its assignment
		(*blk).NewStore(expression, global)
	} else {
		local := emt.Locals[varName]
		local.IsSet = true
		emt.Locals[varName] = local

		// emit its assignemnt
		(*blk).NewStore(expression, local.IRLocal)
	}
}

func (emt *Emitter) EmitGotoStatement(blk **ir.Block, stmt boundnodes.BoundGotoStatementNode) {
	(*blk).NewBr(emt.Labels[string(stmt.Label)])
}

func (emt *Emitter) EmitConditionalGotoStatement(blk **ir.Block, stmt boundnodes.BoundConditionalGotoStatementNode) {
	// figure out where to jump
	ifLabel := emt.Labels[string(stmt.IfLabel)]
	elseLabel := emt.Labels[string(stmt.ElseLabel)]

	(*blk).NewCondBr(emt.EmitExpression(blk, stmt.Condition), ifLabel, elseLabel)
}

func (emt *Emitter) EmitBoundExpressionStatement(blk **ir.Block, stmt boundnodes.BoundExpressionStatementNode) {
	expr := emt.EmitExpression(blk, stmt.Expression)

	// if the expressions value is a string is requires cleanup
	if stmt.Expression.Type().IsObject {
		(*blk).Insts = append((*blk).Insts, NewComment("expression value unused -> destroying reference"))
		emt.DestroyReference(blk, expr, "destroying unused expression")
	}
}

func (emt *Emitter) EmitReturnStatement(blk **ir.Block, stmt boundnodes.BoundReturnStatementNode) {
	// return value
	var expression value.Value

	// calculate the return value first (so its not destroyed by the GC)
	if stmt.Expression != nil {
		expression = emt.EmitExpression(blk, stmt.Expression)

		// if the expression is an object, increase its reference count to not have it deleted
		// only do this for variables, as expressions will already have the correct count
		if stmt.Expression.IsPersistent() {
			if stmt.Expression.Type().IsObject {
				emt.CreateReference(blk, expression, "return value copy ["+emt.Function.Name()+"]")
			}
		}
	}

	// --< state-of-the-art garbage collecting >---------------------------
	// 1. go through all locals created up to this point
	// 2. decrement their reference counter
	(*blk).Insts = append((*blk).Insts, NewComment("<ReturnARC>"))
	for name, local := range emt.Locals {

		// if nothing has been assigned yet, there's no need to clean up
		if !local.IsSet {
			continue
		}

		// only clean up things that actually need it (any and string)
		if local.Type.IsObject {
			(*blk).Insts = append((*blk).Insts, NewComment(" -> destroying reference to '%"+name+"'"))
			emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(local.Type), local.IRLocal), "ReturnGC variable '"+local.IRLocal.Ident()+"' (leaving '"+emt.Function.Name()+"')")
		}
	}

	// clean up any parameters as well
	// (passed variables wont be cleaned up as their reference counter has been increased before the call)
	for _, param := range emt.FunctionSym.Parameters {
		if param.Type.IsObject {
			(*blk).Insts = append((*blk).Insts, NewComment(" -> destroying reference to '%"+param.Name+"'"))
			emt.DestroyReference(blk, emt.Function.Params[param.Ordinal], "ReturnGC (parameter) (leaving '"+emt.Function.Name()+"')")
		}
	}
	(*blk).Insts = append((*blk).Insts, NewComment("</ReturnARC>"))

	if stmt.Expression != nil {
		(*blk).NewRet(expression)
	} else {
		(*blk).NewRet(nil)
	}
}

func (emt *Emitter) EmitGarbageCollectionStatement(blk **ir.Block, stmt boundnodes.BoundGarbageCollectionStatementNode) {
	(*blk).Insts = append((*blk).Insts, NewComment("<GC - ARC>"))
	for _, variable := range stmt.Variables {
		// check if the variables type needs to be freed

		if variable.VarType().IsObject {
			varName := emt.Id(variable)
			(*blk).Insts = append((*blk).Insts, NewComment(" -> destroying reference to '%"+varName+"'"))

			emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(variable.VarType()), emt.Locals[varName].IRLocal), "GC statement (end of block)")

			// write NULL to the pointer
			(*blk).NewStore(constant.NewNull(emt.IRTypes(variable.VarType()).(*types.PointerType)), emt.Locals[varName].IRLocal)
		}
	}
	(*blk).Insts = append((*blk).Insts, NewComment("</GC - ARC>"))
}

// </STATEMENTS>---------------------------------------------------------------
// <EXPRESSIONS>---------------------------------------------------------------

func (emt *Emitter) EmitExpression(blk **ir.Block, expr boundnodes.BoundExpressionNode) value.Value {
	switch expr.NodeType() {
	case boundnodes.BoundLiteralExpression:
		return emt.EmitLiteralExpression(blk, expr.(boundnodes.BoundLiteralExpressionNode))

	case boundnodes.BoundVariableExpression:
		return emt.EmitVariableExpression(blk, expr.(boundnodes.BoundVariableExpressionNode))

	case boundnodes.BoundAssignmentExpression:
		return emt.EmitAssignmentExpression(blk, expr.(boundnodes.BoundAssignmentExpressionNode))

	case boundnodes.BoundMakeExpression:
		return emt.EmitMakeExpression(blk, expr.(boundnodes.BoundMakeExpressionNode))

	case boundnodes.BoundMakeArrayExpression:
		return emt.EmitMakeArrayExpression(blk, expr.(boundnodes.BoundMakeArrayExpressionNode))

	case boundnodes.BoundArrayAccessExpression:
		return emt.EmitArrayAccessExpression(blk, expr.(boundnodes.BoundArrayAccessExpressionNode))

	case boundnodes.BoundArrayAssignmentExpression:
		return emt.EmitArrayAssignmentExpression(blk, expr.(boundnodes.BoundArrayAssignmentExpressionNode))

	case boundnodes.BoundUnaryExpression:
		return emt.EmitUnaryExpression(blk, expr.(boundnodes.BoundUnaryExpressionNode))

	case boundnodes.BoundBinaryExpression:
		return emt.EmitBinaryExpression(blk, expr.(boundnodes.BoundBinaryExpressionNode))

	case boundnodes.BoundTernaryExpression:
		return emt.EmitTernaryExpression(blk, expr.(boundnodes.BoundTernaryExpressionNode))

	case boundnodes.BoundCallExpression:
		return emt.EmitCallExpression(blk, expr.(boundnodes.BoundCallExpressionNode))

	case boundnodes.BoundTypeCallExpression:
		return emt.EmitTypeCallExpression(blk, expr.(boundnodes.BoundTypeCallExpressionNode))

	case boundnodes.BoundClassCallExpression:
		return emt.EmitClassCallExpression(blk, expr.(boundnodes.BoundClassCallExpressionNode))

	case boundnodes.BoundConversionExpression:
		return emt.EmitConversionExpression(blk, expr.(boundnodes.BoundConversionExpressionNode))

	case boundnodes.BoundThreadExpression:
		return emt.EmitThreadStatement(blk, expr.(boundnodes.BoundThreadExpressionNode))
	}

	fmt.Println("Unimplemented node: " + expr.NodeType())
	return nil
}

func (emt *Emitter) EmitLiteralExpression(blk **ir.Block, expr boundnodes.BoundLiteralExpressionNode) value.Value {
	switch expr.LiteralType.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(expr.Value.(bool))
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, int64(expr.Value.(int)))
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, float64(expr.Value.(float32)))
	case builtins.String.Fingerprint():
		charPtr := emt.GetStringConstant(blk, expr.Value.(string))
		strObj := emt.CreateObject(blk, emt.Id(builtins.String))
		(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, charPtr)
		return strObj
	}

	fmt.Println("Unknown literal!")

	return nil
}

func (emt *Emitter) EmitVariableExpression(blk **ir.Block, expr boundnodes.BoundVariableExpressionNode) value.Value {
	return emt.EmitVariable(blk, expr.Variable)
}

func (emt *Emitter) EmitVariable(blk **ir.Block, variable symbols.VariableSymbol) value.Value {
	varName := emt.Id(variable)

	// parameters
	if variable.SymbolType() == symbols.Parameter {
		paramSymbol := variable.(symbols.ParameterSymbol)

		if emt.IsInClass {
			return emt.Function.Params[paramSymbol.Ordinal+1]
		} else {
			return emt.Function.Params[paramSymbol.Ordinal]
		}
	}

	if variable.IsGlobal() {
		// if we're in a class we need to load from the struct instead of a global
		if emt.IsInClass {
			ptr := (*blk).NewGetElementPtr(emt.Class.Type, emt.Function.Params[0], CI32(0), CI32(int32(emt.Class.Fields[emt.Id(variable)])))
			return (*blk).NewLoad(emt.IRTypes(variable.VarType()), ptr)

		} else {
			// if we arent we can just get the global
			return (*blk).NewLoad(emt.IRTypes(emt.Globals[varName].Type), emt.Globals[varName].IRGlobal)
		}
	} else {
		return (*blk).NewLoad(emt.IRTypes(emt.Locals[varName].Type), emt.Locals[varName].IRLocal)
	}
}

func (emt *Emitter) EmitAssignmentExpression(blk **ir.Block, expr boundnodes.BoundAssignmentExpressionNode) value.Value {
	varName := emt.Id(expr.Variable)
	expression := emt.EmitExpression(blk, expr.Expression)

	// if the expression is a variable -> increase reference counter
	if expr.Expression.IsPersistent() && expr.Expression.Type().IsObject {
		emt.CreateReference(blk, expression, "variable assignment ["+varName+"]")
	}

	if expr.Variable.IsGlobal() {
		if emt.IsInClass {
			// the location we need to store to
			ptr := (*blk).NewGetElementPtr(emt.Class.Type, emt.Function.Params[0], CI32(0), CI32(int32(emt.Class.Fields[emt.Id(expr.Variable)])))

			// if this variable already contained an object -> destroy the reference
			if expr.Variable.VarType().IsObject {
				emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(expr.Variable.VarType()), ptr), "destroying reference previously stored in '"+varName+"'")
			}

			// assign the value to the structs field
			(*blk).NewStore(expression, ptr)
		} else {
			// if this variable already contained an object -> destroy the reference
			if expr.Variable.VarType().IsObject {
				emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(expr.Variable.VarType()), emt.Globals[varName].IRGlobal), "destroying reference previously stored in '"+varName+"'")
			}

			// assign the value to the global variable
			(*blk).NewStore(expression, emt.Globals[varName].IRGlobal)
		}

	} else {
		// if this variable already contained an object -> destroy there reference
		if expr.Variable.VarType().IsObject {
			emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(expr.Variable.VarType()), emt.Locals[varName].IRLocal), "destroying reference previously stored in '"+varName+"'")
		}

		// assign the value to the local variable
		(*blk).NewStore(expression, emt.Locals[varName].IRLocal)
	}

	// also return the value as this can also be used as an expression
	// if we're working with objects, a new reference has to be counted
	if expr.Variable.VarType().IsObject {
		emt.CreateReference(blk, expression, "assignment value copy (for stuff like a <- b++)")
		return expression
	}

	return expression
}

func (emt *Emitter) EmitMakeExpression(blk **ir.Block, expr boundnodes.BoundMakeExpressionNode) value.Value {
	// emit all of the constructors arguments
	arguments := make([]value.Value, 0)

	for _, arg := range expr.Arguments {
		expression := emt.EmitExpression(blk, arg)

		// if this is an object -> increase its reference counter
		// (only do this for variables)
		if arg.IsPersistent() {
			if arg.Type().Fingerprint() == builtins.String.Fingerprint() ||
				arg.Type().Fingerprint() == builtins.Any.Fingerprint() {
				emt.CreateReference(blk, expression, "copy to be passed into a parameter")
			}
		}

		arguments = append(arguments, expression)
	}

	// create an object of the given type
	obj := emt.CreateObject(blk, emt.Id(expr.BaseType.Type), arguments...)

	// return the object
	return obj
}

func (emt *Emitter) EmitMakeArrayExpression(blk **ir.Block, expr boundnodes.BoundMakeArrayExpressionNode) value.Value {
	// get the array length
	var length value.Value

	if !expr.IsLiteral {
		length = emt.EmitExpression(blk, expr.Length)
	} else {
		length = CI32(int32(len(expr.Literals)))
	}

	// variable for keepong track of the array
	var arrObject value.Value

	// check if this is an object array or a primitive array
	if expr.BaseType.IsObject {
		// create a new object array object
		arrObject = emt.CreateObject(blk, emt.Id(builtins.Array), length)
	} else {
		// get the size of the primitive we want to allocate
		size := emt.SizeOf(blk, expr.BaseType)

		// create a new primitive array object
		arrObject = emt.CreateObject(blk, emt.Id(builtins.PArray), length, size)
	}

	// if this is a literal, load its values
	if expr.IsLiteral {
		for i, literal := range expr.Literals {
			emt.EmitArrayAssignment(blk, arrObject, CI32(int32(i)), literal)
		}
	}

	return arrObject
}

func (emt *Emitter) EmitThreadStatement(blk **ir.Block, stmt boundnodes.BoundThreadExpressionNode) value.Value {

	// get this function's thread wrapper
	wrapper := emt.GetThreadWrapper(stmt.Function)

	// get a pointer to the thread wrapper
	wrapperPointer := (*blk).NewAlloca(wrapper.Typ)
	(*blk).NewStore(wrapper, wrapperPointer)

	// load the pointer again
	pointer := (*blk).NewLoad(wrapper.Typ, wrapperPointer)

	// Not sure how to approach passing arguments to the constructor in CreateObject (class_Action)
	obj := emt.CreateObject(blk, emt.Id(builtins.Thread), pointer, constant.NewNull(types.I8Ptr))

	return obj
}

func (emt *Emitter) EmitArrayAssignmentExpression(blk **ir.Block, expr boundnodes.BoundArrayAssignmentExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// index
	// -----
	index := emt.EmitExpression(blk, expr.Index)

	// assignment
	// ----------
	value := emt.EmitExpression(blk, expr.Value)

	// decide if we should do object or primitive array access
	if expr.Base.Type().SubTypes[0].IsObject {
		// bitcast our pointer to any
		anyValue := (*blk).NewBitCast(value, emt.IRTypes(builtins.Any))

		// call the array's set element function
		(*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["SetElement"], base, index, anyValue)

		// if the element wasnt a variable -> decrease its reference counter
		if !expr.Value.IsPersistent() && expr.Value.Type().IsObject {
			emt.DestroyReference(blk, value, "array assignment cleanup")
		}

		// return a copy of the value (i really don't think this is necessary but oh well)
		emt.CreateReference(blk, value, "assign value copy (array assignment)")

		return value
	} else {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		(*blk).NewStore(value, castedPtr)

		// return a copy of the value
		return value
	}
}

func (emt *Emitter) EmitArrayAccessExpression(blk **ir.Block, expr boundnodes.BoundArrayAccessExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// load the index
	index := emt.EmitExpression(blk, expr.Index)

	// do the access
	// -------------

	// decide if we should do object or primitive array access
	if expr.Base.Type().SubTypes[0].IsObject {
		// call the array's get element function
		element := (*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["GetElement"], base, index)

		originalType := expr.Base.Type().SubTypes[0]

		// if the original type is an array it'll need a little help to pick the correct type
		if originalType.Name == builtins.Array.Name {
			if originalType.SubTypes[0].IsObject {
				originalType = builtins.Array
			} else {
				originalType = builtins.PArray
			}
		}

		// bitcast our pointer to its original class
		castedElement := (*blk).NewBitCast(element, types.NewPointer(emt.Classes[emt.Id(originalType)].Type))

		return castedElement
	} else {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		// load the value
		val := (*blk).NewLoad(emt.IRTypes(expr.Base.Type().SubTypes[0]), castedPtr)
		return val
	}
}

func (emt *Emitter) EmitUnaryExpression(blk **ir.Block, expr boundnodes.BoundUnaryExpressionNode) value.Value {
	expression := emt.EmitExpression(blk, expr.Expression)

	switch expr.Op.OperatorKind {
	case boundnodes.Identity:
		return expression
	case boundnodes.Negation:
		// int negation   -> 0 - value
		// float negation -> fneg value
		if expr.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFNeg(expression)
		}

	case boundnodes.LogicalNegation:
		cmp := (*blk).NewICmp(enum.IPredNE, expression, CI32(0))
		xor := (*blk).NewXor(cmp, CB1(true))
		return xor
	}

	fmt.Println("Unknown Unary!")
	return nil
}

func (emt *Emitter) EmitBinaryExpression(blk **ir.Block, expr boundnodes.BoundBinaryExpressionNode) value.Value {
	left := emt.EmitExpression(blk, expr.Left)
	right := emt.EmitExpression(blk, expr.Right)

	switch expr.Op.OperatorKind {
	case boundnodes.Addition:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			newStr := (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Concat"], left, right)

			// if left and right aren't variables (meaning they are already memory managed)
			// decrease their reference count
			if !expr.Left.IsPersistent() {
				emt.DestroyReference(blk, left, "string concat cleanup (left)")
			}

			if !expr.Right.IsPersistent() {
				emt.DestroyReference(blk, right, "string concat cleanup (right)")
			}

			return newStr
		}

	case boundnodes.Subtraction:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSub(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSub(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFSub(left, right)
		}

	case boundnodes.Multiplication:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFMul(left, right)
		}

	case boundnodes.Division:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFDiv(left, right)
		}

	case boundnodes.Modulus:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFRem(left, right)
		}

	case boundnodes.BitwiseAnd:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewAnd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewAnd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewAnd(left, right)
		}

	case boundnodes.BitwiseOr:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewOr(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewOr(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewOr(left, right)
		}

	case boundnodes.BitwiseXor:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewXor(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewXor(left, right)
		}

	case boundnodes.Equals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// compare left and right using the string class' equal function
			result := (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Equal"], left, right)

			// if left and right aren't variables (meaning they are already memory managed)
			// free() them
			if !expr.Left.IsPersistent() {
				emt.DestroyReference(blk, left, "string compare cleanup (left)")
			}

			if !expr.Right.IsPersistent() {
				emt.DestroyReference(blk, right, "string compare cleanup (right)")
			}

			// return the result
			return result
		}

	case boundnodes.NotEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredONE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// compare left and right using the string class' equal function
			result := (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Equal"], left, right)

			// if left and right aren't variables (meaning they are already memory managed)
			// free() them
			if !expr.Left.IsPersistent() {
				emt.DestroyReference(blk, left, "string compare cleanup (left)")
			}

			if !expr.Right.IsPersistent() {
				emt.DestroyReference(blk, right, "string compare cleanup (right)")
			}

			// to check if they are unequal, negate the result
			return (*blk).NewICmp(enum.IPredEQ, result, CI32(0))
		}

	case boundnodes.Greater:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGT, left, right)
		}

	case boundnodes.GreaterOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGE, left, right)
		}

	case boundnodes.Less:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOLT, left, right)
		}

	case boundnodes.LessOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOLE, left, right)
		}

	case boundnodes.LogicalAnd:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewAnd(left, right)
		}

	case boundnodes.LogicalOr:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewOr(left, right)
		}

	}

	fmt.Println("Unknown Binary!")
	fmt.Println(expr.Op.OperatorKind)
	fmt.Println(left)
	fmt.Println(right)
	return nil
}

func (emt *Emitter) EmitTernaryExpression(blk **ir.Block, expr boundnodes.BoundTernaryExpressionNode) value.Value {
	// phew okay well here we go...
	// ----------------------------

	// generate the blocks we need
	IfBlock := (*blk).Parent.NewBlock(string(expr.IfLabel))
	ElseBlock := (*blk).Parent.NewBlock(string(expr.ElseLabel))
	EndBlock := (*blk).Parent.NewBlock(string(expr.EndLabel))

	// emit our temp variable in the root block of the function
	varName := emt.Id(expr.Tmp)
	local := emt.Function.Blocks[0].NewAlloca(emt.IRTypes(expr.Tmp.VarType()))
	local.SetName(varName)
	emt.Locals[varName] = Local{IRLocal: local, IRBlock: (*blk), Type: expr.Tmp.VarType()}
	emt.EmitVariableDeclaration(blk, expr.Tmp, true)

	// emit the conditional jump
	cond := emt.EmitExpression(blk, expr.Condition)
	(*blk).NewCondBr(cond, IfBlock, ElseBlock)

	// emit the IF block
	emt.EmitAssignment(&IfBlock, expr.Tmp, expr.If)
	IfBlock.NewBr(EndBlock)

	// emit the ELSE block
	emt.EmitAssignment(&ElseBlock, expr.Tmp, expr.Else)
	ElseBlock.NewBr(EndBlock)

	// set the current working block to the END block
	(*blk) = EndBlock

	// return the variable as the exressions value
	return emt.EmitVariable(blk, expr.Tmp)
}

func (emt *Emitter) EmitCallExpression(blk **ir.Block, expr boundnodes.BoundCallExpressionNode) value.Value {
	arguments := make([]value.Value, 0)

	for _, arg := range expr.Arguments {
		expression := emt.EmitExpression(blk, arg)

		// if this is an object -> increase its reference counter
		// (only do this for variables)
		if arg.IsPersistent() {
			if arg.Type().Fingerprint() == builtins.String.Fingerprint() ||
				arg.Type().Fingerprint() == builtins.Any.Fingerprint() {
				emt.CreateReference(blk, expression, "copy to be passed into a parameter")
			}
		}

		arguments = append(arguments, expression)
	}

	functionName := emt.Id(expr.Function)

	var call *ir.InstCall

	if emt.IsInClass && !expr.Function.BuiltIn {
		// prepend the "$me" parameter to the given arguments
		arguments = append([]value.Value{emt.Function.Params[0]}, arguments...)
		call = (*blk).NewCall(emt.Classes[emt.Id(emt.ClassSym.Type)].Functions[functionName], arguments...)
	} else {
		call = (*blk).NewCall(emt.Functions[functionName].IRFunction, arguments...)
	}

	// if this is an external function it doesn't implement the garbage collector
	// meaning we have to clean up its arguments ourselves
	if expr.Function.BuiltIn {
		for i, arg := range arguments {
			if expr.Arguments[i].Type().IsObject {
				emt.DestroyReference(blk, arg, "ReturnARC of system function '"+functionName+"'")
			}
		}
	}

	return call
}

func (emt *Emitter) EmitTypeCallExpression(blk **ir.Block, expr boundnodes.BoundTypeCallExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	switch expr.Function.Fingerprint() {
	case builtins.GetLength.Fingerprint():
		// call the get length function on the string
		return (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetLength"], base)

	case builtins.Substring.Fingerprint():
		start := emt.EmitExpression(blk, expr.Arguments[0])
		length := emt.EmitExpression(blk, expr.Arguments[1])

		// call the substring function on the string
		return (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Substring"], base, start, length)
	case builtins.GetArrayLength.Fingerprint():
		// call the get length function on the array

		// object arrays
		if expr.Base.Type().SubTypes[0].IsObject {
			return (*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["GetLength"], base)
		} else {
			// primitive arrays
			return (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetLength"], base)
		}

	case builtins.Push.Fingerprint():
		element := emt.EmitExpression(blk, expr.Arguments[0])

		// Push() for object arrays
		// call the array's push function
		(*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["Push"], base, element)

		// if the element wasn't a variable -> decrease its reference counter
		if !expr.Arguments[0].IsPersistent() {
			emt.DestroyReference(blk, element, "array push cleanup")
		}

		// no need for a return value, this is a void
		return nil

	case builtins.PPush.Fingerprint():
		element := emt.EmitExpression(blk, expr.Arguments[0])

		// Push() for primitive arrays
		// grow the array and get the elem pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["Grow"], base)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		(*blk).NewStore(element, castedPtr)

		// no need for a return value, this is a void
		return nil

	case builtins.Kill.Fingerprint():
		return (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Kill"], base)

	case builtins.Start.Fingerprint():
		return (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Start"], base)

	case builtins.Join.Fingerprint():
		return (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Join"], base)
	}

	fmt.Println("Unknown TypeCall!")
	return nil
}

func (emt *Emitter) EmitClassCallExpression(blk **ir.Block, expr boundnodes.BoundClassCallExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// emit all arguments
	args := make([]value.Value, 0)
	args = append(args, base)
	for _, arg := range expr.Arguments {
		args = append(args, emt.EmitExpression(blk, arg))
	}

	return (*blk).NewCall(emt.Classes[emt.Id(expr.Base.Type())].Functions[emt.Id(expr.Function)], args...)
}

func (emt *Emitter) EmitConversionExpression(blk **ir.Block, expr boundnodes.BoundConversionExpressionNode) value.Value {
	value := emt.EmitExpression(blk, expr.Expression)

	// TODO: this is kind of an ugly function, maybe it can be implemented a little nicer

	if expr.ToType.Fingerprint() == builtins.Any.Fingerprint() {

		// if this is a string we only need to change the pointer type as it's already an object
		if expr.Expression.Type().IsObject {
			// change the pointer type
			return (*blk).NewBitCast(value, emt.IRTypes(builtins.Any))

		} else {
			// if it's not a string it needs to be boxed
			boxedValue := emt.Box(blk, value, expr.Expression.Type())
			return (*blk).NewBitCast(boxedValue, emt.IRTypes(builtins.Any))
		}

		// to string conversion
	} else if expr.ToType.Fingerprint() == builtins.String.Fingerprint() {
		switch expr.Expression.Type().Fingerprint() {
		case builtins.Any.Fingerprint():
			// change the pointer type from any to string
			return (*blk).NewBitCast(value, emt.IRTypes(builtins.String))
		case builtins.Bool.Fingerprint():
			trueStr := emt.GetStringConstant(blk, "true")
			falseStr := emt.GetStringConstant(blk, "false")

			strObj := emt.CreateObject(blk, emt.Id(builtins.String))
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, (*blk).NewSelect(value, trueStr, falseStr))

			return strObj
		case builtins.Int.Fingerprint(), builtins.Byte.Fingerprint():
			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%d"), value)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%d"), value)

			// create a new string object
			strObj := emt.CreateObject(blk, emt.Id(builtins.String))
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj

		case builtins.Float.Fingerprint():
			// convert float to double, idk why but it doesnt work without it
			double := (*blk).NewFPExt(value, types.Double)

			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%g"), double)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%g"), double)

			// create a new string object
			strObj := emt.CreateObject(blk, emt.Id(builtins.String))
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj
		}

		// to bool conversions
	} else if expr.ToType.Fingerprint() == builtins.Bool.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			// see if the string we got is equal to "true"
			result := (*blk).NewCall(emt.CFuncs["strcmp"], (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], value), emt.GetStringConstant(blk, "true"))

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "string to bool conversion cleanup")
			}

			// to check if they are equal, check if the result is 0
			return (*blk).NewICmp(enum.IPredEQ, result, CI32(0))

		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// bitcast to boxed bool
			boxedBool := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Bool)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Bool)].Functions["GetValue"], boxedBool)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to bool conversion cleanup")
			}

			return primitive
		}

		// to int conversion
	} else if expr.ToType.Fingerprint() == builtins.Int.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			result := (*blk).NewCall(emt.CFuncs["atoi"], (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], value))

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "string to int conversion cleanup")
			}

			return result

		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// bitcast to boxed int
			boxedInt := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Int)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Int)].Functions["GetValue"], boxedInt)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to int conversion cleanup")
			}

			return primitive
		} else if expr.Expression.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			// extend the byte to an int
			result := (*blk).NewSExt(value, emt.IRTypes(builtins.Int))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			// extend the byte to an int
			result := (*blk).NewFPToSI(value, emt.IRTypes(builtins.Int))
			return result
		}
	} else if expr.ToType.Fingerprint() == builtins.Byte.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			// truncate the int to a byte
			result := (*blk).NewTrunc(value, emt.IRTypes(builtins.Byte))
			return result
		}
	} else if expr.ToType.Fingerprint() == builtins.Float.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			result := (*blk).NewCall(emt.CFuncs["atof"], (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], value))

			// convert the result from a double to a float
			floatRes := (*blk).NewFPTrunc(result, types.Float)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "string to float conversion cleanup")
			}

			return floatRes

		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// bitcast to boxed float
			boxedFloat := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Float)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Float)].Functions["GetValue"], boxedFloat)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to float conversion cleanup")
			}

			return primitive
		} else if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Float))

			return result
		}
	} else if expr.ToType.Name == builtins.Array.Name {
		if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// object arrays
			if expr.ToType.SubTypes[0].IsObject {
				// change the pointer type
				return (*blk).NewBitCast(value, emt.IRTypes(builtins.Array))
			} else {
				// change the pointer type
				return (*blk).NewBitCast(value, emt.IRTypes(builtins.PArray))
			}
		}
	}

	fmt.Println("Unknown Conversion!")
	return nil
}

// </EXPRESSIONS>--------------------------------------------------------------
// <UTILS>---------------------------------------------------------------------

func (emt *Emitter) DefaultConstant(blk **ir.Block, typ symbols.TypeSymbol) constant.Constant {
	switch typ.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(false)
	case builtins.Byte.Fingerprint():
		return constant.NewInt(types.I8, 0)
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, 0)
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, 0)
	case builtins.String.Fingerprint():
		return constant.NewNull(emt.IRTypes(builtins.String).(*types.PointerType))
	}

	if typ.Name == builtins.Array.Name {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	if typ.IsUserDefined {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	fmt.Println("Unknown Constant!")
	return nil
}

func (emt *Emitter) GetStringConstant(blk **ir.Block, literal string) value.Value {

	// check if this literal has already been created
	val, ok := emt.StrConstants[literal]
	if ok {
		return (*blk).NewGetElementPtr(types.NewArray(uint64(len(literal)+1), types.I8), val, CI32(0), CI32(0))
	}

	// add a null byte at the end
	str := literal + "\x00"

	// create a global to store our literal
	global := emt.Module.NewGlobalDef(fmt.Sprintf(".str.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
	global.Immutable = true
	emt.StrNameCounter++

	pointer := (*blk).NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), global, CI32(0), CI32(0))
	emt.StrConstants[literal] = global

	return pointer
}

func (emt *Emitter) GetConstantStringConstant(literal string) constant.Constant {
	// add a null byte at the end
	str := literal + "\x00"

	// create a global to store our literal
	global := emt.Module.NewGlobalDef(fmt.Sprintf(".str.c.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
	global.Immutable = true
	emt.StrNameCounter++
	pointer := constant.NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), global, CIC32(0), CIC32(0))
	emt.StrConstants[literal] = global

	return pointer
}

func (emt *Emitter) GetThreadWrapper(source symbols.FunctionSymbol) *ir.Func {
	wrapper, ok := emt.FunctionWrappers[emt.Id(source)]

	// if there's already a wrapper for this function, return it
	if ok {
		return wrapper
	}

	// if there's no wrapper -> create one
	newWrapper := emt.Module.NewFunc(emt.Id(source)+"_ThreadWrapper", types.I8Ptr, ir.NewParam("param", types.I8Ptr))

	// create a root block
	root := newWrapper.NewBlock("")

	// add its instructions
	root.NewCall(emt.Functions[emt.Id(source)].IRFunction)
	root.NewRet(constant.NewNull(types.I8Ptr))

	// register the wrapper if we need to use it again later
	emt.FunctionWrappers[emt.Id(source)] = newWrapper

	// return the new wrapper
	return newWrapper
}

func (emt *Emitter) CreateObject(blk **ir.Block, typ string, args ...value.Value) value.Value {
	size := (*blk).NewGetElementPtr(emt.Classes[typ].Type, constant.NewNull(types.NewPointer(emt.Classes[typ].Type)), CI32(1))
	sizeInt := (*blk).NewPtrToInt(size, types.I32)

	// create space for the instance
	instance := (*blk).NewBitCast((*blk).NewCall(emt.CFuncs["malloc"], sizeInt), types.NewPointer(emt.Classes[typ].Type))

	// get pointer to instance
	instancePointer := (*blk).NewGetElementPtr(emt.Classes[typ].Type, instance, CI32(0))

	// contructor arguments
	arguments := []value.Value{instancePointer}
	arguments = append(arguments, args...)

	// call the constructor
	(*blk).NewCall(emt.Classes[typ].Constructor, arguments...)

	// create reference
	emt.CreateReference(blk, instancePointer, "initial instance")

	return instancePointer
}

func (emt *Emitter) SizeOf(blk **ir.Block, typ symbols.TypeSymbol) value.Value {
	// calculate the size by calculating the position of the second array element in an array starting at NULL
	size := (*blk).NewGetElementPtr(emt.IRTypes(typ), constant.NewNull(types.NewPointer(emt.IRTypes(typ))), CI32(1))

	// convert the size to an integer
	sizeInt := (*blk).NewPtrToInt(size, types.I32)

	return sizeInt
}

func (emt *Emitter) Box(blk **ir.Block, val value.Value, typ symbols.TypeSymbol) value.Value {
	// boxing is the act of "objectifying" primitive types
	// (like an int or bool)

	// create a new object and give it the primitive to be "capsuled"
	obj := emt.CreateObject(blk, emt.Id(typ), val)

	return obj
}

// ARC FUNCTIONS
func (emt *Emitter) CreateReference(blk **ir.Block, expr value.Value, comment string) {
	if verboseARC {
		emt.CreateReferenceVerbose(blk, expr, emt.GetStringConstant(blk, comment))
	} else {
		emt.CreateReferenceNormal(blk, expr)
	}
}

func (emt *Emitter) DestroyReference(blk **ir.Block, expr value.Value, comment string) {
	if verboseARC {
		emt.DestroyReferenceVerbose(blk, expr, emt.GetStringConstant(blk, comment))
	} else {
		emt.DestroyReferenceNormal(blk, expr)
	}
}

// NORMAL ARC FUNCTIONS
func (emt *Emitter) CreateReferenceNormal(blk **ir.Block, expr value.Value) {
	// bitcast the expression to an Any-Pointer
	// (meaning we dont change any data, we only change the pointer type)
	any := (*blk).NewBitCast(expr, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type))
	(*blk).NewCall(emt.ArcFuncs["RegisterReference"], any)
}

func (emt *Emitter) DestroyReferenceNormal(blk **ir.Block, expr value.Value) {
	// bitcast the expression to an Any-Pointer
	// (meaning we dont change any data, we only change the pointer type)
	any := (*blk).NewBitCast(expr, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type))
	(*blk).NewCall(emt.ArcFuncs["UnregisterReference"], any)
}

// DEBUG ARC FUNCTIONS
func (emt *Emitter) CreateReferenceVerbose(blk **ir.Block, expr value.Value, comment value.Value) {
	// bitcast the expression to an Any-Pointer
	// (meaning we dont change any data, we only change the pointer type)
	any := (*blk).NewBitCast(expr, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type))
	(*blk).NewCall(emt.ArcFuncs["RegisterReferenceVerbose"], any, comment)
}

func (emt *Emitter) DestroyReferenceVerbose(blk **ir.Block, expr value.Value, comment value.Value) {
	// bitcast the expression to an Any-Pointer
	// (meaning we dont change any data, we only change the pointer type)
	any := (*blk).NewBitCast(expr, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type))
	(*blk).NewCall(emt.ArcFuncs["UnregisterReferenceVerbose"], any, comment)
}

func (emt *Emitter) EmitVariableDeclaration(blk **ir.Block, varibale symbols.LocalVariableSymbol, isTmp bool) {
	varName := emt.Id(varibale)
	expression := emt.DefaultConstant(blk, varibale.VarType())

	local := emt.Locals[varName]
	local.IsSet = true
	emt.Locals[varName] = local

	// emit its assignemnt
	(*blk).NewStore(expression, local.IRLocal)

	// if this is a temp, add it to the list for cleanup
	if isTmp {
		emt.Temps = append(emt.Temps, varName)
	}
}

func (emt *Emitter) EmitAssignment(blk **ir.Block, variable symbols.LocalVariableSymbol, val boundnodes.BoundExpressionNode) {
	varName := emt.Id(variable)
	expression := emt.EmitExpression(blk, val)

	// if the expression is a variable -> increase reference counter
	if val.IsPersistent() && val.Type().IsObject {
		emt.CreateReference(blk, expression, "variable assignment ["+varName+"]")
	}

	if variable.IsGlobal() {
		// if this variable already contained an object -> destroy the reference
		if variable.VarType().IsObject {
			emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(variable.VarType()), emt.Globals[varName].IRGlobal), "destroying reference previously stored in '"+varName+"'")
		}

		// assign the vTernaralue to the global variable
		(*blk).NewStore(expression, emt.Globals[varName].IRGlobal)

	} else {
		// if this variable already contained an object -> destroy there reference
		if variable.VarType().IsObject {
			emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(variable.VarType()), emt.Locals[varName].IRLocal), "destroying reference previously stored in '"+varName+"'")
		}

		// assign the value to the local variable
		(*blk).NewStore(expression, emt.Locals[varName].IRLocal)
	}
}

func (emt *Emitter) EmitArrayAssignment(blk **ir.Block, base value.Value, index value.Value, literalValue boundnodes.BoundExpressionNode) {
	// assignment
	// ----------
	value := emt.EmitExpression(blk, literalValue)

	// decide if we should do object or primitive array access
	if literalValue.Type().IsObject {
		// bitcast our pointer to any
		anyValue := (*blk).NewBitCast(value, emt.IRTypes(builtins.Any))

		// call the array's set element function
		(*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["SetElement"], base, index, anyValue)

		// if the element wasnt a variable -> decrease its reference counter
		if !literalValue.IsPersistent() && literalValue.Type().IsObject {
			emt.DestroyReference(blk, value, "array assignment cleanup")
		}

		// return a copy of the value (i really don't think this is necessary but oh well)
		//emt.CreateReference(blk, value, "assign value copy (array assignment)")

		//return value
	} else {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(literalValue.Type())))

		(*blk).NewStore(value, castedPtr)

		// return a copy of the value
		//return value
	}
}

// </UTILS>--------------------------------------------------------------------
//<HELPERS>--------------------------------------------------------------------

// Ternary (because its kinda nice)
func tern(cond bool, str1 string, str2 string) string {
	if cond {
		return str1
	} else {
		return str2
	}
}

// a little function to get the relevant name of a symbol (name or fingerprint)
func (emt *Emitter) Id(sym symbols.Symbol) string {
	return tern(emt.UseFingerprints, sym.Fingerprint(), sym.SymbolName())
}

// another little shortcut to create an integer constant
func CI32(val int32) value.Value {
	return constant.NewInt(types.I32, int64(val))
}

func CIC32(val int32) constant.Constant {
	return constant.NewInt(types.I32, int64(val))
}

// b o o l
func CB1(val bool) value.Value {
	return constant.NewBool(val)
}

//</HELPERS>-------------------------------------------------------------------
