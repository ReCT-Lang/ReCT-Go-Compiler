package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundVariableExpressionNode struct {
	BoundExpressionNode

	InMain        bool
	Variable      symbols.VariableSymbol
	UnboundSource nodes.SyntaxNode
}

func (BoundVariableExpressionNode) NodeType() BoundType { return BoundVariableExpression }

func (node BoundVariableExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundVariableExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
}

func (node BoundVariableExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundVariableExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundVariableExpressionNode) Type() symbols.TypeSymbol { return node.Variable.VarType() }

func CreateBoundVariableExpressionNode(variable symbols.VariableSymbol, inMain bool, src nodes.SyntaxNode) BoundVariableExpressionNode {
	return BoundVariableExpressionNode{
		Variable:      variable,
		InMain:        inMain,
		UnboundSource: src,
	}
}
