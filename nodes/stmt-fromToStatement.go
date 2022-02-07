package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// WhileStatementNode joke comments get old after awhile
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
