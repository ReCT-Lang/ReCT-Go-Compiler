package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundThreadExpressionNode struct {
	BoundExpressionNode

	Function symbols.FunctionSymbol
}

func (BoundThreadExpressionNode) NodeType() BoundType { return BoundThreadExpression }
func (node BoundThreadExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThreadExpressionNode")
	node.Function.Print(indent)
}

func (BoundThreadExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThreadExpressionNode) Type() symbols.TypeSymbol { return builtins.Thread }

func CreateBoundThreadExpressionNode(function symbols.FunctionSymbol) BoundThreadExpressionNode {
	return BoundThreadExpressionNode{
		Function: function,
	}
}
