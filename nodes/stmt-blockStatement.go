package nodes

import "ReCT-Go-Compiler/lexer"

// basic global statement member
type BlockStatementNode struct {
	StatementNode

	OpenBrace  lexer.Token
	Statments  []StatementNode
	CloseBrace lexer.Token
}

// implement node type from interface
func (node *BlockStatementNode) NodeType() NodeType { return BlockStatement }

// "constructor" / ooga booga OOP cave man brain
func CreateBlockStatementNode(openBrace lexer.Token, statements []StatementNode, closeBrace lexer.Token) BlockStatementNode {
	return BlockStatementNode{
		OpenBrace:  openBrace,
		Statments:  statements,
		CloseBrace: closeBrace,
	}
}
