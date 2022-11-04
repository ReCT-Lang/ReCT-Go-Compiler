package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundLabelStatementNode struct {
	BoundStatementNode

	Label     BoundLabel
	BoundSpan print.TextSpan
}

// implement the interface
func (BoundLabelStatementNode) NodeType() BoundType { return BoundLabelStatement }
func (node BoundLabelStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundLabelStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

func (node BoundLabelStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundLabelStatementNode(label BoundLabel, span print.TextSpan) BoundLabelStatementNode {
	return BoundLabelStatementNode{
		Label:     label,
		BoundSpan: span,
	}
}
