package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundAssignmentExpressionNode struct {
	BoundExpressionNode

	InMain bool

	Variable      symbols.VariableSymbol
	Expression    BoundExpressionNode
	UnboundSource nodes.SyntaxNode
}

func (BoundAssignmentExpressionNode) NodeType() BoundType { return BoundAssignmentExpression }

func (node BoundAssignmentExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (node BoundAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundAssignmentExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (BoundAssignmentExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundAssignmentExpressionNode) Type() symbols.TypeSymbol { return node.Expression.Type() }

func CreateBoundAssignmentExpressionNode(variable symbols.VariableSymbol, expression BoundExpressionNode, inMain bool, src nodes.SyntaxNode) BoundAssignmentExpressionNode {
	return BoundAssignmentExpressionNode{
		Variable:      variable,
		Expression:    expression,
		InMain:        inMain,
		UnboundSource: src,
	}
}
