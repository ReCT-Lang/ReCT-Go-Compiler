package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundVariableDeclarationStatementNode struct {
	BoundStatementNode

	Variable    symbols.VariableSymbol
	Initializer BoundExpressionNode

	UnboundSource nodes.SyntaxNode
}

// implement the interface
func (BoundVariableDeclarationStatementNode) NodeType() BoundType { return BoundVariableDeclaration }
func (node BoundVariableDeclarationStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundVariableDeclarationStatementNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Initializer: ")
	if node.Initializer != nil {
		node.Initializer.Print(indent + "    ")
	} else {
		fmt.Println(indent + "      none")
	}
}

func (node BoundVariableDeclarationStatementNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// constructor
func CreateBoundVariableDeclarationStatementNode(variable symbols.VariableSymbol, init BoundExpressionNode, src nodes.SyntaxNode) BoundVariableDeclarationStatementNode {
	return BoundVariableDeclarationStatementNode{
		Variable:      variable,
		Initializer:   init,
		UnboundSource: src,
	}
}
