package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
	"os"
)

// basic global statement member
type LiteralExpressionNode struct {
	ExpressionNode

	LiteralToken lexer.Token
	LiteralValue interface{}
}

// implement node type from interface
func (LiteralExpressionNode) NodeType() NodeType { return LiteralExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node LiteralExpressionNode) Position() (int, int, int) {
	length := node.LiteralToken.Column

	switch node.LiteralValue.(type) {
	case string:
		length += len(node.LiteralValue.(string))
	case bool:
		length += len(fmt.Sprintf("%t", node.LiteralValue.(bool)))
	case int:
		length += len(fmt.Sprintf("%d", node.LiteralValue.(int)))
	case float32:
		length += len(fmt.Sprintf("%f", node.LiteralValue.(float32)))
	default:
		print.PrintC(
			print.Red,
			fmt.Sprintf("ERROR: Unknown type symbol \"%s\" debug: (BoundLiteralExpressionNode line 40)", node.LiteralToken.Value),
		)
		os.Exit(1) // shrug
	}

	return node.LiteralToken.Line, node.LiteralToken.Column, length
}

// node print function
func (node LiteralExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ LiteralExpressionNode")
	fmt.Printf("%s  └ Value: %s\n", indent, node.LiteralToken.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreateLiteralExpressionNode(tok lexer.Token) LiteralExpressionNode {
	return LiteralExpressionNode{
		LiteralToken: tok,
		LiteralValue: tok.RealValue,
	}
}
