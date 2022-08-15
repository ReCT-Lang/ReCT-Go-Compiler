package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type ClassFieldAccessExpressionNode struct {
	ExpressionNode

	Base            ExpressionNode
	FieldIdentifier lexer.Token
}

// implement node type from interface
func (ClassFieldAccessExpressionNode) NodeType() NodeType { return ClassFieldAccessExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ClassFieldAccessExpressionNode) Position() (int, int, int) {
	// im so sorry tokorv i nuked this one
	//length := len(node.Identifier.Value) + len(node.CallIdentifier.Value) + 2 // +2 for the ->
	//for _, arg := range node.Arguments {
	//	_, _, argLength := arg.Position()
	//	length += argLength + 2 // +2 for spaces and commas
	//}
	return 0, 0, 0 //node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node ClassFieldAccessExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ClassFieldAccessExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Printf("%s  └ FieldIdentifier: %s\n", indent, node.FieldIdentifier.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreateClassFieldAccessExpressionNode(base ExpressionNode, fieldId lexer.Token) ClassFieldAccessExpressionNode {
	return ClassFieldAccessExpressionNode{
		Base:            base,
		FieldIdentifier: fieldId,
	}
}
