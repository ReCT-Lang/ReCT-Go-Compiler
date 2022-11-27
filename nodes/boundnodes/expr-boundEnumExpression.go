package boundnodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

// basic global statement member
type BoundEnumExpressionNode struct {
	BoundExpressionNode

	Value         int
	Enum          symbols.EnumSymbol
	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundEnumExpressionNode) NodeType() BoundType { return BoundEnumExpression }

// node print function
func (node BoundEnumExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundEnumExpressionNode")
}

func (node BoundEnumExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundEnumExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundEnumExpressionNode) Type() symbols.TypeSymbol { return node.Enum.Type }

// Doubt this is right
func CreateBoundEnumExpressionNode(val int, enm symbols.EnumSymbol, src nodes.SyntaxNode) BoundEnumExpressionNode {
	return BoundEnumExpressionNode{
		Value:         val,
		Enum:          enm,
		UnboundSource: src,
	}
}
