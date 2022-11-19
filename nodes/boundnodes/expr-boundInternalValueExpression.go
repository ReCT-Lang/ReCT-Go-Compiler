package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"github.com/llir/llvm/ir/value"
)

// basic global statement member
type BoundInternalValueExpressionNode struct {
	BoundExpressionNode

	Value     value.Value
	ValueType symbols.TypeSymbol
}

// implement node type from interface
func (BoundInternalValueExpressionNode) NodeType() BoundType { return BoundInternalValueExpression }

// node print function
func (node BoundInternalValueExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundInternalValueExpressionNode")
}

func (node BoundInternalValueExpressionNode) Span() print.TextSpan {
	return print.TextSpan{}
}

func (BoundInternalValueExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundInternalValueExpressionNode) Type() symbols.TypeSymbol { return node.ValueType }

// Doubt this is right
func CreateBoundInternalValueExpressionNode(val value.Value, typ symbols.TypeSymbol) BoundInternalValueExpressionNode {
	return BoundInternalValueExpressionNode{
		Value:     val,
		ValueType: typ,
	}
}
