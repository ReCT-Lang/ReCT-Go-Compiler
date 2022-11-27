package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundTypeCallExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Function  symbols.TypeFunctionSymbol
	Arguments []BoundExpressionNode

	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundTypeCallExpressionNode) NodeType() BoundType { return BoundTypeCallExpression }

// implement the expression node interface
func (node BoundTypeCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func (node BoundTypeCallExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// node print function
func (node BoundTypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundTypeCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundTypeCallExpressionNode) IsPersistent() bool { return false }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundTypeCallExpressionNode(
	base BoundExpressionNode,
	callId symbols.TypeFunctionSymbol,
	args []BoundExpressionNode,
	src nodes.SyntaxNode,
) BoundTypeCallExpressionNode {
	return BoundTypeCallExpressionNode{
		Base:          base,
		Function:      callId,
		Arguments:     args,
		UnboundSource: src,
	}
}
