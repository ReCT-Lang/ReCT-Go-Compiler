package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundArrayAssignmentExpressionNode struct {
	BoundExpressionNode

	Variable symbols.VariableSymbol
	Index    BoundExpressionNode
	Value    BoundExpressionNode
}

func (BoundArrayAssignmentExpressionNode) NodeType() BoundType { return BoundArrayAssignmentExpression }

func (node BoundArrayAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundArrayAssignmentExpression")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
	fmt.Println(indent + "  └ Value: ")
	node.Value.Print(indent + "    ")
}

func (BoundArrayAssignmentExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundArrayAssignmentExpressionNode) Type() symbols.TypeSymbol {
	return node.Variable.VarType().SubTypes[0]
}

func CreateBoundArrayAssignmentExpressionNode(variable symbols.VariableSymbol, index BoundExpressionNode, value BoundExpressionNode) BoundArrayAssignmentExpressionNode {
	return BoundArrayAssignmentExpressionNode{
		Variable: variable,
		Index:    index,
		Value:    value,
	}
}
