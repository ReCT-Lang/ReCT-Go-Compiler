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
	Statements []StatementNode
	CloseBrace lexer.Token
}

// implement node type from interface
func (BlockStatementNode) NodeType() NodeType { return BlockStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
// BlockStatementNode is a bit weird because it gets the length of all the statements...
// Oh well, I'm sure nothing wacky can happen from this! - tokorv
func (node BlockStatementNode) Position() (int, int, int) {
	length := 2 // OpenBrace + CloseBrace
	for _, s := range node.Statements {
		_, _, stmtLength := s.Position()
		length += stmtLength
	}
	return node.OpenBrace.Line, node.OpenBrace.Column, length
}

// node print function
func (node BlockStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BlockStatementNode")
	fmt.Printf("%s  └ OpenBrace: %s\n", indent, node.OpenBrace.Kind)
	fmt.Printf("%s  └ CloseBrace: %s\n", indent, node.CloseBrace.Kind)
	fmt.Println(indent + "  └ Statements: ")

	for _, stmt := range node.Statements {
		stmt.Print(indent + "    ")
	}

}

// "constructor" / ooga booga OOP cave man brain
func CreateBlockStatementNode(openBrace lexer.Token, statements []StatementNode, closeBrace lexer.Token) BlockStatementNode {
	return BlockStatementNode{
		OpenBrace:  openBrace,
		Statements: statements,
		CloseBrace: closeBrace,
	}
}
