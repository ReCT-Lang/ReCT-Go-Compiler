package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundCallExpressionNode struct {
	BoundExpressionNode

	Function  symbols.FunctionSymbol
	Arguments []BoundExpressionNode
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

func (BoundCallExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundCallExpressionNode(function symbols.FunctionSymbol, args []BoundExpressionNode) BoundCallExpressionNode {
	return BoundCallExpressionNode{
		Function:  function,
		Arguments: args,
	}
}
