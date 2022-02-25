package nodes

import (
	"ReCT-Go-Compiler/print"
	"fmt"
)

type ArrayAccessExpressionNode struct {
	ExpressionNode

	Base  ExpressionNode
	Index ExpressionNode
}

// implement node type from interface
func (ArrayAccessExpressionNode) NodeType() NodeType { return ArrayAccessExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ArrayAccessExpressionNode) Position() (int, int, int) {
	//length := len(node.Identifier.Value) + 2 // +2 for spaces and stuff
	//_, _, exprLength := node.Index.Position()
	//length += exprLength
	return 0, 0, 0 //node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node ArrayAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ArrayAccessExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Println(indent + "  └ Index: ")
	node.Index.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateArrayAccessExpressionNode(base ExpressionNode, index ExpressionNode) ArrayAccessExpressionNode {
	return ArrayAccessExpressionNode{
		Base:  base,
		Index: index,
	}
}
