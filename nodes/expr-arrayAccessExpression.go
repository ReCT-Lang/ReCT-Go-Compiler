package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type ArrayAccessExpressionNode struct {
	ExpressionNode

	Base           ExpressionNode
	Index          ExpressionNode
	ClosingBracket lexer.Token
}

// implement node type from interface
func (ArrayAccessExpressionNode) NodeType() NodeType { return ArrayAccessExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ArrayAccessExpressionNode) Span() print.TextSpan {
	return node.Base.Span().SpanBetween(node.ClosingBracket.Span)
}

// node print function
func (node ArrayAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ArrayAccessExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateArrayAccessExpressionNode(base ExpressionNode, index ExpressionNode) ArrayAccessExpressionNode {
	return ArrayAccessExpressionNode{
		Base:  base,
		Index: index,
	}
}
