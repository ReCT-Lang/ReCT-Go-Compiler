package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type MakeStructExpressionNode struct {
	ExpressionNode

	MakeKeyword  lexer.Token
	ClosingToken lexer.Token

	Type          lexer.Token
	LiteralValues []ExpressionNode
}

// implement node type from interface
func (MakeStructExpressionNode) NodeType() NodeType { return MakeStructExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node MakeStructExpressionNode) Span() print.TextSpan {
	return node.MakeKeyword.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node MakeStructExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ MakeStructExpressionNode")
	fmt.Println(indent + "  └ Type: " + node.Type.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreateMakeStructExpressionNode(typ lexer.Token, literals []ExpressionNode, makeKw lexer.Token, closing lexer.Token) MakeStructExpressionNode {
	return MakeStructExpressionNode{
		Type:          typ,
		LiteralValues: literals,
		MakeKeyword:   makeKw,
		ClosingToken:  closing,
	}
}
