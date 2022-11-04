package nodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type TernaryExpressionNode struct {
	ExpressionNode

	Condition ExpressionNode
	If        ExpressionNode
	Else      ExpressionNode
}

// implement node type from interface
func (TernaryExpressionNode) NodeType() NodeType { return TernaryExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node TernaryExpressionNode) Span() print.TextSpan {
	return node.Condition.Span().SpanBetween(node.Else.Span())
}

// node print function
func (node TernaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ TernaryExpressionNode")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ If: ")
	node.If.Print(indent + "    ")
	fmt.Println(indent + "  └ Else: ")
	node.Else.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateTernaryExpressionNode(condition ExpressionNode, ifexpr ExpressionNode, elseexpr ExpressionNode) TernaryExpressionNode {
	return TernaryExpressionNode{
		Condition: condition,
		If:        ifexpr,
		Else:      elseexpr,
	}
}
