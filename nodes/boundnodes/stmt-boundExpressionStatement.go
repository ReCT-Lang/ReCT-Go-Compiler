package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundExpressionStatementNode struct {
	BoundStatementNode

	Expression BoundExpressionNode
}

// implement the interface
func (BoundExpressionStatementNode) NodeType() BoundType { return BoundExpressionStatement }
func (node BoundExpressionStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundExpressionStatementNode")
	fmt.Println(indent + "  └ Expression:")
	node.Expression.Print(indent + "    ")
}

// constructor
func CreateBoundExpressionStatementNode(expr BoundExpressionNode) BoundExpressionStatementNode {
	return BoundExpressionStatementNode{
		Expression: expr,
	}
}
