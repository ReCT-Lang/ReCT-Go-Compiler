package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type BlockStatementNode struct {
	StatementNode

	OpenBrace  lexer.Token
	Statments  []StatementNode
	CloseBrace lexer.Token
}

// implement node type from interface
func (node BlockStatementNode) NodeType() NodeType { return BlockStatement }

// node print function
func (node BlockStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BlockStatementNode")
	fmt.Printf("%s  └ OpenBrace: %s\n", indent, node.OpenBrace.Kind)
	fmt.Printf("%s  └ CloseBrace: %s\n", indent, node.CloseBrace.Kind)
	fmt.Println(indent + "  └ Statements: ")

	for _, stmt := range node.Statments {
		stmt.Print(indent + "    ")
	}

}

// "constructor" / ooga booga OOP cave man brain
func CreateBlockStatementNode(openBrace lexer.Token, statements []StatementNode, closeBrace lexer.Token) BlockStatementNode {
	return BlockStatementNode{
		OpenBrace:  openBrace,
		Statments:  statements,
		CloseBrace: closeBrace,
	}
}
