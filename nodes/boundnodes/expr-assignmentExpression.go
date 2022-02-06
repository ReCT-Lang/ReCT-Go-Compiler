package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundAssignmentExpressionNode struct {
	BoundExpressionNode

	Variable   symbols.VariableSymbol
	Expression BoundExpressionNode
	Type       symbols.TypeSymbol
}

func (BoundAssignmentExpressionNode) NodeType() BoundType { return BoundAssignmentExpression }
func (node BoundAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundAssignmentExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func CreateBoundAssignmentExpressionNode(variable symbols.VariableSymbol, expression BoundExpressionNode) BoundAssignmentExpressionNode {
	return BoundAssignmentExpressionNode{
		Variable:   variable,
		Expression: expression,
		Type:       expression.Type(),
	}
}
