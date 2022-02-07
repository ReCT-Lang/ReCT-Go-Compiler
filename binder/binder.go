package binder

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"os"
)

type Binder struct {
	MemberScope    Scope
	Scopes         []Scope
	ActiveScope    *Scope
	FunctionSymbol symbols.FunctionSymbol

	LabelCounter   int
	BreakLabels    []boundnodes.BoundLabel
	ContinueLabels []boundnodes.BoundLabel
}

// helpers for the label stacks
func (bin *Binder) PushLabels(breakLabel boundnodes.BoundLabel, continueLabel boundnodes.BoundLabel) {
	bin.BreakLabels = append(bin.BreakLabels, breakLabel)
	bin.ContinueLabels = append(bin.ContinueLabels, continueLabel)
}

func (bin *Binder) PopLabels() {
	bin.BreakLabels[len(bin.BreakLabels)-1] = ""
	bin.ContinueLabels[len(bin.ContinueLabels)-1] = ""
	bin.BreakLabels = bin.BreakLabels[:len(bin.BreakLabels)-1]
	bin.ContinueLabels = bin.ContinueLabels[:len(bin.ContinueLabels)-1]
}

func (bin *Binder) GetLabels() (boundnodes.BoundLabel, boundnodes.BoundLabel) {
	return bin.BreakLabels[len(bin.BreakLabels)-1],
		bin.ContinueLabels[len(bin.ContinueLabels)-1]
}

// helpers for the scope list
func (bin *Binder) PushScope(s Scope) {
	bin.Scopes = append(bin.Scopes, s)
	bin.ActiveScope = &bin.Scopes[len(bin.Scopes)-1]
}

func (bin *Binder) PopScope() {
	bin.ActiveScope = bin.ActiveScope.Parent
}

// constructor
func CreateBinder(parent Scope, functionSymbol symbols.FunctionSymbol) *Binder {
	binder := Binder{
		MemberScope:    CreateScope(&parent),
		Scopes:         make([]Scope, 0),
		FunctionSymbol: functionSymbol,
		LabelCounter:   0,
		BreakLabels:    make([]boundnodes.BoundLabel, 0),
		ContinueLabels: make([]boundnodes.BoundLabel, 0),
	}

	binder.ActiveScope = &binder.MemberScope
	return &binder
}

// binder action

// <STATEMENTS> ---------------------------------------------------------------
func (bin *Binder) BindStatement(stmt nodes.StatementNode) boundnodes.BoundStatementNode {
	result := bin.BindStatementInternal(stmt)

	// only specific expressions are allowed to be used as statements
	// like function calls and variable assignments
	if result.NodeType() == boundnodes.BoundExpressionStatement {
		exprStmt := result.(boundnodes.BoundExpressionStatementNode)
		allowed := exprStmt.Expression.NodeType() == boundnodes.BoundErrorExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundCallExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundAssignmentExpression

		if !allowed {
			print.PrintC(print.Red, "Only call and assignment expressions are allowed to be used as statements!")
			os.Exit(-1)
		}
	}

	return result
}

func (bin *Binder) BindStatementInternal(stmt nodes.StatementNode) boundnodes.BoundStatementNode {
	switch stmt.NodeType() {
	case nodes.BlockStatement:
		return bin.BindBlockStatement(stmt.(nodes.BlockStatementNode))
	case nodes.VariableDeclaration:
		return bin.BindVariableDeclaration(stmt.(nodes.VariableDeclarationStatementNode))
	case nodes.IfStatement:
		return bin.BindIfStatement(stmt.(nodes.IfStatementNode))
	case nodes.ReturnStatement:
		return bin.BindReturnStatement(stmt.(nodes.ReturnStatementNode))
	case nodes.ForStatement:
		return bin.BindForStatement(stmt.(nodes.ForStatementNode))
	}

	print.PrintC(print.Red, "Unexpected statement node! Got: '"+string(stmt.NodeType())+"'")
	os.Exit(-1)
	return nil
}

func (bin *Binder) BindBlockStatement(stmt nodes.BlockStatementNode) boundnodes.BoundBlockStatementNode {
	// array of our new and improved bound statements
	statements := make([]boundnodes.BoundStatementNode, 0)

	for _, statement := range stmt.Statements {
		statements = append(statements, bin.BindStatement(statement))
	}

	return boundnodes.CreateBoundBlockStatementNode(statements)
}

func (bin *Binder) BindVariableDeclaration(stmt nodes.VariableDeclarationStatementNode) boundnodes.BoundVariableDeclarationStatementNode {
	// find out if this should be a global var or not
	isGlobal := stmt.Keyword.Kind == lexer.SetKeyword
	typeClause, clauseExists := bin.BindTypeClause(stmt.TypeClause)

	initializer := bin.BindExpression(stmt.Initializer)

	variableType := initializer.Type()
	if clauseExists {
		variableType = typeClause
	}

	variable := bin.BindVariableCreation(stmt.Identifier, false, isGlobal, variableType)
	convertedInitializer := bin.BindConversion(initializer, variableType, false)

	return boundnodes.CreateBoundVariableDeclarationStatementNode(variable, convertedInitializer)
}

func (bin *Binder) BindIfStatement(stmt nodes.IfStatementNode) boundnodes.BoundIfStatementNode {
	condition := bin.BindExpression(stmt.Condition)
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false)

	thenStatement := bin.BindStatement(stmt.ThenStatement)
	elseStatement := bin.BindElseClause(stmt.ElseClause)

	return boundnodes.CreateBoundIfStatementNode(convertedCondition, thenStatement, elseStatement)
}

func (bin *Binder) BindElseClause(clause nodes.ElseClauseNode) boundnodes.BoundStatementNode {
	if !clause.ClauseIsSet {
		return nil
	}

	return bin.BindStatement(clause.ElseStatement)
}

func (bin *Binder) BindReturnStatement(stmt nodes.ReturnStatementNode) boundnodes.BoundReturnStatementNode {
	var expression boundnodes.BoundExpressionNode = nil

	if stmt.Expression != nil {
		expression = bin.BindExpression(stmt.Expression)
	}

	// if we're not in any function
	if !bin.FunctionSymbol.Exists {
		print.PrintC(print.Red, "Cannot return when outside of a function!")
		os.Exit(-1)
	}

	// if we are in a function but the return type is void and we are trying to return smth
	if bin.FunctionSymbol.Exists &&
		bin.FunctionSymbol.Type.Fingerprint() == builtins.Void.Fingerprint() &&
		expression != nil {
		print.PrintC(print.Red, "Cannot return a value inside a void function!")
		os.Exit(-1)
	}

	return boundnodes.CreateBoundReturnStatementNode(expression)
}

func (bin *Binder) BindForStatement(stmt nodes.ForStatementNode) boundnodes.BoundForStatementNode {
	bin.PushScope(CreateScope(bin.ActiveScope))

	variable := bin.BindVariableDeclaration(stmt.Initaliser)

	condition := bin.BindExpression(stmt.Condition)
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false)

	updation := bin.BindStatement(stmt.Updation)

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()

	return boundnodes.CreateBoundForStatementNode(variable.Variable, convertedCondition, updation, body, breakLabel, continueLabel)
}

func (bin *Binder) BindLoopBody(stmt nodes.StatementNode) (boundnodes.BoundStatementNode, boundnodes.BoundLabel, boundnodes.BoundLabel) {
	bin.LabelCounter++

	breakLabel := boundnodes.BoundLabel(fmt.Sprintf("break%d", bin.LabelCounter))
	continueLabel := boundnodes.BoundLabel(fmt.Sprintf("continue%d", bin.LabelCounter))

	bin.PushLabels(breakLabel, continueLabel)
	loopBody := bin.BindStatement(stmt)
	bin.PopLabels()

	return loopBody, breakLabel, continueLabel
}

// </STATEMENTS> --------------------------------------------------------------
// <EXPRESSIONS> --------------------------------------------------------------

func (bin *Binder) BindExpression(expr nodes.ExpressionNode) boundnodes.BoundExpressionNode {
	switch expr.NodeType() {
	case nodes.LiteralExpression:
		return bin.BindLiteralExpression(expr.(nodes.LiteralExpressionNode))
	default:
		print.PrintC(print.Red, "Not implemented!")
		os.Exit(-1)
		return nil
	}
}

func (bin *Binder) BindLiteralExpression(expr nodes.LiteralExpressionNode) boundnodes.BoundLiteralExpressionNode {
	return boundnodes.CreateBoundLiteralExpressionNode(expr.LiteralValue)
}

// </EXPRESSIONS> -------------------------------------------------------------
// <SYMBOLS> ------------------------------------------------------------------

func (bin *Binder) BindVariableCreation(id lexer.Token, isReadOnly bool, isGlobal bool, varType symbols.TypeSymbol) symbols.VariableSymbol {
	var variable symbols.VariableSymbol

	if isGlobal {
		variable = symbols.CreateGlobalVariableSymbol(id.Value, isReadOnly, varType)
	} else {
		variable = symbols.CreateLocalVariableSymbol(id.Value, isReadOnly, varType)
	}

	if !bin.ActiveScope.TryDeclareSymbol(variable) {
		print.PrintC(print.Red, "Couldn't declare variable '"+id.Value+"'! Seems like a variable with this name has already been declared!")
		os.Exit(-1)
	}

	return variable
}

// </SYMBOLS> -----------------------------------------------------------------
// <IDEK> ---------------------------------------------------------------------

func (bin *Binder) BindTypeClause(tc nodes.TypeClauseNode) (symbols.TypeSymbol, bool) {
	// if this type clause doesnt actually exist
	if !tc.ClauseIsSet {
		return symbols.TypeSymbol{}, false
	}

	typ, _ := LookupType(tc.TypeIdentifier.Value)
	return typ, true
}

// </IDEK> --------------------------------------------------------------------
// <TYPES> --------------------------------------------------------------------

func (bin *Binder) BindConversion(expr boundnodes.BoundExpressionNode, to symbols.TypeSymbol, allowExplicit bool) boundnodes.BoundExpressionNode {
	conversionType := ClassifyConversion(expr.Type(), to)

	if !conversionType.Exists {
		print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'!")
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsExplicit && !allowExplicit {
		print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'! (An explicit conversion exists. Are you missing a cast?)")
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsIdentity {
		return expr
	}

	return boundnodes.CreateBoundConversionExpressionNode(to, expr)
}

func LookupType(name string) (symbols.TypeSymbol, bool) {
	switch name {
	case "void":
		return builtins.Void, true
	case "bool":
		return builtins.Bool, true
	case "int":
		return builtins.Int, true
	case "float":
		return builtins.Float, true
	case "string":
		return builtins.String, true
	case "any":
		return builtins.Any, true
	default:
		print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
		os.Exit(-1)

		return symbols.TypeSymbol{}, false
	}
}

// </TYPES> -------------------------------------------------------------------
