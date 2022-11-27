package boundnodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
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

func (node BoundInternalValueExpressionNode) Source() nodes.SyntaxNode {
	return nil
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
