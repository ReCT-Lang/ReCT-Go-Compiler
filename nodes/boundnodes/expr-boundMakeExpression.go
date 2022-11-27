package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundMakeExpressionNode struct {
	BoundExpressionNode

	BaseType  symbols.ClassSymbol
	Arguments []BoundExpressionNode

	UnboundSource nodes.SyntaxNode
}

func (BoundMakeExpressionNode) NodeType() BoundType { return BoundMakeExpression }

func (node BoundMakeExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundMakeExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.BaseType.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, v := range node.Arguments {
		v.Print(indent + "   ")
	}
}

func (node BoundMakeExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundMakeExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundMakeExpressionNode) Type() symbols.TypeSymbol {
	return node.BaseType.Type
}

func CreateBoundMakeExpressionNode(baseType symbols.ClassSymbol, args []BoundExpressionNode, src nodes.SyntaxNode) BoundMakeExpressionNode {
	return BoundMakeExpressionNode{
		BaseType:      baseType,
		Arguments:     args,
		UnboundSource: src,
	}
}
