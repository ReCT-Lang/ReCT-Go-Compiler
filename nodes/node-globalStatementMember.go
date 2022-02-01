package nodes

import "ReCT-Go-Compiler/print"

// basic global statement member
type GlobalStatementMember struct {
	MemberNode
	Statement StatementNode
}

// implement node type from interface
func (node GlobalStatementMember) NodeType() NodeType { return GlobalStatement }

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
