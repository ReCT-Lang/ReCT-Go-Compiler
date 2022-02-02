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
