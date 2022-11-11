package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type ReferenceExpressionNode struct {
	ExpressionNode

	RefKeyword lexer.Token
	Expression NameExpressionNode
}

// implement node type from interface
func (ReferenceExpressionNode) NodeType() NodeType { return ReferenceExpression }

func (node ReferenceExpressionNode) Span() print.TextSpan {
	return node.RefKeyword.Span.SpanBetween(node.Expression.Span())
}

// node print function
func (node ReferenceExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ ReferenceExpressionNode")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateReferenceExpressionNode(kw lexer.Token, expr NameExpressionNode) ReferenceExpressionNode {
	return ReferenceExpressionNode{
		RefKeyword: kw,
		Expression: expr,
	}
}
