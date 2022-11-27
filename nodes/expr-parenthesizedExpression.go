package nodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// aaaaaa
type ParenthesisedExpressionNode struct {
	ExpressionNode

	OpenParenthesis   lexer.Token
	Expression        ExpressionNode
	ClosedParenthesis lexer.Token
}

// implement node type from interface
func (ParenthesisedExpressionNode) NodeType() NodeType { return ParenthesisedExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ParenthesisedExpressionNode) Span() print.TextSpan {
	return node.OpenParenthesis.Span.SpanBetween(node.ClosedParenthesis.Span)
}

// node print function
func (node ParenthesisedExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ParenthesisedExpressionNode")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateParenthesisedExpressionNode(expression ExpressionNode, open lexer.Token, closed lexer.Token) ParenthesisedExpressionNode {
	return ParenthesisedExpressionNode{
		Expression:        expression,
		OpenParenthesis:   open,
		ClosedParenthesis: closed,
	}
}
