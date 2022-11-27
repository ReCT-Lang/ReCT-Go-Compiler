package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundIfStatementNode struct {
	BoundStatementNode

	Condition     BoundExpressionNode
	ThenStatement BoundStatementNode
	ElseStatement BoundStatementNode

	UnboundSource nodes.SyntaxNode
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

func (node BoundIfStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundIfStatementNode(cond BoundExpressionNode, thenStmt BoundStatementNode, elseStmt BoundStatementNode, src nodes.SyntaxNode) BoundIfStatementNode {
	return BoundIfStatementNode{
		Condition:     cond,
		ThenStatement: thenStmt,
		ElseStatement: elseStmt,
		UnboundSource: src,
	}
}
