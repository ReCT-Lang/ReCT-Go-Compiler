package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type LambdaExpressionNode struct {
	ExpressionNode

	LambdaKeyword lexer.Token
	Parameters    []ParameterNode
	TypeClause    TypeClauseNode
	Body          BlockStatementNode
}

// implement node type from interface
func (LambdaExpressionNode) NodeType() NodeType { return LambdaExpression }

func (node LambdaExpressionNode) Span() print.TextSpan {
	span := node.LambdaKeyword.Span.SpanBetween(node.Body.Span())
	return span
}

// node print function
func (node LambdaExpressionNode) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- LambdaExpressionNode")

	fmt.Println(indent + "  └ Parameters: ")
	for _, param := range node.Parameters {
		param.Print(indent + "    ")
	}

	if !node.TypeClause.ClauseIsSet {
		fmt.Printf("%s  └ TypeClause: none\n", indent)
	} else {
		fmt.Println(indent + "  └ TypeClause: ")
		node.TypeClause.Print(indent + "    ")
	}

	fmt.Println(indent + "  └ Body: ")
	node.Body.Print(indent + "    ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateLambdaExpressionNode(kw lexer.Token, params []ParameterNode, typeClause TypeClauseNode, body BlockStatementNode) LambdaExpressionNode {
	return LambdaExpressionNode{
		LambdaKeyword: kw,
		Parameters:    params,
		TypeClause:    typeClause,
		Body:          body,
	}
}
