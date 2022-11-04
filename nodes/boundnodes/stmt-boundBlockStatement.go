package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundBlockStatementNode struct {
	BoundStatementNode

	Statements []BoundStatementNode
	BoundSpan  print.TextSpan
}

// implement the interface
func (BoundBlockStatementNode) NodeType() BoundType { return BoundBlockStatement }
func (node BoundBlockStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundBlockStatementNode")
	fmt.Println(indent + "  └ Statements: ")

	for _, stmt := range node.Statements {
		stmt.Print(indent + "    ")
	}
}

func (node BoundBlockStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundBlockStatementNode(stmts []BoundStatementNode, span print.TextSpan) BoundBlockStatementNode {
	return BoundBlockStatementNode{
		Statements: stmts,
		BoundSpan:  span,
	}
}
