package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// FromToStatementNode joke comments get old after awhile
type FromToStatementNode struct {
	StatementNode

	Keyword    lexer.Token
	Identifier lexer.Token
	LowerBound ExpressionNode
	UpperBound ExpressionNode
	Statement  StatementNode
}

// NodeType Copy + Paste again
func (FromToStatementNode) NodeType() NodeType { return FromToStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
// FromToStatementNode we don't do the Statement because it can be super long (i.e., a block statement)
func (node FromToStatementNode) Span() print.TextSpan {
	return node.Keyword.Span.SpanBetween(node.Statement.Span())
}

// Print Prints beautiful stuff in console
func (node FromToStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ FromToStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)
	fmt.Println(indent + "  └ Lower Bound: ")
	node.LowerBound.Print(indent + "    ")
	fmt.Println(indent + "  └ Upper Bound: ")
	node.UpperBound.Print(indent + "    ")
	fmt.Println(indent + "  └ Statement: ")
	node.Statement.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateFromToStatementNode(keyword lexer.Token, id lexer.Token, lower ExpressionNode, upper ExpressionNode, statement StatementNode) FromToStatementNode {
	return FromToStatementNode{
		Keyword:    keyword,
		Identifier: id,
		LowerBound: lower,
		UpperBound: upper,
		Statement:  statement,
	}
}
