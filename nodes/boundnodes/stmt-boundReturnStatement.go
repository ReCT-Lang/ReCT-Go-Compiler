package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundReturnStatementNode struct {
	BoundStatementNode

	Expression BoundExpressionNode
	BoundSpan  print.TextSpan
}

// implement the interface
func (BoundReturnStatementNode) NodeType() BoundType { return BoundReturnStatement }
func (node BoundReturnStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundReturnStatementNode")
	if node.Expression == nil {
		fmt.Println(indent + "  └ Expression: none")
	} else {
		fmt.Println(indent + "  └ Expression:")
		node.Expression.Print(indent + "    ")
	}
}

func (node BoundReturnStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundReturnStatementNode(expr BoundExpressionNode, span print.TextSpan) BoundReturnStatementNode {
	return BoundReturnStatementNode{
		Expression: expr,
		BoundSpan:  span,
	}
}
