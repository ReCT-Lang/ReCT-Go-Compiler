package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundThisExpressionNode struct {
	BoundExpressionNode

	Class     symbols.ClassSymbol
	BoundSpan print.TextSpan
}

func (BoundThisExpressionNode) NodeType() BoundType { return BoundThisExpression }

func (node BoundThisExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThisExpressionNode")
}

func (node BoundThisExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundThisExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThisExpressionNode) Type() symbols.TypeSymbol { return node.Class.Type }

func CreateBoundThisExpressionNode(class symbols.ClassSymbol, span print.TextSpan) BoundThisExpressionNode {
	return BoundThisExpressionNode{
		Class:     class,
		BoundSpan: span,
	}
}
