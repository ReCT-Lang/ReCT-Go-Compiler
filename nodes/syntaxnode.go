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
type NodeType string

const (
	// i am basing these objects off of the rect 1.0 source
	// -> https://github.com/RedCubeDev-ByteSpace/ReCT/tree/834776cbf0ad97da0e6441835f1bc19d903f115b/ReCT/CodeAnalysis/Syntax

	// Members
	// -------
	GlobalStatement     NodeType = "Global Statement"
	FunctionDeclaration NodeType = "Function Declaration"

	// General
	// -------
	Parameter  NodeType = "Parameter"
	TypeClause NodeType = "Type Clause"

	// Statements
	// ----------
	BlockStatement      NodeType = "Block Statement"
	VariableDeclaration NodeType = "Variable Declaration"
	IfStatement         NodeType = "If Statement"
	ElseClause          NodeType = "Else Clause"
	ReturnStatement     NodeType = "Return Statement"
	ForStatement        NodeType = "For Statement"
	WhileStatement      NodeType = "While Statement"
	BreakStatement      NodeType = "Break Statement"
	ContinueStatement   NodeType = "Continue Statement"
	FromToStatement     NodeType = "FromTo Statement"
	ExpressionStatement NodeType = "Expression Statement"

	// Expressions
	// -----------
	LiteralExpression       NodeType = "Literal Expression"
	ParenthesisedExpression NodeType = "Parenthesised Expression"
	NameExpression          NodeType = "Name Expression"
	AssignmentExpression    NodeType = "Assignment Expression"
	CallExpression          NodeType = "Call Expression"
	UnaryExpression         NodeType = "Unary Expression"
	BinaryExpression        NodeType = "Binary Expression"
	TypeCallExpression      NodeType = "TypeCall Expression"
)
