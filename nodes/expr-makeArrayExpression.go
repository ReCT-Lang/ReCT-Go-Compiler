package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type MakeArrayExpressionNode struct {
	ExpressionNode

	IsLiteral bool

	MakeKeyword  lexer.Token
	ClosingToken lexer.Token

	Type          TypeClauseNode
	Length        ExpressionNode
	LiteralValues []ExpressionNode
}

// implement node type from interface
func (MakeArrayExpressionNode) NodeType() NodeType { return MakeArrayExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node MakeArrayExpressionNode) Span() print.TextSpan {
	return node.MakeKeyword.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node MakeArrayExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ MakeArrayExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.Type.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateMakeArrayExpressionNode(typ TypeClauseNode, length ExpressionNode, makeKw lexer.Token, closing lexer.Token) MakeArrayExpressionNode {
	return MakeArrayExpressionNode{
		Type:         typ,
		Length:       length,
		IsLiteral:    false,
		MakeKeyword:  makeKw,
		ClosingToken: closing,
	}
}
func CreateMakeArrayExpressionNodeLiteral(typ TypeClauseNode, literals []ExpressionNode, makeKw lexer.Token, closing lexer.Token) MakeArrayExpressionNode {
	return MakeArrayExpressionNode{
		Type:          typ,
		LiteralValues: literals,
		IsLiteral:     true,
		MakeKeyword:   makeKw,
		ClosingToken:  closing,
	}
}
