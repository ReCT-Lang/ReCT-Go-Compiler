package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"os"
)

// basic global statement member
type BoundLiteralExpressionNode struct {
	BoundExpressionNode

	Value interface{}
	Type  symbols.TypeSymbol
}

// implement node type from interface
func (BoundLiteralExpressionNode) NodeType() BoundType { return BoundLiteralExpression }

// node print function
func (node BoundLiteralExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundLiteralExpressionNode")
	fmt.Printf("%s  └ Value: %s\n", indent, node.Value.(string))
	fmt.Println(indent + "  └ Type: ")
	node.Type.Print(indent + "    ")
}

// Doubt this is right
func CreateBoundLiteralExpressionNode(value interface{}) BoundLiteralExpressionNode {
	var _type symbols.TypeSymbol
	switch value.(type) {
	case string:
		_type = builtins.String
	case bool:
		_type = builtins.Bool
	case int:
		_type = builtins.Int
	default:
		print.PrintC(
			print.Red,
			fmt.Sprintf("ERROR: Uknown type symbol \"%s\" debug: (BoundLiteralExpressionNode line 40)", value.(string)),
		)
		os.Exit(1) // shrug
	}
	return BoundLiteralExpressionNode{
		Value: value,
		Type:  _type,
	}
}
