package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type ClassDeclarationMember struct {
	MemberNode

	Identifier lexer.Token
	Members    []MemberNode
}

// implement node type from interface
func (ClassDeclarationMember) NodeType() NodeType { return ClassDeclaration }

// aaaaAAAAAA
func (node ClassDeclarationMember) Position() (int, int, int) {
	return 0, 9, 0
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
func CreateClassDeclarationMember(id lexer.Token, members []MemberNode) ClassDeclarationMember {
	return ClassDeclarationMember{
		Identifier: id,
		Members:    members,
	}
}
