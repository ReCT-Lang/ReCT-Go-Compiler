package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundGotoStatementNode struct {
	BoundStatementNode

	Label     BoundLabel
	BoundSpan print.TextSpan
}

// implement the interface
func (BoundGotoStatementNode) NodeType() BoundType { return BoundGotoStatement }
func (node BoundGotoStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundGotoStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

func (node BoundGotoStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundGotoStatementNode(label BoundLabel, span print.TextSpan) BoundGotoStatementNode {
	return BoundGotoStatementNode{
		Label:     label,
		BoundSpan: span,
	}
}
