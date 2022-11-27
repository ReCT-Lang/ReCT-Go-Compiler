package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundClassFieldAssignmentExpressionNode struct {
	BoundExpressionNode

	Base  BoundExpressionNode
	Field symbols.VariableSymbol
	Value BoundExpressionNode

	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundClassFieldAssignmentExpressionNode) NodeType() BoundType {
	return BoundClassFieldAssignmentExpression
}

// implement the expression node interface
func (node BoundClassFieldAssignmentExpressionNode) Type() symbols.TypeSymbol {
	return node.Field.VarType()
}

func (node BoundClassFieldAssignmentExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

// node print function
func (node BoundClassFieldAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundClassFieldAssignmentExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Field: ")
	node.Field.Print(indent + "    ")
	fmt.Println(indent + "  └ Value: ")
	node.Value.Print(indent + "    ")
}

func (BoundClassFieldAssignmentExpressionNode) IsPersistent() bool { return true }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundClassFieldAssignmentExpressionNode(base BoundExpressionNode, field symbols.VariableSymbol, value BoundExpressionNode, src nodes.SyntaxNode) BoundClassFieldAssignmentExpressionNode {
	return BoundClassFieldAssignmentExpressionNode{
		Base:          base,
		Field:         field,
		Value:         value,
		UnboundSource: src,
	}
}
