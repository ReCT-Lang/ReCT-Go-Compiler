package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
	"strings"
)

// basic global statement member
type FunctionDeclarationMember struct {
	MemberNode

	Identifier lexer.Token
	Parameters []ParameterNode
	TypeClause TypeClauseNode
	Body       BlockStatementNode
}

// implement node type from interface
func (FunctionDeclarationMember) NodeType() NodeType { return FunctionDeclaration }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
// For FunctionDeclarationMember we don't get the length of the body, only the keyword, name, parameters, and type
func (node FunctionDeclarationMember) Position() (int, int, int) {
	length := len(strings.Split(string(lexer.FunctionKeyword), " ")[0])
	length += len(node.Identifier.Value)
	_, _, typeLength := node.TypeClause.Position()
	length += typeLength
	for _, arg := range node.Parameters {
		_, _, argLength := arg.Position()
		length += argLength + 2 // +2 for spaces and commas
	}
	return node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node FunctionDeclarationMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- FunctionDeclarationMember")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Kind)

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

	fmt.Println(indent + "  └ Body: ")
	node.Body.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateFunctionDeclarationMember(id lexer.Token, params []ParameterNode, typeClause TypeClauseNode, body BlockStatementNode) FunctionDeclarationMember {
	return FunctionDeclarationMember{
		Identifier: id,
		Parameters: params,
		TypeClause: typeClause,
		Body:       body,
	}
}
