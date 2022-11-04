package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ReturnStatementNode like: return "Yo mama"; there, get rect.
type BreakStatementNode struct {
	StatementNode

	Keyword lexer.Token
}

// NodeType Copy + Paste
func (BreakStatementNode) NodeType() NodeType { return BreakStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node BreakStatementNode) Span() print.TextSpan {
	return node.Keyword.Span
}

// Print Prints beautiful stuff in console
func (node BreakStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BreakStatement")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateBreakStatement(keyword lexer.Token) BreakStatementNode {
	return BreakStatementNode{
		Keyword: keyword,
	}
}
