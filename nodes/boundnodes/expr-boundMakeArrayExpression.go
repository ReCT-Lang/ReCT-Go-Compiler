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
}

func (BoundMakeArrayExpressionNode) NodeType() BoundType { return BoundMakeArrayExpression }

func (node BoundMakeArrayExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundMakeArrayExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.BaseType.Print(indent + "    ")
	//fmt.Println(indent + "  └ Length: ")
	//node.Length.Print(indent + "    ")
}

func (BoundMakeArrayExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundMakeArrayExpressionNode) Type() symbols.TypeSymbol {
	return symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{node.BaseType}, true)
}

func CreateBoundMakeArrayExpressionNode(baseType symbols.TypeSymbol, length BoundExpressionNode) BoundMakeArrayExpressionNode {
	return BoundMakeArrayExpressionNode{
		BaseType:  baseType,
		Length:    length,
		IsLiteral: false,
	}
}

func CreateBoundMakeArrayExpressionNodeLiteral(baseType symbols.TypeSymbol, literals []BoundExpressionNode) BoundMakeArrayExpressionNode {
	return BoundMakeArrayExpressionNode{
		BaseType:  baseType,
		Literals:  literals,
		IsLiteral: true,
	}
}
