package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// ReturnStatementNode like: return "Yo mama"; there, get rect.
type ReturnStatementNode struct {
	StatementNode

	Keyword    lexer.Token
	Expression ExpressionNode
}

// NodeType Copy + Paste
func (ReturnStatementNode) NodeType() NodeType { return ReturnStatement }

// Print Prints beautiful stuff in console
func (node ReturnStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ ReturnStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain - Same -_-
func CreateReturnStatementNode(keyword lexer.Token, expression ExpressionNode) ReturnStatementNode {
	return ReturnStatementNode{
		Keyword:    keyword,
		Expression: expression,
	}
}
