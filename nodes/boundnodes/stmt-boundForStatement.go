package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundForStatementNode struct {
	BoundLoopStatementNode

	Variable  BoundVariableDeclarationStatementNode
	Condition BoundExpressionNode
	Action    BoundStatementNode

	Body          BoundStatementNode
	BreakLabel    BoundLabel
	ContinueLabel BoundLabel

	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundForStatementNode) NodeType() BoundType { return BoundForStatement }
func (node BoundForStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundForStatementNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Action: ")
	node.Action.Print(indent + "    ")
	fmt.Println(indent + "  └ Body: ")
	node.Body.Print(indent + "    ")

	fmt.Printf("%s  └ BreakLabel: %s\n", indent, node.BreakLabel)
	fmt.Printf("%s  └ ContinueLabel: %s\n", indent, node.ContinueLabel)
}

func (node BoundForStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundForStatementNode) LoopBreakLabel() BoundLabel    { return node.BreakLabel }
func (node BoundForStatementNode) LoopContinueLabel() BoundLabel { return node.ContinueLabel }

// constructor
func CreateBoundForStatementNode(variable BoundVariableDeclarationStatementNode, cond BoundExpressionNode, action BoundStatementNode, body BoundStatementNode, breakLabel BoundLabel, continueLabel BoundLabel, src nodes.SyntaxNode) BoundForStatementNode {
	return BoundForStatementNode{
		Variable:      variable,
		Condition:     cond,
		Action:        action,
		Body:          body,
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
		UnboundSource: src,
	}
}
