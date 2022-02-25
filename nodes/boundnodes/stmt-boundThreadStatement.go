package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundThreadStatementNode struct {
	BoundStatementNode

	Function symbols.FunctionSymbol
}

func (BoundThreadStatementNode) NodeType() BoundType { return BoundThreadStatement }
func (node BoundThreadStatementNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThreadStatementNode")
	node.Function.Print(indent)
}

func (BoundThreadStatementNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThreadStatementNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundThreadStatementNode(function symbols.FunctionSymbol) BoundThreadStatementNode {
	return BoundThreadStatementNode{
		Function: function,
	}
}
