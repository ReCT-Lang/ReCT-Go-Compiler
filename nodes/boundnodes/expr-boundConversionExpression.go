package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundConversionExpressionNode struct {
	BoundExpressionNode

	Expression BoundExpressionNode
	ToType     symbols.TypeSymbol
	BoundSpan  print.TextSpan
}

func (BoundConversionExpressionNode) NodeType() BoundType { return BoundConversionExpression }

func (node BoundConversionExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundConversionExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.ToType.Print(indent + "    ")
	fmt.Println(indent + "  └ Expression: ")
	node.Expression.Print(indent + "    ")
}

func (node BoundConversionExpressionNode) Span() print.TextSpan {
	return node.BoundSpan
}

func (node BoundConversionExpressionNode) IsPersistent() bool {
	// object -> object ---> persistent, if the object is
	if node.Expression.Type().IsObject && node.ToType.IsObject {
		return node.Expression.IsPersistent()
	}

	// object -> primitive ---> not persistent, primitives dont need cleanup
	if node.Expression.Type().IsObject && !node.ToType.IsObject {
		return false
	}

	// primitive -> object ---> never persistent, objects are created and need cleanup
	// (a converted object can be made persistent by handing it to a variable for management)
	if !node.Expression.Type().IsObject && node.ToType.IsObject {
		return false
	}

	return false
}

// implement the expression node interface
func (node BoundConversionExpressionNode) Type() symbols.TypeSymbol { return node.ToType }

func CreateBoundConversionExpressionNode(_type symbols.TypeSymbol, expression BoundExpressionNode, span print.TextSpan) BoundConversionExpressionNode {
	return BoundConversionExpressionNode{
		ToType:     _type,
		Expression: expression,
		BoundSpan:  span,
	}
}
