package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ForStatementNode for(var i = 0; i < 10; i++) Print("Hello");
type ForStatementNode struct {
	StatementNode

	Keyword    lexer.Token
	Initaliser VariableDeclarationStatementNode
	Condition  ExpressionNode
	Updation   StatementNode
	Statement  StatementNode
}

// NodeType Copy + Paste
func (ForStatementNode) NodeType() NodeType { return ForStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
// ForStatementNode we don't do the Statement because it can be super long (i.e., a block statement)
func (node ForStatementNode) Position() (int, int, int) {
	length := len(node.Keyword.Value)
	_, _, initLength := node.Initaliser.Position()
	_, _, condLength := node.Condition.Position()
	_, _, UpdaLength := node.Updation.Position()
	length += initLength + condLength + UpdaLength

	return node.Keyword.Line, node.Keyword.Column, length
}

// Print Prints beautiful stuff in console
func (node ForStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ForStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Println(indent + "  └ Initaliser: ")
	node.Initaliser.Print(indent + "    ")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Updation: ")
	node.Updation.Print(indent + "    ")
	fmt.Println(indent + "  └ Statement: ")
	node.Statement.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateForStatementNode(keyword lexer.Token, initaliser VariableDeclarationStatementNode, condition ExpressionNode, updation ExpressionNode, statement StatementNode) ForStatementNode {
	return ForStatementNode{
		Keyword:    keyword,
		Initaliser: initaliser,
		Condition:  condition,
		Updation:   updation,
		Statement:  statement,
	}
}
