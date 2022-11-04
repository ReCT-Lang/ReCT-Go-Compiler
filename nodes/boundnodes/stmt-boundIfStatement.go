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

	BoundSpan print.TextSpan
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

func (node BoundIfStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundIfStatementNode(cond BoundExpressionNode, thenStmt BoundStatementNode, elseStmt BoundStatementNode, span print.TextSpan) BoundIfStatementNode {
	return BoundIfStatementNode{
		Condition:     cond,
		ThenStatement: thenStmt,
		ElseStatement: elseStmt,
		BoundSpan:     span,
	}
}
