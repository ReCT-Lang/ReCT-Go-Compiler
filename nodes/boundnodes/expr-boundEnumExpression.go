package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

// basic global statement member
type BoundEnumExpressionNode struct {
	BoundExpressionNode

	Value     int
	Enum      symbols.EnumSymbol
	BoundSpan print.TextSpan
}

// implement node type from interface
func (BoundEnumExpressionNode) NodeType() BoundType { return BoundEnumExpression }

// node print function
func (node BoundEnumExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundEnumExpressionNode")
}

func (node BoundEnumExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundEnumExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundEnumExpressionNode) Type() symbols.TypeSymbol { return node.Enum.Type }

// Doubt this is right
func CreateBoundEnumExpressionNode(val int, enm symbols.EnumSymbol, span print.TextSpan) BoundEnumExpressionNode {
	return BoundEnumExpressionNode{
		Value:     val,
		Enum:      enm,
		BoundSpan: span,
	}
}
