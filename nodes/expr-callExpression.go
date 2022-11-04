package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type CallExpressionNode struct {
	ExpressionNode

	Identifier lexer.Token
	Arguments  []ExpressionNode

	CastingType TypeClauseNode // if this call is actually a complex cast

	ClosingParenthesis lexer.Token
}

// implement node type from interface
func (CallExpressionNode) NodeType() NodeType { return CallExpression }

func (node CallExpressionNode) Span() print.TextSpan {
	if node.CastingType.ClauseIsSet {
		return node.CastingType.Span().SpanBetween(node.ClosingParenthesis.Span)
	} else {
		return node.Identifier.Span.SpanBetween(node.ClosingParenthesis.Span)
	}
}

// node print function
func (node CallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ CallExpressionNode")
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)

	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateCallExpressionNode(id lexer.Token, args []ExpressionNode, castClause TypeClauseNode, parenthesis lexer.Token) CallExpressionNode {
	return CallExpressionNode{
		Identifier:         id,
		Arguments:          args,
		CastingType:        castClause,
		ClosingParenthesis: parenthesis,
	}
}
