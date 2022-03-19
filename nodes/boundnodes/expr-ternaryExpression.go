package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundTernaryExpressionNode struct {
	BoundExpressionNode

	Condition BoundExpressionNode
	If        BoundExpressionNode
	Else      BoundExpressionNode
}

func (BoundTernaryExpressionNode) NodeType() BoundType { return BoundTernaryExpression }

func (node BoundTernaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundTernaryExpressionNode")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ If: ")
	node.If.Print(indent + "    ")
	fmt.Println(indent + "  └ Else: ")
	node.Else.Print(indent + "    ")
}

func (node BoundTernaryExpressionNode) IsPersistent() bool {
	return node.If.IsPersistent() || node.Else.IsPersistent()
}

// implement the expression node interface
func (node BoundTernaryExpressionNode) Type() symbols.TypeSymbol { return node.If.Type() }

func CreateBoundTernaryExpressionNode(cond BoundExpressionNode, left BoundExpressionNode, right BoundExpressionNode) BoundTernaryExpressionNode {
	return BoundTernaryExpressionNode{
		Condition: cond,
		If:        left,
		Else:      right,
	}
}
