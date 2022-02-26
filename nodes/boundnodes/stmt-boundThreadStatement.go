package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type BoundThreadStatementNode struct {
	BoundExpressionNode

	Function symbols.FunctionSymbol
}

func (BoundThreadStatementNode) NodeType() BoundType { return BoundThreadStatement }
func (node BoundThreadStatementNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundThreadStatementNode")
	node.Function.Print(indent)
}

func (BoundThreadStatementNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundThreadStatementNode) Type() symbols.TypeSymbol { return builtins.Action }

func CreateBoundThreadStatementNode(function symbols.FunctionSymbol) BoundThreadStatementNode {
	return BoundThreadStatementNode{
		Function: function,
	}
}
