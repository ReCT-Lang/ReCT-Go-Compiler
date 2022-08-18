package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

type MakeExpressionNode struct {
	ExpressionNode

	Package   *lexer.Token
	BaseType  lexer.Token
	Arguments []ExpressionNode
}

// implement node type from interface
func (MakeExpressionNode) NodeType() NodeType { return MakeExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node MakeExpressionNode) Position() (int, int, int) {
	// TODO: aaaaAAAAAA (im sorry tokorv i cant deal with this rn lmao)
	return 0, 0, 0
}

// node print function
func (node MakeExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ MakeExpressionNode")
	fmt.Println(indent + "  └ Type: " + node.BaseType.Value)
	fmt.Println(indent + "  └ Arguments: ")
	for _, v := range node.Arguments {
		v.Print(indent + "    ")
	}
}

// "constructor" / ooga booga OOP cave man brain
func CreateMakeExpressionNode(pack *lexer.Token, typ lexer.Token, args []ExpressionNode) MakeExpressionNode {
	return MakeExpressionNode{
		Package:   pack,
		BaseType:  typ,
		Arguments: args,
	}
}
