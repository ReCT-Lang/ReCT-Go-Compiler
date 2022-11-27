package nodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type ThisExpressionNode struct {
	ExpressionNode

	ThisKeyword lexer.Token
}

// implement node type from interface
func (ThisExpressionNode) NodeType() NodeType { return ThisExpression }

func (node ThisExpressionNode) Span() print.TextSpan {
	return node.ThisKeyword.Span
}

// node print function
func (node ThisExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ThisExpressionNode")
}

// "constructor" / ooga booga OOP cave man brain
func CreateThisExpressionNode(kw lexer.Token) ThisExpressionNode {
	return ThisExpressionNode{
		ThisKeyword: kw,
	}
}
