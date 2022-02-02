package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type VariableDeclarationStatementNode struct {
	StatementNode

	Keyword     lexer.Token
	Identifier  lexer.Token
	AssignToken lexer.Token
	Initializer ExpressionNode
}

// implement node type from interface
func (VariableDeclarationStatementNode) NodeType() NodeType { return VariableDeclaration }

// node print function
func (node VariableDeclarationStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ VariableDeclarationStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
	fmt.Println(indent + "  └ Initializer: ")
	node.Initializer.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateVariableDeclarationStatementNode(kw lexer.Token, id lexer.Token, assign lexer.Token, init ExpressionNode) VariableDeclarationStatementNode {
	return VariableDeclarationStatementNode{
		Keyword:     kw,
		Identifier:  id,
		AssignToken: assign,
		Initializer: init,
	}
}
