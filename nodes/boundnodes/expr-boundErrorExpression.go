package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundErrorExpressionNode struct {
	BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundErrorExpressionNode) NodeType() BoundType { return BoundErrorExpression }

func (node BoundErrorExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundErrorExpressionNode")
}

func (node BoundErrorExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundErrorExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundErrorExpressionNode) Type() symbols.TypeSymbol { return builtins.Error }

func CreateBoundErrorExpressionNode(src nodes.SyntaxNode) BoundErrorExpressionNode {
	return BoundErrorExpressionNode{
		UnboundSource: src,
	}
}
