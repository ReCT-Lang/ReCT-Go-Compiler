package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundTernaryExpressionNode struct {
	BoundExpressionNode

	Condition BoundExpressionNode
	If        BoundExpressionNode
	Else      BoundExpressionNode
	Tmp       symbols.LocalVariableSymbol

	IfLabel   BoundLabel
	ElseLabel BoundLabel
	EndLabel  BoundLabel

	UnboundSource nodes.SyntaxNode
}

func (BoundTernaryExpressionNode) NodeType() BoundType { return BoundTernaryExpression }

func (node BoundTernaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundTernaryExpressionNode")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ If: ")
	node.If.Print(indent + "    ")
	fmt.Println(indent + "  └ Else: ")
	node.Else.Print(indent + "    ")
}

func (node BoundTernaryExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundTernaryExpressionNode) IsPersistent() bool {
	return node.If.IsPersistent() || node.Else.IsPersistent()
}

// implement the expression node interface
func (node BoundTernaryExpressionNode) Type() symbols.TypeSymbol { return node.If.Type() }

func CreateBoundTernaryExpressionNode(cond BoundExpressionNode, left BoundExpressionNode, right BoundExpressionNode, tmp symbols.LocalVariableSymbol, src nodes.SyntaxNode) BoundTernaryExpressionNode {
	return BoundTernaryExpressionNode{
		Condition:     cond,
		If:            left,
		Else:          right,
		Tmp:           tmp,
		UnboundSource: src,
	}
}
