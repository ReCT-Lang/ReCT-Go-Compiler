package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type BoundFunctionExpressionNode struct {
	BoundExpressionNode

	InClass       symbols.ClassSymbol
	Function      symbols.FunctionSymbol
	UnboundSource nodes.SyntaxNode
}

func (BoundFunctionExpressionNode) NodeType() BoundType { return BoundFunctionExpression }

func (node BoundFunctionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundFunctionExpressionNode")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
}

func (node BoundFunctionExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundFunctionExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundFunctionExpressionNode) Type() symbols.TypeSymbol {
	// create cool typesymbol
	subtypes := make([]symbols.TypeSymbol, 0)

	// if this function is inside a class -> add the obj as the first parameter (bc thats how it be)
	if node.InClass.Exists {
		subtypes = append(subtypes, node.InClass.Type)
	}

	// [prm1, prm2, returnType]
	for _, parameter := range node.Function.Parameters {
		subtypes = append(subtypes, parameter.Type)
	}
	subtypes = append(subtypes, node.Function.Type)

	return symbols.CreateTypeSymbol("action", subtypes, false, false, false, symbols.PackageSymbol{}, nil)
}

func CreateBoundFunctionExpressionNode(function symbols.FunctionSymbol, src nodes.SyntaxNode) BoundFunctionExpressionNode {
	return BoundFunctionExpressionNode{
		Function:      function,
		UnboundSource: src,
	}
}

func CreateBoundFunctionInClassExpressionNode(function symbols.FunctionSymbol, class symbols.ClassSymbol, src nodes.SyntaxNode) BoundFunctionExpressionNode {
	return BoundFunctionExpressionNode{
		Function:      function,
		InClass:       class,
		UnboundSource: src,
	}
}
