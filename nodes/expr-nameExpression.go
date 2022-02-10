package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// aaaaaa, b even
type NameExpressionNode struct {
	ExpressionNode

	Identifier lexer.Token
}

// implement node type from interface
func (NameExpressionNode) NodeType() NodeType { return NameExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node NameExpressionNode) Position() (int, int, int) {
	return node.Identifier.Line, node.Identifier.Column, len(node.Identifier.Value)
}

// node print function
func (node NameExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ NameExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreateNameExpressionNode(id lexer.Token) NameExpressionNode {
	return NameExpressionNode{
		Identifier: id,
	}
}
