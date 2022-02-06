package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundConversionExpressionNode struct {
	BoundExpressionNode

	Expression BoundExpressionNode
	Type       symbols.TypeSymbol
}

func (BoundConversionExpressionNode) NodeType() BoundType { return BoundConversionExpression }

func (node BoundConversionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundConversionExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.Type.Print(indent + "    ")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func CreateBoundConversionExpressionNode(_type symbols.TypeSymbol, expression BoundExpressionNode) BoundConversionExpressionNode {
	return BoundConversionExpressionNode{
		Type:       _type,
		Expression: expression,
	}
}
