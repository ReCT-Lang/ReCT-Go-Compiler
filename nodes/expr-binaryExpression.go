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

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node BinaryExpressionNode) Position() (int, int, int) {
	// If you come across and error to do with the operator or something you can still access its position
	// via Operator.Line and Operator.Column
	line, column, leftLength := node.Left.Position()
	_, _, rightLength := node.Right.Position()
	return line, column, leftLength + rightLength + 2 // 2 for the operator btw
}

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
