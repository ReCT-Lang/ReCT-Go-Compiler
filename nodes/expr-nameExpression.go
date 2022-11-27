package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// aaaaaa, b even
type NameExpressionNode struct {
	ExpressionNode

	InMain     bool
	Identifier lexer.Token
}

// implement node type from interface
func (NameExpressionNode) NodeType() NodeType { return NameExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node NameExpressionNode) Span() print.TextSpan {
	return node.Identifier.Span
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

func CreateMainNameExpressionNode(id lexer.Token) NameExpressionNode {
	return NameExpressionNode{
		Identifier: id,
		InMain:     true,
	}
}
