package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundTypeCallExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Function  symbols.TypeFunctionSymbol
	Arguments []BoundExpressionNode

	BoundSpan print.TextSpan
}

// implement node type from interface
func (BoundTypeCallExpressionNode) NodeType() BoundType { return BoundTypeCallExpression }

// implement the expression node interface
func (node BoundTypeCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func (node BoundTypeCallExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

// node print function
func (node BoundTypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundTypeCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundTypeCallExpressionNode) IsPersistent() bool { return false }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundTypeCallExpressionNode(
	base BoundExpressionNode,
	callId symbols.TypeFunctionSymbol,
	args []BoundExpressionNode,
	span print.TextSpan,
) BoundTypeCallExpressionNode {
	return BoundTypeCallExpressionNode{
		Base:      base,
		Function:  callId,
		Arguments: args,
		BoundSpan: span,
	}
}
