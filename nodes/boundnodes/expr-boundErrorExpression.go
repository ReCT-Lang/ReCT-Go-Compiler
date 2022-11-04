package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundErrorExpressionNode struct {
	BoundExpressionNode
	BoundSpan print.TextSpan
}

func (BoundErrorExpressionNode) NodeType() BoundType { return BoundErrorExpression }

func (node BoundErrorExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundErrorExpressionNode")
}

func (node BoundErrorExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundErrorExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundErrorExpressionNode) Type() symbols.TypeSymbol { return builtins.Error }

func CreateBoundErrorExpressionNode(span print.TextSpan) BoundErrorExpressionNode {
	return BoundErrorExpressionNode{
		BoundSpan: span,
	}
}
