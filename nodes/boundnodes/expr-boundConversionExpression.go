package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundConversionExpressionNode struct {
	BoundExpressionNode

	Expression BoundExpressionNode
	ToType     symbols.TypeSymbol
}

func (BoundConversionExpressionNode) NodeType() BoundType { return BoundConversionExpression }

func (node BoundConversionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundConversionExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.ToType.Print(indent + "    ")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// implement the expression node interface
func (node BoundConversionExpressionNode) Type() symbols.TypeSymbol { return node.ToType }

func CreateBoundConversionExpressionNode(_type symbols.TypeSymbol, expression BoundExpressionNode) BoundConversionExpressionNode {
	return BoundConversionExpressionNode{
		ToType:     _type,
		Expression: expression,
	}
}
