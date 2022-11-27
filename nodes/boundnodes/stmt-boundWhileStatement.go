package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type BoundWhileStatementNode struct {
	BoundLoopStatementNode

	Condition     BoundExpressionNode
	Body          BoundStatementNode
	BreakLabel    BoundLabel
	ContinueLabel BoundLabel

	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundWhileStatementNode) NodeType() BoundType { return BoundWhileStatement }
func (node BoundWhileStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundWhileStatementNode")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Body: ")
	node.Body.Print(indent + "    ")

	fmt.Printf("%s  └ BreakLabel: %s\n", indent, node.BreakLabel)
	fmt.Printf("%s  └ ContinueLabel: %s\n", indent, node.ContinueLabel)
}

func (node BoundWhileStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundWhileStatementNode) LoopBreakLabel() BoundLabel    { return node.BreakLabel }
func (node BoundWhileStatementNode) LoopContinueLabel() BoundLabel { return node.ContinueLabel }

// constructor
func CreateBoundWhileStatementNode(cond BoundExpressionNode, body BoundStatementNode, breakLabel BoundLabel, continueLabel BoundLabel, src nodes.SyntaxNode) BoundWhileStatementNode {
	return BoundWhileStatementNode{
		Condition:     cond,
		Body:          body,
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
		UnboundSource: src,
	}
}
