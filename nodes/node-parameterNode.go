package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// basic global statement member
type ParameterNode struct {
	SyntaxNode

	Identifier lexer.Token
	TypeClause TypeClauseNode
}

// implement node type from interface
func (ParameterNode) NodeType() NodeType { return Parameter }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ParameterNode) Span() print.TextSpan {
	return node.Identifier.Span.SpanBetween(node.TypeClause.Span())
}

// node print function
func (node ParameterNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ParameterNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)

	if !node.TypeClause.ClauseIsSet {
		fmt.Printf("%s  └ TypeClause: none\n", indent)
	} else {
		fmt.Println(indent + "  └ TypeClause: ")
		node.TypeClause.Print(indent + "    ")
	}

}

// "constructor" / ooga booga OOP cave man brain
func CreateParameterNode(id lexer.Token, typeClause TypeClauseNode) ParameterNode {
	return ParameterNode{
		Identifier: id,
		TypeClause: typeClause,
	}
}
