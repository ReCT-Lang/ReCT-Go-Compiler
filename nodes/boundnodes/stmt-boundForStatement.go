package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundForStatementNode struct {
	BoundLoopStatementNode

	Variable  symbols.VariableSymbol
	Condition BoundExpressionNode
	Action    BoundStatementNode

	Body          BoundStatementNode
	BreakLabel    BoundLabel
	ContinueLabel BoundLabel
}

// implement the interface
func (BoundForStatementNode) NodeType() BoundType { return BoundWhileStatement }
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

func (node BoundForStatementNode) LoopBreakLabel() BoundLabel    { return node.BreakLabel }
func (node BoundForStatementNode) LoopContinueLabel() BoundLabel { return node.ContinueLabel }

// constructor
func CreateBoundForStatementNode(variable symbols.VariableSymbol, cond BoundExpressionNode, action BoundStatementNode, body BoundStatementNode, breakLabel BoundLabel, continueLabel BoundLabel) BoundForStatementNode {
	return BoundForStatementNode{
		Variable:      variable,
		Condition:     cond,
		Action:        action,
		Body:          body,
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
	}
}
