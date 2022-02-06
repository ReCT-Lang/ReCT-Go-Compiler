package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundVariableExpressionNode struct {
	BoundExpressionNode

	Variable symbols.VariableSymbol
	Type     symbols.TypeSymbol
}

func (BoundVariableExpressionNode) NodeType() BoundType { return BoundVariableExpression }

func (node BoundVariableExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundVariableExpressionNode")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
}

func CreateBoundVariableExpressionNode(variable symbols.VariableSymbol) BoundVariableExpressionNode {
	return BoundVariableExpressionNode{
		Variable: variable,
		Type:     variable.VarType(),
	}
}
