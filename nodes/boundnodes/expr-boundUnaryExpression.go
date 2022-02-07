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
}

func (BoundUnaryExpressionNode) NodeType() BoundType { return BoundUnaryExpression }

func (node BoundUnaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundUnaryExpressionNode")
	fmt.Printf("%s  └ Operator: %s\n", indent, node.Op.OperatorKind)
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// implement the expression node interface
func (node BoundUnaryExpressionNode) Type() symbols.TypeSymbol { return node.Op.ResultType }

func CreateBoundUnaryExpressionNode(op BoundUnaryOperator, expression BoundExpressionNode) BoundUnaryExpressionNode {
	return BoundUnaryExpressionNode{
		Op:         op,
		Expression: expression,
	}
}
