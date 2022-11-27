package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type StructDeclarationMember struct {
	MemberNode

	StructKeyword lexer.Token
	Identifier    lexer.Token
	Fields        []ParameterNode
	ClosingToken  lexer.Token
}

// implement node type from interface
func (StructDeclarationMember) NodeType() NodeType { return StructDeclaration }

func (node StructDeclarationMember) Span() print.TextSpan {
	return node.StructKeyword.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node StructDeclarationMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- ClassDeclarationMember")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Kind)
}

// "constructor" / ooga booga OOP cave man brain
func CreateStructDeclarationMember(kw lexer.Token, id lexer.Token, fields []ParameterNode, closing lexer.Token) StructDeclarationMember {
	return StructDeclarationMember{
		StructKeyword: kw,
		Identifier:    id,
		Fields:        fields,
		ClosingToken:  closing,
	}
}
