package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type TypeCallExpressionNode struct {
	ExpressionNode

	Base           ExpressionNode
	CallIdentifier lexer.Token
	Arguments      []ExpressionNode
}

// implement node type from interface
func (TypeCallExpressionNode) NodeType() NodeType { return TypeCallExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node TypeCallExpressionNode) Position() (int, int, int) {
	// im so sorry tokorv i nuked this one
	//length := len(node.Identifier.Value) + len(node.CallIdentifier.Value) + 2 // +2 for the ->
	//for _, arg := range node.Arguments {
	//	_, _, argLength := arg.Position()
	//	length += argLength + 2 // +2 for spaces and commas
	//}
	return 0, 0, 0 //node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node TypeCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ TypeCallExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Printf("%s  └ CallIdentifier: %s\n", indent, node.CallIdentifier.Value)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateTypeCallExpressionNode(base ExpressionNode, callId lexer.Token, args []ExpressionNode) TypeCallExpressionNode {
	return TypeCallExpressionNode{
		Base:           base,
		CallIdentifier: callId,
		Arguments:      args,
	}
}
