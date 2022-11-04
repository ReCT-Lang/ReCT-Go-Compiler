package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type PackageCallExpressionNode struct {
	ExpressionNode

	Package    lexer.Token
	Identifier lexer.Token
	Arguments  []ExpressionNode

	ClosingToken lexer.Token
}

// implement node type from interface
func (PackageCallExpressionNode) NodeType() NodeType { return PackageCallExpression }

func (node PackageCallExpressionNode) Span() print.TextSpan {
	return node.Package.Span.SpanBetween(node.ClosingToken.Span)
}

// node print function
func (node PackageCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ PackageCallExpressionNode")
	fmt.Printf("%s  └ Package: %s\n", indent, node.Identifier.Value)
	fmt.Printf("%s  └ Identifier: %s\n", indent, node.Identifier.Value)

	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreatePackageCallExpressionNode(pck lexer.Token, id lexer.Token, args []ExpressionNode, closing lexer.Token) PackageCallExpressionNode {
	return PackageCallExpressionNode{
		Package:      pck,
		Identifier:   id,
		Arguments:    args,
		ClosingToken: closing,
	}
}
