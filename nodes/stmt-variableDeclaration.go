package nodes

import "ReCT-Go-Compiler/lexer"

// basic global statement member
type VariableDeclarationStatementNode struct {
	StatementNode

	Keyword     lexer.Token
	Identifier  lexer.Token
	AssignToken lexer.Token
	Initializer ExpressionNode
}

// implement node type from interface
func (node *VariableDeclarationStatementNode) NodeType() NodeType { return VariableDeclaration }

// "constructor" / ooga booga OOP cave man brain
func CreateVariableDeclarationStatementNode(kw lexer.Token, id lexer.Token, assign lexer.Token, init ExpressionNode) VariableDeclarationStatementNode {
	return VariableDeclarationStatementNode{
		Keyword:     kw,
		Identifier:  id,
		AssignToken: assign,
		Initializer: init,
	}
}
