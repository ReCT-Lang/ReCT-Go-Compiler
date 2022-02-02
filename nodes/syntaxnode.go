package nodes

// very cool interface for creating syntax nodes
type SyntaxNode interface {
	NodeType() NodeType
	Print(indent string)
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

// very cool interface for creating expressions
// (craaaaazy ikr)
type ExpressionNode interface {
	SyntaxNode
}

// cool node type Enum straight up stolen from ReCT v1.0
type NodeType int

const (
	// commented stuff has yet to be implemented
	// i am basing these objects off of the rect 1.0 source
	// -> https://github.com/RedCubeDev-ByteSpace/ReCT/tree/834776cbf0ad97da0e6441835f1bc19d903f115b/ReCT/CodeAnalysis/Syntax

	// Members
	// -------
	GlobalStatement NodeType = iota

	// Statements
	// ----------
	BlockStatement
	VariableDeclaration
	IfStatement
	ElseClause
	ReturnStatement
	ForStatement
	WhileStatement
	BreakStatement
	ContinueStatement
	FromToStatement
	ExpressionStatement

	// Expressions
	// -----------
	LiteralExpression
	ParenthesisedExpression
	NameExpression
	AssignmentExpression
	CallExpression

	// TODO (Parsing)
	UnaryExpression
	BinaryExpression

	/*
			 * // Members
			 * // -------
			 * FunctionDeclaration
			 *
			 * // General nodes (that dont qualify as statements)
			 * // ----------------------------------------------
			 * Parameter
			 * TypeClause
			 *
			 * // Statements
			 * // -----------
			 *    all done!
			 *
			 * // Expressions
			 * // -----------
		     * UnaryExpression
			 * BinaryExpression
	*/
)
