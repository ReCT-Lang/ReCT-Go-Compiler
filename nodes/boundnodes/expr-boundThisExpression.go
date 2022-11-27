package boundnodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundThisExpressionNode struct {
	BoundExpressionNode

	Class         symbols.ClassSymbol
	UnboundSource nodes.SyntaxNode
}

func (BoundThisExpressionNode) NodeType() BoundType { return BoundThisExpression }

func (node BoundThisExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThisExpressionNode")
}

func (node BoundThisExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundThisExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThisExpressionNode) Type() symbols.TypeSymbol { return node.Class.Type }

func CreateBoundThisExpressionNode(class symbols.ClassSymbol, src nodes.SyntaxNode) BoundThisExpressionNode {
	return BoundThisExpressionNode{
		Class:         class,
		UnboundSource: src,
	}
}
