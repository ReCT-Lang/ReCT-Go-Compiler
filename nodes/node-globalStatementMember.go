package nodes

// basic global statement member
type GlobalStatementMember struct {
	MemberNode
	Statement StatementNode
}

// implement node type from interface
func (node *GlobalStatementMember) NodeType() NodeType { return GlobalStatement }

// "constructor" / ooga booga OOP cave man brain
func CreateGlobalStatementMember(stmt StatementNode) GlobalStatementMember {
	return GlobalStatementMember{
		Statement: stmt,
	}
}
