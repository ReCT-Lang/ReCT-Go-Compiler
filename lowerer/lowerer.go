package lowerer

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

var labelCounter int = 0

func GenerateLabel() boundnodes.BoundLabel {
	labelCounter++
	return boundnodes.BoundLabel(fmt.Sprintf("Label%d", labelCounter))
}

func Lower(functionSymbol symbols.FunctionSymbol, stmt boundnodes.BoundStatementNode) boundnodes.BoundBlockStatementNode {
	result := RewriteStatement(stmt)
	return Flatten(functionSymbol, result)
}

func Flatten(functionSymbol symbols.FunctionSymbol, stmt boundnodes.BoundStatementNode) boundnodes.BoundBlockStatementNode {
	statements := make([]boundnodes.BoundStatementNode, 0)
	stack := make([]boundnodes.BoundStatementNode, 0)

	push := func(stmt boundnodes.BoundStatementNode) {
		stack = append(stack, stmt)
	}

	pop := func() boundnodes.BoundStatementNode {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return element
	}

	push(stmt)

	for len(stack) > 0 {
		current := pop()

		if current.NodeType() == boundnodes.BoundBlockStatement {
			// push all elements onto the stack in reverse order
			currentBlock := current.(boundnodes.BoundBlockStatementNode)
			for i := len(currentBlock.Statements) - 1; i >= 0; i-- {
				push(currentBlock.Statements[i])
			}
		} else {
			statements = append(statements, current)
		}
	}

	if functionSymbol.Type.Fingerprint() == builtins.Void.Fingerprint() {
		if len(statements) == 0 || CanFallThrough(statements[len(statements)-1]) {
			statements = append(statements, boundnodes.CreateBoundReturnStatementNode(nil))
		}
	}

	return boundnodes.CreateBoundBlockStatementNode(statements)
}

func CanFallThrough(stmt boundnodes.BoundStatementNode) bool {
	return stmt.NodeType() != boundnodes.BoundReturnStatement &&
		stmt.NodeType() != boundnodes.BoundGotoStatement
}

func RewriteStatement(stmt boundnodes.BoundStatementNode) boundnodes.BoundStatementNode {
	switch stmt.NodeType() {
	case boundnodes.BoundBlockStatement:
		return RewriteBlockStatement(stmt.(boundnodes.BoundBlockStatementNode))
	case boundnodes.BoundVariableDeclaration:
		return RewriteVariableDeclaration(stmt.(boundnodes.BoundVariableDeclarationStatementNode))
	case boundnodes.BoundIfStatement:
		return RewriteIfStatement(stmt.(boundnodes.BoundIfStatementNode))
	case boundnodes.BoundWhileStatement:
		return RewriteWhileStatement(stmt.(boundnodes.BoundWhileStatementNode))
	case boundnodes.BoundForStatement:
		return RewriteForStatement(stmt.(boundnodes.BoundForStatementNode))
	case boundnodes.BoundFromToStatement:
		return RewriteFromToStatement(stmt.(boundnodes.BoundFromToStatementNode))
	case boundnodes.BoundLabelStatement:
		return RewriteLabelStatement(stmt.(boundnodes.BoundLabelStatementNode))
	case boundnodes.BoundGotoStatement:
		return RewriteGotoStatement(stmt.(boundnodes.BoundGotoStatementNode))
	case boundnodes.BoundConditionalGotoStatement:
		return RewriteConditionalGotoStatement(stmt.(boundnodes.BoundConditionalGotoStatementNode))
	case boundnodes.BoundReturnStatement:
		return RewriteReturnStatement(stmt.(boundnodes.BoundReturnStatementNode))
	case boundnodes.BoundExpressionStatement:
		return RewriteExpressionStatement(stmt.(boundnodes.BoundExpressionStatementNode))
	default:
		return stmt
	}
}

func RewriteBlockStatement(stmt boundnodes.BoundBlockStatementNode) boundnodes.BoundBlockStatementNode {
	rewrittenStatements := make([]boundnodes.BoundStatementNode, 0)

	for _, statement := range stmt.Statements {
		rewrittenStatements = append(rewrittenStatements, RewriteStatement(statement))
	}

	return boundnodes.CreateBoundBlockStatementNode(rewrittenStatements)
}

func RewriteVariableDeclaration(stmt boundnodes.BoundVariableDeclarationStatementNode) boundnodes.BoundVariableDeclarationStatementNode {
	initializer := RewriteExpression(stmt.Initializer)
	return boundnodes.CreateBoundVariableDeclarationStatementNode(stmt.Variable, initializer)
}

func RewriteIfStatement(stmt boundnodes.BoundIfStatementNode) boundnodes.BoundIfStatementNode {
	condition := RewriteExpression(stmt.Condition)
	thenStatement := RewriteStatement(stmt.ThenStatement)
	var elseStatement boundnodes.BoundStatementNode = nil
	if stmt.ElseStatement != nil {
		elseStatement = RewriteStatement(stmt.ElseStatement)
	}

	return boundnodes.CreateBoundIfStatementNode(condition, thenStatement, elseStatement)
}

func RewriteWhileStatement(stmt boundnodes.BoundWhileStatementNode) boundnodes.BoundWhileStatementNode {
	condition := RewriteExpression(stmt.Condition)
	body := RewriteStatement(stmt.Body)

	return boundnodes.CreateBoundWhileStatementNode(condition, body, stmt.BreakLabel, stmt.ContinueLabel)
}

func RewriteForStatement(stmt boundnodes.BoundForStatementNode) boundnodes.BoundForStatementNode {
	variable := RewriteStatement(stmt.Variable).(boundnodes.BoundVariableDeclarationStatementNode)
	condition := RewriteExpression(stmt.Condition)
	action := RewriteStatement(stmt.Action)

	body := RewriteStatement(stmt.Body)
	return boundnodes.CreateBoundForStatementNode(variable, condition, action, body, stmt.BreakLabel, stmt.ContinueLabel)
}

func RewriteFromToStatement(stmt boundnodes.BoundFromToStatementNode) boundnodes.BoundFromToStatementNode {
	lowerBound := RewriteExpression(stmt.LowerBound)
	upperBound := RewriteExpression(stmt.UpperBound)
	body := RewriteStatement(stmt.Body)

	return boundnodes.CreateBoundFromToStatementNode(stmt.Variable, lowerBound, upperBound, body, stmt.BreakLabel, stmt.ContinueLabel)
}

func RewriteLabelStatement(stmt boundnodes.BoundLabelStatementNode) boundnodes.BoundLabelStatementNode {
	return stmt
}

func RewriteGotoStatement(stmt boundnodes.BoundGotoStatementNode) boundnodes.BoundGotoStatementNode {
	return stmt
}

func RewriteConditionalGotoStatement(stmt boundnodes.BoundConditionalGotoStatementNode) boundnodes.BoundConditionalGotoStatementNode {
	condition := RewriteExpression(stmt.Condition)
	return boundnodes.CreateBoundConditionalGotoStatementNode(condition, stmt.Label, stmt.JumpIfTrue)
}

func RewriteReturnStatement(stmt boundnodes.BoundReturnStatementNode) boundnodes.BoundReturnStatementNode {
	var expression boundnodes.BoundExpressionNode = nil
	if stmt.Expression != nil {
		expression = RewriteExpression(stmt.Expression)
	}

	return boundnodes.CreateBoundReturnStatementNode(expression)
}

func RewriteExpressionStatement(stmt boundnodes.BoundExpressionStatementNode) boundnodes.BoundExpressionStatementNode {
	expression := RewriteExpression(stmt.Expression)
	return boundnodes.CreateBoundExpressionStatementNode(expression)
}

func RewriteExpression(expr boundnodes.BoundExpressionNode) boundnodes.BoundExpressionNode {
	switch expr.NodeType() {
	case boundnodes.BoundErrorExpression:
		return RewriteErrorExpression(expr.(boundnodes.BoundErrorExpressionNode))
	case boundnodes.BoundLiteralExpression:
		return RewriteLiteralExpression(expr.(boundnodes.BoundLiteralExpressionNode))
	case boundnodes.BoundVariableExpression:
		return RewriteVariableExpression(expr.(boundnodes.BoundVariableExpressionNode))
	case boundnodes.BoundAssignmentExpression:
		return RewriteAssignmentExpression(expr.(boundnodes.BoundAssignmentExpressionNode))
	case boundnodes.BoundUnaryExpression:
		return RewriteUnaryExpression(expr.(boundnodes.BoundUnaryExpressionNode))
	case boundnodes.BoundBinaryExpression:
		return RewriteBinaryExpression(expr.(boundnodes.BoundBinaryExpressionNode))
	case boundnodes.BoundCallExpression:
		return RewriteCallExpression(expr.(boundnodes.BoundCallExpressionNode))
	case boundnodes.BoundConversionExpression:
		return RewriteConversionExpression(expr.(boundnodes.BoundConversionExpressionNode))
	default:
		return expr
	}
}

func RewriteErrorExpression(expr boundnodes.BoundErrorExpressionNode) boundnodes.BoundErrorExpressionNode {
	return expr
}

func RewriteLiteralExpression(expr boundnodes.BoundLiteralExpressionNode) boundnodes.BoundLiteralExpressionNode {
	return expr
}

func RewriteVariableExpression(expr boundnodes.BoundVariableExpressionNode) boundnodes.BoundVariableExpressionNode {
	return expr
}

func RewriteAssignmentExpression(expr boundnodes.BoundAssignmentExpressionNode) boundnodes.BoundAssignmentExpressionNode {
	expression := RewriteExpression(expr.Expression)
	return boundnodes.CreateBoundAssignmentExpressionNode(expr.Variable, expression)
}

func RewriteUnaryExpression(expr boundnodes.BoundUnaryExpressionNode) boundnodes.BoundUnaryExpressionNode {
	operand := RewriteExpression(expr.Expression)
	return boundnodes.CreateBoundUnaryExpressionNode(expr.Op, operand)
}

func RewriteBinaryExpression(expr boundnodes.BoundBinaryExpressionNode) boundnodes.BoundBinaryExpressionNode {
	left := RewriteExpression(expr.Left)
	right := RewriteExpression(expr.Right)
	return boundnodes.CreateBoundBinaryExpressionNode(left, expr.Op, right)
}

func RewriteCallExpression(expr boundnodes.BoundCallExpressionNode) boundnodes.BoundCallExpressionNode {
	rewrittenArgs := make([]boundnodes.BoundExpressionNode, 0)

	for _, arg := range expr.Arguments {
		rewrittenArgs = append(rewrittenArgs, RewriteExpression(arg))
	}

	return boundnodes.CreateBoundCallExpressionNode(expr.Function, rewrittenArgs)
}

func RewriteConversionExpression(expr boundnodes.BoundConversionExpressionNode) boundnodes.BoundConversionExpressionNode {
	expression := RewriteExpression(expr.Expression)

	return boundnodes.CreateBoundConversionExpressionNode(expr.ToType, expression)
}
