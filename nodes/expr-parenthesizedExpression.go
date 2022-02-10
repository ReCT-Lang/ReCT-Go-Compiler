package nodes

import (
	"ReCT-Go-Compiler/print"
)

// aaaaaa
type ParenthesisedExpressionNode struct {
	ExpressionNode

	Expression ExpressionNode
}

// implement node type from interface
func (ParenthesisedExpressionNode) NodeType() NodeType { return ParenthesisedExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ParenthesisedExpressionNode) Position() (int, int, int) {
	line, column, length := node.Expression.Position()
	return line, column + 1, length + 1 // +1s for the parentheses
}

// node print function
func (node ParenthesisedExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ParenthesisedExpressionNode")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateParenthesisedExpressionNode(expression ExpressionNode) ParenthesisedExpressionNode {
	return ParenthesisedExpressionNode{
		Expression: expression,
	}
}
