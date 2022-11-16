package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type ExternalFunctionDeclarationMember struct {
	MemberNode

	ExternalKeyword lexer.Token
	Identifier      lexer.Token
	Parameters      []ParameterNode
	ClosingToken    lexer.Token
	TypeClause      TypeClauseNode
	IsVariadic      bool
	IsAdapted       bool
}

// implement node type from interface
func (ExternalFunctionDeclarationMember) NodeType() NodeType { return ExternalFunctionDeclaration }

func (node ExternalFunctionDeclarationMember) Span() print.TextSpan {
	span := node.ExternalKeyword.Span.SpanBetween(node.ClosingToken.Span)
	if node.TypeClause.ClauseIsSet {
		span = span.SpanBetween(node.TypeClause.Span())
	}

	return span
}

// node print function
func (node ExternalFunctionDeclarationMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- FunctionDeclarationMember")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Kind)
	fmt.Printf("%s  └ IsVariadic: %t\n", indent, node.IsVariadic)
	fmt.Printf("%s  └ IsAdapted: %t\n", indent, node.IsAdapted)

	fmt.Println(indent + "  └ Parameters: ")
	for _, param := range node.Parameters {
		param.Print(indent + "    ")
	}

	if !node.TypeClause.ClauseIsSet {
		fmt.Printf("%s  └ TypeClause: none\n", indent)
	} else {
		fmt.Println(indent + "  └ TypeClause: ")
		node.TypeClause.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateExternalFunctionDeclarationMember(kw lexer.Token, id lexer.Token, params []ParameterNode, typeClause TypeClauseNode, closing lexer.Token, variadic bool, adapted bool) ExternalFunctionDeclarationMember {
	return ExternalFunctionDeclarationMember{
		ExternalKeyword: kw,
		Identifier:      id,
		Parameters:      params,
		TypeClause:      typeClause,
		ClosingToken:    closing,
		IsVariadic:      variadic,
		IsAdapted:       adapted,
	}
}
