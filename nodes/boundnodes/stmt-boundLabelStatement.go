package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundLabelStatementNode struct {
	BoundStatementNode

	Label BoundLabel
}

// implement the interface
func (BoundLabelStatementNode) NodeType() BoundType { return BoundLabelStatement }
func (node BoundLabelStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundLabelStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

// constructor
func CreateBoundLabelStatementNode(label BoundLabel) BoundLabelStatementNode {
	return BoundLabelStatementNode{
		Label: label,
	}
}
