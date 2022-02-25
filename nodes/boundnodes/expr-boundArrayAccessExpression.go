package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundArrayAccessExpressionNode struct {
	BoundExpressionNode

	Base  BoundExpressionNode
	Index BoundExpressionNode
}

func (BoundArrayAccessExpressionNode) NodeType() BoundType { return BoundArrayAccessExpression }

func (node BoundArrayAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundArrayAccessExpression")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
}

func (BoundArrayAccessExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundArrayAccessExpressionNode) Type() symbols.TypeSymbol {
	return node.Base.Type().SubTypes[0]
}

func CreateBoundArrayAccessExpressionNode(base BoundExpressionNode, index BoundExpressionNode) BoundArrayAccessExpressionNode {
	return BoundArrayAccessExpressionNode{
		Base:  base,
		Index: index,
	}
}
