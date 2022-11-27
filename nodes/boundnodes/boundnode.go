package boundnodes

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/symbols"
)

// incredibly cool interface for creating bound nodes
type BoundNode interface {
	NodeType() BoundType
	Print(indent string)
	Source() nodes.SyntaxNode
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

	// Statements
	BoundBlockStatement           BoundType = "BoundBlockStatement"
	BoundVariableDeclaration      BoundType = "BoundVariableDeclaration"
	BoundIfStatement              BoundType = "BoundIfStatement"
	BoundWhileStatement           BoundType = "BoundWhileStatement"
	BoundForStatement             BoundType = "BoundForStatement"
	BoundFromToStatement          BoundType = "BoundFromToStatement"
	BoundLabelStatement           BoundType = "BoundLabelStatement"
	BoundGotoStatement            BoundType = "BoundGotoStatement"
	BoundConditionalGotoStatement BoundType = "BoundConditionalGotoStatement"
	BoundReturnStatement          BoundType = "BoundReturnStatement"
	BoundExpressionStatement      BoundType = "BoundExpressionStatement"

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
	BoundArrayAccessExpression          BoundType = "BoundArrayAccessExpression"
	BoundArrayAssignmentExpression      BoundType = "BoundArrayAssignmentExpression"
	BoundMakeExpression                 BoundType = "BoundMakeExpression"
	BoundMakeArrayExpression            BoundType = "BoundMakeArrayExpression"
	BoundFunctionExpression             BoundType = "BoundFunctionExpression"
	BoundTernaryExpression              BoundType = "BoundTernaryExpression"
	BoundReferenceExpression            BoundType = "BoundReferenceExpression"
	BoundDereferenceExpression          BoundType = "BoundDereferenceExpression"
	BoundMakeStructExpression           BoundType = "BoundMakeStructExpression"
	BoundLambdaExpression               BoundType = "BoundLambdaExpression"
	BoundThisExpression                 BoundType = "BoundThisExpression"
	BoundInternalValueExpression        BoundType = "BoundInternalValueExpression"
	BoundEnumExpression                 BoundType = "BoundEnumExpression"
)
