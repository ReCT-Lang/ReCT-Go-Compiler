package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundFunctionExpressionNode struct {
	BoundExpressionNode

	Function  symbols.FunctionSymbol
	BoundSpan print.TextSpan
}

func (BoundFunctionExpressionNode) NodeType() BoundType { return BoundFunctionExpression }

func (node BoundFunctionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundFunctionExpressionNode")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
}

func (node BoundFunctionExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundFunctionExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundFunctionExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundFunctionExpressionNode(function symbols.FunctionSymbol, span print.TextSpan) BoundFunctionExpressionNode {
	return BoundFunctionExpressionNode{
		Function:  function,
		BoundSpan: span,
	}
}
