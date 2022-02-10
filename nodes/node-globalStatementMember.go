package nodes

import (
	"ReCT-Go-Compiler/print"
)

// basic global statement member
type GlobalStatementMember struct {
	MemberNode
	Statement StatementNode
}

// implement node type from interface
func (GlobalStatementMember) NodeType() NodeType { return GlobalStatement }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node GlobalStatementMember) Position() (int, int, int) {
	return node.Statement.Position()
}

// node print function
func (node GlobalStatementMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- GlobalStatementMember")
	node.Statement.Print(indent + "  ")
}

// "constructor" / ooga booga OOP cave man brain
func CreateGlobalStatementMember(stmt StatementNode) GlobalStatementMember {
	return GlobalStatementMember{
		Statement: stmt,
	}
}
