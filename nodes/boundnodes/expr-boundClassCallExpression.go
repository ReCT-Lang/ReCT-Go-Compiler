package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundClassCallExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Function  symbols.FunctionSymbol
	Arguments []BoundExpressionNode
}

// implement node type from interface
func (BoundClassCallExpressionNode) NodeType() BoundType { return BoundClassCallExpression }

// implement the expression node interface
func (node BoundClassCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

// node print function
func (node BoundClassCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundClassCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundClassCallExpressionNode) IsPersistent() bool { return false }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundClassCallExpressionNode(
	base BoundExpressionNode,
	callId symbols.FunctionSymbol,
	args []BoundExpressionNode,
) BoundClassCallExpressionNode {
	return BoundClassCallExpressionNode{
		Base:      base,
		Function:  callId,
		Arguments: args,
	}
}
