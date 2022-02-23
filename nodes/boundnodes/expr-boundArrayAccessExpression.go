package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundArrayAccessExpressionNode struct {
	BoundExpressionNode

	Variable symbols.VariableSymbol
	Index    BoundExpressionNode
}

func (BoundArrayAccessExpressionNode) NodeType() BoundType { return BoundArrayAccessExpression }

func (node BoundArrayAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundArrayAccessExpression")
	fmt.Println(indent + "  └ Variable: ")
	node.Variable.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
}

func (BoundArrayAccessExpressionNode) IsPersistent() bool { return true }

// implement the expression node interface
func (node BoundArrayAccessExpressionNode) Type() symbols.TypeSymbol {
	return node.Variable.VarType().SubTypes[0]
}

func CreateBoundArrayAccessExpressionNode(variable symbols.VariableSymbol, index BoundExpressionNode) BoundArrayAccessExpressionNode {
	return BoundArrayAccessExpressionNode{
		Variable: variable,
		Index:    index,
	}
}
