package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type BoundExpressionStatementNode struct {
	BoundStatementNode

	Expression    BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundExpressionStatementNode) NodeType() BoundType { return BoundExpressionStatement }
func (node BoundExpressionStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundExpressionStatementNode")
	fmt.Println(indent + "  └ Expression:")
	node.Expression.Print(indent + "    ")
}

func (node BoundExpressionStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundExpressionStatementNode(expr BoundExpressionNode, src nodes.SyntaxNode) BoundExpressionStatementNode {
	return BoundExpressionStatementNode{
		Expression:    expr,
		UnboundSource: src,
	}
}
