package nodes

import "ReCT-Go-Compiler/lexer"

// basic global statement member
type LiteralExpressionNode struct {
	ExpressionNode

	LiteralToken lexer.Token
	LiteralValue interface{}
}

// implement node type from interface
func (node *LiteralExpressionNode) NodeType() NodeType { return LiteralExpression }

// "constructor" / ooga booga OOP cave man brain
func CreateLiteralExpressionNode(tok lexer.Token) LiteralExpressionNode {
	return LiteralExpressionNode{
		LiteralToken: tok,
		LiteralValue: tok.Value,
	}
}
