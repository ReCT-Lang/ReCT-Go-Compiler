package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundGarbageCollectionStatementNode struct {
	BoundStatementNode

	Variables []symbols.VariableSymbol
	BoundSpan print.TextSpan
}

// implement the interface
func (BoundGarbageCollectionStatementNode) NodeType() BoundType {
	return BoundGarbageCollectionStatement
}
func (node BoundGarbageCollectionStatementNode) Print(indent string) {
	print.PrintC(print.Green, indent+"└ BoundGarbageCollectionStatementNode")
	fmt.Println(indent + "  └ Variables: ")

	for _, variable := range node.Variables {
		variable.Print(indent + "    ")
	}
}

func (node BoundGarbageCollectionStatementNode) Span() print.TextSpan {
	return node.BoundSpan
}

// constructor
func CreateBoundGarbageCollectionStatementNode(variables []symbols.VariableSymbol, span print.TextSpan) BoundGarbageCollectionStatementNode {
	return BoundGarbageCollectionStatementNode{
		Variables: variables,
		BoundSpan: span,
	}
}
