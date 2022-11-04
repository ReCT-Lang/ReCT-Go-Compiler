package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundClassFieldAccessExpressionNode struct {
	BoundExpressionNode

	Base      BoundExpressionNode
	Field     symbols.VariableSymbol
	BoundSpan print.TextSpan
}

// implement node type from interface
func (BoundClassFieldAccessExpressionNode) NodeType() BoundType {
	return BoundClassFieldAccessExpression
}

// implement the expression node interface
func (node BoundClassFieldAccessExpressionNode) Type() symbols.TypeSymbol {
	return node.Field.VarType()
}

func (node BoundClassFieldAccessExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

// node print function
func (node BoundClassFieldAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundClassFieldAccessExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Field: ")
	node.Field.Print(indent + "    ")
}

func (BoundClassFieldAccessExpressionNode) IsPersistent() bool { return true }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundClassFieldAccessExpressionNode(base BoundExpressionNode, field symbols.VariableSymbol, span print.TextSpan) BoundClassFieldAccessExpressionNode {
	return BoundClassFieldAccessExpressionNode{
		Base:      base,
		Field:     field,
		BoundSpan: span,
	}
}
