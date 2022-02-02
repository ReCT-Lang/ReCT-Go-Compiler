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
