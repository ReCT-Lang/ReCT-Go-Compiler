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
	FunctionLocals map[string]map[string]Local
	StrConstants   map[string]value.Value
	StrNameCounter int

	// local things for this current function
	Function    *ir.Func
	FunctionSym symbols.FunctionSymbol
	Locals      map[string]Local
	Labels      map[string]*ir.Block
}

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:         program,
		Module:          ir.NewModule(),
		UseFingerprints: useFingerprints,
		Globals:         make(map[string]Global),
		Functions:       make(map[string]Function),
		CFunctions:      make(map[string]*ir.Func),
		FunctionLocals:  make(map[string]map[string]Local),
		StrConstants:    make(map[string]value.Value),
	}

	emitter.EmitBuiltInFunctions()

	// declare all function names
	for _, fnc := range emitter.Program.Functions {
		if !fnc.Symbol.BuiltIn {
			function := Function{IRFunction: emitter.EmitFunction(fnc.Symbol, fnc.Body), BoundFunction: fnc}
			emitter.Functions[function.IRFunction.Name()] = function
		}
	}

	// emit main function first
	mainName := tern(emitter.UseFingerprints, program.MainFunction.Fingerprint(), program.MainFunction.Name)
	emitter.FunctionSym = emitter.Functions[mainName].BoundFunction.Symbol
	emitter.EmitBlockStatement(emitter.Functions[mainName].IRFunction, emitter.Functions[mainName].BoundFunction.Body)

	// emit function bodies
	for _, fnc := range emitter.Functions {
		if !fnc.BoundFunction.Symbol.BuiltIn && fnc.BoundFunction.Symbol.Fingerprint() != program.MainFunction.Fingerprint() {
			emitter.FunctionSym = fnc.BoundFunction.Symbol
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

			varName := tern(emt.UseFingerprints, declStatement.Variable.Fingerprint(), declStatement.Variable.SymbolName())

			// create local variable
			local := root.NewAlloca(IRTypes[declStatement.Variable.VarType().Fingerprint()])
			local.SetName(varName)

			// save it for referencing later
			locals[varName] = Local{IRLocal: local, IRBlock: root, Type: declStatement.Variable.VarType()}
		}
	}

	// store this for later
	emt.FunctionLocals[functionName] = locals

	return function
}

func (emt *Emitter) EmitBlockStatement(fnc *ir.Func, body boundnodes.BoundBlockStatementNode) {
	// set up our environment
	emt.Function = fnc
	emt.Locals = emt.FunctionLocals[fnc.Name()]
	emt.Labels = make(map[string]*ir.Block)

	// load the root block
	// each label statement will create a new block and store it in this variable
	currentBlock := fnc.Blocks[0]

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
		global := emt.Module.NewGlobalDef(varName, emt.DefaultConstant(blk, stmt.Variable.VarType()))

		// save it for referencing later
		emt.Globals[varName] = Global{IRGlobal: global, Type: stmt.Variable.VarType()}

		// emit its assignment
		blk.NewStore(expression, global)
	} else {
		local := emt.Locals[varName]
		local.IsSet = true
		emt.Locals[varName] = local

		// emit its assignemnt
		blk.NewStore(expression, local.IRLocal)
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
	expr := emt.EmitExpression(blk, stmt.Expression)

	// if the expressions value is a string is requires cleanup
	if stmt.Expression.Type().Fingerprint() == builtins.String.Fingerprint() {
		blk.Insts = append(blk.Insts, NewComment("expression value unused -> cleanup required"))
		blk.NewCall(emt.CFunctions["free"], expr)
	}
}

func (emt *Emitter) EmitReturnStatement(blk *ir.Block, stmt boundnodes.BoundReturnStatementNode) {
	// --< state of the art garbage collecting >---------------------------
	// 1. go through all locals created up to this point
	// 2. check if they need freeing
	// 3. if they do, free them
	blk.Insts = append(blk.Insts, NewComment("<ReturnGC>"))
	for name, local := range emt.Locals {
		if !local.IsSet {
			continue
		}

		if local.Type.Fingerprint() == builtins.String.Fingerprint() {
			blk.Insts = append(blk.Insts, NewComment(" -> utterly obliterating '%"+name+"'"))
			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(IRTypes[local.Type.Fingerprint()], local.IRLocal))
		}
	}

	// free any parameters as well
	for _, param := range emt.FunctionSym.Parameters {
		if param.Type.Fingerprint() == builtins.String.Fingerprint() {
			blk.Insts = append(blk.Insts, NewComment(" -> utterly obliterating '%"+param.Name+"'"))
			blk.NewCall(emt.CFunctions["free"], emt.Function.Params[param.Ordinal])
		}
	}
	blk.Insts = append(blk.Insts, NewComment("</ReturnGC>"))

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
	blk.Insts = append(blk.Insts, NewComment("<GC>"))
	for _, variable := range stmt.Variables {
		// check if the variables type needs to be freed

		if variable.VarType().Fingerprint() == builtins.String.Fingerprint() {
			varName := tern(emt.UseFingerprints, variable.Fingerprint(), variable.SymbolName())
			blk.Insts = append(blk.Insts, NewComment(" -> destroy variable '%"+varName+"'"))

			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(types.I8Ptr, emt.Locals[varName].IRLocal))
			blk.NewStore(constant.NewNull(types.I8Ptr), emt.Locals[varName].IRLocal)
		}
	}
	blk.Insts = append(blk.Insts, NewComment("</GC>"))
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
		return emt.GetStringConstant(blk, expr.Value.(string))
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
		// if this variable already contained a string -> clean that one up
		if expr.Variable.VarType().Fingerprint() == builtins.String.Fingerprint() {
			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(IRTypes[expr.Variable.VarType().Fingerprint()], emt.Globals[varName].IRGlobal))
		}

		// assign the value to the global variable
		blk.NewStore(expression, emt.Globals[varName].IRGlobal)

	} else {
		// if this variable already contained a string -> clean that one up
		if expr.Variable.VarType().Fingerprint() == builtins.String.Fingerprint() {
			blk.NewCall(emt.CFunctions["free"], blk.NewLoad(IRTypes[expr.Variable.VarType().Fingerprint()], emt.Locals[varName].IRLocal))
		}

		// assign the value to the local variable
		blk.NewStore(expression, emt.Locals[varName].IRLocal)
	}

	// also return the value as this can also be used as an expression
	// if we're working with strings the value returned here needs to be a copy
	if expr.Variable.VarType().Fingerprint() == builtins.String.Fingerprint() {
		return emt.CopyStringNoFree(blk, expression)
	}

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
			// figure out how long our left and right are
			leftLen := blk.NewCall(emt.CFunctions["strlen"], left)
			rightLen := blk.NewCall(emt.CFunctions["strlen"], right)

			// allocate a new buffer for the concatination to go into
			newStr := blk.NewCall(emt.CFunctions["malloc"], blk.NewAdd(blk.NewAdd(leftLen, rightLen), CI32(1)))

			// copy over the left string
			blk.NewCall(emt.CFunctions["strcpy"], newStr, left)

			// concat the other side into it
			blk.NewCall(emt.CFunctions["strcat"], newStr, right)

			// if left and right arent variables (meaning they are already memory managed)
			// free() them
			if expr.Left.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Left.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], left)
			}

			if expr.Right.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Right.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], right)
			}

			return newStr
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
			// compare left and right using strcmp
			result := blk.NewCall(emt.CFunctions["strcmp"], left, right)

			// if left and right arent variables (meaning they are already memory managed)
			// free() them
			if expr.Left.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Left.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], left)
			}

			if expr.Right.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Right.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], right)
			}

			// to check if they are equal, check if the result is 0
			return blk.NewICmp(enum.IPredEQ, result, CI32(0))
		}

	case boundnodes.NotEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return blk.NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return blk.NewFCmp(enum.FPredONE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return blk.NewICmp(enum.IPredNE, left, right)

		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			// compare left and right using strcmpint
			result := blk.NewCall(emt.CFunctions["strcmp"], left, right)

			// if left and right arent variables (meaning they are already memory managed)
			// free() them
			if expr.Left.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Left.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], left)
			}

			if expr.Right.NodeType() != boundnodes.BoundVariableExpression &&
				expr.Right.NodeType() != boundnodes.BoundLiteralExpression {
				blk.NewCall(emt.CFunctions["free"], right)
			}

			// to check if they are unequal, check if the result is not 0
			return blk.NewICmp(enum.IPredNE, result, CI32(0))
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
		expression := emt.EmitExpression(blk, arg)

		// if this is a string, make a copy before handing it over
		if arg.Type().Fingerprint() == builtins.String.Fingerprint() {
			expression = emt.CopyString(blk, expression, arg)
		}

		arguments = append(arguments, expression)
	}

	functionName := tern(emt.UseFingerprints, expr.Function.Fingerprint(), expr.Function.Name)

	call := blk.NewCall(emt.Functions[functionName].IRFunction, arguments...)

	// if this is an external function it doesnt implement the garbage collector
	// meaning we have to clean up its arguments outselves
	if expr.Function.BuiltIn {
		for i, arg := range arguments {
			if expr.Arguments[i].Type().Fingerprint() == builtins.String.Fingerprint() {
				blk.NewCall(emt.CFunctions["free"], arg)
			}
		}
	}

	return call
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

func (emt *Emitter) CopyStringNoFree(blk *ir.Block, expression value.Value) value.Value {
	// copy over the string
	newStr := blk.NewCall(emt.CFunctions["uStringCopy"], expression)
	return newStr
}

func (emt *Emitter) DefaultConstant(blk *ir.Block, typ symbols.TypeSymbol) constant.Constant {
	switch typ.Fingerprint() {
	case builtins.Bool.Fingerprint():
		return constant.NewBool(false)
	case builtins.Int.Fingerprint():
		return constant.NewInt(types.I32, 0)
	case builtins.Float.Fingerprint():
		return constant.NewFloat(types.Float, 0)
	case builtins.String.Fingerprint():
		return constant.NewNull(types.I8Ptr)
	}

	return nil
}

func (emt *Emitter) GetStringConstant(blk *ir.Block, literal string) value.Value {
	// check if this literal has already been created
	val, ok := emt.StrConstants[literal]
	if ok {
		return val
	}

	// add a null byte at the end
	str := literal + "\x00"

	// create a global to store our literal
	global := emt.Module.NewGlobalDef(fmt.Sprintf(".str.%d", emt.StrNameCounter), constant.NewCharArrayFromString(str))
	global.Immutable = true
	emt.StrNameCounter++

	pointer := blk.NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), global, CI32(0), CI32(0))
	emt.StrConstants[literal] = pointer

	return pointer
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
