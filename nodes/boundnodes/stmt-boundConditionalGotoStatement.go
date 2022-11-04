package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundConditionalGotoStatementNode struct {
	BoundStatementNode

	Condition BoundExpressionNode
	IfLabel   BoundLabel
	ElseLabel BoundLabel

	BoundSpan print.TextSpan
}

// implement the interface
func (BoundConditionalGotoStatementNode) NodeType() BoundType { return BoundConditionalGotoStatement }
func (node BoundConditionalGotoStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundConditionalGotoStatementNode")
	fmt.Println(indent + "  └ Condition:")
	node.Condition.Print(indent + "    ")
	fmt.Printf("%s  └ IfLabel: %s\n", indent, node.IfLabel)
	fmt.Printf("%s  └ ElseLabel: %s\n", indent, node.ElseLabel)
}

func (node BoundConditionalGotoStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundConditionalGotoStatementNode(condition BoundExpressionNode, ifLabel BoundLabel, elseLabel BoundLabel, span print.TextSpan) BoundConditionalGotoStatementNode {
	return BoundConditionalGotoStatementNode{
		Condition: condition,
		IfLabel:   ifLabel,
		ElseLabel: elseLabel,
		BoundSpan: span,
	}
}
