package nodes

// very cool interface for creating syntax nodes
type SyntaxNode interface {
	NodeType() NodeType
	// only type atm, might contain more stuff like text-location later
}

// very cool interface for creating statements
// (ik this isnt very revolutionary but i just like to organise stuff)
type StatementNode interface {
	SyntaxNode
}

// very cool interface for creating members
// (again, organisation.)
type MemberNode interface {
	SyntaxNode
}

// cool node type Enum straight up stolen from ReCT v1.0
type NodeType int

const (
	// non-commented stuff has been implemented already

	// Members
	// -------
	GlobalStatement NodeType = iota

	// Statements
	// ----------
	BlockStatement
	VariableDeclaration
	IfStatement
	ElseClause

	/*
			 * // Members
			 * // -------
			 * FunctionDeclaration
			 *
			 * // General nodes (that dont qualify as statement)
			 * // ----------------------------------------------
			 * Parameter
			 * TypeClause
			 *
			 * // Statements
			 * // ----------
			 * WhileStatement
			 * ForStatement
			 * BreakStatement
			 * ContinueStatement
			 * ReturnStatement
			 * ExpressionStatement
			 *
			 * // Expressions
			 * // -----------
			 * LiteralExpression
		     * NameExpression
		     * UnaryExpression
		     * BinaryExpression
		     * ParenthesizedExpression
		     * AssignmentExpression
		     * CallExpression
		     * EndKeyword
		     * EditVariableToken
		     * SingleEditVariableToken
		     * FromKeyword
		     * FromToStatement
	*/
)
