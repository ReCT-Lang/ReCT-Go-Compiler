package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type TypeClauseNode struct {
	SyntaxNode

	ClauseIsSet bool

	Package        *lexer.Token
	TypeIdentifier lexer.Token
	SubClauses     []TypeClauseNode
	ClosingBracket lexer.Token
}

// implement node type from interface
func (TypeClauseNode) NodeType() NodeType { return TypeClause }

func (node TypeClauseNode) Span() print.TextSpan {
	if node.ClauseIsSet {
		span := node.TypeIdentifier.Span
		span = span.SpanBetween(node.ClosingBracket.Span)

		return span
	} else {
		return print.TextSpan{}
	}
}

// node print function
func (node TypeClauseNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ TypeClauseNode")
	fmt.Printf("%s  └ Type: %s\n", indent, node.TypeIdentifier.Value)

}

// "constructor" / ooga booga OOP cave man brain
func CreateTypeClauseNode(pack *lexer.Token, id lexer.Token, subtypes []TypeClauseNode, bracket lexer.Token) TypeClauseNode {
	return TypeClauseNode{
		ClauseIsSet:    true,
		Package:        pack,
		TypeIdentifier: id,
		SubClauses:     subtypes,
		ClosingBracket: bracket,
	}
}
