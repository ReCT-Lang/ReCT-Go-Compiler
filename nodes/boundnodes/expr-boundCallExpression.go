package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundCallExpressionNode struct {
	BoundExpressionNode

	InMain bool

	Function      symbols.FunctionSymbol
	Arguments     []BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundCallExpressionNode) NodeType() BoundType { return BoundCallExpression }

func (node BoundCallExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundCallExpressionNode")
	node.Function.Print(indent)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundCallExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundCallExpressionNode(function symbols.FunctionSymbol, args []BoundExpressionNode, inMain bool, src nodes.SyntaxNode) BoundCallExpressionNode {
	return BoundCallExpressionNode{
		Function:      function,
		Arguments:     args,
		InMain:        inMain,
		UnboundSource: src,
	}
}
