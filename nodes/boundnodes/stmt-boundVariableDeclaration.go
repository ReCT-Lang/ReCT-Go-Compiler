package boundnodes

//import (
//	"ReCT-Go-Compiler/print"
//	"ReCT-Go-Compiler/symbols"
//	"fmt"
//)

//type BoundVariableDeclarationStatementNode struct {
//	BoundStatementNode

//	Variable    symbols.VariableSymbol
//	Initializer BoundExpressionNode
//}

//// implement the interface
//func (BoundVariableDeclarationStatementNode) NodeType() BoundType { return BoundVariableDeclaration }
//func (node BoundVariableDeclarationStatementNode) Print(indent string) {
//	print.PrintC(print.Green, indent+"└ BlockStatementNode")
//	fmt.Println(indent + "  └ Statements: ")

//	for _, stmt := range node.Statements {
//		stmt.Print(indent + "    ")
//	}
//}

//// constructor
//func CreateBoundVariableDeclarationStatementNode(stmts []BoundStatementNode) BoundVariableDeclarationStatementNode {
//	return BoundVariableDeclarationStatementNode{
//		Statements: stmts,
//	}
//}
