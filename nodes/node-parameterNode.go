package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type ParameterNode struct {
	SyntaxNode

	Identifier lexer.Token
	TypeClause TypeClauseNode
}

// implement node type from interface
func (ParameterNode) NodeType() NodeType { return Parameter }

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
