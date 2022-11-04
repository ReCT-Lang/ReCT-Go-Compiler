package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundThreadExpressionNode struct {
	BoundExpressionNode

	Function  symbols.FunctionSymbol
	BoundSpan print.TextSpan
}

func (BoundThreadExpressionNode) NodeType() BoundType { return BoundThreadExpression }
func (node BoundThreadExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThreadExpressionNode")
	node.Function.Print(indent)
}

func (node BoundThreadExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundThreadExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThreadExpressionNode) Type() symbols.TypeSymbol { return builtins.Thread }

func CreateBoundThreadExpressionNode(function symbols.FunctionSymbol, span print.TextSpan) BoundThreadExpressionNode {
	return BoundThreadExpressionNode{
		Function:  function,
		BoundSpan: span,
	}
}
