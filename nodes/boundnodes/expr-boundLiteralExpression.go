package boundnodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/builtins"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
	"os"
)

// basic global statement member
type BoundLiteralExpressionNode struct {
	BoundExpressionNode

	Value         interface{}
	LiteralType   symbols.TypeSymbol
	UnboundSource nodes.SyntaxNode
}

// implement node type from interface
func (BoundLiteralExpressionNode) NodeType() BoundType { return BoundLiteralExpression }

// node print function
func (node BoundLiteralExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundLiteralExpressionNode")

	if node.LiteralType.Fingerprint() == builtins.String.Fingerprint() {
		fmt.Printf("%s  └ Value: %s\n", indent, node.Value.(string))

	} else if node.LiteralType.Fingerprint() == builtins.Int.Fingerprint() {
		fmt.Printf("%s  └ Value: %d\n", indent, node.Value.(int))

	} else if node.LiteralType.Fingerprint() == builtins.Bool.Fingerprint() {
		fmt.Printf("%s  └ Value: %t\n", indent, node.Value.(bool))
	}
	fmt.Println(indent + "  └ Type: ")
	node.LiteralType.Print(indent + "    ")
}

func (node BoundLiteralExpressionNode) Source() nodes.SyntaxNode {
	return node.UnboundSource
}

func (BoundLiteralExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundLiteralExpressionNode) Type() symbols.TypeSymbol { return node.LiteralType }

// Doubt this is right
func CreateBoundLiteralExpressionNode(expr nodes.LiteralExpressionNode, src nodes.SyntaxNode) BoundLiteralExpressionNode {
	var _type symbols.TypeSymbol
	switch expr.LiteralValue.(type) {
	case string:
		if expr.IsNative {
			_type = symbols.CreateTypeSymbol("pointer", []symbols.TypeSymbol{builtins.Byte}, false, false, false, symbols.PackageSymbol{}, nil)
		} else {
			_type = builtins.String
		}
	case bool:
		_type = builtins.Bool
	case int, int32:
		_type = builtins.Int
	case float32:
		_type = builtins.Float
	default:
		print.PrintC(
			print.Red,
			fmt.Sprintf("ERROR: Uknown type symbol \"%s\" debug: (BoundLiteralExpressionNode line 40)", expr.LiteralValue.(string)),
		)
		os.Exit(1) // shrug
	}
	return BoundLiteralExpressionNode{
		Value:         expr.LiteralValue,
		LiteralType:   _type,
		UnboundSource: src,
	}
}

func CreateBoundLiteralExpressionNodeFromValue(value interface{}, src nodes.SyntaxNode) BoundLiteralExpressionNode {
	var _type symbols.TypeSymbol
	switch value.(type) {
	case string:
		_type = builtins.String
	case bool:
		_type = builtins.Bool
	case int, int32:
		_type = builtins.Int
	case int64:
		_type = builtins.Long
	case byte:
		_type = builtins.Byte
	case float32:
		_type = builtins.Float
	default:
		print.PrintC(
			print.Red,
			fmt.Sprintf("ERROR: Uknown type symbol \"%s\" debug: (BoundLiteralExpressionNode line 40)", value.(string)),
		)
		os.Exit(1) // shrug
	}
	return BoundLiteralExpressionNode{
		Value:         value,
		LiteralType:   _type,
		UnboundSource: src,
	}
}
