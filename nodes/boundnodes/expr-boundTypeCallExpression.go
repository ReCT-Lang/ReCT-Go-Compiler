package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundTypeCallExpressionNode struct {
	BoundExpressionNode

	Variable  symbols.VariableSymbol
	Function  symbols.TypeFunctionSymbol
	Arguments []BoundExpressionNode
}

// implement node type from interface
func (BoundTypeCallExpressionNode) NodeType() BoundType { return BoundTypeCallExpression }

// node print function
func (node BoundTypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundTypeCallExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent)
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreatBoundTypeCallExpressionNode(
	id symbols.VariableSymbol,
	callId symbols.TypeFunctionSymbol,
	args []BoundExpressionNode,
) BoundTypeCallExpressionNode {
	return BoundTypeCallExpressionNode{
		Variable:  id,
		Function:  callId,
		Arguments: args,
	}
}
