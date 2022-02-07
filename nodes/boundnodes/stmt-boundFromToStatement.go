package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundFromToStatementNode struct {
	BoundLoopStatementNode

	Variable      symbols.VariableSymbol
	LowerBound    BoundExpressionNode
	UpperBound    BoundExpressionNode
	Body          BoundStatementNode
	BreakLabel    BoundLabel
	ContinueLabel BoundLabel
}

// implement the interface
func (BoundFromToStatementNode) NodeType() BoundType { return BoundFromToStatement }
func (node BoundFromToStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundFromToStatementNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ LowerBound: ")
	node.LowerBound.Print(indent + "    ")
	fmt.Println(indent + "  └ UpperBound: ")
	node.UpperBound.Print(indent + "    ")
	fmt.Println(indent + "  └ Body: ")
	node.Body.Print(indent + "    ")

	fmt.Printf("%s  └ BreakLabel: %s\n", indent, node.BreakLabel)
	fmt.Printf("%s  └ ContinueLabel: %s\n", indent, node.ContinueLabel)
}

func (node BoundFromToStatementNode) LoopBreakLabel() BoundLabel    { return node.BreakLabel }
func (node BoundFromToStatementNode) LoopContinueLabel() BoundLabel { return node.ContinueLabel }

// constructor
func CreateBoundFromToStatementNode(variable symbols.VariableSymbol, lower BoundExpressionNode, upper BoundExpressionNode, body BoundStatementNode, breakLabel BoundLabel, continueLabel BoundLabel) BoundFromToStatementNode {
	return BoundFromToStatementNode{
		Variable:      variable,
		LowerBound:    lower,
		UpperBound:    upper,
		Body:          body,
		BreakLabel:    breakLabel,
		ContinueLabel: continueLabel,
	}
}
