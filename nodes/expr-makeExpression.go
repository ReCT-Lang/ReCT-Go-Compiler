package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type MakeExpressionNode struct {
	ExpressionNode

	MakeKeyword  lexer.Token
	ClosingToken lexer.Token

	Package   *lexer.Token
	BaseType  lexer.Token
	Arguments []ExpressionNode
}

// implement node type from interface
func (MakeExpressionNode) NodeType() NodeType { return MakeExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node MakeExpressionNode) Span() print.TextSpan {
	return node.MakeKeyword.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node MakeExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ MakeExpressionNode")
	fmt.Println(indent + "  └ Type: " + node.BaseType.Value)
	fmt.Println(indent + "  └ Arguments: ")
	for _, v := range node.Arguments {
		v.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateMakeExpressionNode(pack *lexer.Token, typ lexer.Token, args []ExpressionNode, makeKw lexer.Token, closing lexer.Token) MakeExpressionNode {
	return MakeExpressionNode{
		Package:      pack,
		BaseType:     typ,
		Arguments:    args,
		MakeKeyword:  makeKw,
		ClosingToken: closing,
	}
}
