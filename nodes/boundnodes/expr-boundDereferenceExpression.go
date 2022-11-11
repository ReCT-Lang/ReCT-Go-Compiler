package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundDereferenceExpressionNode struct {
	BoundExpressionNode

	Expression BoundExpressionNode
	BoundSpan  print.TextSpan
}

func (BoundDereferenceExpressionNode) NodeType() BoundType { return BoundDereferenceExpression }

func (node BoundDereferenceExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (node BoundDereferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (node BoundDereferenceExpressionNode) IsPersistent() bool { return node.Expression.IsPersistent() }

// implement the expression node interface
func (node BoundDereferenceExpressionNode) Type() symbols.TypeSymbol {
	return node.Expression.Type().SubTypes[0]
}

func CreateBoundDereferenceExpressionNode(expression BoundExpressionNode, span print.TextSpan) BoundDereferenceExpressionNode {
	return BoundDereferenceExpressionNode{
		Expression: expression,
		BoundSpan:  span,
	}
}
