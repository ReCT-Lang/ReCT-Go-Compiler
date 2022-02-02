package nodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type ExpressionStatementNode struct {
	StatementNode

	Expression ExpressionNode
}

// implement node type from interface
func (ExpressionStatementNode) NodeType() NodeType { return ExpressionStatement }

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
