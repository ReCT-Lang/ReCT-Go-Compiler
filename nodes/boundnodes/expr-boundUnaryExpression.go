package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundUnaryExpressionNode struct {
	BoundExpressionNode

	Op         BoundUnaryOperator
	Expression BoundExpressionNode
	Type       symbols.TypeSymbol
}

func (BoundUnaryExpressionNode) NodeType() BoundType { return BoundUnaryExpression }

func (node BoundUnaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundUnaryExpressionNode")
	fmt.Printf("%s  └ Operator: %d", indent, node.Op.OperatorKind)
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func CreateBoundUnaryExpressionNode(op BoundUnaryOperator, expression BoundExpressionNode) BoundUnaryExpressionNode {
	return BoundUnaryExpressionNode{
		Op:         op,
		Expression: expression,
	}
}
