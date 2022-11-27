package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type BoundReturnStatementNode struct {
	BoundStatementNode

	Expression    BoundExpressionNode
	UnboundSource nodes.SyntaxNode
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

func (node BoundReturnStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundReturnStatementNode(expr BoundExpressionNode, src nodes.SyntaxNode) BoundReturnStatementNode {
	return BoundReturnStatementNode{
		Expression:    expr,
		UnboundSource: src,
	}
}
