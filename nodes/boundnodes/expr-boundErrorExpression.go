package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundErrorExpressionNode struct {
	BoundExpressionNode
}

func (BoundErrorExpressionNode) NodeType() BoundType { return BoundErrorExpression }

func (node BoundErrorExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundErrorExpressionNode")
}

func (BoundErrorExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundErrorExpressionNode) Type() symbols.TypeSymbol { return builtins.Error }

func CreateBoundErrorExpressionNode() BoundErrorExpressionNode {
	return BoundErrorExpressionNode{}
}
