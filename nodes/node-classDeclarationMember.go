package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type ClassDeclarationMember struct {
	MemberNode

	ClassKeyword lexer.Token
	Identifier   lexer.Token
	Members      []MemberNode
	ClosingToken lexer.Token
}

// implement node type from interface
func (ClassDeclarationMember) NodeType() NodeType { return ClassDeclaration }

func (node ClassDeclarationMember) Span() print.TextSpan {
	return node.ClassKeyword.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node ClassDeclarationMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- ClassDeclarationMember")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Kind)

	fmt.Println(indent + "  └ Members: ")
	for _, mem := range node.Members {
		mem.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateClassDeclarationMember(kw lexer.Token, id lexer.Token, members []MemberNode, closing lexer.Token) ClassDeclarationMember {
	return ClassDeclarationMember{
		ClassKeyword: kw,
		Identifier:   id,
		Members:      members,
		ClosingToken: closing,
	}
}
