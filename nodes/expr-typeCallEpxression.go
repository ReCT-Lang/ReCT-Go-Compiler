package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type TypeCallExpressionNode struct {
	ExpressionNode

	Base           ExpressionNode
	CallIdentifier lexer.Token
	Arguments      []ExpressionNode
	ClosingToken   lexer.Token
}

// implement node type from interface
func (TypeCallExpressionNode) NodeType() NodeType { return TypeCallExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node TypeCallExpressionNode) Span() print.TextSpan {
	return node.Base.Span().SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node TypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ TypeCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Printf("%s  └ CallIdentifier: %s\n", indent, node.CallIdentifier.Value)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateTypeCallExpressionNode(base ExpressionNode, callId lexer.Token, args []ExpressionNode, closing lexer.Token) TypeCallExpressionNode {
	return TypeCallExpressionNode{
		Base:           base,
		CallIdentifier: callId,
		Arguments:      args,
		ClosingToken:   closing,
	}
}
