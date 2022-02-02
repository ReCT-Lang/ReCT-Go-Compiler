package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type BinaryExpressionNode struct {
	ExpressionNode

	Left     ExpressionNode
	Operator lexer.Token
	Right    ExpressionNode
}

// implement node type from interface
func (BinaryExpressionNode) NodeType() NodeType { return BinaryExpression }

// node print function
func (node BinaryExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BinaryExpressionNode")
	fmt.Printf("%s  └ Operator: %s\n", indent, node.Operator.Kind)
	fmt.Println(indent + "  └ Left: ")
	node.Left.Print(indent + "    ")
	fmt.Println(indent + "  └ Right: ")
	node.Right.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateBinaryExpressionNode(op lexer.Token, left ExpressionNode, right ExpressionNode) BinaryExpressionNode {
	return BinaryExpressionNode{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}
