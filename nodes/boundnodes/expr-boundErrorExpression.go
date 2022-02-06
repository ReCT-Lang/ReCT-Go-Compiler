package boundnodes

import (
	"ReCT-Go-Compiler/print"
)

type BoundErrorExpressionNode struct {
	BoundExpressionNode
}

func (BoundErrorExpressionNode) NodeType() BoundType { return BoundErrorExpression }

func (node BoundErrorExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundErrorExpressionNode")
}

func CreateBoundErrorExpressionNode() BoundErrorExpressionNode {
	return BoundErrorExpressionNode{}
}
