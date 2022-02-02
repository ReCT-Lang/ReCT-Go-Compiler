package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
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
