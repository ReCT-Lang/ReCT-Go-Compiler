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

	// referenced C functions
	CFunctions map[string]*ir.Func

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
		CFunctions:      make(map[string]*ir.Func),
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

	case boundnodes.BoundGarbageCollectionStatement:
		emt.EmitGarbageCollectionStatement(blk, stmt.(boundnodes.BoundGarbageCollectionStatementNode))
	}
}

// </FUNCTIONS>----------------------------------------------------------------
// <STATEMENTS>----------------------------------------------------------------

func (emt *Emitter) EmitVariableDeclarationStatement(blk *ir.Block, stmt boundnodes.BoundVariableDeclarationStatementNode) {
	varName := tern(emt.UseFingerprints, stmt.Variable.Fingerprint(), stmt.Variable.SymbolName())
	expression := emt.EmitExpression(blk, stmt.Initializer)

	// if the value is a string, copy it (dont just pass the pointer)
	if stmt.Initializer.Type().Fingerprint() == builtins.String.Fingerprint() {
		expression = emt.CopyString(blk, expression, stmt.Initializer)
	}

	if stmt.Variable.IsGlobal() {
		// create a new global
		global := emt.Module.NewGlobal(varName, IRTypes[stmt.Variable.VarType().Fingerprint()])

		// save it for referencing later
		emt.Globals[varName] = Global{IRGlobal: global, Type: stmt.Variable.VarType()}

		// emit its assignment
		blk.NewStore(expression, global)
	} else {
		// create local variable
		local := blk.NewAlloca(IRTypes[stmt.Variable.VarType().Fingerprint()])
		local.SetName(varName)

		// save it for referencing later
		emt.Locals[varName] = Local{IRLocal: local, IRBlock: blk, Type: stmt.Variable.VarType()}

		// emit its assignemnt
		blk.NewStore(expression, local)
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
	// --< state of the art garbage collecting >---------------------------
	// 1. go through all locals created up to this point
	// 2. check if they need freeing
	// 3. if they do, free them
	for _, local := range emt.Locals {
		if local.Type.Fingerprint() == builtins.String.Fingerprint() {
			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(IRTypes[local.Type.Fingerprint()], local.IRLocal))
		}
	}

	if stmt.Expression != nil {
		expression := emt.EmitExpression(blk, stmt.Expression)

		// if the expression is a string, copy it to unlink it from any variable
		if stmt.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
			expression = emt.CopyString(blk, expression, stmt.Expression)
		}

		blk.NewRet(expression)
	} else {
		blk.NewRet(nil)
	}
}

func (emt *Emitter) EmitGarbageCollectionStatement(blk *ir.Block, stmt boundnodes.BoundGarbageCollectionStatementNode) {
	for _, variable := range stmt.Variables {
		// check if the variables type needs to be freed

		if variable.VarType().Fingerprint() == builtins.String.Fingerprint() {
			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(types.I8Ptr, emt.Locals[variable.Fingerprint()].IRLocal))
		}
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

	case boundnodes.BoundBinaryExpression:
		return emt.EmitBinaryExpression(blk, expr.(boundnodes.BoundBinaryExpressionNode))

	case boundnodes.BoundCallExpression:
		return emt.EmitCallExpression(blk, expr.(boundnodes.BoundCallExpressionNode))

	case boundnodes.BoundTypeCallExpression:
		return emt.EmitTypeCallExpression(blk, expr.(boundnodes.BoundTypeCallExpressionNode))

	case boundnodes.BoundConversionExpression:
		return emt.EmitConversionExpression(blk, expr.(boundnodes.BoundConversionExpressionNode))
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
		// add a null byte at the end
		str := expr.Value.(string) + "\x00"

		// create a global to store our literal
		global := emt.Module.NewGlobalDef(fmt.Sprintf(".str.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
		global.Immutable = true
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

	// if the value is a string, copy it (dont just pass the pointer)
	if expr.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
		expression = emt.CopyString(blk, expression, expr.Expression)
	}

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

func (emt *Emitter) EmitBinaryExpression(blk *ir.Block, expr boundnodes.BoundBinaryExpressionNode) value.Value {
	left := emt.EmitExpression(blk, expr.Left)
	right := emt.EmitExpression(blk, expr.Right)

	switch expr.Op.OperatorKind {
	case boundnodes.Addition:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFAdd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// TODO: string concat
		}

	case boundnodes.Subtraction:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewSub(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFSub(left, right)
		}

	case boundnodes.Multiplication:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewMul(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFMul(left, right)
		}

	case boundnodes.Division:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewSDiv(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFDiv(left, right)
		}

	case boundnodes.Modulus:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewSRem(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFRem(left, right)
		}

	case boundnodes.BitwiseAnd:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewAnd(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewAnd(left, right)
		}

	case boundnodes.BitwiseOr:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewOr(left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewOr(left, right)
		}

	case boundnodes.BitwiseXor:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewXor(left, right)
		}

	case boundnodes.Equals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredOEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewICmp(enum.IPredEQ, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// TODO: string equality
		}

	case boundnodes.NotEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredONE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// TODO: string inequality
		}

	case boundnodes.Greater:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredSGT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredOGT, left, right)
		}

	case boundnodes.GreaterOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredSGE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredOGE, left, right)
		}

	case boundnodes.Less:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredSLT, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredOLT, left, right)
		}

	case boundnodes.LessOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredSLE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredOLE, left, right)
		}

	case boundnodes.LogicalAnd:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewAnd(left, right)
		}

	case boundnodes.LogicalOr:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewOr(left, right)
		}

	}

	print.PrintCF(print.Red, "UNIMPLEMENTED BINARY EXPRESSION: %s : %s => %s", expr.Left, expr.Right, expr.Op.OperatorKind)
	os.Exit(-1)

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

func (emt *Emitter) EmitConversionExpression(blk *ir.Block, expr boundnodes.BoundConversionExpressionNode) value.Value {

	print.PrintC(print.Red, "conversions arent implemented yet!")
	os.Exit(-1)

	return nil
}

// </EXPRESSIONS>--------------------------------------------------------------
// <UTILS>---------------------------------------------------------------------

func (emt *Emitter) CopyString(blk *ir.Block, expression value.Value, source boundnodes.BoundExpressionNode) value.Value {
	// copy over the string
	newStr := blk.NewCall(emt.CFunctions["uStringCopy"], expression)

	// if this isnt another variable, free the old buffer
	if source.NodeType() != boundnodes.BoundVariableExpression &&
		source.NodeType() != boundnodes.BoundLiteralExpression {
		blk.NewCall(emt.CFunctions["free"], expression)
	}

	return newStr
}

// </UTILS>--------------------------------------------------------------------

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
