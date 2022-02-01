package nodes

import "ReCT-Go-Compiler/lexer"

// basic global statement member
type IfStatementNode struct {
	StatementNode

	IfKeyword     lexer.Token
	Condition     ExpressionNode
	ThenStatement StatementNode
	ElseClause    ElseClauseNode
}

// implement node type from interface
func (node *IfStatementNode) NodeType() NodeType { return IfStatement }

// "constructor" / ooga booga OOP cave man brain
func CreateIfStatementNode(kw lexer.Token, cond ExpressionNode, then StatementNode, elseClause ElseClauseNode) IfStatementNode {
	return IfStatementNode{
		IfKeyword:     kw,
		Condition:     cond,
		ThenStatement: then,
		ElseClause:    elseClause,
	}
}
