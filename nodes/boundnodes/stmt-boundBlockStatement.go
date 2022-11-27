package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type BoundBlockStatementNode struct {
	BoundStatementNode

	Statements    []BoundStatementNode
	UnboundSource nodes.SyntaxNode
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

func (node BoundBlockStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundBlockStatementNode(stmts []BoundStatementNode, src nodes.SyntaxNode) BoundBlockStatementNode {
	return BoundBlockStatementNode{
		Statements:    stmts,
		UnboundSource: src,
	}
}
