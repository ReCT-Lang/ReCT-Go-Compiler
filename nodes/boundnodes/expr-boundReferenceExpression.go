package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundReferenceExpressionNode struct {
	BoundExpressionNode

	Expression BoundExpressionNode
	BoundSpan  print.TextSpan
}

func (BoundReferenceExpressionNode) NodeType() BoundType { return BoundReferenceExpression }

func (node BoundReferenceExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (node BoundReferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (node BoundReferenceExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundReferenceExpressionNode) Type() symbols.TypeSymbol {
	return symbols.CreateTypeSymbol("pointer", []symbols.TypeSymbol{node.Expression.Type()}, false, false, false)
}

func CreateBoundReferenceExpressionNode(expression BoundExpressionNode, span print.TextSpan) BoundReferenceExpressionNode {
	return BoundReferenceExpressionNode{
		Expression: expression,
		BoundSpan:  span,
	}
}
