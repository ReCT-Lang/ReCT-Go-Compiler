package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundClassCallExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Function  symbols.FunctionSymbol
	Arguments []BoundExpressionNode

	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundClassCallExpressionNode) NodeType() BoundType { return BoundClassCallExpression }

// implement the expression node interface
func (node BoundClassCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func (node BoundClassCallExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// node print function
func (node BoundClassCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundClassCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundClassCallExpressionNode) IsPersistent() bool { return false }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundClassCallExpressionNode(
	base BoundExpressionNode,
	callId symbols.FunctionSymbol,
	args []BoundExpressionNode,
	src nodes.SyntaxNode,
) BoundClassCallExpressionNode {
	return BoundClassCallExpressionNode{
		Base:          base,
		Function:      callId,
		Arguments:     args,
		UnboundSource: src,
	}
}
