package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundBinaryExpressionNode struct {
	BoundExpressionNode

	Left          BoundExpressionNode
	Op            BoundBinaryOperator
	Right         BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundBinaryExpressionNode) NodeType() BoundType { return BoundBinaryExpression }

func (node BoundBinaryExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundBinaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundBinaryExpressionNode")
	fmt.Println(indent + "  └ Left: ")
	node.Left.Print(indent + "    ")
	fmt.Printf("%s  └ Operator: %s\n", indent, node.Op.OperatorKind)
	fmt.Println(indent + "  └ Right: ")
	node.Right.Print(indent + "    ")
}

func (BoundBinaryExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundBinaryExpressionNode) Type() symbols.TypeSymbol { return node.Op.ResultType }

func CreateBoundBinaryExpressionNode(left BoundExpressionNode, op BoundBinaryOperator, right BoundExpressionNode, src nodes.SyntaxNode) BoundBinaryExpressionNode {
	return BoundBinaryExpressionNode{
		Left:          left,
		Op:            op,
		Right:         right,
		UnboundSource: src,
	}
}
