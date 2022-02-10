package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"os"

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

	// global variables
	Globals        map[string]Global
	Functions      map[string]Function
	StrNameCounter int

	// local things for this current function
	Function *ir.Func
	Locals   map[string]Local
	Labels   map[string]*ir.Block
}

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:         program,
		Module:          ir.NewModule(),
		UseFingerprints: useFingerprints,
		Globals:         make(map[string]Global),
		Functions:       make(map[string]Function),
	}

	emitter.EmitBuiltInFunctions()

	// declare all function names
	for _, fnc := range emitter.Program.Functions {
		if !fnc.Symbol.BuiltIn {
			function := Function{IRFunction: emitter.EmitFunction(fnc.Symbol, fnc.Body), BoundFunction: fnc}
			emitter.Functions[function.IRFunction.Name()] = function
		}
	}

	// emit function bodies
	for _, fnc := range emitter.Functions {
		if !fnc.BoundFunction.Symbol.BuiltIn {
			emitter.EmitBlockStatement(fnc.IRFunction, fnc.BoundFunction.Body)
		}
	}

	return emitter.Module
}

// <FUNCTIONS>-----------------------------------------------------------------
func (emt *Emitter) EmitFunction(sym symbols.FunctionSymbol, body boundnodes.BoundBlockStatementNode) *ir.Func {
	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := tern(emt.UseFingerprints, param.Fingerprint(), param.Name)

		// create it
		params = append(params, ir.NewParam(paramName, IRTypes[param.Type.Fingerprint()]))
	}

	// figure out the return type
	returnType := IRTypes[sym.Type.Fingerprint()]

	// the function name
	functionName := tern(emt.UseFingerprints, sym.Fingerprint(), sym.Name)

	// create an IR function definition
	function := emt.Module.NewFunc(functionName, returnType, params...)

	return function
}

func (emt *Emitter) EmitBlockStatement(fnc *ir.Func, body boundnodes.BoundBlockStatementNode) {
	// set up our environment
	emt.Function = fnc
	emt.Locals = make(map[string]Local)
	emt.Labels = make(map[string]*ir.Block)

	// create a root block
	// each label statement will create a new block and store it in this variable
	currentBlock := fnc.NewBlock("")

	// go through the body and register all label blocks
	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// create a new block when we encounter a label
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			emt.Labels[string(labelStatement.Label)] = fnc.NewBlock(string(labelStatement.Label))
		}
	}

	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// if we encounter a laber statement -> switch to the label's block
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			currentBlock = emt.Labels[string(labelStatement.Label)]
		} else {
			emt.EmitStatement(currentBlock, stmt)
		}
	}
}

func (emt *Emitter) EmitStatement(blk *ir.Block, stmt boundnodes.BoundStatementNode) {
	// emit a statement to the current block
	switch stmt.NodeType() {
	case boundnodes.BoundVariableDeclaration:
		emt.EmitVariableDeclarationStatement(blk, stmt.(boundnodes.BoundVariableDeclarationStatementNode))

	case boundnodes.BoundGotoStatement:
		emt.EmitGotoStatement(blk, stmt.(boundnodes.BoundGotoStatementNode))

	case boundnodes.BoundConditionalGotoStatement:
		emt.EmitConditionalGotoStatement(blk, stmt.(boundnodes.BoundConditionalGotoStatementNode))

	case boundnodes.BoundExpressionStatement:
		emt.EmitBoundExpressionStatement(blk, stmt.(boundnodes.BoundExpressionStatementNode))

	case boundnodes.BoundReturnStatement:
		emt.EmitReturnStatement(blk, stmt.(boundnodes.BoundReturnStatementNode))
	}
}

// </FUNCTIONS>----------------------------------------------------------------
// <STATEMENTS>----------------------------------------------------------------

func (emt *Emitter) EmitVariableDeclarationStatement(blk *ir.Block, stmt boundnodes.BoundVariableDeclarationStatementNode) {
	varName := tern(emt.UseFingerprints, stmt.Variable.Fingerprint(), stmt.Variable.SymbolName())

	if stmt.Variable.IsGlobal() {
		// create a new global
		global := emt.Module.NewGlobal(varName, IRTypes[stmt.Variable.VarType().Fingerprint()])

		// save it for referencing later
		emt.Globals[varName] = Global{IRGlobal: global, Type: stmt.Variable.VarType()}

		// emit its assignment
		blk.NewStore(emt.EmitExpression(blk, stmt.Initializer), global)
	} else {
		// create local variable
		local := ir.NewAlloca(IRTypes[stmt.Variable.VarType().Fingerprint()])
		local.SetName(varName)

		// save it for referencing later
		emt.Locals[varName] = Local{IRLocal: local, IRBlock: blk, Type: stmt.Variable.VarType()}

		// emit its assignemnt
		blk.NewStore(emt.EmitExpression(blk, stmt.Initializer), local)
	}
}

func (emt *Emitter) EmitGotoStatement(blk *ir.Block, stmt boundnodes.BoundGotoStatementNode) {
	blk.NewBr(emt.Labels[string(stmt.Label)])
}

func (emt *Emitter) EmitConditionalGotoStatement(blk *ir.Block, stmt boundnodes.BoundConditionalGotoStatementNode) {
	// figure out where to jump
	ifLabel := emt.Labels[string(stmt.IfLabel)]
	elseLabel := emt.Labels[string(stmt.ElseLabel)]

	blk.NewCondBr(emt.EmitExpression(blk, stmt.Condition), ifLabel, elseLabel)
}

func (emt *Emitter) EmitBoundExpressionStatement(blk *ir.Block, stmt boundnodes.BoundExpressionStatementNode) {
	emt.EmitExpression(blk, stmt.Expression)
}

func (emt *Emitter) EmitReturnStatement(blk *ir.Block, stmt boundnodes.BoundReturnStatementNode) {
	if stmt.Expression != nil {
		blk.NewRet(emt.EmitExpression(blk, stmt.Expression))
	} else {
		blk.NewRet(nil)
	}
}

// </STATEMENTS>---------------------------------------------------------------
// <EXPRESSIONS>---------------------------------------------------------------

func (emt *Emitter) EmitExpression(blk *ir.Block, expr boundnodes.BoundExpressionNode) value.Value {
	switch expr.NodeType() {
	case boundnodes.BoundLiteralExpression:
		return emt.EmitLiteralExpression(blk, expr.(boundnodes.BoundLiteralExpressionNode))

	case boundnodes.BoundVariableExpression:
		return emt.EmitVariableExpression(blk, expr.(boundnodes.BoundVariableExpressionNode))

	case boundnodes.BoundAssignmentExpression:
		return emt.EmitAssignmentExpression(blk, expr.(boundnodes.BoundAssignmentExpressionNode))

	case boundnodes.BoundUnaryExpression:
		return emt.EmitUnaryExpression(blk, expr.(boundnodes.BoundUnaryExpressionNode))

	case boundnodes.BoundCallExpression:
		return emt.EmitCallExpression(blk, expr.(boundnodes.BoundCallExpressionNode))

	case boundnodes.BoundTypeCallExpression:
		return emt.EmitTypeCallExpression(blk, expr.(boundnodes.BoundTypeCallExpressionNode))

	}

	fmt.Println("Unimplemented node: " + expr.NodeType())
	return nil
}

func (emt *Emitter) EmitLiteralExpression(blk *ir.Block, expr boundnodes.BoundLiteralExpressionNode) value.Value {
	switch expr.LiteralType.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(expr.Value.(bool))
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, int64(expr.Value.(int)))
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, float64(expr.Value.(float32)))
	case builtins.String.Fingerprint():
		str := expr.Value.(string)
		global := emt.Module.NewGlobalDef(fmt.Sprintf(".str.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
		global.IsConstant()
		emt.StrNameCounter++
		return blk.NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), global, CI32(0), CI32(0))
	}

	return nil
}

func (emt *Emitter) EmitVariableExpression(blk *ir.Block, expr boundnodes.BoundVariableExpressionNode) value.Value {
	varName := tern(emt.UseFingerprints, expr.Variable.Fingerprint(), expr.Variable.SymbolName())

	// parameters
	if expr.Variable.SymbolType() == symbols.Parameter {
		paramSymbol := expr.Variable.(symbols.ParameterSymbol)
		return emt.Function.Params[paramSymbol.Ordinal]
	}

	if expr.Variable.IsGlobal() {
		return blk.NewLoad(IRTypes[emt.Globals[varName].Type.Fingerprint()], emt.Globals[varName].IRGlobal)
	} else {
		return blk.NewLoad(IRTypes[emt.Locals[varName].Type.Fingerprint()], emt.Locals[varName].IRLocal)
	}
}

func (emt *Emitter) EmitAssignmentExpression(blk *ir.Block, expr boundnodes.BoundAssignmentExpressionNode) value.Value {
	varName := tern(emt.UseFingerprints, expr.Variable.Fingerprint(), expr.Variable.SymbolName())
	expression := emt.EmitExpression(blk, expr.Expression)

	if expr.Variable.IsGlobal() {
		// assign the value to the global variable
		blk.NewStore(expression, emt.Globals[varName].IRGlobal)
	} else {
		// assign the value to the local variable
		blk.NewStore(expression, emt.Locals[varName].IRLocal)
	}

	// also return the value as this can also be used as an expression
	return expression
}

func (emt *Emitter) EmitUnaryExpression(blk *ir.Block, expr boundnodes.BoundUnaryExpressionNode) value.Value {
	expression := emt.EmitExpression(blk, expr.Expression)

	switch expr.Op.OperatorKind {
	case boundnodes.Identity:
		return expression
	case boundnodes.Negation:
		// int negation   -> 0 - value
		// float negation -> fneg value
		if expr.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewSub(CI32(0), expression)

		} else if expr.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFNeg(expression)
		}

	case boundnodes.LogicalNegation:
		cmp := blk.NewICmp(enum.IPredNE, expression, CI32(0))
		xor := blk.NewXor(cmp, CB1(true))
		return xor
	}

	return nil
}

func (emt *Emitter) EmitCallExpression(blk *ir.Block, expr boundnodes.BoundCallExpressionNode) value.Value {
	arguments := make([]value.Value, 0)

	for _, arg := range expr.Arguments {
		arguments = append(arguments, emt.EmitExpression(blk, arg))
	}

	functionName := tern(emt.UseFingerprints, expr.Function.Fingerprint(), expr.Function.Name)

	return blk.NewCall(emt.Functions[functionName].IRFunction, arguments...)
}

func (emt *Emitter) EmitTypeCallExpression(blk *ir.Block, expr boundnodes.BoundTypeCallExpressionNode) value.Value {

	print.PrintC(print.Red, "type calls arent implemented yet!")
	os.Exit(-1)

	return nil
}

// </EXPRESSIONS>--------------------------------------------------------------
// <BUILTINS>------------------------------------------------------------------

func (emt *Emitter) EmitBuiltInFunctions() {
	// printf link
	printf := emt.Module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true

	// PrintfI32 -------------
	PrintfI32Name := tern(emt.UseFingerprints, builtins.PrintfI32.Fingerprint(), builtins.PrintfI32.Name)
	PrintfI32 := emt.Module.NewFunc(PrintfI32Name, types.Void, ir.NewParam("message", types.I8Ptr), ir.NewParam("slotin", types.I32))

	emt.Functions[PrintfI32Name] = Function{IRFunction: PrintfI32, BoundFunction: binder.BoundFunction{Symbol: builtins.PrintfI32}}
	printfI32Body := PrintfI32.NewBlock("")

	printfI32Body.NewCall(printf, PrintfI32.Params[0], PrintfI32.Params[1])
	printfI32Body.NewRet(nil)
}

// </BUILTINS>-----------------------------------------------------------------

// yes
func tern(cond bool, str1 string, str2 string) string {
	if cond {
		return str1
	} else {
		return str2
	}
}

func btern(cond bool, v1 *ir.Block, v2 *ir.Block) *ir.Block {
	if cond {
		return v1
	} else {
		return v2
	}
}

func itern(cond bool, v1 interface{}, v2 interface{}) interface{} {
	if cond {
		return v1
	} else {
		return v2
	}
}

func CI32(val int32) value.Value {
	return constant.NewInt(types.I32, int64(val))
}

func CB1(val bool) value.Value {
	return constant.NewBool(val)
}
