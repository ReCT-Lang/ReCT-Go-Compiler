package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ThreadStatementNode I was not told of this before I tried to implement it
type ThreadExpressionNode struct {
	StatementNode

	Keyword    lexer.Token
	Expression NameExpressionNode
}

// NodeType Copy + Paste
func (ThreadExpressionNode) NodeType() NodeType { return ThreadExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient. (this sucks gonna change it one day)
func (node ThreadExpressionNode) Position() (int, int, int) {
	return 0, 0, 0 // l8r
}

// Print Prints beautiful stuff in console
func (node ThreadExpressionNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ThreadExpressionNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)

	if node.Expression.ExpressionNode == nil {
		fmt.Printf("%s  └ NameExpression: none\n", indent)
	} else {
		fmt.Println(indent + "  └ NameExpression: ")
		node.Expression.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateThreadExpressionNode(keyword lexer.Token, expression NameExpressionNode) ThreadExpressionNode {
	return ThreadExpressionNode{
		Keyword:    keyword,
		Expression: expression,
	}
}
