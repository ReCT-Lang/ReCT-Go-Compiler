package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// WhileStatementNode joke comments get old after awhile
type FromToStatementNode struct {
	StatementNode

	Keyword     lexer.Token
	Initialiser ExpressionNode // Assignment Expression
	Condition   ExpressionNode
	Statement   StatementNode
}

// NodeType Copy + Paste again
func (FromToStatementNode) NodeType() NodeType { return WhileStatement }

// Print Prints beautiful stuff in console
func (node FromToStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ FromToStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Println(indent + "  └ Initialiser: ")
	node.Initialiser.Print(indent + "    ")
	fmt.Println(indent + "  └ Condition: ")
	node.Condition.Print(indent + "    ")
	fmt.Println(indent + "  └ Statement: ")
	node.Statement.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateFromToStatementNode(keyword lexer.Token, initialiser, condition ExpressionNode, statement StatementNode) FromToStatementNode {
	return FromToStatementNode{
		Keyword:     keyword,
		Initialiser: initialiser,
		Condition:   condition,
		Statement:   statement,
	}
}
