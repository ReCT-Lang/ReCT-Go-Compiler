package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type CallExpressionNode struct {
	ExpressionNode

	InMain bool

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

func CreateMainCallExpressionNode(id lexer.Token, args []ExpressionNode, castClause TypeClauseNode, parenthesis lexer.Token) CallExpressionNode {
	return CallExpressionNode{
		Identifier:         id,
		Arguments:          args,
		CastingType:        castClause,
		ClosingParenthesis: parenthesis,
		InMain:             true,
	}
}
