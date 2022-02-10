package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type ElseClauseNode struct {
	SyntaxNode

	ClauseIsSet   bool
	ElseKeyword   lexer.Token
	ElseStatement StatementNode
}

// implement node type from interface
func (ElseClauseNode) NodeType() NodeType { return ElseClause }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node ElseClauseNode) Position() (int, int, int) {
	_, _, stmtLength := node.ElseStatement.Position()
	return node.ElseKeyword.Line, node.ElseKeyword.Column, len(node.ElseKeyword.Value) + stmtLength
}

// node print function
func (node ElseClauseNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ElseClauseNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.ElseKeyword.Kind)
	fmt.Println(indent + "  └ Statement: ")
	node.ElseStatement.Print(indent + "    ")

}

// "constructor" / ooga booga OOP cave man brain
func CreateElseClauseNode(kw lexer.Token, stmt StatementNode) ElseClauseNode {
	return ElseClauseNode{
		ElseKeyword:   kw,
		ElseStatement: stmt,
		ClauseIsSet:   true,
	}
}
