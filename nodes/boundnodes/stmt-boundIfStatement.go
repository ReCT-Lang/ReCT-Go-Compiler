package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundIfStatementNode struct {
	BoundStatementNode

	Condition     BoundExpressionNode
	ThenStatement BoundStatementNode
	ElseStatement BoundStatementNode
}

// implement the interface
func (BoundIfStatementNode) NodeType() BoundType { return BoundIfStatement }
func (node BoundIfStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundIfStatementNode")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ ThenStatement: ")
	node.ThenStatement.Print(indent + "    ")

	if node.ElseStatement != nil {
		fmt.Println(indent + "  └ ElseStatement: ")
		node.ThenStatement.Print(indent + "    ")
	} else {
		fmt.Println(indent + "  └ ElseStatement: none")
	}
}

// constructor
func CreateBoundIfStatementNode(cond BoundExpressionNode, thenStmt BoundStatementNode, elseStmt BoundStatementNode) BoundIfStatementNode {
	return BoundIfStatementNode{
		Condition:     cond,
		ThenStatement: thenStmt,
		ElseStatement: elseStmt,
	}
}
