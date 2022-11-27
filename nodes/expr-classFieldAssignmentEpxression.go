package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type ClassFieldAssignmentExpressionNode struct {
	ExpressionNode

	Base            ExpressionNode
	FieldIdentifier lexer.Token
	Value           ExpressionNode
}

// implement node type from interface
func (ClassFieldAssignmentExpressionNode) NodeType() NodeType { return ClassFieldAssignmentExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ClassFieldAssignmentExpressionNode) Span() print.TextSpan {
	return node.Base.Span().SpanBetween(node.Value.Span())
}

// node print function
func (node ClassFieldAssignmentExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ClassFieldAssignmentExpressionNode")
	fmt.Println(indent + "  └ Base: ")
	node.Base.Print(indent + "    ")
	fmt.Printf("%s  └ FieldIdentifier: %s\n", indent, node.FieldIdentifier.Value)
	fmt.Println(indent + "  └ Value: ")
	node.Value.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateClassFieldAssignmentExpressionNode(base ExpressionNode, fieldId lexer.Token, value ExpressionNode) ClassFieldAssignmentExpressionNode {
	return ClassFieldAssignmentExpressionNode{
		Base:            base,
		FieldIdentifier: fieldId,
		Value:           value,
	}
}
