package nodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type MakeArrayExpressionNode struct {
	ExpressionNode

	IsLiteral     bool
	Type          TypeClauseNode
	Length        ExpressionNode
	LiteralValues []ExpressionNode
}

// implement node type from interface
func (MakeArrayExpressionNode) NodeType() NodeType { return MakeArrayExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node MakeArrayExpressionNode) Position() (int, int, int) {
	// TODO: aaaaAAAAAA (im sorry tokorv i cant deal with this rn lmao)
	return 0, 0, 0
}

// node print function
func (node MakeArrayExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ MakeArrayExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.Type.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateMakeArrayExpressionNode(typ TypeClauseNode, length ExpressionNode) MakeArrayExpressionNode {
	return MakeArrayExpressionNode{
		Type:      typ,
		Length:    length,
		IsLiteral: false,
	}
}
func CreateMakeArrayExpressionNodeLiteral(typ TypeClauseNode, literals []ExpressionNode) MakeArrayExpressionNode {
	return MakeArrayExpressionNode{
		Type:          typ,
		LiteralValues: literals,
		IsLiteral:     true,
	}
}
