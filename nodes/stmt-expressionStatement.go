package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// basic global statement member
type ExpressionStatementNode struct {
	StatementNode

	Expression ExpressionNode
}

// implement node type from interface
func (ExpressionStatementNode) NodeType() NodeType { return ExpressionStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ExpressionStatementNode) Span() print.TextSpan {
	return node.Expression.Span()
}

// node print function
func (node ExpressionStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ExpressionStatementNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")

}

// "constructor" / ooga booga OOP cave man brain
func CreateExpressionStatementNode(expr ExpressionNode) ExpressionStatementNode {
	return ExpressionStatementNode{
		Expression: expr,
	}
}
