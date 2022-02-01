package nodes

import "ReCT-Go-Compiler/lexer"

// basic global statement member
type ElseClauseNode struct {
	SyntaxNode

	ElseKeyword   lexer.Token
	ElseStatement StatementNode
}

// implement node type from interface
func (node *ElseClauseNode) NodeType() NodeType { return ElseClause }

// "constructor" / ooga booga OOP cave man brain
func CreateElseClauseNode(kw lexer.Token, stmt StatementNode) ElseClauseNode {
	return ElseClauseNode{
		ElseKeyword:   kw,
		ElseStatement: stmt,
	}
}
