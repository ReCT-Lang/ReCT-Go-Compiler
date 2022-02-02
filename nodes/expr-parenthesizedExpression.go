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
