package boundnodes

import "ReCT-Go-Compiler/symbols"

// incredibly cool interface for creating bound nodes
type BoundNode interface {
	NodeType() BoundType
	Print(indent string)
}

type BoundStatementNode interface {
	BoundNode
}

type BoundLoopStatementNode interface {
	BoundStatementNode
	LoopBreakLabel() BoundLabel
	LoopContinueLabel() BoundLabel
}

type BoundLabel string

type BoundExpressionNode interface {
	BoundNode
	Type() symbols.TypeSymbol
}

// enum for all our node types
type BoundType int

const (
	// based off of rect 1.0 source
	// -> https://github.com/RedCubeDev-ByteSpace/ReCT/tree/834776cbf0ad97da0e6441835f1bc19d903f115b/ReCT/CodeAnalysis/Binding

	// Statements
	BoundBlockStatement BoundType = iota
	BoundVariableDeclaration
	BoundIfStatement
	BoundWhileStatement
	BoundForStatement
	BoundFromToStatement
	BoundLabelStatement
	BoundGotoStatement
	BoundConditionalGotoStatement
	BoundReturnStatement
	BoundExpressionStatement

	// Expressions
	BoundErrorExpression
	BoundLiteralExpression
	BoundVariableExpression
	BoundAssignmentExpression
	BoundUnaryExpression
	BoundBinaryExpression
	BoundCallExpression
	BoundConversionExpression
	BoundTypeCallExpression
)
