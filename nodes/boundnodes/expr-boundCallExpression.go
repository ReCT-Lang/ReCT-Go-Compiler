package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundCallExpressionNode struct {
	BoundExpressionNode

	Function  symbols.FunctionSymbol
	Arguments []symbols.ParameterSymbol
}

func (BoundCallExpressionNode) NodeType() BoundType { return BoundCallExpression }
func (node BoundCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundCallExpressionNode")
	node.Function.Print(indent)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func CreateCallExpressionNode(function symbols.FunctionSymbol, args []symbols.ParameterSymbol) BoundCallExpressionNode {
	return BoundCallExpressionNode{
		Function:  function,
		Arguments: args,
	}
}
