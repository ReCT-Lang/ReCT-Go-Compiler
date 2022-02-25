package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundFunctionExpressionNode struct {
	BoundExpressionNode

	Function symbols.FunctionSymbol
}

func (BoundFunctionExpressionNode) NodeType() BoundType { return BoundFunctionExpression }

func (node BoundFunctionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundFunctionExpressionNode")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
}

func (BoundFunctionExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundFunctionExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundFunctionExpressionNode(function symbols.FunctionSymbol) BoundFunctionExpressionNode {
	return BoundFunctionExpressionNode{
		Function: function,
	}
}
