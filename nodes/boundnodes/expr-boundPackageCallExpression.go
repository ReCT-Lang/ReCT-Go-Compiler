package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundPackageCallExpressionNode struct {
	BoundExpressionNode

	Package   symbols.PackageSymbol
	Function  symbols.FunctionSymbol
	Arguments []BoundExpressionNode
}

func (BoundPackageCallExpressionNode) NodeType() BoundType { return BoundPackageCallExpression }
func (node BoundPackageCallExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundPackageCallExpressionNode")
	node.Package.Print(indent)
	node.Function.Print(indent)
	fmt.Println(indent + "  └ Arguments: ")
	for _, arg := range node.Arguments {
		arg.Print(indent + "    ")
	}
}

func (BoundPackageCallExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundPackageCallExpressionNode) Type() symbols.TypeSymbol { return node.Function.Type }

func CreateBoundPackageCallExpressionNode(pack symbols.PackageSymbol, function symbols.FunctionSymbol, args []BoundExpressionNode) BoundPackageCallExpressionNode {
	return BoundPackageCallExpressionNode{
		Package:   pack,
		Function:  function,
		Arguments: args,
	}
}
