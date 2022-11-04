package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ReturnStatementNode like: return "Yo mama"; there, get rect.
type ReturnStatementNode struct {
	StatementNode

	Keyword    lexer.Token
	Expression ExpressionNode
}

// NodeType Copy + Paste
func (ReturnStatementNode) NodeType() NodeType { return ReturnStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ReturnStatementNode) Span() print.TextSpan {
	return node.Keyword.Span.SpanBetween(node.Expression.Span())
}

// Print Prints beautiful stuff in console
func (node ReturnStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ReturnStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)

	if node.Expression == nil {
		fmt.Printf("%s  └ Expression: none\n", indent)
	} else {
		fmt.Println(indent + "  └ Expression: ")
		node.Expression.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateReturnStatementNode(keyword lexer.Token, expression ExpressionNode) ReturnStatementNode {
	return ReturnStatementNode{
		Keyword:    keyword,
		Expression: expression,
	}
}
