package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
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
func (node ContinueStatementNode) Span() print.TextSpan {
	return node.Keyword.Span
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
