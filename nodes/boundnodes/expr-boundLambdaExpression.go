package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundLambdaExpressionNode struct {
	BoundExpressionNode

	Function      symbols.FunctionSymbol
	Body          BoundBlockStatementNode
	UnboundSource nodes.SyntaxNode
}

func (BoundLambdaExpressionNode) NodeType() BoundType { return BoundLambdaExpression }

func (node BoundLambdaExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundLambdaExpressionNode")
	fmt.Println(indent + "  └ Symbol: ")
	node.Function.Print(indent + "    ")
}

func (node BoundLambdaExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundLambdaExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundLambdaExpressionNode) Type() symbols.TypeSymbol {
	// create cool typesymbol
	subtypes := make([]symbols.TypeSymbol, 0)

	// [prm1, prm2, returnType]
	for _, parameter := range node.Function.Parameters {
		subtypes = append(subtypes, parameter.Type)
	}
	subtypes = append(subtypes, node.Function.Type)

	return symbols.CreateTypeSymbol("action", subtypes, false, false, false, symbols.PackageSymbol{}, nil)
}

func CreateBoundLambdaExpressionNode(function symbols.FunctionSymbol, body BoundBlockStatementNode, src nodes.SyntaxNode) BoundLambdaExpressionNode {
	return BoundLambdaExpressionNode{
		Function:      function,
		Body:          body,
		UnboundSource: src,
	}
}
