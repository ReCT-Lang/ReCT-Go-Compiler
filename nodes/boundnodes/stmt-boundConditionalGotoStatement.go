package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundConditionalGotoStatementNode struct {
	BoundStatementNode

	Condition  BoundExpressionNode
	Label      BoundLabel
	JumpIfTrue bool
}

// implement the interface
func (BoundConditionalGotoStatementNode) NodeType() BoundType { return BoundConditionalGotoStatement }
func (node BoundConditionalGotoStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundConditionalGotoStatementNode")
	fmt.Println(indent + "  └ Condition:")
	node.Condition.Print(indent + "    ")
	fmt.Printf("%s  └ Label: %s", indent, node.Label)
	fmt.Printf("%s  └ JumpIfTrue: %t", indent, node.JumpIfTrue)
}

// constructor
func CreateBoundConditionalGotoStatementNode(condition BoundExpressionNode, label BoundLabel, jumpIfTrue bool) BoundConditionalGotoStatementNode {
	return BoundConditionalGotoStatementNode{
		Condition:  condition,
		Label:      label,
		JumpIfTrue: jumpIfTrue,
	}
}
