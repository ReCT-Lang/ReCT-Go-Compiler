package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type TypeClauseNode struct {
	SyntaxNode

	ClauseIsSet    bool
	TypeIdentifier lexer.Token
}

// implement node type from interface
func (TypeClauseNode) NodeType() NodeType { return TypeClause }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node TypeClauseNode) Position() (int, int, int) {
	return node.TypeIdentifier.Line, node.TypeIdentifier.Column, len(node.TypeIdentifier.Value)
}

// node print function
func (node TypeClauseNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ TypeClauseNode")
	fmt.Printf("%s  └ Type: %s\n", indent, node.TypeIdentifier.Value)

}

// "constructor" / ooga booga OOP cave man brain
func CreateTypeClauseNode(id lexer.Token) TypeClauseNode {
	return TypeClauseNode{
		TypeIdentifier: id,
		ClauseIsSet:    true,
	}
}
