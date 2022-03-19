package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundVariableDeclarationStatementNode struct {
	BoundStatementNode

	Variable    symbols.VariableSymbol
	Initializer BoundExpressionNode
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

// constructor
func CreateBoundVariableDeclarationStatementNode(variable symbols.VariableSymbol, init BoundExpressionNode) BoundVariableDeclarationStatementNode {
	return BoundVariableDeclarationStatementNode{
		Variable:    variable,
		Initializer: init,
	}
}
