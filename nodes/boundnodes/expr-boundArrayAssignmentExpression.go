package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundArrayAssignmentExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Index     BoundExpressionNode
	Value     BoundExpressionNode
	IsPointer bool

	UnboundSource nodes.SyntaxNode
}

func (BoundArrayAssignmentExpressionNode) NodeType() BoundType { return BoundArrayAssignmentExpression }

func (node BoundArrayAssignmentExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundArrayAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundArrayAssignmentExpression")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
	fmt.Println(indent + "  └ Value: ")
	node.Value.Print(indent + "    ")
}

func (BoundArrayAssignmentExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundArrayAssignmentExpressionNode) Type() symbols.TypeSymbol {
	return node.Base.Type().SubTypes[0]
}

func CreateBoundArrayAssignmentExpressionNode(base BoundExpressionNode, index BoundExpressionNode, value BoundExpressionNode, isPointer bool, src nodes.SyntaxNode) BoundArrayAssignmentExpressionNode {
	return BoundArrayAssignmentExpressionNode{
		Base:          base,
		Index:         index,
		Value:         value,
		IsPointer:     isPointer,
		UnboundSource: src,
	}
}
