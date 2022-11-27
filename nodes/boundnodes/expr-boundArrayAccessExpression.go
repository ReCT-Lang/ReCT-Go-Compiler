package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundArrayAccessExpressionNode struct {
	BoundExpressionNode

	Base          BoundExpressionNode
	Index         BoundExpressionNode
	IsPointer     bool
	UnboundSource nodes.SyntaxNode
}

func (BoundArrayAccessExpressionNode) NodeType() BoundType { return BoundArrayAccessExpression }

func (node BoundArrayAccessExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundArrayAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundArrayAccessExpression")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
}

func (BoundArrayAccessExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundArrayAccessExpressionNode) Type() symbols.TypeSymbol {
	return node.Base.Type().SubTypes[0]
}

func CreateBoundArrayAccessExpressionNode(base BoundExpressionNode, index BoundExpressionNode, isPointer bool, src nodes.SyntaxNode) BoundArrayAccessExpressionNode {
	return BoundArrayAccessExpressionNode{
		Base:          base,
		Index:         index,
		IsPointer:     isPointer,
		UnboundSource: src,
	}
}
