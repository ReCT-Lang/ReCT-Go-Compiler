package boundnodes

import (
	print2 "ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

// incredibly cool interface for creating bound nodes
type BoundNode interface {
	NodeType() BoundType
	Print(indent string)
	Span() print2.TextSpan
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
	IsPersistent() bool
}

// enum for all our node types
type BoundType string

const (
	// based off of rect 1.0 source
	// -> https://github.com/RedCubeDev-ByteSpace/ReCT/tree/834776cbf0ad97da0e6441835f1bc19d903f115b/ReCT/CodeAnalysis/Binding

	// Statements
	BoundBlockStatement             BoundType = "BoundBlockStatement"
	BoundVariableDeclaration        BoundType = "BoundVariableDeclaration"
	BoundIfStatement                BoundType = "BoundIfStatement"
	BoundWhileStatement             BoundType = "BoundWhileStatement"
	BoundForStatement               BoundType = "BoundForStatement"
	BoundFromToStatement            BoundType = "BoundFromToStatement"
	BoundLabelStatement             BoundType = "BoundLabelStatement"
	BoundGotoStatement              BoundType = "BoundGotoStatement"
	BoundConditionalGotoStatement   BoundType = "BoundConditionalGotoStatement"
	BoundReturnStatement            BoundType = "BoundReturnStatement"
	BoundExpressionStatement        BoundType = "BoundExpressionStatement"
	BoundGarbageCollectionStatement BoundType = "BoundGarbageCollectionStatement"

	// Expressions
	BoundErrorExpression                BoundType = "BoundErrorExpression"
	BoundLiteralExpression              BoundType = "BoundLiteralExpression"
	BoundVariableExpression             BoundType = "BoundVariableExpression"
	BoundAssignmentExpression           BoundType = "BoundAssignmentExpression"
	BoundUnaryExpression                BoundType = "BoundUnaryExpression"
	BoundBinaryExpression               BoundType = "BoundBinaryExpression"
	BoundCallExpression                 BoundType = "BoundCallExpression"
	BoundPackageCallExpression          BoundType = "BoundPackageCallExpression"
	BoundConversionExpression           BoundType = "BoundConversionExpression"
	BoundTypeCallExpression             BoundType = "BoundTypeCallExpression"
	BoundClassCallExpression            BoundType = "BoundClassCallExpression"
	BoundClassFieldAccessExpression     BoundType = "BoundClassFieldAccessExpression"
	BoundClassFieldAssignmentExpression BoundType = "BoundClassFieldAssignmentExpression"
	BoundClassDestructionExpression     BoundType = "BoundClassDestructionExpression"
	BoundArrayAccessExpression          BoundType = "BoundArrayAccessExpression"
	BoundArrayAssignmentExpression      BoundType = "BoundArrayAssignmentExpression"
	BoundMakeExpression                 BoundType = "BoundMakeExpression"
	BoundMakeArrayExpression            BoundType = "BoundMakeArrayExpression"
	BoundFunctionExpression             BoundType = "BoundFunctionExpression"
	BoundThreadExpression               BoundType = "BoundThreadExpression"
	BoundTernaryExpression              BoundType = "BoundTernaryExpression"
	BoundReferenceExpression            BoundType = "BoundReferenceExpression"
	BoundDereferenceExpression          BoundType = "BoundDereferenceExpression"
	BoundMakeStructExpression           BoundType = "BoundMakeStructExpression"
)
