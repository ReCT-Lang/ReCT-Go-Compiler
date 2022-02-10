package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type VariableEditorExpressionNode struct {
	ExpressionNode

	Identifier   lexer.Token
	Operator     lexer.Token
	Expression   ExpressionNode
	IsSingleStep bool // things like ++ or --
}

// implement node type from interface
func (VariableEditorExpressionNode) NodeType() NodeType { return VariableEditorExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node VariableEditorExpressionNode) Position() (int, int, int) {
	length := len(node.Operator.Value) + len(node.Operator.Value) + 2 // +2 for spaces and stuff
	_, _, exprLength := node.Expression.Position()
	length += exprLength
	return node.Identifier.Line, node.Identifier.Column, length
}

// node print function
func (node VariableEditorExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ VariableEditorExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
	fmt.Printf("%s  └ Operator: %s\n", indent, node.Operator.Kind)
	fmt.Printf("%s  └ IsSingleStep: %t\n", indent, node.IsSingleStep)

	if node.Expression != nil {
		fmt.Println(indent + "  └ Expression: ")
		node.Expression.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateVariableEditorExpressionNode(id lexer.Token, op lexer.Token, expr ExpressionNode, singleStep bool) VariableEditorExpressionNode {
	return VariableEditorExpressionNode{
		Identifier:   id,
		Operator:     op,
		Expression:   expr,
		IsSingleStep: singleStep,
	}
}
