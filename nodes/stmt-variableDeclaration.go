package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type VariableDeclarationStatementNode struct {
	StatementNode

	Keyword     lexer.Token
	TypeClause  TypeClauseNode
	Identifier  lexer.Token
	AssignToken lexer.Token
	Initializer ExpressionNode
}

// implement node type from interface
func (VariableDeclarationStatementNode) NodeType() NodeType { return VariableDeclaration }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node VariableDeclarationStatementNode) Span() print.TextSpan {
	span := node.Keyword.Span.SpanBetween(node.Identifier.Span)

	if node.Initializer != nil {
		span = span.SpanBetween(node.Initializer.Span())
	}

	return span
}

// node print function
func (node VariableDeclarationStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ VariableDeclarationStatementNode")
	fmt.Printf("%s  └ Keyword: %s\n", indent, node.Keyword.Kind)

	if !node.TypeClause.ClauseIsSet {
		fmt.Printf("%s  └ TypeClause: none\n", indent)
	} else {
		fmt.Println(indent + "  └ TypeClause: ")
		node.TypeClause.Print(indent + "    ")
	}

	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)

	if node.Initializer != nil {
		fmt.Println(indent + "  └ Initializer: ")
		node.Initializer.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateVariableDeclarationStatementNode(kw lexer.Token, typeClause TypeClauseNode, id lexer.Token, assign lexer.Token, init ExpressionNode) VariableDeclarationStatementNode {
	return VariableDeclarationStatementNode{
		Keyword:     kw,
		TypeClause:  typeClause,
		Identifier:  id,
		AssignToken: assign,
		Initializer: init,
	}
}
