package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundFunctionExpressionNode struct {
	BoundExpressionNode

	Function  symbols.FunctionSymbol
	BoundSpan print.TextSpan
}

func (BoundFunctionExpressionNode) NodeType() BoundType { return BoundFunctionExpression }

func (node BoundFunctionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundFunctionExpressionNode")
	fmt.Println(indent + "  └ Function: ")
	node.Function.Print(indent + "    ")
}

func (node BoundFunctionExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (BoundFunctionExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundFunctionExpressionNode) Type() symbols.TypeSymbol {
	// create cool typesymbol
	subtypes := make([]symbols.TypeSymbol, 0)

	// [prm1, prm2, returnType]
	for _, parameter := range node.Function.Parameters {
		subtypes = append(subtypes, parameter.Type)
	}
	subtypes = append(subtypes, node.Function.Type)

	return symbols.CreateTypeSymbol("action", subtypes, false, false)
}

func CreateBoundFunctionExpressionNode(function symbols.FunctionSymbol, span print.TextSpan) BoundFunctionExpressionNode {
	return BoundFunctionExpressionNode{
		Function:  function,
		BoundSpan: span,
	}
}
