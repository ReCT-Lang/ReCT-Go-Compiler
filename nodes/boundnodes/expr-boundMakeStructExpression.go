package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundMakeStructExpressionNode struct {
	BoundExpressionNode

	StructType symbols.TypeSymbol
	Literals   []BoundExpressionNode

	BoundSpan print.TextSpan
}

func (BoundMakeStructExpressionNode) NodeType() BoundType { return BoundMakeStructExpression }

func (node BoundMakeStructExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundMakeStructExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.StructType.Print(indent + "    ")
	//fmt.Println(indent + "  └ Length: ")
	//node.Length.Print(indent + "    ")
}

func (node BoundMakeStructExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundMakeStructExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundMakeStructExpressionNode) Type() symbols.TypeSymbol {
	return node.StructType
}

func CreateBoundMakeStructExpressionNode(structType symbols.TypeSymbol, literals []BoundExpressionNode, span print.TextSpan) BoundMakeStructExpressionNode {
	return BoundMakeStructExpressionNode{
		StructType: structType,
		Literals:   literals,
		BoundSpan:  span,
	}
}
