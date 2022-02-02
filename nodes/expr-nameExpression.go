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
