package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// WhileStatementNode joke comments get old after awhile
type WhileStatementNode struct {
	StatementNode

	Keyword   lexer.Token
	Condition ExpressionNode
	Statement StatementNode
}

// NodeType Copy + Paste again
func (WhileStatementNode) NodeType() NodeType { return WhileStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
// We don't process the statement for WhileStatementNode because it do be chonky
func (node WhileStatementNode) Span() print.TextSpan {
	return node.Keyword.Span.SpanBetween(node.Statement.Span())
}

// Print Prints beautiful stuff in console
func (node WhileStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ WhileStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Statement: ")
	node.Statement.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateWhileStatementNode(keyword lexer.Token, condition ExpressionNode, statement StatementNode) WhileStatementNode {
	return WhileStatementNode{
		Keyword:   keyword,
		Condition: condition,
		Statement: statement,
	}
}
