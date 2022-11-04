package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundVariableExpressionNode struct {
	BoundExpressionNode

	Variable  symbols.VariableSymbol
	BoundSpan print.TextSpan
}

func (BoundVariableExpressionNode) NodeType() BoundType { return BoundVariableExpression }

func (node BoundVariableExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundVariableExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
}

func (node BoundVariableExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundVariableExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundVariableExpressionNode) Type() symbols.TypeSymbol { return node.Variable.VarType() }

func CreateBoundVariableExpressionNode(variable symbols.VariableSymbol, span print.TextSpan) BoundVariableExpressionNode {
	return BoundVariableExpressionNode{
		Variable:  variable,
		BoundSpan: span,
	}
}
