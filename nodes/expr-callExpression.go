package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type CallExpressionNode struct {
	ExpressionNode

	Identifier lexer.Token
	Arguments  []ExpressionNode
}

// implement node type from interface
func (CallExpressionNode) NodeType() NodeType { return CallExpression }

func (node CallExpressionNode) Position() (int, int, int) {
	length := len(node.Identifier.Value) + 2
	for _, arg := range node.Arguments {
		_, _, argLength := arg.Position()
		length += argLength + 2
	}
	return node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node CallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ CallExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)

	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateCallExpressionNode(id lexer.Token, args []ExpressionNode) CallExpressionNode {
	return CallExpressionNode{
		Identifier: id,
		Arguments:  args,
	}
}
