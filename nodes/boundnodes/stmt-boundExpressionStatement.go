package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundExpressionStatementNode struct {
	BoundStatementNode

	Expression BoundExpressionNode
	BoundSpan  print.TextSpan
}

// implement the interface
func (BoundExpressionStatementNode) NodeType() BoundType { return BoundExpressionStatement }
func (node BoundExpressionStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundExpressionStatementNode")
	fmt.Println(indent + "  └ Expression:")
	node.Expression.Print(indent + "    ")
}

func (node BoundExpressionStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundExpressionStatementNode(expr BoundExpressionNode, span print.TextSpan) BoundExpressionStatementNode {
	return BoundExpressionStatementNode{
		Expression: expr,
		BoundSpan:  span,
	}
}
