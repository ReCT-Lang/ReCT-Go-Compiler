package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundGotoStatementNode struct {
	BoundStatementNode

	Label         BoundLabel
	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundGotoStatementNode) NodeType() BoundType { return BoundGotoStatement }
func (node BoundGotoStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundGotoStatementNode")
	fmt.Printf("%s  └ Label: %s\n", indent, node.Label)
}

func (node BoundGotoStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundGotoStatementNode(label BoundLabel, src nodes.SyntaxNode) BoundGotoStatementNode {
	return BoundGotoStatementNode{
		Label:         label,
		UnboundSource: src,
	}
}
