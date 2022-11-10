package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type DereferenceExpressionNode struct {
	ExpressionNode

	RefKeyword lexer.Token
	Expression ExpressionNode
}

// implement node type from interface
func (DereferenceExpressionNode) NodeType() NodeType { return ReferenceExpression }

func (node DereferenceExpressionNode) Span() print.TextSpan {
	return node.RefKeyword.Span.SpanBetween(node.Expression.Span())
}

// node print function
func (node DereferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func GetDereferenceExpressionNode(kw lexer.Token, expr ExpressionNode) DereferenceExpressionNode {
	return DereferenceExpressionNode{
		RefKeyword: kw,
		Expression: expr,
	}
}
