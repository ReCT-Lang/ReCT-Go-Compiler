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
func (TypeCallExpressionNode) NodeType() NodeType { return CallExpression }

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
