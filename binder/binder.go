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

	if binder.FunctionSymbol.Exists {
		for _, param := range binder.FunctionSymbol.Parameters {
			binder.ActiveScope.TryDeclareSymbol(param)
		}
	}

	return &binder
}

// binder action

// <MEMBERS> -----------------------------------------------------------------

func (bin *Binder) BindFunctionDeclaration(mem nodes.FunctionDeclarationMember) {
	boundParameters := make([]symbols.ParameterSymbol, 0)

	for i, param := range mem.Parameters {
		pName := param.Identifier.Value
		pType, _ := bin.BindTypeClause(param.TypeClause)

		// check if we've registered this param name before
		for i, p := range boundParameters {
			if p.Name == pName {
				// I haven't done bound nodes yet so I just get the syntax node parameter using the same index
				line, column, length := mem.Parameters[i].Position() // Very hacky :/
				print.Error(
					"BINDER",
					print.DuplicateParameterError,
					line,
					column,
					length,
					// Kind of a hacky way of getting the values and positions needed for the error
					"a parameter with the name \"%s\" already exists for function \"%s\"!",
					pName,
					mem.Identifier.Value,
				)
				os.Exit(-1)
			}
		}

		boundParameters = append(boundParameters, symbols.CreateParameterSymbol(pName, i, pType))
	}

	returnType, exists := bin.BindTypeClause(mem.TypeClause)
	if !exists {
		returnType = builtins.Void
	}

	functionSymbol := symbols.CreateFunctionSymbol(mem.Identifier.Value, boundParameters, returnType, mem)
	if !bin.ActiveScope.TryDeclareSymbol(functionSymbol) {
		//print.PrintC(print.Red, "Function '"+functionSymbol.Name+"' could not be defined! Seems like a function with the same name alredy exists!")
		line, column, length := mem.Position()
		print.Error(
			"BINDER",
			print.DuplicateFunctionError,
			line,
			column,
			length,
			"a function with the name \"%s\" already exists! \"%s\" could not be defined!",
			functionSymbol.Name,
			functionSymbol.Name,
		)
		os.Exit(-1)
	}
}

// </MEMBERS> ----------------------------------------------------------------
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
			//print.PrintC(print.Red, "Only call and assignment expressions are allowed to be used as statements!")
			line, column, length := stmt.Position()
			print.Error(
				"BINDER",
				print.UnexpectedExpressionStatementError,
				line,
				column,
				length,
				"cannot use \"%s\" as statement, only call and assignment expressions can be used as statements!",
				exprStmt,
			)
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
	case nodes.WhileStatement:
		return bin.BindWhileStatement(stmt.(nodes.WhileStatementNode))
	case nodes.FromToStatement:
		return bin.BindFromToStatement(stmt.(nodes.FromToStatementNode))
	case nodes.BreakStatement:
		return bin.BindBreakStatement(stmt.(nodes.BreakStatementNode))
	case nodes.ContinueStatement:
		return bin.BindContinueStatement(stmt.(nodes.ContinueStatementNode))
	case nodes.ExpressionStatement:
		return bin.BindExpressionStatement(stmt.(nodes.ExpressionStatementNode))
	}

	// print.PrintC(print.Red, "Unexpected statement node! Got: '"+string(stmt.NodeType())+"'")
	line, column, length := stmt.Position()
	print.Error(
		"Parser",
		print.UnknownStatementError,
		line,
		column,
		length,
		"\"%s\" Statement found. This was unexpected!",
	)
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
		//print.PrintC(print.Red, "Cannot return when outside of a function!")
		line, column, length := stmt.Position()
		print.Error(
			"BINDER",
			print.OutsideReturnError,
			line,
			column,
			length,
			"cannot use \"%s\" outside of a function!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	// if we are in a function but the return type is void, and we are trying to return something
	if bin.FunctionSymbol.Exists &&
		bin.FunctionSymbol.Type.Fingerprint() == builtins.Void.Fingerprint() &&
		expression != nil {
		//print.PrintC(print.Red, "Cannot return a value inside a void function!")
		line, column, length := stmt.Position()
		print.Error(
			"BINDER",
			print.VoidReturnError,
			line,
			column,
			length,
			"cannot use \"%s\" inside of a void function!",
			stmt.Keyword.Value,
		)
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

	return boundnodes.CreateBoundForStatementNode(variable, convertedCondition, updation, body, breakLabel, continueLabel)
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

func (bin *Binder) BindWhileStatement(stmt nodes.WhileStatementNode) boundnodes.BoundWhileStatementNode {
	bin.PushScope(CreateScope(bin.ActiveScope))

	condition := bin.BindExpression(stmt.Condition)
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false)

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()
	return boundnodes.CreateBoundWhileStatementNode(convertedCondition, body, breakLabel, continueLabel)
}

func (bin *Binder) BindFromToStatement(stmt nodes.FromToStatementNode) boundnodes.BoundFromToStatementNode {
	bin.PushScope(CreateScope(bin.ActiveScope))

	variable := bin.BindVariableCreation(stmt.Identifier, true, false, builtins.Int)
	lowerBound := bin.BindExpression(stmt.LowerBound)
	upperBound := bin.BindExpression(stmt.UpperBound)

	if lowerBound.Type().Fingerprint() != builtins.Int.Fingerprint() {
		line, column, length := stmt.LowerBound.Position()
		print.Error(
			"BINDER",
			print.UnexpectedNonIntegerValueError,
			line,
			column,
			length,
			"FromTo statement was expecting an integer value but instead got \"%s\"!\n",
			lowerBound.Type().Name,
		)
		os.Exit(-1)
	} else if upperBound.Type().Fingerprint() != builtins.Int.Fingerprint() {
		line, column, length := stmt.UpperBound.Position()
		print.Error(
			"BINDER",
			print.UnexpectedNonIntegerValueError,
			line,
			column,
			length,
			"FromTo statement was expecting an integer value but instead got \"%s\"!\n",
			upperBound.Type().Name,
		)
		os.Exit(-1)
	}

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()
	return boundnodes.CreateBoundFromToStatementNode(variable, lowerBound, upperBound, body, breakLabel, continueLabel)
}

func (bin *Binder) BindBreakStatement(stmt nodes.BreakStatementNode) boundnodes.BoundGotoStatementNode {
	// if we're not in any loop
	if len(bin.BreakLabels) == 0 {
		//print.PrintC(print.Red, "Cannot use break statement outside a loop!")
		line, column, length := stmt.Position()
		print.Error(
			"BINDER",
			print.OutsideBreakError,
			line,
			column,
			length,
			"cannot use \"%s\" outside of a loop!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	breakLabel, _ := bin.GetLabels()
	return boundnodes.CreateBoundGotoStatementNode(breakLabel)
}

func (bin *Binder) BindContinueStatement(stmt nodes.ContinueStatementNode) boundnodes.BoundGotoStatementNode {
	// if we're not in any loop
	if len(bin.BreakLabels) == 0 {
		//print.PrintC(print.Red, "Cannot use continue statement outside a loop!")
		line, column, length := stmt.Position()
		print.Error(
			"BINDER",
			print.OutsideContinueError,
			line,
			column,
			length,
			"cannot use \"%s\" outside of a loop!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	_, continueLabel := bin.GetLabels()
	return boundnodes.CreateBoundGotoStatementNode(continueLabel)
}

func (bin *Binder) BindExpressionStatement(stmt nodes.ExpressionStatementNode) boundnodes.BoundExpressionStatementNode {
	expression := bin.BindExpression(stmt.Expression)
	return boundnodes.CreateBoundExpressionStatementNode(expression)
}

// </STATEMENTS> --------------------------------------------------------------
// <EXPRESSIONS> --------------------------------------------------------------

func (bin *Binder) BindExpression(expr nodes.ExpressionNode) boundnodes.BoundExpressionNode {
	switch expr.NodeType() {
	case nodes.LiteralExpression:
		return bin.BindLiteralExpression(expr.(nodes.LiteralExpressionNode))
	case nodes.ParenthesisedExpression:
		return bin.BindParenthesisedExpression(expr.(nodes.ParenthesisedExpressionNode))
	case nodes.NameExpression:
		return bin.BindNameExpression(expr.(nodes.NameExpressionNode))
	case nodes.AssignmentExpression:
		return bin.BindAssignmentExpression(expr.(nodes.AssignmentExpressionNode))
	case nodes.VariableEditorExpression:
		return bin.BindVariableEditorExpression(expr.(nodes.VariableEditorExpressionNode))
	case nodes.CallExpression:
		return bin.BindCallExpression(expr.(nodes.CallExpressionNode))
	case nodes.UnaryExpression:
		return bin.BindUnaryExpression(expr.(nodes.UnaryExpressionNode))
	case nodes.TypeCallExpression:
		return bin.BindTypeCallExpression(expr.(nodes.TypeCallExpressionNode))
	case nodes.BinaryExpression:
		return bin.BindBinaryExpression(expr.(nodes.BinaryExpressionNode))
	default:
		//print.PrintC(print.Red, "Not implemented!")
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.NotImplementedError,
			line,
			column,
			length,
			"\"%s\" is not implemented yet! (cringe)",
			expr.NodeType(),
		)
		os.Exit(-1)
		return nil
	}
}

func (bin *Binder) BindLiteralExpression(expr nodes.LiteralExpressionNode) boundnodes.BoundLiteralExpressionNode {
	return boundnodes.CreateBoundLiteralExpressionNode(expr.LiteralValue)
}

func (bin *Binder) BindParenthesisedExpression(expr nodes.ParenthesisedExpressionNode) boundnodes.BoundExpressionNode {
	return bin.BindExpression(expr.Expression)
}

func (bin *Binder) BindNameExpression(expr nodes.NameExpressionNode) boundnodes.BoundVariableExpressionNode {
	variable := bin.BindVariableReference(expr.Identifier.Value)
	return boundnodes.CreateBoundVariableExpressionNode(variable)
}

func (bin *Binder) BindAssignmentExpression(expr nodes.AssignmentExpressionNode) boundnodes.BoundAssignmentExpressionNode {
	variable := bin.BindVariableReference(expr.Identifier.Value)
	expression := bin.BindExpression(expr.Expression)
	convertedExpression := bin.BindConversion(expression, variable.VarType(), false)

	return boundnodes.CreateBoundAssignmentExpressionNode(variable, convertedExpression)
}

func (bin *Binder) BindVariableEditorExpression(expr nodes.VariableEditorExpressionNode) boundnodes.BoundAssignmentExpressionNode {
	// bind the variable
	variable := bin.BindVariableReference(expr.Identifier.Value)

	// create a placeholder expression of value 1
	var expression boundnodes.BoundExpressionNode = boundnodes.CreateBoundLiteralExpressionNode(1)

	// if we have an expression given, use it instead
	if expr.Expression != nil {
		expression = bin.BindExpression(expr.Expression)
	}

	// bind the operator and binary expression for our operation
	operator := boundnodes.BindBinaryOperator(expr.Operator.Kind, variable.VarType(), expression.Type())

	// check if the operator is valid
	if !operator.Exists {
		//print.PrintC(print.Red, "Binary operator '"+expr.Operator.Value+"' is not defined for types '"+variable.VarType().Name+"' and '"+expression.Type().Name+"'!")
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.BinaryOperatorTypeError,
			line,
			column,
			length,
			"the use of binary operator \"%s\" with types \"%s\" and \"%s\" is undefined!",
			expr.Operator.Value,
			variable.VarType().Name,
			expression.Type().Name,
		)
		os.Exit(-1)
	}

	binaryExpression := boundnodes.CreateBoundBinaryExpressionNode(
		boundnodes.CreateBoundVariableExpressionNode(variable),
		operator,
		expression,
	)

	// return it as an assignment
	return boundnodes.CreateBoundAssignmentExpressionNode(variable, binaryExpression)
}

func (bin *Binder) BindTypeCallExpression(expr nodes.TypeCallExpressionNode) boundnodes.BoundTypeCallExpressionNode {
	variable := bin.BindVariableReference(expr.Identifier.Value)

	// This line will error out because CallIdentifier.RealValue is nil (interface)
	// I've replaced it with CallIdentifier.Value which seems to do the trick.
	function := bin.LookupTypeFunction(expr.CallIdentifier.Value) // Should be a string anyway
	if function.OriginType.Name != variable.VarType().Name {
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.IncorrectTypeFunctionCallError,
			line,
			column,
			length,
			"the use of builtin function \"%s\" on \"%s\" datatype is undefined!",
			function.Name,
			variable.VarType().Name,
		)
		os.Exit(-1)
	}

	// bind all given arguments
	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	// make sure we got the right number of arguments
	if len(boundArguments) != len(function.Parameters) {
		//print.PrintCF(print.Red, "Type function '%s' expects %d arguments, got %d!", function.Name, len(function.Parameters), len(boundArguments))
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			line,
			column,
			length,
			"type function \"%s\" expects %d arguments but got %d!",
			function.Name,
			variable.VarType().Name,
		)
		os.Exit(-1)
	}

	// make sure all arguments are the right type
	for i, arg := range boundArguments {
		boundArguments[i] = bin.BindConversion(arg, function.Parameters[i].VarType(), false)
	}

	return boundnodes.CreateBoundTypeCallExpressionNode(variable, function, boundArguments)
}

func (bin *Binder) BindCallExpression(expr nodes.CallExpressionNode) boundnodes.BoundExpressionNode {
	// check if this is a cast
	typeSymbol, exists := LookupType(expr.Identifier.Value, true)
	if exists && len(expr.Arguments) == 1 {
		// bind the expression and return a conversion
		expression := bin.BindExpression(expr.Arguments[0])
		return bin.BindConversion(expression, typeSymbol, true)
	}

	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	symbol := bin.ActiveScope.TryLookupSymbol(expr.Identifier.Value)
	if symbol == nil ||
		symbol.SymbolType() != symbols.Function {
		print.PrintC(print.Red, "Cannot find function '"+expr.Identifier.Value+"'!")
		os.Exit(-1)
	}

	functionSymbol := symbol.(symbols.FunctionSymbol)
	if len(boundArguments) != len(functionSymbol.Parameters) {
		//fmt.Printf("%sFunction '%s' expects %d arguments, got %d!%s\n", print.ERed, functionSymbol.Name, len(functionSymbol.Parameters), len(boundArguments), print.EReset)
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			line,
			column,
			length,
			"type function \"%s\" expects %d arguments but got %d!",
			expr.Identifier,
			len(functionSymbol.Parameters),
			len(expr.Arguments),
		)
		os.Exit(-1)
	}

	for i := 0; i < len(boundArguments); i++ {
		boundArguments[i] = bin.BindConversion(boundArguments[i], functionSymbol.Parameters[i].VarType(), false)
	}

	return boundnodes.CreateBoundCallExpressionNode(functionSymbol, boundArguments)
}

func (bin *Binder) BindUnaryExpression(expr nodes.UnaryExpressionNode) boundnodes.BoundUnaryExpressionNode {
	operand := bin.BindExpression(expr.Operand)
	op := boundnodes.BindUnaryOperator(expr.Operator.Kind, operand.Type())

	if !op.Exists {
		//print.PrintC(print.Red, "Unary operator '"+expr.Operator.Value+"' is not defined for type '"+operand.Type().Name+"'!")
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.UnaryOperatorTypeError,
			line,
			column,
			length,
			"the use of unary operator \"%s\" with type \"%s\" is undefined!",
			expr.Operator.Value,
			operand.Type().Name,
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundUnaryExpressionNode(op, operand)
}

func (bin *Binder) BindBinaryExpression(expr nodes.BinaryExpressionNode) boundnodes.BoundBinaryExpressionNode {
	left := bin.BindExpression(expr.Left)
	right := bin.BindExpression(expr.Right)
	op := boundnodes.BindBinaryOperator(expr.Operator.Kind, left.Type(), right.Type())

	if !op.Exists {
		//print.PrintC(print.Red, "Binary operator '"+expr.Operator.Value+"' is not defined for types '"+left.Type().Name+"' and '"+right.Type().Name+"'!")
		line, column, length := expr.Position()
		print.Error(
			"BINDER",
			print.BinaryOperatorTypeError,
			line,
			column,
			length,
			"the use of binary operator \"%s\" with types \"%s\" and \"%s\" is undefined!",
			expr.Operator.Value,
			left.Type().Name,
			right.Type().Name,
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundBinaryExpressionNode(left, op, right)
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
		//print.PrintC(print.Red, "Couldn't declare variable '"+id.Value+"'! Seems like a variable with this name has already been declared!")
		print.Error(
			"BINDER",
			print.DuplicateVariableDeclarationError,
			id.Line,
			id.Column,
			len(id.Value)+3+2+len(variable.VarType().Name), // Probably wrong, but it works - that's the tokorv guarantee
			"Variable \"%s\" could not be declared! Variable with this name has already been declared!",
			id.Value,
		)
		os.Exit(-1)
	}

	return variable
}

func (bin *Binder) BindVariableReference(name string) symbols.VariableSymbol {
	variable := bin.ActiveScope.TryLookupSymbol(name)

	if variable == nil ||
		!(variable.SymbolType() == symbols.GlobalVariable ||
			variable.SymbolType() == symbols.LocalVariable ||
			variable.SymbolType() == symbols.Parameter) {
		//print.PrintC(print.Red, "Could not find variable '"+name+"'!")
		print.Error(
			"BINDER",
			print.UndefinedVariableReferenceError,
			0,
			0, // We have no data for where this variable reference is?
			0,
			"Could not find variable \"%s\"! Are you sure it exists?",
			name,
		)
		os.Exit(-1)
	}

	return variable.(symbols.VariableSymbol)
}

// </SYMBOLS> -----------------------------------------------------------------
// <IDEK> ---------------------------------------------------------------------

func (bin *Binder) BindTypeClause(tc nodes.TypeClauseNode) (symbols.TypeSymbol, bool) {
	// if this type clause doesnt actually exist
	if !tc.ClauseIsSet {
		return symbols.TypeSymbol{}, false
	}

	typ, _ := LookupType(tc.TypeIdentifier.Value, false)
	return typ, true
}

func (bin *Binder) LookupTypeFunction(name string) symbols.TypeFunctionSymbol {
	switch name {
	case "GetLength":
		return builtins.GetLength
	case "Substring":
		return builtins.Substring
	default:
		/*print.PrintC(
			print.Red,
			fmt.Sprintf("Could not find builtin TypeFunctionSymbol \"%s\"!", name),
		)*/
		print.Error(
			"BINDER",
			print.TypeFunctionDoesNotExistError,
			0,
			0, // needs extra data added and passed into the function
			0, // Probably wrong, but it works - that's the tokorv guarantee
			"Could not find builtin TypeFunctionSymbol \"%s\"!",
			name,
		)
		os.Exit(-1)
	}
	return symbols.TypeFunctionSymbol{}
}

// </IDEK> --------------------------------------------------------------------
// <TYPES> --------------------------------------------------------------------

func (bin *Binder) BindConversion(expr boundnodes.BoundExpressionNode, to symbols.TypeSymbol, allowExplicit bool) boundnodes.BoundExpressionNode {
	conversionType := ClassifyConversion(expr.Type(), to)

	if !conversionType.Exists {
		//print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'!")
		print.Error(
			"BINDER",
			print.ConversionError,
			0,
			0, // needs extra data added and passed into the function
			0, // Probably wrong, but it works - that's the tokorv guarantee
			"Cannot convert type \"%s\" to \"%s\"!",
			expr.Type().Name,
			to.Name,
		)
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsExplicit && !allowExplicit {
		//print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'! (An explicit conversion exists. Are you missing a cast?)")
		print.Error(
			"BINDER",
			print.ExplicitConversionError,
			0,
			0, // needs extra data added and passed into the function
			0, // Probably wrong, but it works - that's the tokorv guarantee
			"Cannot convert type \"%s\" to \"%s\"! (An explicit conversion exists. Are you missing a cast?)",
			expr.Type().Name,
			to.Name,
		)
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsIdentity {
		return expr
	}

	return boundnodes.CreateBoundConversionExpressionNode(to, expr)
}

func LookupType(name string, canFail bool) (symbols.TypeSymbol, bool) {
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
		if !canFail {
			//print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
			print.Error(
				"BINDER",
				print.ExplicitConversionError,
				0,
				0, // needs extra data added and passed into the function
				0, // Probably wrong, but it works - that's the tokorv guarantee
				"Couldn't find datatype \"%s\"! Are you sure it exists?",
				name,
			)
			os.Exit(-1)
		}

		return symbols.TypeSymbol{}, false
	}
}

// </TYPES> -------------------------------------------------------------------
