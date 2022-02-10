package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type TypeCallExpressionNode struct {
	ExpressionNode

	Identifier     lexer.Token
	CallIdentifier lexer.Token
	Arguments      []ExpressionNode
}

// implement node type from interface
func (TypeCallExpressionNode) NodeType() NodeType { return TypeCallExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node TypeCallExpressionNode) Position() (int, int, int) {
	length := len(node.Identifier.Value) + len(node.CallIdentifier.Value) + 2 // +2 for the ->
	for _, arg := range node.Arguments {
		_, _, argLength := arg.Position()
		length += argLength + 2 // +2 for spaces and commas
	}
	return node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node TypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ TypeCallExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
	fmt.Printf("%s  └ CallIdentifier: %s\n", indent, node.CallIdentifier.Value)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateTypeCallExpressionNode(id, callId lexer.Token, args []ExpressionNode) TypeCallExpressionNode {
	return TypeCallExpressionNode{
		Identifier:     id,
		CallIdentifier: callId,
		Arguments:      args,
	}
}
