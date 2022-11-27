package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundClassFieldAccessExpressionNode struct {
	BoundExpressionNode

	Base          BoundExpressionNode
	Field         symbols.VariableSymbol
	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundClassFieldAccessExpressionNode) NodeType() BoundType {
	return BoundClassFieldAccessExpression
}

// implement the expression node interface
func (node BoundClassFieldAccessExpressionNode) Type() symbols.TypeSymbol {
	return node.Field.VarType()
}

func (node BoundClassFieldAccessExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
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
func CreateBoundClassFieldAccessExpressionNode(base BoundExpressionNode, field symbols.VariableSymbol, src nodes.SyntaxNode) BoundClassFieldAccessExpressionNode {
	return BoundClassFieldAccessExpressionNode{
		Base:          base,
		Field:         field,
		UnboundSource: src,
	}
}
