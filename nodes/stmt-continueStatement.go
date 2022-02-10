package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ReturnStatementNode like: return "Yo mama"; there, get rect.
type ContinueStatementNode struct {
	StatementNode

	Keyword lexer.Token
}

// NodeType Copy + Paste
func (ContinueStatementNode) NodeType() NodeType { return ContinueStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ContinueStatementNode) Position() (int, int, int) {
	return node.Keyword.Line, node.Keyword.Column, len(node.Keyword.Value)
}

// Print Prints beautiful stuff in console
func (node ContinueStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ReturnStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateContinueStatement(keyword lexer.Token) ContinueStatementNode {
	return ContinueStatementNode{
		Keyword: keyword,
	}
}
