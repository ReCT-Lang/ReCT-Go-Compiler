package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"strings"

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
	ExcFuncs map[string]*ir.Func

	// referenced classes
	Classes map[string]*Class

	// referenced structs
	Structs map[string]*Struct

	// referenced packages
	Packages map[string]*Package

	// global variables
	Globals          map[string]Global
	Functions        map[string]Function
	FunctionWrappers map[string]*ir.Func
	FunctionLocals   map[string]map[string]Local
	StrConstants     map[string]value.Value
	StrNameCounter   int

	// local things for this current class
	Class    *Class
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

var VerboseARC = false
var EmitDebugInfo = true

var CompileAsPackage bool
var PackageName string

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:          program,
		Module:           ir.NewModule(),
		UseFingerprints:  useFingerprints,
		Globals:          make(map[string]Global),
		Functions:        make(map[string]Function),
		CFuncs:           make(map[string]*ir.Func),
		ArcFuncs:         make(map[string]*ir.Func),
		ExcFuncs:         make(map[string]*ir.Func),
		FunctionLocals:   make(map[string]map[string]Local),
		StrConstants:     make(map[string]value.Value),
		Classes:          make(map[string]*Class),
		Structs:          make(map[string]*Struct),
		FunctionWrappers: make(map[string]*ir.Func),
		Temps:            make([]string, 0),
		Packages:         make(map[string]*Package),
	}

	emitter.EmitBuiltInFunctions()

	if EmitDebugInfo {
		emitter.InitDbg()
	}

	// import all package functions and classes
	for _, pck := range emitter.Program.Packages {
		emitter.ImportPackage(pck)
	}

	// declare all struct structs
	for _, stc := range emitter.Program.Structs {
		emitter.EmitStruct(stc)
	}

	// declare all class structs
	for _, cls := range emitter.Program.Classes {
		emitter.EmitClass(cls)
	}

	// populate all struct structs
	for _, stc := range emitter.Program.Structs {
		emitter.PopulateStruct(emitter.Structs[emitter.Id(stc.Type)], stc)
	}

	// populate all class structs
	for _, cls := range emitter.Program.Classes {
		emitter.PopulateClass(emitter.Classes[emitter.Id(cls.Symbol.Type)], cls)

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
		if !fnc.Symbol.BuiltIn && !(CompileAsPackage && fnc.Symbol.Name == "main") {
			function := Function{IRFunction: emitter.EmitFunction(fnc.Symbol, fnc.Body), BoundFunction: fnc}
			functionName := emitter.Id(fnc.Symbol)
			emitter.Functions[functionName] = function
		}
	}

	// adapt any external functions that might need it
	emitter.Adapt()

	// declare all external functions
	for _, fnc := range emitter.Program.ExternalFunctions {
		function := Function{IRFunction: emitter.EmitExternalFunction(fnc), BoundFunction: binder.BoundFunction{Symbol: fnc}}
		functionName := emitter.Id(fnc)
		emitter.Functions[functionName] = function
	}

	// emit main function first
	if !CompileAsPackage {
		mainName := emitter.Id(program.MainFunction)
		emitter.FunctionSym = emitter.Functions[mainName].BoundFunction.Symbol
		emitter.EmitBlockStatement(emitter.Functions[mainName].BoundFunction.Symbol, emitter.Functions[mainName].IRFunction, emitter.Functions[mainName].BoundFunction.Body)
	}

	// emit function bodies
	for _, fnc := range emitter.Functions {
		if !fnc.BoundFunction.Symbol.BuiltIn && fnc.BoundFunction.Symbol.Fingerprint() != program.MainFunction.Fingerprint() && !fnc.BoundFunction.Symbol.External {
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
				emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor.Blocks = emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor.Blocks[1:]
				emitter.Classes[emitter.Id(cls.Symbol.Type)].Destructor.Blocks[0].LocalID = 0

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

// <STRUCTS>-------------------------------------------------------------------
func (emt *Emitter) EmitStruct(stc symbols.StructSymbol) {
	// create the class object to keep track of things
	emt.Structs[emt.Id(stc.Type)] = &Struct{
		Name:   stc.Name,
		Type:   &types.StructType{},
		Symbol: stc,
	}
}

func (emt *Emitter) PopulateStruct(stc *Struct, bstc symbols.StructSymbol) {
	// ---------------------------------------------------------------------
	// create the llvm type
	// ---------------------------------------------------------------------

	// list of our type's fields
	stcFields := make([]types.Type, 0)
	stcFieldMap := make(map[string]int)

	// add our types one by one
	for i, field := range bstc.Fields {
		stcFieldMap[emt.Id(field)] = i
		stcFields = append(stcFields, emt.IRTypes(field.VarType()))
	}

	// update the fields inside the struct
	stc.Type.(*types.StructType).Fields = append(stc.Type.(*types.StructType).Fields, stcFields...)
	emt.Structs[emt.Id(bstc.Type)].Fields = stcFieldMap

	// create the struct definition
	emt.Module.NewTypeDef("struct."+bstc.Name, stc.Type)
}

// </STRUCTS>------------------------------------------------------------------
// <CLASSES>-------------------------------------------------------------------
func (emt *Emitter) EmitClass(cls binder.BoundClass) {
	// create the class object to keep track of things
	emt.Classes[emt.Id(cls.Symbol.Type)] = &Class{
		Name: cls.Symbol.Name,
		Type: &types.StructType{},
	}

	// if this is a package
	if CompileAsPackage {
		// create a constant which holds all the field names
		fieldNameStrings := make([]constant.Constant, 0)
		for _, fld := range cls.Symbol.Fields {
			fieldNameStrings = append(fieldNameStrings, emt.GetConstantStringConstant(fld.SymbolName()))
		}

		fieldNameArr := constant.NewArray(types.NewArray(uint64(len(fieldNameStrings)), types.I8Ptr), fieldNameStrings...)
		emt.Module.NewGlobalDef(cls.Symbol.Name+"_Fields_Const", fieldNameArr)
	}
}

func (emt *Emitter) PopulateClass(cls *Class, bcls binder.BoundClass) {
	// create the class' vTable type
	//clsvTable := emt.Module.NewTypeDef("struct."+bcls.Symbol.Name+"_vTable",
	//	types.NewStruct(
	//		types.NewPointer(emt.Classes[emt.Id(builtins.Any)].vTable),
	//		types.I8Ptr,
	//		types.NewPointer(types.NewFunc(types.Void, types.I8Ptr)),
	//	),
	//)

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

	clsFields = append(clsFields, emt.Classes[emt.Id(builtins.Any)].vTable) // standard vTable
	clsFields = append(clsFields, types.I32)                                // ARC counter

	// add our types one by one
	for i, field := range bcls.Symbol.Fields {
		clsFieldMap[emt.Id(field)] = i + 2
		clsFields = append(clsFields, emt.IRTypes(field.VarType()))
	}

	// update the fields inside the struct
	cls.Type.(*types.StructType).Fields = append(cls.Type.(*types.StructType).Fields, clsFields...)

	// create the class' struct definition
	emt.Module.NewTypeDef("struct.class_"+bcls.Symbol.Name, cls.Type)

	// ---------------------------------------------------------------------
	// create the destructor
	// ---------------------------------------------------------------------
	destructor := emt.EmitClassDestructor(cls, bcls, clsFieldMap)

	// create the vTable constant
	vtc := constant.NewStruct(
		emt.Classes[emt.Id(builtins.Any)].vTable.(*types.StructType), //clsvtable

		constant.NewNull(types.I8Ptr),
		emt.GetConstantStringConstant(bcls.Symbol.Name),
		destructor,
		constant.NewNull(types.I8Ptr),
	)
	clsvConstant := emt.Module.NewGlobalDef(bcls.Symbol.Name+"_vTable_Const", vtc)

	// ---------------------------------------------------------------------
	// create the constructor
	// ---------------------------------------------------------------------
	constructor := emt.EmitClassConstructor(cls, bcls, clsvConstant, clsFieldMap)

	// store all of this info in the class object
	emt.Classes[emt.Id(bcls.Symbol.Type)].vTable = emt.Classes[emt.Id(builtins.Any)].vTable //clsvtable
	emt.Classes[emt.Id(bcls.Symbol.Type)].vConstant = clsvConstant

	emt.Classes[emt.Id(bcls.Symbol.Type)].Constructor = constructor
	emt.Classes[emt.Id(bcls.Symbol.Type)].Destructor = destructor

	emt.Classes[emt.Id(bcls.Symbol.Type)].Functions = make(map[string]*ir.Func)
	emt.Classes[emt.Id(bcls.Symbol.Type)].Fields = clsFieldMap
}

func (emt *Emitter) EmitClassConstructor(cls *Class, bcls binder.BoundClass, clsvConstant value.Value, clsFieldMap map[string]int) *ir.Func {
	// ---------------------------------------------------------------------
	// create the constructor
	// ---------------------------------------------------------------------

	// create an IR function definition for the constructor
	clsParams := make([]*ir.Param, 0)
	clsParams = append(clsParams, ir.NewParam("me", types.NewPointer(cls.Type)))

	// look for an explicit constructor
	var constructorFunction *binder.BoundFunction
	for _, fnc := range bcls.Functions {
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

	constructor := emt.Module.NewFunc(bcls.Symbol.Name+"_public_Constructor", types.Void, clsParams...)

	// create a basic constructor in IR
	// --------------------------------

	// get the objects own reference
	croot := constructor.NewBlock("")
	clsMePtr := croot.NewAlloca(types.NewPointer(cls.Type))
	croot.NewStore(clsParams[0], clsMePtr)
	//clsMe := croot.NewLoad(types.NewPointer(cls.Type), clsMePtr)

	// set the vTable to the vTable constant
	//clsMyVTable := croot.NewGetElementPtr(cls.Type, clsMe, CI32(0), CI32(0))
	//croot.NewStore(clsvConstant, clsMyVTable)

	// set the reference count to 0
	//clsMyRefCount := croot.NewGetElementPtr(cls.Type, clsMe, CI32(0), CI32(1))
	//croot.NewStore(CI32(0), clsMyRefCount)

	// store a NULL in any object fields, to tell the ARC that they are empty
	for _, field := range bcls.Symbol.Fields {
		if field.VarType().IsObject {
			ptr := croot.NewGetElementPtr(cls.Type, constructor.Params[0], CI32(0), CI32(int32(clsFieldMap[emt.Id(field)])))
			croot.NewStore(constant.NewNull(emt.IRTypes(field.VarType()).(*types.PointerType)), ptr)
		}
	}

	// create locals array for constructor
	// if the constructor has local variables they need to be present in there
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
		emt.FunctionLocals[bcls.Symbol.Name+"_private_"+emt.Id(constructorFunction.Symbol)] = locals
	}

	// done, constructor constructed
	croot.NewRet(nil)

	return constructor
}

func (emt *Emitter) EmitClassDestructor(cls *Class, bcls binder.BoundClass, clsFieldMap map[string]int) *ir.Func {
	// ---------------------------------------------------------------------
	// create the destructor
	// ---------------------------------------------------------------------

	// create an IR function definition for the destructor
	destructor := emt.Module.NewFunc(bcls.Symbol.Name+"_public_Die", types.Void, ir.NewParam("obj", types.I8Ptr))

	// create a new root block
	root := destructor.NewBlock("")

	// bitcast the given void* into a class pointer
	clsPtr := root.NewBitCast(destructor.Params[0], types.NewPointer(cls.Type))

	// generate cleanup for all object members
	root.Insts = append(root.Insts, NewComment("<DieARC>"))

	for _, field := range bcls.Symbol.Fields {
		if field.VarType().IsObject {
			root.Insts = append(root.Insts, NewComment(fmt.Sprintf("-> destroying reference to '%s [Field %d]'", emt.Id(field), clsFieldMap[emt.Id(field)])))

			// get the field's pointer
			ptr := root.NewGetElementPtr(cls.Type, clsPtr, CI32(0), CI32(int32(clsFieldMap[emt.Id(field)])))

			// load the pointer
			obj := root.NewLoad(emt.IRTypes(field.VarType()), ptr)

			// decrement its ARC reference counter
			emt.DestroyReference(&root, obj, "")
		}
	}

	root.Insts = append(root.Insts, NewComment("</DieARC>"))

	// look for an explicit destructor
	var destructorFunction *binder.BoundFunction
	for _, fnc := range bcls.Functions {
		if fnc.Symbol.Name == "Die" {
			destructorFunction = &fnc
			break
		}
	}

	// create a second block for any local variable defintions
	// if there are none this empty block will be optimized away by llvm
	decl := destructor.NewBlock("$decl")
	root.NewBr(decl)

	// create locals array for destructor
	// if the destructor has local variables they need to be present in there
	if destructorFunction != nil {
		// create locals array
		locals := make(map[string]Local)

		// create all needed locals in the decl block so GC can trash them anywhere
		for _, stmt := range destructorFunction.Body.Statements {
			if stmt.NodeType() == boundnodes.BoundVariableDeclaration {
				declStatement := stmt.(boundnodes.BoundVariableDeclarationStatementNode)

				if declStatement.Variable.IsGlobal() {
					continue
				}

				varName := emt.Id(declStatement.Variable)

				// create local variable
				local := decl.NewAlloca(emt.IRTypes(declStatement.Variable.VarType()))
				local.SetName(varName)

				// save it for referencing later
				locals[varName] = Local{IRLocal: local, IRBlock: decl, Type: declStatement.Variable.VarType()}
			}
		}

		// store this for later
		emt.FunctionLocals[bcls.Symbol.Name+"_private_"+emt.Id(destructorFunction.Symbol)] = locals
	}

	decl.NewRet(nil)

	return destructor
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
	functionName := tern(CompileAsPackage, PackageName+"_"+sym.Name, emt.Id(sym))
	irName := tern(sym.Fingerprint() == emt.Program.MainFunction.Fingerprint(), "main", functionName)

	// create an IR function definition
	function := emt.Module.NewFunc(irName, returnType, params...)

	// create a root block
	root := function.NewBlock("")

	// create locals array
	locals := make(map[string]Local)

	// create a local copy of all parameters
	// this is necessary because otherwise they would be read-only
	for _, param := range sym.Parameters {
		varName := emt.Id(param)

		// create local variable
		local := root.NewAlloca(emt.IRTypes(param.VarType()))
		local.SetName("L" + varName)

		if param.VarType().IsUserDefined && !param.VarType().IsObject {
			local.Align = 1
		}

		// store the parameters value
		root.NewStore(function.Params[param.Ordinal], local)

		// save it for referencing later
		locals[varName] = Local{IRLocal: local, IRBlock: root, Type: param.VarType()}

	}

	// create all needed locals in the root block so GC can trash them anywhere
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

			if declStatement.Variable.VarType().IsUserDefined && !declStatement.Variable.VarType().IsObject {
				local.Align = 1
			}

			// save it for referencing later
			locals[varName] = Local{IRLocal: local, IRBlock: root, Type: declStatement.Variable.VarType()}

			// debuuuuuuugging
			//if EmitDebugInfo {
			//}
		}
	}

	// store this for later
	emt.FunctionLocals[functionName] = locals

	return function
}

func (emt *Emitter) EmitExternalFunction(sym symbols.FunctionSymbol) *ir.Func {
	// check if this function is already imported
	exists := false
	var fnc *ir.Func
	for _, f := range emt.Module.Funcs {
		if f.Name() == sym.Name {
			exists = true
			fnc = f
			break
		}
	}

	// already imported, just copy reference
	if exists {
		return fnc
	}

	hasStruct := false

	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := emt.Id(param)

		// create it
		prm := ir.NewParam(paramName, emt.IRTypes(param.Type))

		// structs need some special treatment
		if param.Type.IsUserDefined && !param.Type.IsObject {

			//prm.Attrs = append(prm.Attrs, ir.Byval{Typ: emt.IRTypes(param.Type)})
			prm.Typ = types.NewPointer(emt.IRTypes(param.Type))
			hasStruct = true
		}

		// store it
		params = append(params, prm)
	}

	// figure out the return type
	returnType := emt.IRTypes(sym.Type)

	// if the return type is a struct -> bointer
	if sym.Type.IsUserDefined && !sym.Type.IsObject {
		returnType = types.NewPointer(returnType)
		hasStruct = true
	}

	// the function name
	irName := sym.Name

	// if this gamer got adapted, call the adapted function instead
	if sym.Adapted {
		irName += "$ADAPTED"
	}

	// if this function has struct but isnt adapted -> warning (that shit may not work)
	if hasStruct && !sym.Adapted {
		print.Warning(
			"EMITTER",
			print.ExternalCAdapterWarning,
			sym.Declaration.Span(),
			"This external function is using structs but is not using c_adapter. This will likely cause compatibility issues!",
		)
	}

	// if this function has no structs but is adapted -> warning (that shit is redundant)
	if !hasStruct && sym.Adapted {
		print.Warning(
			"EMITTER",
			print.ExternalCAdapterWarning,
			sym.Declaration.Span(),
			"This external function is not using any structs but is using c_adapter. This is unnecessary and redundant.",
		)
	}

	// create an IR function definition
	function := emt.Module.NewFunc(irName, returnType, params...)

	if sym.Variadic {
		function.Sig.Variadic = true
	}

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

	// create a local copy of all parameters
	// this is necessary because otherwise they would be read-only
	for _, param := range sym.Parameters {
		varName := emt.Id(param)

		// create local variable
		local := root.NewAlloca(emt.IRTypes(param.VarType()))
		local.SetName("L" + varName)

		// store the parameters value
		root.NewStore(function.Params[param.Ordinal+1], local)

		// save it for referencing later
		locals[varName] = Local{IRLocal: local, IRBlock: root, Type: param.VarType()}

	}

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

		// if this is a struct we'll have to load it first
		if stmt.Initializer.Type().IsUserDefined && !stmt.Initializer.Type().IsObject {
			expression = (*blk).NewLoad(emt.IRTypes(stmt.Initializer.Type()), expression)
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
		local, ok := emt.Locals[varName]

		if !ok {
			fmt.Println("Broken local! Failed lookup for " + varName)
		}

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

	// if this isnt a package, handle end-of-function cleanup
	if !CompileAsPackage {
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
	}

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

	// cheeky comments
	if EmitDebugInfo {
		(*blk).Insts = append((*blk).Insts, NewComment("<"+string(expr.NodeType())+">"))
	}

	// return value
	var val value.Value

	switch expr.NodeType() {
	case boundnodes.BoundLiteralExpression:
		val = emt.EmitLiteralExpression(blk, expr.(boundnodes.BoundLiteralExpressionNode))

	case boundnodes.BoundVariableExpression:
		val = emt.EmitVariableExpression(blk, expr.(boundnodes.BoundVariableExpressionNode))

	case boundnodes.BoundAssignmentExpression:
		val = emt.EmitAssignmentExpression(blk, expr.(boundnodes.BoundAssignmentExpressionNode))

	case boundnodes.BoundMakeExpression:
		val = emt.EmitMakeExpression(blk, expr.(boundnodes.BoundMakeExpressionNode))

	case boundnodes.BoundMakeArrayExpression:
		val = emt.EmitMakeArrayExpression(blk, expr.(boundnodes.BoundMakeArrayExpressionNode))

	case boundnodes.BoundMakeStructExpression:
		val = emt.EmitMakeStructExpression(blk, expr.(boundnodes.BoundMakeStructExpressionNode))

	case boundnodes.BoundArrayAccessExpression:
		val = emt.EmitArrayAccessExpression(blk, expr.(boundnodes.BoundArrayAccessExpressionNode))

	case boundnodes.BoundArrayAssignmentExpression:
		val = emt.EmitArrayAssignmentExpression(blk, expr.(boundnodes.BoundArrayAssignmentExpressionNode))

	case boundnodes.BoundUnaryExpression:
		val = emt.EmitUnaryExpression(blk, expr.(boundnodes.BoundUnaryExpressionNode))

	case boundnodes.BoundBinaryExpression:
		val = emt.EmitBinaryExpression(blk, expr.(boundnodes.BoundBinaryExpressionNode))

	case boundnodes.BoundTernaryExpression:
		val = emt.EmitTernaryExpression(blk, expr.(boundnodes.BoundTernaryExpressionNode))

	case boundnodes.BoundCallExpression:
		val = emt.EmitCallExpression(blk, expr.(boundnodes.BoundCallExpressionNode))

	case boundnodes.BoundPackageCallExpression:
		val = emt.EmitPackageCallExpression(blk, expr.(boundnodes.BoundPackageCallExpressionNode))

	case boundnodes.BoundTypeCallExpression:
		val = emt.EmitTypeCallExpression(blk, expr.(boundnodes.BoundTypeCallExpressionNode))

	case boundnodes.BoundClassCallExpression:
		val = emt.EmitClassCallExpression(blk, expr.(boundnodes.BoundClassCallExpressionNode))

	case boundnodes.BoundClassFieldAccessExpression:
		val = emt.EmitClassFieldAccessExpression(blk, expr.(boundnodes.BoundClassFieldAccessExpressionNode))

	case boundnodes.BoundClassFieldAssignmentExpression:
		val = emt.EmitClassFieldAssignmentExpression(blk, expr.(boundnodes.BoundClassFieldAssignmentExpressionNode))

	case boundnodes.BoundClassDestructionExpression:
		val = emt.EmitClassDestructionExpression(blk, expr.(boundnodes.BoundClassDestructionExpressionNode))

	case boundnodes.BoundConversionExpression:
		val = emt.EmitConversionExpression(blk, expr.(boundnodes.BoundConversionExpressionNode))

	case boundnodes.BoundReferenceExpression:
		val = emt.EmitReferenceExpression(blk, expr.(boundnodes.BoundReferenceExpressionNode))

	case boundnodes.BoundDereferenceExpression:
		val = emt.EmitDereferenceExpression(blk, expr.(boundnodes.BoundDereferenceExpressionNode))

	case boundnodes.BoundLambdaExpression:
		val = emt.EmitLambdaExpression(blk, expr.(boundnodes.BoundLambdaExpressionNode))

	case boundnodes.BoundFunctionExpression:
		val = emt.EmitFunctionExpression(blk, expr.(boundnodes.BoundFunctionExpressionNode))

	case boundnodes.BoundThisExpression:
		val = emt.EmitThisExpression(blk, expr.(boundnodes.BoundThisExpressionNode))

	default:
		fmt.Println("Unimplemented node: " + expr.NodeType())
		return nil
	}

	if EmitDebugInfo {
		(*blk).Insts = append((*blk).Insts, NewComment("</"+string(expr.NodeType())+">"))
	}

	return val
}

func (emt *Emitter) EmitUnloadedReference(blk **ir.Block, expr boundnodes.BoundExpressionNode) (value.Value, value.Value) {
	switch expr.NodeType() {
	case boundnodes.BoundVariableExpression:
		exp := expr.(boundnodes.BoundVariableExpressionNode)
		return emt.EmitVariablePtr(blk, exp.Variable), nil
	case boundnodes.BoundClassFieldAccessExpression:
		exp := expr.(boundnodes.BoundClassFieldAccessExpressionNode)
		if exp.Base.Type().IsObject {
			return emt.EmitClassFieldAccessExpressionRef(blk, exp.Base, exp.Field)
		} else {
			return emt.EmitStructFieldAccessExpressionRef(blk, exp.Base, exp.Field), nil
		}
	case boundnodes.BoundClassFieldAssignmentExpression:
		exp := expr.(boundnodes.BoundClassFieldAssignmentExpressionNode)
		if exp.Base.Type().IsObject {
			return emt.EmitClassFieldAccessExpressionRef(blk, exp.Base, exp.Field)
		} else {
			return emt.EmitStructFieldAccessExpressionRef(blk, exp.Base, exp.Field), nil
		}
	case boundnodes.BoundArrayAccessExpression:
		exp := expr.(boundnodes.BoundArrayAccessExpressionNode)
		if exp.IsPointer {
			return emt.EmitPointerAccessRef(blk, exp), nil
		} else {
			return emt.EmitArrayAccessRef(blk, exp), nil
		}
	}

	fmt.Println("Unknown persistent expression! (EmitUnloadedReference)")

	return nil, nil
}

func (emt *Emitter) EmitLiteralExpression(blk **ir.Block, expr boundnodes.BoundLiteralExpressionNode) value.Value {
	switch expr.LiteralType.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(expr.Value.(bool))
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, int64(expr.Value.(int)))
	case builtins.Byte.Fingerprint():
		return constant.NewInt(types.I8, int64(expr.Value.(byte)))
	case builtins.Long.Fingerprint():
		return constant.NewInt(types.I64, int64(expr.Value.(int)))
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, float64(expr.Value.(float32)))
	case builtins.String.Fingerprint():
		charPtr := emt.GetStringConstant(blk, expr.Value.(string))
		strObj := emt.CreateObject(blk, builtins.String)
		(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, charPtr)
		return strObj
	default:
		// native strings (aka byte pointers)
		if expr.Type().Name == builtins.Pointer.Name && expr.Type().SubTypes[0].Fingerprint() == builtins.Byte.Fingerprint() {
			return emt.GetStringConstant(blk, expr.Value.(string))
		}
	}

	fmt.Println("Unknown literal!")

	return nil
}

func (emt *Emitter) EmitVariableExpression(blk **ir.Block, expr boundnodes.BoundVariableExpressionNode) value.Value {
	return emt.EmitVariable(blk, expr.Variable)
}

func (emt *Emitter) EmitVariable(blk **ir.Block, variable symbols.VariableSymbol) value.Value {
	varName := emt.Id(variable)

	//// parameters
	//if variable.SymbolType() == symbols.Parameter {
	//	paramSymbol := variable.(symbols.ParameterSymbol)
	//
	//	if emt.IsInClass {
	//		return emt.Function.Params[paramSymbol.Ordinal+1]
	//	} else {
	//		return emt.Function.Params[paramSymbol.Ordinal]
	//	}
	//}

	if variable.IsGlobal() {
		// if we're in a class we need to load from the struct instead of a global
		if emt.IsInClass {
			mePtr := value.Value(emt.Function.Params[0])

			// if we're in a Destructor, the "mePtr" will be a generic void*
			// that means it will need to be converted
			if emt.FunctionSym.Name == "Die" && len(emt.Function.Params) == 1 {
				mePtr = (*blk).NewBitCast(mePtr, types.NewPointer(emt.Class.Type))
			}

			ptr := (*blk).NewGetElementPtr(emt.Class.Type, mePtr, CI32(0), CI32(int32(emt.Class.Fields[emt.Id(variable)])))

			// if what we're accessing is a struct, do not dereference it
			if variable.VarType().IsUserDefined && !variable.VarType().IsObject {
				return ptr
			}

			return (*blk).NewLoad(emt.IRTypes(variable.VarType()), ptr)

		} else {
			// if we aren't we can just get the global
			// --------------------------------------

			// if what we're accessing is a struct, do not dereference it
			if variable.VarType().IsUserDefined && !variable.VarType().IsObject {
				return emt.Globals[varName].IRGlobal
			}

			// non-structs
			return (*blk).NewLoad(emt.IRTypes(emt.Globals[varName].Type), emt.Globals[varName].IRGlobal)
		}
	} else {
		// if what we're accessing is a struct, do not dereference it
		if variable.VarType().IsUserDefined && !variable.VarType().IsObject {
			return emt.Locals[varName].IRLocal
		}

		// non-structs
		return (*blk).NewLoad(emt.IRTypes(emt.Locals[varName].Type), emt.Locals[varName].IRLocal)
	}
}

func (emt *Emitter) EmitVariablePtr(blk **ir.Block, variable symbols.VariableSymbol) value.Value {
	varName := emt.Id(variable)

	//// parameters
	//if variable.SymbolType() == symbols.Parameter {
	//	paramSymbol := variable.(symbols.ParameterSymbol)
	//
	//	if emt.IsInClass {
	//		return emt.Function.Params[paramSymbol.Ordinal+1]
	//	} else {
	//		return emt.Function.Params[paramSymbol.Ordinal]
	//	}
	//}

	if variable.IsGlobal() {
		// if we're in a class we need to load from the struct instead of a global
		if emt.IsInClass {
			mePtr := value.Value(emt.Function.Params[0])

			// if we're in a Destructor, the "mePtr" will be a generic void*
			// that means it will need to be converted
			if emt.FunctionSym.Name == "Die" && len(emt.Function.Params) == 1 {
				mePtr = (*blk).NewBitCast(mePtr, types.NewPointer(emt.Class.Type))
			}

			ptr := (*blk).NewGetElementPtr(emt.Class.Type, mePtr, CI32(0), CI32(int32(emt.Class.Fields[emt.Id(variable)])))
			return ptr

		} else {
			// if we arent we can just get the global
			return emt.Globals[varName].IRGlobal
		}
	} else {
		return emt.Locals[varName].IRLocal
	}
}

func (emt *Emitter) EmitAssignmentExpression(blk **ir.Block, expr boundnodes.BoundAssignmentExpressionNode) value.Value {
	varName := emt.Id(expr.Variable)
	expression := emt.EmitExpression(blk, expr.Expression)

	// if the expression is a variable -> increase reference counter
	if expr.Expression.IsPersistent() && expr.Expression.Type().IsObject {
		emt.CreateReference(blk, expression, "variable assignment ["+varName+"]")
	}

	// if this is a struct we'll have to load it first
	if expr.Expression.Type().IsUserDefined && !expr.Expression.Type().IsObject {
		expression = (*blk).NewLoad(emt.IRTypes(expr.Expression.Type()), expression)
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
	obj := emt.CreateObject(blk, expr.BaseType.Type, arguments...)

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
		arrObject = emt.CreateObject(blk, expr.Type(), length)
	} else {
		// get the size of the primitive we want to allocate
		size := emt.SizeOf(blk, expr.BaseType)

		// construct the correct type symbol
		typ := expr.Type()
		typ.Name = "parray"

		// create a new primitive array object
		arrObject = emt.CreateObject(blk, typ, length, size)
	}

	// if this is a literal, load its values
	if expr.IsLiteral {
		for i, literal := range expr.Literals {
			emt.EmitArrayAssignment(blk, arrObject, CI32(int32(i)), literal)
		}
	}

	// bitcast to typed array type
	arrObject = (*blk).NewBitCast(arrObject, emt.IRTypes(symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{expr.BaseType}, true, false)))

	return arrObject
}

func (emt *Emitter) EmitMakeStructExpression(blk **ir.Block, expr boundnodes.BoundMakeStructExpressionNode) value.Value {

	// create a space to store our values in
	stc := emt.Function.Blocks[0].NewAlloca(emt.Structs[emt.Id(expr.StructType)].Type)
	//stc.Align = 1

	// go through all fields in our struct
	for i, field := range emt.Structs[emt.Id(expr.StructType)].Symbol.Fields {
		fieldIndex := emt.Structs[emt.Id(expr.StructType)].Fields[field.Fingerprint()]
		// pointer of this field
		//     (*blk).NewGetElementPtr(emt.Structs[emt.Id(base.Type())].Type, basePtr, CI32(0), CI32(int32(fieldIndex)))
		ptr := (*blk).NewGetElementPtr(emt.Structs[emt.Id(expr.StructType)].Type, stc, CI32(0), CI32(int32(fieldIndex)))
		ptr.InBounds = true

		//(*blk).NewCall(emt.CFuncs["printf"], emt.GetConstantStringConstant("%ld\n"), (*blk).NewPtrToInt(ptr, types.I64))

		// if we're still below the index to which we assigned up to -> assign the given value
		if i < len(expr.Literals) {
			val := emt.EmitExpression(blk, expr.Literals[i])

			// if this is a struct we'll have to load it first
			if field.VarType().IsUserDefined && !field.VarType().IsObject {
				val = (*blk).NewLoad(emt.IRTypes(field.VarType()), val)
			}

			(*blk).NewStore(val, ptr)

			// if we're out of values -> use default ones
		} else {
			(*blk).NewStore(emt.DefaultConstant(blk, field.VarType()), ptr)
		}
	}

	ld := (*blk).NewLoad(emt.IRTypes(expr.StructType), stc)
	//ld.Align = 1
	return ld
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

	// if this is a struct we'll have to load it first
	if expr.Value.Type().IsUserDefined && !expr.Value.Type().IsObject {
		value = (*blk).NewLoad(emt.IRTypes(expr.Value.Type()), value)
	}

	// is this actually a sneaky pointer access?
	if expr.IsPointer {
		return emt.EmitPointerAssignmentExpression(blk, expr, base, index, value)
	}

	// bitcast the base into a generic array type
	if expr.Base.Type().SubTypes[0].IsObject {
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.Array)].Type))
	} else {
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type))
	}

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

	} else {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		(*blk).NewStore(value, castedPtr)
	}

	// if base isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "array assignment base cleanup")
	}

	return value
}

func (emt *Emitter) EmitArrayAccessExpression(blk **ir.Block, expr boundnodes.BoundArrayAccessExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)
	// load the index
	index := emt.EmitExpression(blk, expr.Index)

	// is this actually a sneaky pointer access?
	if expr.IsPointer {
		return emt.EmitPointerAccessExpression(blk, expr, base, index)
	}

	// bitcast the base into a generic array type
	if expr.Base.Type().SubTypes[0].IsObject {
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.Array)].Type))
	} else {
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type))
	}

	// do the access
	// -------------
	var value value.Value

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
		value = (*blk).NewBitCast(element, types.NewPointer(emt.Classes[emt.Id(originalType)].Type))
	} else {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		// load the value
		value = (*blk).NewLoad(emt.IRTypes(expr.Base.Type().SubTypes[0]), castedPtr)
	}

	// if base isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "array access base cleanup")
	}

	return value
}

func (emt *Emitter) EmitArrayAccessRef(blk **ir.Block, expr boundnodes.BoundArrayAccessExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)
	// load the index
	index := emt.EmitExpression(blk, expr.Index)

	// bitcast the base into a generic array type
	base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type))

	// do the access
	// -------------

	// decide if we should do object or primitive array access
	if !expr.Base.Type().SubTypes[0].IsObject {
		// get the elements pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetElementPtr"], base, index)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		return castedPtr
	} else {
		fmt.Println("how tf did this even happen?? (EmitArrayAccessRef received object type (this should be literally impossible))")
		return nil
	}
}

func (emt *Emitter) EmitPointerAccessExpression(blk **ir.Block, expr boundnodes.BoundArrayAccessExpressionNode, base value.Value, index value.Value) value.Value {
	// get the elements pointer
	elementPtr := (*blk).NewGetElementPtr(emt.IRTypes(expr.Type()), base, index)

	// if what we're accessing is a struct, do not dereference it
	if expr.Type().IsUserDefined && !expr.Type().IsObject {
		return elementPtr
	}

	// load the value
	val := (*blk).NewLoad(emt.IRTypes(expr.Type()), elementPtr)
	return val
}

func (emt *Emitter) EmitPointerAccessRef(blk **ir.Block, expr boundnodes.BoundArrayAccessExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)
	// load the index
	index := emt.EmitExpression(blk, expr.Index)

	// get the elements pointer
	elementPtr := (*blk).NewGetElementPtr(emt.IRTypes(expr.Type()), base, index)
	return elementPtr
}

func (emt *Emitter) EmitPointerAssignmentExpression(blk **ir.Block, expr boundnodes.BoundArrayAssignmentExpressionNode, base value.Value, index value.Value, value value.Value) value.Value {
	// get the elements pointer
	elementPtr := (*blk).NewGetElementPtr(emt.IRTypes(expr.Type()), base, index)

	// we storin
	(*blk).NewStore(value, elementPtr)

	// return a copy of the value
	return value
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

		} else if expr.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFNeg(expression)
		} else if expr.Type().Fingerprint() == builtins.Double.Fingerprint() {
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

		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewAdd(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewAdd((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
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

		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewSub(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFSub(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewSub(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewSub(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewSub((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFSub(left, right)
		}

	case boundnodes.Multiplication:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewMul((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFMul(left, right)
		}

	case boundnodes.Division:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSDiv(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewUDiv((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewSDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewUDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewUDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFDiv(left, right)
		}

	case boundnodes.Modulus:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewSRem(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewURem((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewSRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewSRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewURem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewURem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFRem(left, right)
		}

	case boundnodes.BitwiseAnd:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewAnd(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewAnd((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewAnd(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewAnd(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewAnd(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewAnd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewAnd(left, right)
		}

	case boundnodes.BitwiseOr:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewOr(left, right)

		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewOr((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())

		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewOr(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewOr(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewOr(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewOr(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return (*blk).NewOr(left, right)
		}

	case boundnodes.BitwiseXor:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewXor(left, right)
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewXor((*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewXor(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewXor(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewXor(left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewXor(left, right)
		}

	case boundnodes.Equals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredEQ, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOEQ, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
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
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredNE, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredONE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
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
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredUGT, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredUGT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredUGT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGT, left, right)
		}

	case boundnodes.GreaterOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGE, left, right)
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredUGE, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSGE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredUGE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredUGE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOGE, left, right)
		}

	case boundnodes.Less:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLT, left, right)
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredULT, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredULT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredULT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOLT, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOLT, left, right)
		}

	case boundnodes.LessOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLE, left, right)
		} else if expr.Left.Type().Name == builtins.Pointer.Name {
			return (*blk).NewIntToPtr((*blk).NewICmp(enum.IPredULE, (*blk).NewPtrToInt(left, types.I64), (*blk).NewPtrToInt(right, types.I64)), left.Type())
		} else if expr.Left.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewICmp(enum.IPredSLE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewICmp(enum.IPredULE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewICmp(enum.IPredULE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return (*blk).NewFCmp(enum.FPredOLE, left, right)
		} else if expr.Left.Type().Fingerprint() == builtins.Double.Fingerprint() {
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
			if arg.Type().IsObject {
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

func (emt *Emitter) EmitPackageCallExpression(blk **ir.Block, expr boundnodes.BoundPackageCallExpressionNode) value.Value {
	arguments := make([]value.Value, 0)

	for _, arg := range expr.Arguments {
		expression := emt.EmitExpression(blk, arg)

		// if this is an object -> increase its reference counter
		// (only do this for variables)
		if arg.IsPersistent() {
			if arg.Type().IsObject {
				emt.CreateReference(blk, expression, "copy to be passed into a parameter")
			}
		}

		arguments = append(arguments, expression)
	}

	functionName := emt.Id(expr.Function)
	call := (*blk).NewCall(emt.Packages[emt.Id(expr.Package)].Functions[functionName], arguments...)

	// this is an external function so it doesn't implement the garbage collector
	// meaning we have to clean up its arguments ourselves
	for i, arg := range arguments {
		if expr.Arguments[i].Type().IsObject {
			emt.DestroyReference(blk, arg, "ReturnARC of external function '"+functionName+"'")
		}
	}

	return call
}

func (emt *Emitter) EmitTypeCallExpression(blk **ir.Block, expr boundnodes.BoundTypeCallExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// if this is an object type -> do a null check before calling
	if expr.Base.Type().IsObject {
		(*blk).NewCall(emt.ExcFuncs["ThrowIfNull"], (*blk).NewBitCast(base, types.I8Ptr))
	}

	var val value.Value

	switch expr.Function.Fingerprint() {
	case builtins.GetLength.Fingerprint():
		// call the get length function on the string
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetLength"], base)

	case builtins.GetBuffer.Fingerprint():
		// call the get length function on the string
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], base)

	case builtins.Substring.Fingerprint():
		start := emt.EmitExpression(blk, expr.Arguments[0])
		length := emt.EmitExpression(blk, expr.Arguments[1])

		// call the substring function on the string
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Substring"], base, start, length)
	case builtins.GetArrayLength.Fingerprint():
		// call the get length function on the array

		// object arrays
		if expr.Base.Type().SubTypes[0].IsObject {
			base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.Array)].Type))
			val = (*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["GetLength"], base)
		} else {
			// primitive arrays
			base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type))
			val = (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["GetLength"], base)
		}

	case builtins.Push.Fingerprint():
		element := emt.EmitExpression(blk, expr.Arguments[0])

		// sneaky lil bitcast
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.Array)].Type))

		// Push() for object arrays
		// call the array's push function
		(*blk).NewCall(emt.Classes[emt.Id(builtins.Array)].Functions["Push"], base, element)

		// if the element wasn't a variable -> decrease its reference counter
		if !expr.Arguments[0].IsPersistent() {
			emt.DestroyReference(blk, element, "array push cleanup")
		}

		// no need for a return value, this is a void
		val = nil

	case builtins.PPush.Fingerprint():
		element := emt.EmitExpression(blk, expr.Arguments[0])

		// sneakier liler bitcast
		base = (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type))

		// Push() for primitive arrays
		// grow the array and get the elem pointer
		elementPtr := (*blk).NewCall(emt.Classes[emt.Id(builtins.PArray)].Functions["Grow"], base)

		// bitcast the pointer to our type
		castedPtr := (*blk).NewBitCast(elementPtr, types.NewPointer(emt.IRTypes(expr.Base.Type().SubTypes[0])))

		(*blk).NewStore(element, castedPtr)

		// no need for a return value, this is a void
		val = nil

	case builtins.Kill.Fingerprint():
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Kill"], base)

	case builtins.Start.Fingerprint():
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Start"], base)

	case builtins.Join.Fingerprint():
		val = (*blk).NewCall(emt.Classes[emt.Id(builtins.Thread)].Functions["Join"], base)
	default:
		// the funky ones:
		// (these cant be identified by their fingerprint because its generated procedurally)
		if expr.Function.Name == builtins.Run.Name {
			arguments := make([]value.Value, 0)

			// emit some quirky params
			for i := range expr.Function.Parameters {
				arg := emt.EmitExpression(blk, expr.Arguments[i])

				// if this is an object -> increase its reference counter
				// (only do this for variables)
				if expr.Arguments[i].IsPersistent() {
					if expr.Arguments[i].Type().IsObject {
						emt.CreateReference(blk, arg, "copy to be passed into a parameter")
					}
				}

				arguments = append(arguments, arg)
			}

			// call the function
			//fnc := (*blk).NewLoad(emt.IRTypes(expr.Base.Type()), base)
			val = (*blk).NewCall(base, arguments...)

			// memory cleanup
			for i, arg := range arguments {
				if expr.Arguments[i].Type().IsObject {
					emt.DestroyReference(blk, arg, "ReturnARC of lambda call ->Run()")
				}
			}

			break
		}

		fmt.Println("Unknown TypeCall!")
		return nil
	}

	// if base isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if expr.Base.Type().IsObject && !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "type call base cleanup")
	}

	return val

}

func (emt *Emitter) EmitClassCallExpression(blk **ir.Block, expr boundnodes.BoundClassCallExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// run a null check on the base
	(*blk).NewCall(emt.ExcFuncs["ThrowIfNull"], (*blk).NewBitCast(base, types.I8Ptr))

	// emit all arguments
	args := make([]value.Value, 0)
	args = append(args, base)
	for _, arg := range expr.Arguments {
		args = append(args, emt.EmitExpression(blk, arg))
	}

	call := (*blk).NewCall(emt.Classes[emt.Id(expr.Base.Type())].Functions[emt.Id(expr.Function)], args...)

	// if base isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "class call base cleanup")
	}

	return call
}

func (emt *Emitter) EmitClassDestructionExpression(blk **ir.Block, expr boundnodes.BoundClassDestructionExpressionNode) value.Value {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, expr.Base)

	// run a null check on the base
	//(*blk).NewCall(emt.ExcFuncs["ThrowIfNull"], (*blk).NewBitCast(base, types.I8Ptr))

	// destroy the object (mercilessly)
	any := (*blk).NewBitCast(base, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type))
	(*blk).NewCall(emt.ArcFuncs["DestroyObject"], any)

	return (*blk).NewPtrToInt((*blk).NewBitCast(base, types.I8Ptr), types.I32)
}

func (emt *Emitter) EmitClassFieldAccessExpression(blk **ir.Block, expr boundnodes.BoundClassFieldAccessExpressionNode) value.Value {
	// okay but like, is this a struct?
	if !expr.Base.Type().IsObject {
		return emt.EmitStructFieldAccessExpression(blk, expr)
	}

	ptr, base := emt.EmitClassFieldAccessExpressionRef(blk, expr.Base, expr.Field)

	// if value isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "class field access base cleanup")
	}

	return (*blk).NewLoad(emt.IRTypes(expr.Field.VarType()), ptr)
}

func (emt *Emitter) EmitClassFieldAccessExpressionRef(blk **ir.Block, _base boundnodes.BoundExpressionNode, field symbols.VariableSymbol) (value.Value, value.Value) {
	// load the base value
	// -------------------
	base := emt.EmitExpression(blk, _base)

	// run a null check on the base
	(*blk).NewCall(emt.ExcFuncs["ThrowIfNull"], (*blk).NewBitCast(base, types.I8Ptr))

	// look up the field's index
	fieldIndex := emt.Classes[emt.Id(_base.Type())].Fields[field.Fingerprint()]

	ptr := (*blk).NewGetElementPtr(emt.Classes[emt.Id(_base.Type())].Type, base, CI32(0), CI32(int32(fieldIndex)))
	return ptr, base
}

func (emt *Emitter) EmitStructFieldAccessExpression(blk **ir.Block, expr boundnodes.BoundClassFieldAccessExpressionNode) value.Value {
	ptr := emt.EmitStructFieldAccessExpressionRef(blk, expr.Base, expr.Field)

	// if what we're accessing is a struct, do not dereference it
	if expr.Type().IsUserDefined && !expr.Type().IsObject {
		return ptr
	}

	return (*blk).NewLoad(emt.IRTypes(expr.Field.VarType()), ptr)
}

func (emt *Emitter) EmitStructFieldAccessExpressionRef(blk **ir.Block, base boundnodes.BoundExpressionNode, field symbols.VariableSymbol) value.Value {
	// look up the field's index
	fieldIndex := emt.Structs[emt.Id(base.Type())].Fields[field.Fingerprint()]

	var basePtr value.Value
	var baseExp value.Value
	var val value.Value

	if base.IsPersistent() {
		basePtr, baseExp = emt.EmitUnloadedReference(blk, base)

		if basePtr != nil {
			goto FINISH
		}
	}

	// Fallback in case we couldnt resolve a ref
	// create a local alloca to read out the struct info (this sucks)
	basePtr = (*blk).NewAlloca(emt.IRTypes(base.Type()))
	val = emt.EmitExpression(blk, base)
	(*blk).NewStore(val, basePtr)

FINISH:
	gep := (*blk).NewGetElementPtr(emt.Structs[emt.Id(base.Type())].Type, basePtr, CI32(0), CI32(int32(fieldIndex)))

	// if base isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !base.IsPersistent() {
		emt.DestroyReference(blk, baseExp, "array assignment base cleanup")
	}

	return gep
}

func (emt *Emitter) EmitClassFieldAssignmentExpression(blk **ir.Block, expr boundnodes.BoundClassFieldAssignmentExpressionNode) value.Value {
	// okay but like, is this a struct?
	if !expr.Base.Type().IsObject {
		return emt.EmitStructFieldAssignmentExpression(blk, expr)
	}

	// assignment value
	// ----------------
	value := emt.EmitExpression(blk, expr.Value)

	// if this is a struct we'll have to load it first
	if expr.Value.Type().IsUserDefined && !expr.Value.Type().IsObject {
		value = (*blk).NewLoad(emt.IRTypes(expr.Value.Type()), value)
	}

	// if the expression is a variable -> increase reference counter
	if expr.Value.IsPersistent() && expr.Value.Type().IsObject {
		emt.CreateReference(blk, value, "field assignment ["+expr.Field.SymbolName()+"]")
	}

	// the location we need to store to
	ptr, base := emt.EmitClassFieldAccessExpressionRef(blk, expr.Base, expr.Field)

	// if this variable already contained an object -> destroy the reference
	if expr.Field.VarType().IsObject {
		emt.DestroyReference(blk, (*blk).NewLoad(emt.IRTypes(expr.Field.VarType()), ptr), "destroying reference previously stored in '"+expr.Field.SymbolName()+"'")
	}

	// assign the value to the structs field
	(*blk).NewStore(value, ptr)

	// return a copy of the value (i really don't think this is necessary but oh well)
	if expr.Value.Type().IsObject {
		emt.CreateReference(blk, value, "assign value copy (field assignment)")
	}

	// if value isn't a variable (meaning its already memory managed)
	// decrement its reference counter
	if !expr.Base.IsPersistent() {
		emt.DestroyReference(blk, base, "class field assignment cleanup")
	}

	return value
}

func (emt *Emitter) EmitStructFieldAssignmentExpression(blk **ir.Block, expr boundnodes.BoundClassFieldAssignmentExpressionNode) value.Value {

	value := emt.EmitExpression(blk, expr.Value)

	// if this is a struct we'll have to load it first
	if expr.Value.Type().IsUserDefined && !expr.Value.Type().IsObject {
		value = (*blk).NewLoad(emt.IRTypes(expr.Value.Type()), value)
	}

	ptr := emt.EmitStructFieldAccessExpressionRef(blk, expr.Base, expr.Field)

	// assign the value to the structs field
	(*blk).NewStore(value, ptr)
	return value
}

func (emt *Emitter) EmitConversionExpression(blk **ir.Block, expr boundnodes.BoundConversionExpressionNode) value.Value {
	value := emt.EmitExpression(blk, expr.Expression)

	// TODO: this is kind of an ugly function, maybe it can be implemented a little nicer

	if expr.ToType.Fingerprint() == builtins.Any.Fingerprint() {

		// if this is an object we only need to change the pointer type as it's already an object
		if expr.Expression.Type().IsObject || expr.Expression.Type().Name == "pointer" {
			// change the pointer type
			return (*blk).NewBitCast(value, emt.IRTypes(builtins.Any))

		} else {
			// if it's not a string it needs to be boxed
			// -----------------------------------------

			typ := expr.Expression.Type()

			// if the thing we're boxing is a pointer -> convert it to int first
			if expr.Expression.Type().Name == "pointer" ||
				expr.Expression.Type().Name == "action" {
				value = (*blk).NewPtrToInt(value, types.I64)
				typ = builtins.Long
			}

			boxedValue := emt.Box(blk, value, typ)
			return (*blk).NewBitCast(boxedValue, emt.IRTypes(builtins.Any))
		}

		// to string conversion
	} else if expr.ToType.Fingerprint() == builtins.String.Fingerprint() {
		switch expr.Expression.Type().Fingerprint() {
		case builtins.Any.Fingerprint():
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.String, value)

			// change the pointer type from any to string
			return (*blk).NewBitCast(value, emt.IRTypes(builtins.String))
		case builtins.Bool.Fingerprint():
			trueStr := emt.GetStringConstant(blk, "true")
			falseStr := emt.GetStringConstant(blk, "false")

			strObj := emt.CreateObject(blk, builtins.String)
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
			strObj := emt.CreateObject(blk, builtins.String)
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj

		case builtins.Long.Fingerprint():
			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%ld"), value)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%ld"), value)

			// create a new string object
			strObj := emt.CreateObject(blk, builtins.String)
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj

		case builtins.UInt.Fingerprint():
			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%u"), value)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%u"), value)

			// create a new string object
			strObj := emt.CreateObject(blk, builtins.String)
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj

		case builtins.ULong.Fingerprint():
			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%lu"), value)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%lu"), value)

			// create a new string object
			strObj := emt.CreateObject(blk, builtins.String)
			(*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["Load"], strObj, newStr)
			(*blk).NewCall(emt.CFuncs["free"], newStr)

			return strObj

		case builtins.Float.Fingerprint(), builtins.Double.Fingerprint():
			// convert float to double, idk why but it doesnt work without it
			double := (*blk).NewFPExt(value, types.Double)

			// find out how much space we need to allocate
			len := (*blk).NewCall(emt.CFuncs["snprintf"], constant.NewNull(types.I8Ptr), CI32(0), emt.GetStringConstant(blk, "%g"), double)

			// allocate space for the new string
			newStr := (*blk).NewCall(emt.CFuncs["malloc"], (*blk).NewAdd(len, CI32(1)))

			// convert the float
			(*blk).NewCall(emt.CFuncs["snprintf"], newStr, (*blk).NewAdd(len, CI32(1)), emt.GetStringConstant(blk, "%g"), double)

			// create a new string object
			strObj := emt.CreateObject(blk, builtins.String)
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
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Bool, value)

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
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Int, value)

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
		} else if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			// truncate the long to an int
			result := (*blk).NewTrunc(value, emt.IRTypes(builtins.Int))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return value
		} else if expr.Expression.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			// truncate the long to an int
			result := (*blk).NewTrunc(value, emt.IRTypes(builtins.Int))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			// extend the byte to an int
			result := (*blk).NewFPToSI(value, emt.IRTypes(builtins.Int))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Double.Fingerprint() {
			// extend the byte to an int
			result := (*blk).NewFPToSI(value, emt.IRTypes(builtins.Int))
			return result
		} else if expr.Expression.Type().Name == builtins.Pointer.Name {
			return (*blk).NewPtrToInt(value, emt.IRTypes(builtins.Int))
		}
	} else if expr.ToType.Fingerprint() == builtins.Byte.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			// truncate the int to a byte
			result := (*blk).NewTrunc(value, emt.IRTypes(builtins.Byte))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			// truncate the long to a byte
			result := (*blk).NewTrunc(value, emt.IRTypes(builtins.Byte))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Byte, value)

			// bitcast to boxed int
			boxedByte := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Int)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Byte)].Functions["GetValue"], boxedByte)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to byte conversion cleanup")
			}

			return primitive
		}
	} else if expr.ToType.Fingerprint() == builtins.Long.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			// extend the byte to a long
			result := (*blk).NewSExt(value, emt.IRTypes(builtins.Long))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			// extend the int to a long
			result := (*blk).NewSExt(value, emt.IRTypes(builtins.Long))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return value
		} else if expr.Expression.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			// extend the int to a long
			result := (*blk).NewZExt(value, emt.IRTypes(builtins.Long))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Double.Fingerprint() {
			// extend the byte to an int
			result := (*blk).NewFPToSI(value, emt.IRTypes(builtins.Long))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Long, value)

			// bitcast to boxed int
			boxedLong := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Long)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Long)].Functions["GetValue"], boxedLong)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to long conversion cleanup")
			}

			return primitive
		} else if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			result := (*blk).NewCall(emt.CFuncs["atol"], (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], value))

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "string to long conversion cleanup")
			}

			return result
		}
	} else if expr.ToType.Fingerprint() == builtins.UInt.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return value
		} else if expr.Expression.Type().Fingerprint() == builtins.ULong.Fingerprint() {
			return (*blk).NewTrunc(value, emt.IRTypes(builtins.UInt))
		} else if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return (*blk).NewTrunc(value, emt.IRTypes(builtins.UInt))
		}
	} else if expr.ToType.Fingerprint() == builtins.ULong.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			return value
		} else if expr.Expression.Type().Fingerprint() == builtins.UInt.Fingerprint() {
			return (*blk).NewZExt(value, emt.IRTypes(builtins.ULong))
		} else if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return (*blk).NewZExt(value, emt.IRTypes(builtins.ULong))
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
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Float, value)

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
		} else if expr.Expression.Type().Fingerprint() == builtins.Double.Fingerprint() {
			result := (*blk).NewFPTrunc(value, emt.IRTypes(builtins.Float))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Float))
			return result

		} else if expr.Expression.Type().Fingerprint() == builtins.Byte.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Float))
			return result

		} else if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Float))
			return result
		}
	} else if expr.ToType.Fingerprint() == builtins.Double.Fingerprint() {
		if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			result := (*blk).NewCall(emt.CFuncs["atof"], (*blk).NewCall(emt.Classes[emt.Id(builtins.String)].Functions["GetBuffer"], value))

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "string to float conversion cleanup")
			}

			return result

		} else if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Double, value)

			// bitcast to boxed float
			boxedDouble := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Double)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Double)].Functions["GetValue"], boxedDouble)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to float conversion cleanup")
			}

			return primitive
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			result := (*blk).NewFPExt(value, emt.IRTypes(builtins.Double))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Double))
			return result
		} else if expr.Expression.Type().Fingerprint() == builtins.Long.Fingerprint() {
			result := (*blk).NewSIToFP(value, emt.IRTypes(builtins.Double))
			return result
		}
	} else if expr.ToType.Name == builtins.Array.Name {
		if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// object arrays
			if expr.ToType.SubTypes[0].IsObject {
				// make sure this conversion is valid
				emt.EmitValidConversionCheck(blk, expr.ToType, value)

				// change the pointer type
				return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
			} else {
				// make sure this conversion is valid
				emt.EmitValidConversionCheck(blk, expr.ToType, value)

				// change the pointer type
				return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
			}
		}
	} else if expr.ToType.Name == builtins.Pointer.Name {
		if expr.Expression.Type().IsObject {
			// change the pointer type
			return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
		} else if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			// convert the int to a pointer
			return (*blk).NewIntToPtr(value, emt.IRTypes(expr.ToType))
		} else if expr.Expression.Type().Name == builtins.Pointer.Name {
			// beep boop we bit casting
			return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
		}
	} else if expr.ToType.Name == builtins.Action.Name {
		if expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
			// make sure this conversion is valid
			emt.EmitValidConversionCheck(blk, builtins.Long, value)

			// bitcast to boxed long
			boxedLong := (*blk).NewBitCast(value, types.NewPointer(emt.Classes[emt.Id(builtins.Long)].Type))

			// load its value
			primitive := (*blk).NewCall(emt.Classes[emt.Id(builtins.Long)].Functions["GetValue"], boxedLong)

			// if value isn't a variable (meaning its already memory managed)
			// decrement its reference counter
			if !expr.Expression.IsPersistent() {
				emt.DestroyReference(blk, value, "any to long conversion cleanup")
			}

			return (*blk).NewIntToPtr(primitive, emt.IRTypes(expr.ToType))
		}
	}

	if expr.ToType.IsObject {
		fmt.Println(expr.ToType.Fingerprint())
		if expr.Expression.Type().Name == builtins.Pointer.Name {
			// change the pointer type
			return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
		}
	}

	// classes
	if expr.ToType.IsObject && expr.Expression.Type().Fingerprint() == builtins.Any.Fingerprint() {
		// make sure this conversion is valid
		emt.EmitValidConversionCheck(blk, expr.ToType, value)

		return (*blk).NewBitCast(value, emt.IRTypes(expr.ToType))
	}

	fmt.Println("Unknown Conversion!")
	fmt.Println(expr.ToType.IsObject)
	fmt.Println(expr.Expression.Type().Name)
	fmt.Println(expr.Expression.Type().Fingerprint())
	fmt.Println(expr.ToType.Fingerprint())
	return nil
}

func (emt *Emitter) EmitReferenceExpression(blk **ir.Block, expr boundnodes.BoundReferenceExpressionNode) value.Value {
	value := emt.EmitVariablePtr(blk, expr.Expression.(boundnodes.BoundVariableExpressionNode).Variable)
	return value //(*blk).NewIntToPtr((*blk).NewPtrToInt(value, types.I32), emt.IRTypes(expr.Type()))
}

func (emt *Emitter) EmitDereferenceExpression(blk **ir.Block, expr boundnodes.BoundDereferenceExpressionNode) value.Value {
	value := emt.EmitExpression(blk, expr.Expression)
	return (*blk).NewLoad(
		emt.IRTypes(expr.Type()), value)
}

func (emt *Emitter) EmitLambdaExpression(blk **ir.Block, expr boundnodes.BoundLambdaExpressionNode) value.Value {
	// take a snapshot of the function we're in
	// TODO(RedCube): Replace this with a stack
	fnc := emt.Function
	fncsym := emt.FunctionSym
	fnclcs := emt.Locals
	fnctmp := emt.Temps
	fnclbs := emt.Labels

	// lambdas are a lie
	function := emt.EmitFunction(expr.Function, expr.Body)

	// emit the body
	emt.FunctionSym = expr.Function
	emt.EmitBlockStatement(expr.Function, function, expr.Body)

	// remember where tf we were
	emt.Function = fnc
	emt.FunctionSym = fncsym
	emt.Locals = fnclcs
	emt.Temps = fnctmp
	emt.Labels = fnclbs

	// don
	return function
}

func (emt *Emitter) EmitFunctionExpression(blk **ir.Block, expr boundnodes.BoundFunctionExpressionNode) value.Value {

	if emt.IsInClass && !expr.Function.BuiltIn {
		return emt.Classes[emt.Id(emt.ClassSym.Type)].Functions[emt.Id(expr.Function)] // return the lame IR function
	} else {
		return emt.Functions[emt.Id(expr.Function)].IRFunction // return the IR function
	}
}

func (emt *Emitter) EmitThisExpression(blk **ir.Block, expr boundnodes.BoundThisExpressionNode) value.Value {
	return emt.Function.Params[0] // return the reserved parameter
}

// </EXPRESSIONS>--------------------------------------------------------------
// <UTILS>---------------------------------------------------------------------

func (emt *Emitter) EmitValidConversionCheck(blk **ir.Block, typ symbols.TypeSymbol, val value.Value) {
	bas := typ
	bas.SubTypes = make([]symbols.TypeSymbol, 0) // remove subtypes
	(*blk).NewCall(emt.ExcFuncs["ThrowIfInvalidCast"], (*blk).NewBitCast(val, emt.IRTypes(builtins.Any)), (*blk).NewBitCast(emt.Classes[emt.Id(bas)].vConstant, types.NewPointer(emt.Classes[emt.Id(builtins.Any)].vTable)), emt.GetConstantStringConstant(typ.Fingerprint()))
}

func (emt *Emitter) DefaultConstant(blk **ir.Block, typ symbols.TypeSymbol) constant.Constant {
	switch typ.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(false)
	case builtins.Byte.Fingerprint():
		return constant.NewInt(types.I8, 0)
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, 0)
	case builtins.Long.Fingerprint():
		return constant.NewInt(types.I64, 0)
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, 0)
	case builtins.String.Fingerprint():
		return constant.NewNull(emt.IRTypes(builtins.String).(*types.PointerType))
	case builtins.Any.Fingerprint():
		return constant.NewNull(emt.IRTypes(builtins.Any).(*types.PointerType))
	}

	if typ.Name == builtins.Array.Name {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	if typ.Name == builtins.Pointer.Name {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	if typ.Name == builtins.Action.Name {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	if typ.IsUserDefined && typ.IsObject {
		return constant.NewNull(emt.IRTypes(typ).(*types.PointerType))
	}

	if typ.IsUserDefined && !typ.IsObject {
		fields := make([]constant.Constant, 0)

		for _, i := range emt.Structs[emt.Id(typ)].Symbol.Fields {
			fields = append(fields, emt.DefaultConstant(blk, i.VarType()))
		}

		return constant.NewStruct(emt.IRTypes(typ).(*types.StructType), fields...)
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

	// optional prefix
	prefix := ""
	if CompileAsPackage {
		prefix = "." + PackageName
	}

	// create a global to store our literal
	global := emt.Module.NewGlobalDef(fmt.Sprintf(prefix+".str.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
	global.Immutable = true
	emt.StrNameCounter++

	pointer := (*blk).NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), global, CI32(0), CI32(0))
	emt.StrConstants[literal] = global

	return pointer
}

func (emt *Emitter) GetConstantStringConstant(literal string) constant.Constant {
	// add a null byte at the end
	str := literal + "\x00"

	// optional prefix
	prefix := ""
	if CompileAsPackage {
		prefix = "." + PackageName
	}

	// create a global to store our literal
	global := emt.Module.NewGlobalDef(fmt.Sprintf(prefix+".str.c.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
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

func (emt *Emitter) CreateObject(blk **ir.Block, src symbols.TypeSymbol, args ...value.Value) value.Value {
	// make a copy of our source type to not lose any info
	typ := src

	// simplify types with subtypes
	if len(typ.SubTypes) > 0 {
		typ.SubTypes = make([]symbols.TypeSymbol, 0) // empty subtype list
	}

	// class name
	typeName := emt.Id(typ)

	// sizeof the struct
	size := (*blk).NewGetElementPtr(emt.Classes[typeName].Type, constant.NewNull(types.NewPointer(emt.Classes[typeName].Type)), CI32(1))
	sizeInt := (*blk).NewPtrToInt(size, types.I32)

	// create space for the instance
	instance := (*blk).NewBitCast((*blk).NewCall(emt.CFuncs["malloc"], sizeInt), types.NewPointer(emt.Classes[typeName].Type))

	// initialize reference count
	arcCounterPointer := (*blk).NewGetElementPtr(emt.Classes[typeName].Type, instance, CI32(0), CI32(1))
	(*blk).NewStore(CI32(0), arcCounterPointer)

	// load the vTable
	arcVTablePointer := (*blk).NewGetElementPtr(emt.Classes[typeName].Type, instance, CI32(0), CI32(0))
	vTable := emt.GetVtableConstant(src, typeName)
	(*blk).NewStore(vTable, arcVTablePointer)

	// create reference
	emt.CreateReference(blk, instance, "initial instance")

	// constructor arguments
	arguments := []value.Value{instance}
	arguments = append(arguments, args...)

	// call the constructor
	(*blk).NewCall(emt.Classes[typeName].Constructor, arguments...)

	// mmm don
	return instance
}

func (emt *Emitter) GetVtableConstant(src symbols.TypeSymbol, typ string) value.Value {
	fields := make([]constant.Constant, 0)
	template := emt.Classes[typ].vConstant.Init.(*constant.Struct)

	// parent vTable
	fields = append(fields, template.Fields[0])

	// class name
	fields = append(fields, emt.GetConstantStringConstant(strings.ToUpper(src.Name[:1])+src.Name[1:]))

	// die() pointer
	fields = append(fields, template.Fields[2])

	// instance fingerprint
	fields = append(fields, emt.GetConstantStringConstant(src.Fingerprint()))

	return constant.NewStruct(
		template.Typ,
		fields...,
	)
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

	// (or, an action)
	// actions need some special treatment here as they have subtypes and that breaks everything
	if typ.Name == builtins.Action.Name {
		typ = builtins.Action
	}

	// create a new object and give it the primitive to be "capsuled"
	obj := emt.CreateObject(blk, typ, val)

	return obj
}

// ARC FUNCTIONS
func (emt *Emitter) CreateReference(blk **ir.Block, expr value.Value, comment string) {
	if VerboseARC {
		emt.CreateReferenceVerbose(blk, expr, emt.GetStringConstant(blk, comment))
	} else {
		emt.CreateReferenceNormal(blk, expr)
	}
}

func (emt *Emitter) DestroyReference(blk **ir.Block, expr value.Value, comment string) {
	if VerboseARC {
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

	// if this is a struct we'll have to load it first
	if val.Type().IsUserDefined && !val.Type().IsObject {
		expression = (*blk).NewLoad(emt.IRTypes(val.Type()), expression)
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

	// if this is a struct we'll have to load it first
	if literalValue.Type().IsUserDefined && !literalValue.Type().IsObject {
		value = (*blk).NewLoad(emt.IRTypes(literalValue.Type()), value)
	}

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
