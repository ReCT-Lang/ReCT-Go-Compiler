package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type DereferenceExpressionNode struct {
	ExpressionNode

	DerefKeyword lexer.Token
	Expression   ExpressionNode
}

// implement node type from interface
func (DereferenceExpressionNode) NodeType() NodeType { return DereferenceExpression }

func (node DereferenceExpressionNode) Span() print.TextSpan {
	return node.DerefKeyword.Span.SpanBetween(node.Expression.Span())
}

// node print function
func (node DereferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ DereferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateDereferenceExpressionNode(kw lexer.Token, expr ExpressionNode) DereferenceExpressionNode {
	return DereferenceExpressionNode{
		DerefKeyword: kw,
		Expression:   expr,
	}
}
