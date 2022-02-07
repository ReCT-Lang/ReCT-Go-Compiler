package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundGotoStatementNode struct {
	BoundStatementNode

	Label BoundLabel
}

// implement the interface
func (BoundGotoStatementNode) NodeType() BoundType { return BoundGotoStatement }
func (node BoundGotoStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundGotoStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

// constructor
func CreateBoundGotoStatementNode(label BoundLabel) BoundGotoStatementNode {
	return BoundGotoStatementNode{
		Label: label,
	}
}
