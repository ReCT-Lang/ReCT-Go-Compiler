package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundMakeArrayExpressionNode struct {
	BoundExpressionNode

	IsLiteral bool
	BaseType  symbols.TypeSymbol
	Length    BoundExpressionNode
	Literals  []BoundExpressionNode

	BoundSpan print.TextSpan
}

func (BoundMakeArrayExpressionNode) NodeType() BoundType { return BoundMakeArrayExpression }

func (node BoundMakeArrayExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundMakeArrayExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.BaseType.Print(indent + "    ")
	//fmt.Println(indent + "  └ Length: ")
	//node.Length.Print(indent + "    ")
}

func (node BoundMakeArrayExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundMakeArrayExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundMakeArrayExpressionNode) Type() symbols.TypeSymbol {
	return symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{node.BaseType}, true, false, false)
}

func CreateBoundMakeArrayExpressionNode(baseType symbols.TypeSymbol, length BoundExpressionNode) BoundMakeArrayExpressionNode {
	return BoundMakeArrayExpressionNode{
		BaseType:  baseType,
		Length:    length,
		IsLiteral: false,
	}
}

func CreateBoundMakeArrayExpressionNodeLiteral(baseType symbols.TypeSymbol, literals []BoundExpressionNode, span print.TextSpan) BoundMakeArrayExpressionNode {
	return BoundMakeArrayExpressionNode{
		BaseType:  baseType,
		Literals:  literals,
		IsLiteral: true,
		BoundSpan: span,
	}
}
