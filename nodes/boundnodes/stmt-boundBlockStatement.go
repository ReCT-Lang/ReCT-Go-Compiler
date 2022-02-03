package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BoundBlockStatementNode struct {
	BoundStatementNode

	Statements []BoundStatementNode
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

// constructor
func CreateBoundBlockStatementNode(stmts []BoundStatementNode) BoundBlockStatementNode {
	return BoundBlockStatementNode{
		Statements: stmts,
	}
}
