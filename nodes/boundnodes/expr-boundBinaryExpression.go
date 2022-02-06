package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundBinaryExpressionNode struct {
	BoundExpressionNode

	Left  BoundExpressionNode
	Op    BoundBinaryOperator
	Right BoundExpressionNode
	Type  symbols.TypeSymbol
}

func (BoundBinaryExpressionNode) NodeType() BoundType { return BoundBinaryExpression }

func (node BoundBinaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundBinaryExpressionNode")
	fmt.Println(indent + "  └ Left: ")
	node.Left.Print(indent + "    ")
	fmt.Printf("%s  └ Operator: %d", indent, node.Op.OperatorKind)
	fmt.Println(indent + "  └ Right: ")
	node.Right.Print(indent + "    ")
}

func CreateBoundBinaryExpressionNode(left BoundExpressionNode, op BoundBinaryOperator, right BoundExpressionNode) BoundBinaryExpressionNode {
	return BoundBinaryExpressionNode{
		Left:  left,
		Op:    op,
		Right: right,
	}
}
