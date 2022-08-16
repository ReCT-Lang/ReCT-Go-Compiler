package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundClassDestructionExpressionNode struct {
	BoundExpressionNode

	Base BoundExpressionNode
}

// implement node type from interface
func (BoundClassDestructionExpressionNode) NodeType() BoundType {
	return BoundClassDestructionExpression
}

// implement the expression node interface
func (node BoundClassDestructionExpressionNode) Type() symbols.TypeSymbol {
	return builtins.Int
}

// node print function
func (node BoundClassDestructionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundClassDestructionExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
}

func (BoundClassDestructionExpressionNode) IsPersistent() bool { return false }

// "constructor" / ooga booga OOP cave man brain
func CreateBoundClassDestructionExpressionNode(base BoundExpressionNode) BoundClassDestructionExpressionNode {
	return BoundClassDestructionExpressionNode{
		Base: base,
	}
}
