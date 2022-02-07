package evaluator

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"os"
)

type Evaluator struct {
	Program   binder.BoundProgram
	Globals   map[string]interface{}
	Functions map[string]binder.BoundFunction
	Locals    []map[string]interface{}
}

// locals stack helpers
func (evl *Evaluator) PushLocals() {
	evl.Locals = append(evl.Locals, make(map[string]interface{}))
}

func (evl *Evaluator) PopLocals() {
	evl.Locals = evl.Locals[:len(evl.Locals)-1]
}

func (evl *Evaluator) GetLocals() map[string]interface{} {
	return evl.Locals[len(evl.Locals)-1]
}

func (evl *Evaluator) SetLocal(fingerprint string, value interface{}) {
	evl.Locals[len(evl.Locals)-1][fingerprint] = value
}

// variable helper
func (evl *Evaluator) Assign(sym symbols.VariableSymbol, value interface{}) {
	if sym.IsGlobal() {
		evl.Globals[sym.GetFingerprint()] = value
	} else {
		evl.SetLocal(sym.GetFingerprint(), value)
	}
}

// evaluate!
func Evaluate(program binder.BoundProgram) {
	evaluator := Evaluator{
		Program:   program,
		Globals:   make(map[string]interface{}),
		Functions: make(map[string]binder.BoundFunction),
		Locals:    []map[string]interface{}{},
	}

	evaluator.PushLocals()

	for _, fnc := range program.Functions {
		symbol := fnc.Symbol
		evaluator.Functions[symbol.GetFingerprint()] = fnc
	}

	mainFunction := program.MainFunction
	body := evaluator.Functions[mainFunction.GetFingerprint()].Body
	EvaluateStatement(body)
}

func (evl *Evaluator) EvaluateStatement(body boundnodes.BoundBlockStatementNode) interface{} {
	labelIndexes := make(map[boundnodes.BoundLabel]int)

	// look through the entire body and look for any labels
	for i, stmt := range body.Statements {
		// if its a label statement...
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// register it!
			labelIndexes[stmt.(boundnodes.BoundLabelStatementNode).Label] = i
		}
	}

	index := 0

	for index < len(body.Statements) {
		stmt := body.Statements[index]

		switch stmt.NodeType() {
		case boundnodes.BoundVariableDeclaration:
			evl.EvaluateVariableDeclaration(stmt.(boundnodes.BoundVariableDeclarationStatementNode))
			index++
			break

		case boundnodes.BoundExpressionStatement:
			evl.EvaluateExpressionStatement(stmt.(boundnodes.BoundExpressionStatementNode))
			index++
			break

		case boundnodes.BoundGotoStatement:
			gotoStatement := stmt.(boundnodes.BoundGotoStatementNode)
			index = labelIndexes[gotoStatement.Label]
			break

		case boundnodes.BoundConditionalGotoStatement:
			gotoStatement := stmt.(boundnodes.BoundConditionalGotoStatementNode)
			condition := evl.EvaluateExpression(gotoStatement.Condition)

			if condition == gotoStatement.JumpIfTrue {
				index = labelIndexes[gotoStatement.Label]
			} else {
				index++
			}

			break

		case boundnodes.BoundLabelStatement:
			index++
			break

		case boundnodes.BoundReturnStatement:
			returnStatement := stmt.(boundnodes.BoundReturnStatementNode)

			if returnStatement.Expression != nil {
				return evl.EvaluateExpression(returnStatement.Expression)
			}

			return nil
		}
	}

	return nil
}

func (evl *Evaluator) EvaluateVariableDeclaration(stmt boundnodes.BoundVariableDeclarationStatementNode) {
	value := evl.EvaluateExpression(stmt.Initializer)
	evl.Assign(stmt.Variable, value)
}

func (evl *Evaluator) EvaluateExpressionStatement(stmt boundnodes.BoundExpressionStatementNode) {
	evl.EvaluateExpression(stmt.Expression)
}

// all them expressionz
func (evl *Evaluator) EvaluateExpression(expr boundnodes.BoundExpressionNode) interface{} {
	switch expr.NodeType() {
	case boundnodes.BoundLiteralExpression:
		return evl.EvaluateLiteralExpression(expr.(boundnodes.BoundLiteralExpressionNode))

	case boundnodes.BoundVariableExpression:
		return evl.EvaluateVariableExpression(expr.(boundnodes.BoundVariableExpressionNode))

	case boundnodes.BoundAssignmentExpression:
		return evl.EvaluateAssignmentExpression(expr.(boundnodes.BoundAssignmentExpressionNode))

	case boundnodes.BoundUnaryExpression:
		return evl.EvaluateUnaryExpression(expr.(boundnodes.BoundUnaryExpressionNode))
	}

	return nil
}

func (evl *Evaluator) EvaluateLiteralExpression(expr boundnodes.BoundLiteralExpressionNode) interface{} {
	return expr.Value
}

func (evl *Evaluator) EvaluateVariableExpression(expr boundnodes.BoundVariableExpressionNode) interface{} {
	if expr.Variable.IsGlobal() {
		return evl.Globals[expr.Variable.GetFingerprint()]
	} else {
		locals := evl.GetLocals()
		return locals[expr.Variable.GetFingerprint()]
	}
}

func (evl *Evaluator) EvaluateAssignmentExpression(expr boundnodes.BoundAssignmentExpressionNode) interface{} {
	value := evl.EvaluateExpression(expr.Expression)
	evl.Assign(expr.Variable, value)
	return value
}

func (evl *Evaluator) EvaluateUnaryExpression(expr boundnodes.BoundUnaryExpressionNode) interface{} {
	value := evl.EvaluateExpression(expr.Expression)
	switch expr.Op.OperatorKind {

	case boundnodes.Identity:
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return value.(int)
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return value.(float32)
		}

	case boundnodes.Negation:
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return -(value.(int))
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return -(value.(float32))
		}

	case boundnodes.LogicalNegation:
		return !(value.(bool))

	default:
		print.PrintC(print.Red, "Unknown unary operation!")
		os.Exit(-1)
	}

	return nil
}

func (evl *Evaluator) EvaluateBinaryExpression(expr boundnodes.BoundBinaryExpressionNode) interface{} {
	left := evl.EvaluateExpression(expr.Left)
	//right := EvaluateBinaryExpression(expr.Right)

	return left
}

func ()