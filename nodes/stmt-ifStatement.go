package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type IfStatementNode struct {
	StatementNode

	IfKeyword     lexer.Token
	Condition     ExpressionNode
	ThenStatement StatementNode
	ElseClause    ElseClauseNode
}

// implement node type from interface
func (IfStatementNode) NodeType() NodeType { return IfStatement }

// node print function
func (node IfStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ IfStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.IfKeyword.Kind)
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Statement: ")
	node.ThenStatement.Print(indent + "    ")

	if !node.ElseClause.ClauseIsSet {
		fmt.Printf("%s  └ ElseClause: none\n", indent)
	} else {
		fmt.Println(indent + "  └ ElseClause: ")
		node.ElseClause.Print(indent + "    ")
	}

}

// "constructor" / ooga booga OOP cave man brain
func CreateIfStatementNode(kw lexer.Token, cond ExpressionNode, then StatementNode, elseClause ElseClauseNode) IfStatementNode {
	return IfStatementNode{
		IfKeyword:     kw,
		Condition:     cond,
		ThenStatement: then,
		ElseClause:    elseClause,
	}
}
