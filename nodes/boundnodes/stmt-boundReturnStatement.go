package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundReturnStatementNode struct {
	BoundStatementNode

	Expression BoundExpressionNode
}

// implement the interface
func (BoundReturnStatementNode) NodeType() BoundType { return BoundReturnStatement }
func (node BoundReturnStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundReturnStatementNode")
	fmt.Println(indent + "  └ Expression:")
	node.Expression.Print(indent + "    ")
}

// constructor
func CreateBoundReturnStatementNode(expr BoundExpressionNode) BoundReturnStatementNode {
	return BoundReturnStatementNode{
		Expression: expr,
	}
}
