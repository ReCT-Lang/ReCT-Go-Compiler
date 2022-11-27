package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundReferenceExpressionNode struct {
	BoundExpressionNode

	Expression    BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundReferenceExpressionNode) NodeType() BoundType { return BoundReferenceExpression }

func (node BoundReferenceExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundReferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (node BoundReferenceExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundReferenceExpressionNode) Type() symbols.TypeSymbol {
	return symbols.CreateTypeSymbol("pointer", []symbols.TypeSymbol{node.Expression.Type()}, false, false, false, symbols.PackageSymbol{}, nil)
}

func CreateBoundReferenceExpressionNode(expression BoundExpressionNode, src nodes.SyntaxNode) BoundReferenceExpressionNode {
	return BoundReferenceExpressionNode{
		Expression:    expression,
		UnboundSource: src,
	}
}
