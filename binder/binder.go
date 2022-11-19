package binder

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/lowerer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/packager"
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
	ClassSymbol    symbols.ClassSymbol

	LabelCounter   int
	BreakLabels    []boundnodes.BoundLabel
	ContinueLabels []boundnodes.BoundLabel

	PreInitialTypeset []symbols.TypeSymbol
	InClass           bool
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

func (bin *Binder) BindFunctionDeclaration(mem nodes.FunctionDeclarationMember, inClass bool) {
	boundParameters := make([]symbols.ParameterSymbol, 0)

	for i, param := range mem.Parameters {
		pName := param.Identifier.Value
		pType, _ := bin.BindTypeClause(param.TypeClause)

		// check if we've registered this param name before
		for i, p := range boundParameters {
			if p.Name == pName {
				// I haven't done bound nodes yet so I just get the syntax node parameter using the same index
				print.Error(
					"BINDER",
					print.DuplicateParameterError,
					mem.Parameters[i].Span(),
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

	functionSymbol := symbols.CreateFunctionSymbol(mem.Identifier.Value, boundParameters, returnType, mem, mem.IsPublic)

	// make sure reserved functions like Constructor() and Die() meet certain requirements
	if inClass {
		if functionSymbol.Name == "Constructor" && functionSymbol.Public {
			print.Error(
				"BINDER",
				print.IllegalFunctionSignatureError,
				mem.Identifier.Span,
				"reserved function 'Constructor' is not allowed to be public!",
			)
			os.Exit(-1)
		}
	} else {
		if functionSymbol.Name == "main" {
			print.Error(
				"BINDER",
				print.IllegalFunctionSignatureError,
				mem.Identifier.Span,
				"reserved function name 'main' is not allowed to be used by a user defined function!",
			)
			os.Exit(-1)
		}
	}

	if !bin.ActiveScope.TryDeclareSymbol(functionSymbol) {
		//print.PrintC(print.Red, "Function '"+functionSymbol.Name+"' could not be defined! Seems like a function with the same name alredy exists!")
		print.Error(
			"BINDER",
			print.DuplicateFunctionError,
			mem.Identifier.Span,
			"a member with the name \"%s\" already exists! \"%s\" could not be defined!",
			functionSymbol.Name,
			functionSymbol.Name,
		)
		os.Exit(-1)
	}
}

func (bin *Binder) BindExternalFunctionDeclaration(mem nodes.ExternalFunctionDeclarationMember, inClass bool) {
	boundParameters := make([]symbols.ParameterSymbol, 0)

	for i, param := range mem.Parameters {
		pName := param.Identifier.Value
		pType, _ := bin.BindTypeClause(param.TypeClause)

		// check if we've registered this param name before
		for i, p := range boundParameters {
			if p.Name == pName {
				// I haven't done bound nodes yet so I just get the syntax node parameter using the same index
				print.Error(
					"BINDER",
					print.DuplicateParameterError,
					mem.Parameters[i].Span(),
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

	functionSymbol := symbols.CreateExternalFunctionSymbol(
		mem.Identifier.Value,
		nodes.FunctionDeclarationMember{
			FunctionKeyword: mem.ExternalKeyword,
			Identifier:      mem.Identifier,
		},
		boundParameters,
		returnType,
		mem.IsVariadic,
		mem.IsAdapted)

	// make sure reserved functions like Constructor() and Die() meet certain requirements
	if inClass {
		print.Error(
			"BINDER",
			print.InvalidExternalFunctionPlacementError,
			mem.Identifier.Span,
			"all external functions are required to be in the global scope!",
			functionSymbol.Name,
			functionSymbol.Name,
		)
	} else {
		if functionSymbol.Name == "main" {
			print.Error(
				"BINDER",
				print.IllegalFunctionSignatureError,
				mem.Identifier.Span,
				"reserved function name 'main' is not allowed to be used by a user defined function!",
			)
			os.Exit(-1)
		}
	}

	if !bin.ActiveScope.TryDeclareSymbol(functionSymbol) {
		//print.PrintC(print.Red, "Function '"+functionSymbol.Name+"' could not be defined! Seems like a function with the same name alredy exists!")
		print.Error(
			"BINDER",
			print.DuplicateFunctionError,
			mem.Identifier.Span,
			"a member with the name \"%s\" already exists! \"%s\" could not be defined!",
			functionSymbol.Name,
			functionSymbol.Name,
		)
		os.Exit(-1)
	}
}

func (bin *Binder) BindClassDeclaration(mem nodes.ClassDeclarationMember, preInitialTypeset []symbols.TypeSymbol) {
	rootScope := BindRootScope()
	classScope := CreateScope(&rootScope)

	functionDeclarations := make([]nodes.FunctionDeclarationMember, 0)
	globalStatements := make([]nodes.GlobalStatementMember, 0)

	// sort all our members into functions and global statements
	for _, member := range mem.Members {
		if member.NodeType() == nodes.FunctionDeclaration {
			functionDeclarations = append(functionDeclarations, member.(nodes.FunctionDeclarationMember))
		} else if member.NodeType() == nodes.ClassDeclaration {
			print.Error(
				"BINDER",
				print.IllegalNestedClassesError,
				member.Span(),
				// yes
				"Nested classes = illegal! >:(",
			)
			os.Exit(-1)
		} else {
			globalStatements = append(globalStatements, member.(nodes.GlobalStatementMember))
		}
	}

	binder := CreateBinder(classScope, symbols.FunctionSymbol{})
	binder.PreInitialTypeset = preInitialTypeset

	hasConstructor := false

	// declare all our functions
	for _, fnc := range functionDeclarations {
		binder.BindFunctionDeclaration(fnc, true)

		// if this function is a constructor, the class has one
		if fnc.Identifier.Value == "Constructor" {
			hasConstructor = true
		}
	}

	// if the class doenst have a constructor -> create an empty one
	if !hasConstructor {
		binder.BindFunctionDeclaration(nodes.CreateFunctionDeclarationMember(
			lexer.Token{},
			lexer.Token{Kind: lexer.IdToken, Value: "Constructor"},
			make([]nodes.ParameterNode, 0),
			nodes.TypeClauseNode{},
			nodes.CreateBlockStatementNode(lexer.Token{}, make([]nodes.StatementNode, 0), lexer.Token{}),
			false,
		), true)
	}

	// check all our statements, only variable declarations are allowed in here
	for _, stmt := range globalStatements {
		if stmt.Statement.NodeType() != nodes.VariableDeclaration {
			print.Error(
				"BINDER",
				print.InvalidStatementPlacementError,
				stmt.Span(),
				// yes
				"Only variable declarations are allowed in a class' global scope!",
			)
			os.Exit(-1)
		}

		// only public vars can be created here
		if stmt.Statement.(nodes.VariableDeclarationStatementNode).Keyword.Kind != lexer.SetKeyword {
			print.Error(
				"BINDER",
				print.InvalidStatementPlacementError,
				stmt.Span(),
				// yes
				"Only global variable declarations are allowed in a class' global scope!",
			)
			os.Exit(-1)
		}

		// if everything is alright, we can bind the variable
		// this is only done to produce a variable symbol
		binder.BindStatement(stmt.Statement)
	}

	// Build the ClassSymbol
	// ---------------------
	vars := binder.MemberScope.GetAllVariables()
	funcs := binder.MemberScope.GetAllFunctions()

	classSym := symbols.CreateClassSymbol(mem.Identifier.Value, mem, funcs, vars)

	if !bin.ActiveScope.TryDeclareSymbol(classSym) {
		print.Error(
			"BINDER",
			print.DuplicateFunctionError,
			mem.Span(),
			"A member with the name \"%s\" already exists! \"%s\" could not be defined!",
			classSym.Name,
			classSym.Name,
		)
		os.Exit(-1)
	}
}

func (bin *Binder) BindStructDeclaration(mem nodes.StructDeclarationMember, preInitialTypeset []symbols.TypeSymbol) {
	rootScope := BindRootScope()
	classScope := CreateScope(&rootScope)

	binder := CreateBinder(classScope, symbols.FunctionSymbol{})
	binder.PreInitialTypeset = preInitialTypeset

	fields := make([]symbols.VariableSymbol, 0)

	// check all our statements, only variable declarations are allowed in here
	for _, fld := range mem.Fields {

		// type resolving error is ignored, should never happen lol
		fldType, _ := binder.BindTypeClause(fld.TypeClause)

		// check for redefinitions
		for _, field := range fields {
			if fld.Identifier.Value == field.SymbolName() {
				print.Error(
					"BINDER",
					print.DuplicateFunctionError,
					fld.Span(),
					"A field with the name \"%s\" already exists! \"%s\" could not be defined!",
					field.SymbolName(),
					field.SymbolName(),
				)
				os.Exit(-1)
			}
		}

		// store this field
		fields = append(fields, symbols.CreateGlobalVariableSymbol(fld.Identifier.Value, false, fldType))
	}

	// Build the StructSymbol
	// ----------------------

	structSym := symbols.CreateStructSymbol(mem.Identifier.Value, mem, fields)

	if !bin.ActiveScope.TryDeclareSymbol(structSym) {
		print.Error(
			"BINDER",
			print.DuplicateFunctionError,
			mem.Span(),
			"A member with the name \"%s\" already exists! \"%s\" could not be defined!",
			structSym.Name,
			structSym.Name,
		)
		os.Exit(-1)
	}
}

func (bin *Binder) BindPackageReference(mem nodes.PackageReferenceMember) {
	pack := packager.ResolvePackage(mem.Package.Value, mem.Span())
	if !bin.ActiveScope.TryDeclareSymbol(pack) {
		print.Error(
			"BINDER",
			print.DuplicatePackageImportError,
			mem.Span(),
			"a package with the name \"%s\" has already been loaded! \"%s\" could not be referenced!",
			pack.Name,
			pack.Name,
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
			exprStmt.Expression.NodeType() == boundnodes.BoundTypeCallExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundClassCallExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundPackageCallExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundAssignmentExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundArrayAssignmentExpression ||
			exprStmt.Expression.NodeType() == boundnodes.BoundClassFieldAssignmentExpression

		if !allowed {
			//print.PrintC(print.Red, "Only call and assignment expressions are allowed to be used as statements!")
			print.Error(
				"BINDER",
				print.UnexpectedExpressionStatementError,
				stmt.Span(),
				"cannot use \"%s\" as statement, only call and assignment expressions can be used as statements!",
				exprStmt.Expression.NodeType(),
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
	print.Error(
		"BINDER",
		print.UnknownStatementError,
		stmt.Span(),
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

	return boundnodes.CreateBoundBlockStatementNode(statements, stmt.Span())
}

func (bin *Binder) BindVariableDeclaration(stmt nodes.VariableDeclarationStatementNode) boundnodes.BoundVariableDeclarationStatementNode {
	// find out if this should be a global var or not
	isGlobal := stmt.Keyword.Kind == lexer.SetKeyword
	typeClause, clauseExists := bin.BindTypeClause(stmt.TypeClause)

	var initializer boundnodes.BoundExpressionNode
	var convertedInitializer boundnodes.BoundExpressionNode
	var variableType symbols.TypeSymbol

	// if there's an initializer -> bind and use it
	if stmt.Initializer != nil {
		initializer = bin.BindExpression(stmt.Initializer)
		variableType = initializer.Type()
	}

	if clauseExists {
		variableType = typeClause
	}

	// if there's no clause but also no initializer -> throw error!
	if variableType.Name == "" && stmt.Initializer == nil {
		print.Error(
			"BINDER",
			print.IllegalVariableDeclarationError,
			stmt.Span(),
			"Variable declaration is neither given a type, nor an initializer!",
		)
		os.Exit(-1)
	}

	variable := bin.BindVariableCreation(stmt.Identifier, false, isGlobal, variableType)

	if initializer != nil {
		convertedInitializer = bin.BindConversion(initializer, variableType, false, stmt.Initializer.Span())
	}

	return boundnodes.CreateBoundVariableDeclarationStatementNode(variable, convertedInitializer, stmt.Span())
}

func (bin *Binder) BindIfStatement(stmt nodes.IfStatementNode) boundnodes.BoundIfStatementNode {
	condition := bin.BindExpression(stmt.Condition)
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false, stmt.Condition.Span())

	thenStatement := bin.BindStatement(stmt.ThenStatement)
	elseStatement := bin.BindElseClause(stmt.ElseClause)

	return boundnodes.CreateBoundIfStatementNode(convertedCondition, thenStatement, elseStatement, stmt.Span())
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
		print.Error(
			"BINDER",
			print.OutsideReturnError,
			stmt.Span(),
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
		print.Error(
			"BINDER",
			print.VoidReturnError,
			stmt.Span(),
			"cannot use \"%s\" inside of a void function!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundReturnStatementNode(expression, stmt.Span())
}

func (bin *Binder) BindForStatement(stmt nodes.ForStatementNode) boundnodes.BoundForStatementNode {
	bin.PushScope(CreateScope(bin.ActiveScope))

	variable := bin.BindVariableDeclaration(stmt.Initaliser)

	condition := bin.BindExpression(stmt.Condition)
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false, stmt.Condition.Span())

	updation := bin.BindStatement(stmt.Updation)

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()

	return boundnodes.CreateBoundForStatementNode(variable, convertedCondition, updation, body, breakLabel, continueLabel, stmt.Span())
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
	convertedCondition := bin.BindConversion(condition, builtins.Bool, false, stmt.Condition.Span())

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()
	return boundnodes.CreateBoundWhileStatementNode(convertedCondition, body, breakLabel, continueLabel, stmt.Span())
}

func (bin *Binder) BindFromToStatement(stmt nodes.FromToStatementNode) boundnodes.BoundFromToStatementNode {
	bin.PushScope(CreateScope(bin.ActiveScope))

	variable := bin.BindVariableCreation(stmt.Identifier, true, false, builtins.Int)
	lowerBound := bin.BindExpression(stmt.LowerBound)
	upperBound := bin.BindExpression(stmt.UpperBound)

	if lowerBound.Type().Fingerprint() != builtins.Int.Fingerprint() {
		print.Error(
			"BINDER",
			print.UnexpectedNonIntegerValueError,
			stmt.LowerBound.Span(),
			"FromTo statement was expecting an integer value but instead got \"%s\"!\n",
			lowerBound.Type().Name,
		)
		os.Exit(-1)
	} else if upperBound.Type().Fingerprint() != builtins.Int.Fingerprint() {
		print.Error(
			"BINDER",
			print.UnexpectedNonIntegerValueError,
			stmt.UpperBound.Span(),
			"FromTo statement was expecting an integer value but instead got \"%s\"!\n",
			upperBound.Type().Name,
		)
		os.Exit(-1)
	}

	body, breakLabel, continueLabel := bin.BindLoopBody(stmt.Statement)

	bin.PopScope()
	return boundnodes.CreateBoundFromToStatementNode(variable, lowerBound, upperBound, body, breakLabel, continueLabel, stmt.Span())
}

func (bin *Binder) BindBreakStatement(stmt nodes.BreakStatementNode) boundnodes.BoundGotoStatementNode {
	// if we're not in any loop
	if len(bin.BreakLabels) == 0 {
		//print.PrintC(print.Red, "Cannot use break statement outside a loop!")
		print.Error(
			"BINDER",
			print.OutsideBreakError,
			stmt.Span(),
			"cannot use \"%s\" outside of a loop!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	breakLabel, _ := bin.GetLabels()
	return boundnodes.CreateBoundGotoStatementNode(breakLabel, stmt.Span())
}

func (bin *Binder) BindContinueStatement(stmt nodes.ContinueStatementNode) boundnodes.BoundGotoStatementNode {
	// if we're not in any loop
	if len(bin.BreakLabels) == 0 {
		//print.PrintC(print.Red, "Cannot use continue statement outside a loop!")
		print.Error(
			"BINDER",
			print.OutsideContinueError,
			stmt.Span(),
			"cannot use \"%s\" outside of a loop!",
			stmt.Keyword.Value,
		)
		os.Exit(-1)
	}

	_, continueLabel := bin.GetLabels()
	return boundnodes.CreateBoundGotoStatementNode(continueLabel, stmt.Span())
}

func (bin *Binder) BindExpressionStatement(stmt nodes.ExpressionStatementNode) boundnodes.BoundExpressionStatementNode {
	expression := bin.BindExpression(stmt.Expression)
	return boundnodes.CreateBoundExpressionStatementNode(expression, stmt.Span())
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
	case nodes.ArrayAccessExpression:
		return bin.BindArrayAccessExpression(expr.(nodes.ArrayAccessExpressionNode))
	case nodes.ArrayAssignmentExpression:
		return bin.BindArrayAssignmentExpression(expr.(nodes.ArrayAssignmentExpressionNode))
	case nodes.MakeExpression:
		return bin.BindMakeExpression(expr.(nodes.MakeExpressionNode))
	case nodes.MakeArrayExpression:
		return bin.BindMakeArrayExpression(expr.(nodes.MakeArrayExpressionNode))
	case nodes.MakeStructExpression:
		return bin.BindMakeStructExpression(expr.(nodes.MakeStructExpressionNode))
	case nodes.CallExpression:
		return bin.BindCallExpression(expr.(nodes.CallExpressionNode))
	case nodes.PackageCallExpression:
		return bin.BindPackageCallExpression(expr.(nodes.PackageCallExpressionNode))
	case nodes.UnaryExpression:
		return bin.BindUnaryExpression(expr.(nodes.UnaryExpressionNode))
	case nodes.TypeCallExpression:
		return bin.BindTypeCallExpression(expr.(nodes.TypeCallExpressionNode))
	case nodes.ClassFieldAccessExpression:
		return bin.BindClassFieldAccessExpression(expr.(nodes.ClassFieldAccessExpressionNode))
	case nodes.ClassFieldAssignmentExpression:
		return bin.BindClassFieldAssignmentExpression(expr.(nodes.ClassFieldAssignmentExpressionNode))
	case nodes.BinaryExpression:
		return bin.BindBinaryExpression(expr.(nodes.BinaryExpressionNode))
	case nodes.TernaryExpression:
		return bin.BindTernaryExpression(expr.(nodes.TernaryExpressionNode))
	case nodes.ReferenceExpression:
		return bin.BindReferenceExpression(expr.(nodes.ReferenceExpressionNode))
	case nodes.DereferenceExpression:
		return bin.BindDereferenceExpression(expr.(nodes.DereferenceExpressionNode))
	case nodes.LambdaExpression:
		return bin.BindLambdaExpression(expr.(nodes.LambdaExpressionNode))
	case nodes.ThisExpression:
		return bin.BindThisExpression(expr.(nodes.ThisExpressionNode))

	default:
		//print.PrintC(print.Red, "Not implemented!")
		print.Error(
			"BINDER",
			print.NotImplementedError,
			expr.Span(),
			"\"%s\" is not implemented yet! (cringe)",
			expr.NodeType(),
		)
		os.Exit(-1)
		return nil
	}
}

func (bin *Binder) BindLiteralExpression(expr nodes.LiteralExpressionNode) boundnodes.BoundLiteralExpressionNode {
	return boundnodes.CreateBoundLiteralExpressionNode(expr, expr.Span())
}

func (bin *Binder) BindParenthesisedExpression(expr nodes.ParenthesisedExpressionNode) boundnodes.BoundExpressionNode {
	return bin.BindExpression(expr.Expression)
}

func (bin *Binder) BindNameExpression(expr nodes.NameExpressionNode) boundnodes.BoundExpressionNode {
	symbol := bin.ActiveScope.TryLookupSymbol(expr.Identifier.Value)
	// normal variable lookup
	if symbol == nil || symbol.SymbolType() != symbols.Function {
		variable := bin.BindVariableReference(expr.Identifier.Value, expr.Identifier.Span)
		return boundnodes.CreateBoundVariableExpressionNode(variable, expr.Span())

		// funky function lookup
	} else if symbol.SymbolType() == symbols.Function {
		functionSymbol := symbol.(symbols.FunctionSymbol)

		if bin.InClass {
			// we protecc
			// ----------

			// if we are inside a class, dont allow calls to Constructor() and Die()
			if bin.InClass && functionSymbol.Name == "Constructor" {
				print.Error(
					"BINDER",
					print.IllegalConstructorCallError,
					expr.Span(),
					"Lambda reference to Constructor in own class is not allowed!",
				)
				os.Exit(-1)
			}

			return boundnodes.CreateBoundFunctionInClassExpressionNode(functionSymbol, bin.ClassSymbol, expr.Span())
		} else {
			return boundnodes.CreateBoundFunctionExpressionNode(functionSymbol, expr.Span())
		}

		// ah yes, safety
	} else {
		return nil
	}
}

func (bin *Binder) BindAssignmentExpression(expr nodes.AssignmentExpressionNode) boundnodes.BoundAssignmentExpressionNode {
	variable := bin.BindVariableReference(expr.Identifier.Value, expr.Identifier.Span)
	expression := bin.BindExpression(expr.Expression)
	convertedExpression := bin.BindConversion(expression, variable.VarType(), false, expr.Expression.Span())

	return boundnodes.CreateBoundAssignmentExpressionNode(variable, convertedExpression, expr.Span())
}

func (bin *Binder) BindVariableEditorExpression(expr nodes.VariableEditorExpressionNode) boundnodes.BoundAssignmentExpressionNode {
	// bind the variable
	variable := bin.BindVariableReference(expr.Identifier.Value, expr.Identifier.Span)

	// create a placeholder expression of value 1
	var expression boundnodes.BoundExpressionNode = boundnodes.CreateBoundLiteralExpressionNodeFromValue(1, expr.Span())

	// if we have an expression given, use it instead
	if expr.Expression != nil {
		expression = bin.BindExpression(expr.Expression)
	}

	binaryExpression := bin.BindBinaryExpressionInternal(
		boundnodes.CreateBoundVariableExpressionNode(variable, expr.Span()),
		expression,
		expr.Operator.Kind,
	)

	// return it as an assignment
	return boundnodes.CreateBoundAssignmentExpressionNode(variable, binaryExpression, expr.Span())
}

func (bin *Binder) BindArrayAccessExpression(expr nodes.ArrayAccessExpressionNode) boundnodes.BoundArrayAccessExpressionNode {
	// bind the value
	baseExpression := bin.BindExpression(expr.Base)

	// check if the variable is an array
	if baseExpression.Type().Name != "array" && baseExpression.Type().Name != "pointer" {
		print.Error(
			"BINDER",
			print.UnexpectedNonArrayValueError,
			expr.Span(),
			"Trying to Array access non-Array type (%s)", baseExpression.Type().Name,
		)
		os.Exit(-1)
	}

	// bind the index expression
	index := bin.BindExpression(expr.Index)

	// we pointin'?
	isPointer := baseExpression.Type().Name == "pointer"

	// return it as an assignment
	return boundnodes.CreateBoundArrayAccessExpressionNode(baseExpression, index, isPointer, expr.Span())
}

func (bin *Binder) BindArrayAssignmentExpression(expr nodes.ArrayAssignmentExpressionNode) boundnodes.BoundArrayAssignmentExpressionNode {
	// bind the value
	baseExpression := bin.BindExpression(expr.Base)

	// check if the variable is an array
	if baseExpression.Type().Name != "array" && baseExpression.Type().Name != "pointer" {
		print.Error(
			"BINDER",
			print.UnexpectedNonArrayValueError,
			expr.Span(),
			"Trying to Array access non-Array type (%s)", baseExpression.Type().Name,
		)
		os.Exit(-1)
	}

	// bind the index expression
	index := bin.BindExpression(expr.Index)

	// bind the value
	value := bin.BindExpression(expr.Value)

	// check if the value matches the array's type
	if value.Type().Fingerprint() != baseExpression.Type().SubTypes[0].Fingerprint() {
		print.Error(
			"BINDER",
			print.ConversionError,
			expr.Span(),
			"Array assignment types dont match! (trying to put %s into %s-Array)",
			value.Type().Name, baseExpression.Type().SubTypes[0].Name,
		)
		os.Exit(-1)
	}

	// we pointin'?
	isPointer := baseExpression.Type().Name == "pointer"

	// return it as an assignment
	return boundnodes.CreateBoundArrayAssignmentExpressionNode(baseExpression, index, value, isPointer, expr.Span())
}

func (bin *Binder) BindMakeExpression(expr nodes.MakeExpressionNode) boundnodes.BoundMakeExpressionNode {
	// this is not allowed in a class' global scope
	// because at the point in time its bound, constructors doesnt exist yet
	if bin.PreInitialTypeset != nil {
		print.Error(
			"BINDER",
			print.OutsideConstructorCallError,
			expr.Span(),
			"Constructor calls are not allowed in the global scope of a class!",
		)
		os.Exit(-1)
	}

	var baseType symbols.ClassSymbol

	// check if this refers to a package class
	if expr.Package != nil {
		// look up this package
		pack, _ := bin.LookupPackage(expr.Package.Value, false, expr.Package.Span)

		// get the class
		bType, _ := LookupClassInPackage(expr.BaseType.Value, pack, false, expr.Package.Span.SpanBetween(expr.BaseType.Span))
		baseType = bType

	} else {
		// resolve the type symbol
		bType, _ := bin.LookupClass(expr.BaseType.Value, false, expr.BaseType.Span)
		baseType = bType
	}

	// bind the constructors arguments
	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	var constructor *symbols.FunctionSymbol

	// check if class has a constructor
	for _, fnc := range baseType.Functions {
		if fnc.Name == "Constructor" {
			constructor = &fnc
			break
		}
	}

	// if there is, check if the arguments match up
	if constructor != nil {
		// make sure we got the right number of arguments
		if len(boundArguments) != len(constructor.Parameters) {
			//print.PrintCF(print.Red, "Type function '%s' expects %d arguments, got %d!", function.Name, len(function.Parameters), len(boundArguments))
			print.Error(
				"BINDER",
				print.BadNumberOfParametersError,
				expr.Span(),
				"Constructor for class \"%s\" expects %d arguments but got %d!",
				baseType.Name,
				len(constructor.Parameters),
				len(boundArguments),
			)
			os.Exit(-1)
		}

		// make sure all arguments are the right type
		for i, arg := range boundArguments {
			boundArguments[i] = bin.BindConversion(arg, constructor.Parameters[i].VarType(), false, expr.Arguments[i].Span())
		}
	} else {
		// if there is no constructor, make sure we dont have any arguments
		if len(boundArguments) != 0 {
			print.Error(
				"BINDER",
				print.BadNumberOfParametersError,
				expr.Span(),
				"Constructor for class \"%s\" expects %d arguments but got %d!",
				baseType.Name,
				0,
				len(boundArguments),
			)
			os.Exit(-1)
		}
	}

	return boundnodes.CreateBoundMakeExpressionNode(baseType, boundArguments, expr.Span())
}

func (bin *Binder) BindMakeArrayExpression(expr nodes.MakeArrayExpressionNode) boundnodes.BoundMakeArrayExpressionNode {
	// resolve the type symbol
	baseType, _ := bin.BindTypeClause(expr.Type)

	if !expr.IsLiteral {
		// bind the length expression
		length := bin.BindExpression(expr.Length)

		// return the bound node
		return boundnodes.CreateBoundMakeArrayExpressionNode(baseType, length)

	} else {
		literals := make([]boundnodes.BoundExpressionNode, 0)

		// bind all the literals
		for _, literal := range expr.LiteralValues {
			// bind the literal
			boundLiteral := bin.BindExpression(literal)

			// make sure the literal has the correct type
			convertedLiteral := bin.BindConversion(boundLiteral, baseType, false, literal.Span())

			// add the literal to the list
			literals = append(literals, convertedLiteral)
		}

		// return the bound node
		return boundnodes.CreateBoundMakeArrayExpressionNodeLiteral(baseType, literals, expr.Span())
	}
}

func (bin *Binder) BindMakeStructExpression(expr nodes.MakeStructExpressionNode) boundnodes.BoundMakeStructExpressionNode {
	// resolve the type symbol
	structType, _ := bin.LookupStruct(expr.Type.Value, false, expr.Type.Span)

	// is the count of literals we got legit?
	if len(expr.LiteralValues) > len(structType.Fields) {
		print.Error(
			"BINDER",
			print.TooManyStructParametersError,
			expr.Span(),
			"struct type %s can only hold a maximum of %d fields! (got: %d)",
			structType.Name,
			len(structType.Fields),
			len(expr.LiteralValues),
		)
		os.Exit(-1)
	}

	literals := make([]boundnodes.BoundExpressionNode, 0)

	// bind all the literals
	for i, literal := range expr.LiteralValues {
		// bind the literal
		boundLiteral := bin.BindExpression(literal)

		// make sure the literal has the correct type
		convertedLiteral := bin.BindConversion(boundLiteral, structType.Fields[i].VarType(), false, literal.Span())

		// add the literal to the list
		literals = append(literals, convertedLiteral)
	}

	// return the bound node
	return boundnodes.CreateBoundMakeStructExpressionNode(structType.Type, literals, expr.Span())
}

func (bin *Binder) BindTypeCallExpression(expr nodes.TypeCallExpressionNode) boundnodes.BoundExpressionNode {
	baseExpression := bin.BindExpression(expr.Base)

	// if the base type is a class, redirect to BindClassCallExpression
	if baseExpression.Type().IsUserDefined {
		return bin.BindClassCallExpression(expr, baseExpression)
	}

	function := bin.LookupTypeFunction(expr.CallIdentifier.Value, baseExpression.Type(), expr.CallIdentifier.Span) // Should be a string anyway
	if function.OriginType.Name != baseExpression.Type().Name {
		print.Error(
			"BINDER",
			print.IncorrectTypeFunctionCallError,
			expr.Span(),
			"the use of builtin function \"%s\" on \"%s\" datatype is undefined!",
			function.Name,
			baseExpression.Type().Name,
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
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			expr.Span(),
			"type function \"%s\" expects %d arguments but got %d!",
			function.Name,
			len(function.Parameters),
			len(boundArguments),
		)
		os.Exit(-1)
	}

	// make sure all arguments are the right type
	for i, arg := range boundArguments {
		boundArguments[i] = bin.BindConversion(arg, function.Parameters[i].VarType(), false, expr.Arguments[i].Span())
	}

	return boundnodes.CreateBoundTypeCallExpressionNode(baseExpression, function, boundArguments, expr.Span())
}

func (bin *Binder) BindClassCallExpression(expr nodes.TypeCallExpressionNode, baseExpression boundnodes.BoundExpressionNode) boundnodes.BoundExpressionNode {

	// try finding the function meant to be called
	function := bin.LookupClassFunction(expr.CallIdentifier.Value, baseExpression.Type(), expr.Base.Span().SpanBetween(expr.CallIdentifier.Span)) // Should be a string anyway

	// bind all given arguments
	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	// make sure we got the right number of arguments
	if len(boundArguments) != len(function.Parameters) {
		//print.PrintCF(print.Red, "Type function '%s' expects %d arguments, got %d!", function.Name, len(function.Parameters), len(boundArguments))
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			expr.Span(),
			"type function \"%s\" expects %d arguments but got %d!",
			function.Name,
			len(function.Parameters),
			len(boundArguments),
		)
		os.Exit(-1)
	}

	// make sure all arguments are the right type
	for i, arg := range boundArguments {
		boundArguments[i] = bin.BindConversion(arg, function.Parameters[i].VarType(), false, expr.Arguments[i].Span())
	}

	return boundnodes.CreateBoundClassCallExpressionNode(baseExpression, function, boundArguments, expr.Span())
}

func (bin *Binder) BindClassFieldAccessExpression(expr nodes.ClassFieldAccessExpressionNode) boundnodes.BoundClassFieldAccessExpressionNode {
	baseExpression := bin.BindExpression(expr.Base)

	// if the base type is a class, it cant have any fields
	if !baseExpression.Type().IsUserDefined {
		print.Error(
			"BINDER",
			print.InvalidClassAccessError,
			expr.Span(),
			"Can not use field access on non-class '%s'!",
			baseExpression.Type().Name,
		)
		os.Exit(-1)
	}

	var field symbols.VariableSymbol

	// try finding the field meant to be accessed
	// ------------------------------------------

	if baseExpression.Type().IsObject { // classes
		field = bin.LookupClassField(expr.FieldIdentifier.Value, baseExpression.Type(), expr.Base.Span().SpanBetween(expr.FieldIdentifier.Span))
	} else { // structs
		field = bin.LookupStructField(expr.FieldIdentifier.Value, baseExpression.Type(), expr.Base.Span().SpanBetween(expr.FieldIdentifier.Span))
	}

	return boundnodes.CreateBoundClassFieldAccessExpressionNode(baseExpression, field, expr.Span())
}

func (bin *Binder) BindClassFieldAssignmentExpression(expr nodes.ClassFieldAssignmentExpressionNode) boundnodes.BoundClassFieldAssignmentExpressionNode {
	baseExpression := bin.BindExpression(expr.Base)

	// if the base type is a class, it cant have any fields
	if !baseExpression.Type().IsUserDefined {
		print.Error(
			"BINDER",
			print.InvalidClassAccessError,
			expr.Span(),
			"Can not use field access on non-class '%s'!",
			baseExpression.Type().Name,
		)
		os.Exit(-1)
	}

	var field symbols.VariableSymbol

	// try finding the field meant to be accessed
	// ------------------------------------------

	if baseExpression.Type().IsObject { // classes
		field = bin.LookupClassField(expr.FieldIdentifier.Value, baseExpression.Type(), expr.Base.Span().SpanBetween(expr.FieldIdentifier.Span))
	} else { // structs
		field = bin.LookupStructField(expr.FieldIdentifier.Value, baseExpression.Type(), expr.Base.Span().SpanBetween(expr.FieldIdentifier.Span))
	}

	expression := bin.BindExpression(expr.Value)
	convertedExpression := bin.BindConversion(expression, field.VarType(), false, expr.Span())

	return boundnodes.CreateBoundClassFieldAssignmentExpressionNode(baseExpression, field, convertedExpression, expr.Span())
}

func (bin *Binder) BindCallExpression(expr nodes.CallExpressionNode) boundnodes.BoundExpressionNode {
	// check if this is a cast
	// -----------------------

	// check if it's a primitive cast
	typeSymbol, exists := LookupPrimitiveType(expr.Identifier.Value, true, expr.Identifier.Span)

	// if it worked -> create a primitive conversion
	if exists && len(expr.Arguments) == 1 {
		// bind the expression and return a conversion
		expression := bin.BindExpression(expr.Arguments[0])
		return bin.BindConversion(expression, typeSymbol, true, expr.Span())
	}

	// check if it's a class cast
	classSymbol, exists := bin.LookupClass(expr.Identifier.Value, true, expr.Identifier.Span)

	// if it worked -> create a class conversion
	if exists && len(expr.Arguments) == 1 {
		// bind the expression and return a conversion
		expression := bin.BindExpression(expr.Arguments[0])
		return bin.BindConversion(expression, classSymbol.Type, true, expr.Span())
	}

	// check if it's a complex cast
	complexTypeSymbol, exists := bin.LookupType(expr.CastingType, true)

	// if it worked -> create a complex conversion
	if exists && len(expr.Arguments) == 1 {
		// bind the expression and return a conversion
		expression := bin.BindExpression(expr.Arguments[0])
		return bin.BindConversion(expression, complexTypeSymbol, true, expr.Span())
	}

	// normal function calling
	// -----------------------

	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	symbol := bin.ActiveScope.TryLookupSymbol(expr.Identifier.Value)
	if symbol == nil ||
		symbol.SymbolType() != symbols.Function {
		print.Error(
			"BINDER",
			print.UndefinedFunctionCallError,
			expr.Span(),
			"Function \"%s\" does not exist!",
			expr.Identifier.Value,
		)
		os.Exit(-1)
	}

	functionSymbol := symbol.(symbols.FunctionSymbol)
	if len(boundArguments) != len(functionSymbol.Parameters) && !functionSymbol.Variadic {
		//fmt.Printf("%sFunction '%s' expects %d arguments, got %d!%s\n", print.ERed, functionSymbol.Name, len(functionSymbol.Parameters), len(boundArguments), print.EReset)
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			expr.Span(),
			"type function \"%s\" expects %d arguments but got %d!",
			expr.Identifier,
			len(functionSymbol.Parameters),
			len(expr.Arguments),
		)
		os.Exit(-1)
	}

	for i := 0; i < len(functionSymbol.Parameters); i++ {
		boundArguments[i] = bin.BindConversion(boundArguments[i], functionSymbol.Parameters[i].VarType(), false, expr.Arguments[i].Span())
	}

	// if we are inside a class, dont allow calls to Constructor() and Die()
	if bin.InClass && functionSymbol.Name == "Constructor" {
		print.Error(
			"BINDER",
			print.IllegalConstructorCallError,
			expr.Span(),
			"Call to Constructor in own class is not allowed!",
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundCallExpressionNode(functionSymbol, boundArguments, expr.Span())
}

func (bin *Binder) BindPackageCallExpression(expr nodes.PackageCallExpressionNode) boundnodes.BoundExpressionNode {
	// find out what package this is refering to
	pack, _ := bin.LookupPackage(expr.Package.Value, false, expr.Package.Span)

	// check if this is a cast
	// -----------------------

	// check if it's a class cast
	typeSymbol, exists := LookupClassInPackage(expr.Identifier.Value, pack, true, expr.Package.Span.SpanBetween(expr.Identifier.Span))

	// if it worked -> create a conversion
	if exists && len(expr.Arguments) == 1 {
		// bind the expression and return a conversion
		expression := bin.BindExpression(expr.Arguments[0])
		return bin.BindConversion(expression, typeSymbol.Type, true, expr.Span())
	}

	// normal function calling
	// -----------------------

	boundArguments := make([]boundnodes.BoundExpressionNode, 0)
	for _, arg := range expr.Arguments {
		boundArg := bin.BindExpression(arg)
		boundArguments = append(boundArguments, boundArg)
	}

	functionSymbol := LookupFunctionInPackage(expr.Identifier.Value, pack, expr.Package.Span.SpanBetween(expr.Identifier.Span))
	if len(boundArguments) != len(functionSymbol.Parameters) {
		//fmt.Printf("%sFunction '%s' expects %d arguments, got %d!%s\n", print.ERed, functionSymbol.Name, len(functionSymbol.Parameters), len(boundArguments), print.EReset)
		print.Error(
			"BINDER",
			print.BadNumberOfParametersError,
			expr.Span(),
			"function \"%s\" (in package \"%s\") expects %d arguments but got %d!",
			functionSymbol.Name,
			pack.Name,
			len(functionSymbol.Parameters),
			len(expr.Arguments),
		)
		os.Exit(-1)
	}

	for i := 0; i < len(boundArguments); i++ {
		boundArguments[i] = bin.BindConversion(boundArguments[i], functionSymbol.Parameters[i].VarType(), false, expr.Arguments[i].Span())
	}

	return boundnodes.CreateBoundPackageCallExpressionNode(pack, functionSymbol, boundArguments, expr.Span())
}

func (bin *Binder) BindUnaryExpression(expr nodes.UnaryExpressionNode) boundnodes.BoundUnaryExpressionNode {
	operand := bin.BindExpression(expr.Operand)
	op := boundnodes.BindUnaryOperator(expr.Operator.Kind, operand.Type())

	if !op.Exists {
		//print.PrintC(print.Red, "Unary operator '"+expr.Operator.Value+"' is not defined for type '"+operand.Type().Name+"'!")
		print.Error(
			"BINDER",
			print.UnaryOperatorTypeError,
			expr.Span(),
			"the use of unary operator \"%s\" with type \"%s\" is undefined!",
			expr.Operator.Value,
			operand.Type().Name,
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundUnaryExpressionNode(op, operand, expr.Span())
}

func (bin *Binder) BindBinaryExpression(expr nodes.BinaryExpressionNode) boundnodes.BoundBinaryExpressionNode {
	left := bin.BindExpression(expr.Left)
	right := bin.BindExpression(expr.Right)

	return bin.BindBinaryExpressionInternal(left, right, expr.Operator.Kind)
}

func (bin *Binder) BindBinaryExpressionInternal(left boundnodes.BoundExpressionNode, right boundnodes.BoundExpressionNode, opkind lexer.TokenKind) boundnodes.BoundBinaryExpressionNode {
	op := boundnodes.BindBinaryOperator(opkind, left.Type(), right.Type())

	if !op.Exists {
		// if the operation doesn't exist, see if the right side can be casted to the left
		conv := ClassifyConversion(right.Type(), left.Type())

		// if the conversion exists and its allowed to be done, do that (this also allows for explicit conversions)
		if conv.Exists && !conv.IsIdentity {
			right = bin.BindConversion(right, left.Type(), true, right.Span())
			op = boundnodes.BindBinaryOperator(opkind, left.Type(), right.Type())
		}

		// now that we may or may not have converted our right value -> check the operation again#
		if !op.Exists {
			//print.PrintC(print.Red, "Binary operator '"+expr.Operator.Value+"' is not defined for types '"+left.Type().Name+"' and '"+right.Type().Name+"'!")
			print.Error(
				"BINDER",
				print.BinaryOperatorTypeError,
				left.Span().SpanBetween(right.Span()),
				"the use of binary operator \"%s\" with types \"%s\" and \"%s\" is undefined!",
				opkind,
				left.Type().Name,
				right.Type().Name,
			)
			os.Exit(-1)
		}

	}

	return boundnodes.CreateBoundBinaryExpressionNode(left, op, right, left.Span().SpanBetween(right.Span()))
}

func (bin *Binder) BindTernaryExpression(expr nodes.TernaryExpressionNode) boundnodes.BoundTernaryExpressionNode {
	// bind condition
	condition := bin.BindExpression(expr.Condition)

	// the condition needs to be a bool!
	if condition.Type().Fingerprint() != builtins.Bool.Fingerprint() {
		print.Error(
			"BINDER",
			print.TernaryOperatorTypeError,
			expr.Condition.Span(),
			"Condition of ternary operation needs to be of type 'bool'!",
		)
		os.Exit(-1)
	}

	// bind the sides
	left := bin.BindExpression(expr.If)
	right := bin.BindExpression(expr.Else)

	// check if the left and right types are the same
	if left.Type().Fingerprint() != right.Type().Fingerprint() {
		print.Error(
			"BINDER",
			print.TernaryOperatorTypeError,
			expr.Else.Span(),
			"Types of left and right side of ternary need to match!",
		)
		os.Exit(-1)
	}

	// create a temporary variable symbol to keep track of the result
	tmp := symbols.CreateLocalVariableSymbol(symbols.GetTempName(), false, left.Type())

	return boundnodes.CreateBoundTernaryExpressionNode(condition, left, right, tmp, expr.Span())
}

func (bin *Binder) BindReferenceExpression(expr nodes.ReferenceExpressionNode) boundnodes.BoundReferenceExpressionNode {
	// bind the source variable
	variable := bin.BindNameExpression(expr.Expression)
	return boundnodes.CreateBoundReferenceExpressionNode(variable, expr.Span())
}

func (bin *Binder) BindDereferenceExpression(expr nodes.DereferenceExpressionNode) boundnodes.BoundDereferenceExpressionNode {
	// bind the source expression
	src := bin.BindExpression(expr.Expression)

	// make sure this is a pointer type
	if src.Type().Name != builtins.Pointer.Name {
		print.Error(
			"BINDER",
			print.UnexpectedNonPointerValueError,
			expr.Span(),
			"Dereferencing requires a pointer type!",
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundDereferenceExpressionNode(src, expr.Span())
}

func (bin *Binder) BindLambdaExpression(expr nodes.LambdaExpressionNode) boundnodes.BoundLambdaExpressionNode {
	boundParameters := make([]symbols.ParameterSymbol, 0)

	for i, param := range expr.Parameters {
		pName := param.Identifier.Value
		pType, _ := bin.BindTypeClause(param.TypeClause)

		// check if we've registered this param name before
		for i, p := range boundParameters {
			if p.Name == pName {
				// I haven't done bound nodes yet so I just get the syntax node parameter using the same index
				print.Error(
					"BINDER",
					print.DuplicateParameterError,
					expr.Parameters[i].Span(),
					// Kind of a hacky way of getting the values and positions needed for the error
					"a parameter with the name \"%s\" already exists for this anonymous function!",
					pName,
				)
				os.Exit(-1)
			}
		}

		boundParameters = append(boundParameters, symbols.CreateParameterSymbol(pName, i, pType))
	}

	returnType, exists := bin.BindTypeClause(expr.TypeClause)
	if !exists {
		returnType = builtins.Void
	}

	// cool symbol
	functionSymbol := symbols.CreateFunctionSymbol(symbols.GetLambdaName(), boundParameters, returnType, nodes.FunctionDeclarationMember{}, false)

	// b o i n d   f u n c t i o n
	binder := CreateBinder(MainScope, functionSymbol)
	body := binder.BindBlockStatement(expr.Body)
	loweredBody := lowerer.Lower(functionSymbol, body)

	return boundnodes.CreateBoundLambdaExpressionNode(functionSymbol, loweredBody, expr.Span())
}

func (bin *Binder) BindThisExpression(expr nodes.ThisExpressionNode) boundnodes.BoundThisExpressionNode {
	if !bin.InClass {
		// Illegal!
		print.Error(
			"BINDER",
			print.OutsideThisError,
			expr.Span(),
			"'this' expression cannot be used outside of a class!",
		)
		os.Exit(-1)
	}

	return boundnodes.CreateBoundThisExpressionNode(bin.ClassSymbol, expr.Span())
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
			id.Span,
			"Variable \"%s\" could not be declared! Variable with this name has already been declared!",
			id.Value,
		)
		os.Exit(-1)
	}

	return variable
}

func (bin *Binder) BindVariableReference(name string, errorLocation print.TextSpan) symbols.VariableSymbol {
	variable := bin.ActiveScope.TryLookupSymbol(name)

	if variable == nil ||
		!(variable.SymbolType() == symbols.GlobalVariable ||
			variable.SymbolType() == symbols.LocalVariable ||
			variable.SymbolType() == symbols.Parameter) {
		//print.PrintC(print.Red, "Could not find variable '"+name+"'!")
		print.Error(
			"BINDER",
			print.UndefinedVariableReferenceError,
			errorLocation,
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

	// if a package is given
	if tc.Package != nil {
		// look up this package
		pack, _ := bin.LookupPackage(tc.Package.Value, false, tc.Package.Span)

		// get the class
		bType, _ := LookupClassInPackage(tc.TypeIdentifier.Value, pack, false, tc.Span())
		return bType.Type, true
	}

	typ, _ := bin.LookupType(tc, false)
	return typ, true
}

func (bin *Binder) LookupTypeFunction(name string, baseType symbols.TypeSymbol, errorLocation print.TextSpan) symbols.TypeFunctionSymbol {
	switch name {
	case "GetLength":
		if baseType.Fingerprint() == builtins.String.Fingerprint() {
			// string length
			return builtins.GetLength
		} else {
			// array length
			return builtins.GetArrayLength
		}
	case "GetBuffer":
		return builtins.GetBuffer
	case "Substring":
		return builtins.Substring
	case "Push":
		if baseType.Name == builtins.Array.Name {
			if baseType.SubTypes[0].IsObject {
				// push function for object arrays
				return builtins.Push
			} else {
				// push function for primitive arrays
				return builtins.PPush
			}
		}
	case "Start":
		return builtins.Start
	case "Join":
		return builtins.Join
	case "Kill":
		return builtins.Kill
	case "Run":
		// oh boy
		sym := builtins.Run

		// return type is equal to the last subtype
		sym.Type = baseType.SubTypes[len(baseType.SubTypes)-1]

		// add some quirky params
		for i, symbol := range baseType.SubTypes[:len(baseType.SubTypes)-1] {
			sym.Parameters = append(sym.Parameters, symbols.CreateParameterSymbol(fmt.Sprintf("prm_%d", i), i, symbol))
		}

		// we constructed a very nice type func :)
		return sym

	case "RunThread":
		// recycling :)
		sym := builtins.RunThread

		// add some quirky params
		for i, symbol := range baseType.SubTypes[:len(baseType.SubTypes)-1] {
			sym.Parameters = append(sym.Parameters, symbols.CreateParameterSymbol(fmt.Sprintf("prm_%d", i), i, symbol))
		}

		// we constructed a very nice type func :)
		return sym

	default:
		/*print.PrintC(
			print.Red,
			fmt.Sprintf("Could not find builtin TypeFunctionSymbol \"%s\"!", name),
		)*/
		print.Error(
			"BINDER",
			print.TypeFunctionDoesNotExistError,
			errorLocation,
			"Could not find builtin TypeFunctionSymbol \"%s\"!",
			name,
		)
		os.Exit(-1)
	}
	return symbols.TypeFunctionSymbol{}
}

func (bin *Binder) LookupClassFunction(name string, baseType symbols.TypeSymbol, errorLocation print.TextSpan) symbols.FunctionSymbol {
	// try locating the class
	clsSym := bin.ActiveScope.TryLookupSymbol(baseType.Name)

	// if that failed -> look through packages
	if clsSym == nil || clsSym.SymbolType() != symbols.Class {
		for _, pck := range bin.ActiveScope.GetAllPackages() {
			for _, cls := range pck.Classes {
				if cls.Name == baseType.Name {
					clsSym = cls
				}
			}
		}
	}

	// if that failed -> throw an error
	if clsSym == nil || clsSym.SymbolType() != symbols.Class {
		print.Error(
			"BINDER",
			print.UnknownClassError,
			errorLocation,
			"Could not find class \"%s\" in lookup, did something not load correctly?",
			baseType.Name,
		)
		os.Exit(-1)
	}

	// get the symbol as a class symbol
	cls := clsSym.(symbols.ClassSymbol)

	// search through all the class' functions to find the one we're looking for
	for _, fnc := range cls.Functions {
		if fnc.Name == name {
			if !fnc.Public {
				print.Error(
					"BINDER",
					print.FunctionAccessViolationError,
					errorLocation,
					"Function \"%s\" in class \"%s\" is not accessable, is the function intended to be private?",
					name,
					baseType.Name,
				)
				os.Exit(-1)
			}

			return fnc
		}
	}

	print.Error(
		"BINDER",
		print.TypeFunctionDoesNotExistError,
		errorLocation,
		"Could not find function \"%s\" in class \"%s\", does the function exist?",
		name,
		baseType.Name,
	)
	os.Exit(-1)

	return symbols.FunctionSymbol{}
}

func (bin *Binder) LookupClassField(name string, baseType symbols.TypeSymbol, errorLocation print.TextSpan) symbols.VariableSymbol {
	// try locating the class
	clsSym := bin.ActiveScope.TryLookupSymbol(baseType.Name)

	// if that failed -> look through packages
	if clsSym == nil || clsSym.SymbolType() != symbols.Class {
		for _, pck := range bin.ActiveScope.GetAllPackages() {
			for _, cls := range pck.Classes {
				if cls.Name == baseType.Name {
					clsSym = cls
				}
			}
		}
	}

	// if that failed -> throw an error
	if clsSym == nil || clsSym.SymbolType() != symbols.Class {
		print.Error(
			"BINDER",
			print.UnknownClassError,
			errorLocation,
			"Could not find class \"%s\" in lookup, did something not load correctly?",
			baseType.Name,
		)
		os.Exit(-1)
	}

	// get the symbol as a class symbol
	cls := clsSym.(symbols.ClassSymbol)

	// search through all the class' functions to find the one we're looking for
	for _, fld := range cls.Fields {
		if fld.SymbolName() == name {
			return fld
		}
	}

	print.Error(
		"BINDER",
		print.UnknownFieldError,
		errorLocation,
		"Could not find field \"%s\" in class \"%s\", does the field exist?",
		name,
		baseType.Name,
	)
	os.Exit(-1)

	return symbols.LocalVariableSymbol{}
}

func (bin *Binder) LookupStructField(name string, baseType symbols.TypeSymbol, errorLocation print.TextSpan) symbols.VariableSymbol {
	// try locating the class
	stcSym := bin.ActiveScope.TryLookupSymbol(baseType.Name)

	// if that failed -> throw an error
	if stcSym == nil || stcSym.SymbolType() != symbols.Struct {
		print.Error(
			"BINDER",
			print.UnknownStructError,
			errorLocation,
			"Could not find struct \"%s\" in lookup, did something not load correctly?",
			baseType.Name,
		)
		os.Exit(-1)
	}

	// get the symbol as a class symbol
	stc := stcSym.(symbols.StructSymbol)

	// search through all the class' functions to find the one we're looking for
	for _, fld := range stc.Fields {
		if fld.SymbolName() == name {
			return fld
		}
	}

	print.Error(
		"BINDER",
		print.UnknownFieldError,
		errorLocation,
		"Could not find field \"%s\" in struct \"%s\", does the field exist?",
		name,
		baseType.Name,
	)
	os.Exit(-1)

	return symbols.LocalVariableSymbol{}
}

// </IDEK> --------------------------------------------------------------------
// <TYPES> --------------------------------------------------------------------

func (bin *Binder) BindConversion(expr boundnodes.BoundExpressionNode, to symbols.TypeSymbol, allowExplicit bool, errorLocation print.TextSpan) boundnodes.BoundExpressionNode {
	conversionType := ClassifyConversion(expr.Type(), to)

	nameFrom := expr.Type().Name
	nameTo := to.Name

	// if both names are the same -> show the fingerprint instead
	if nameFrom == nameTo {
		nameFrom = expr.Type().Fingerprint()
		nameTo = to.Fingerprint()
	}

	if !conversionType.Exists {
		//print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'!")
		print.Error(
			"BINDER",
			print.ConversionError,
			errorLocation,
			"Cannot convert type \"%s\" to \"%s\"!",
			nameFrom,
			nameTo,
		)
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsExplicit && !allowExplicit {
		//print.PrintC(print.Red, "Cannot convert type '"+expr.Type().Name+"' to '"+to.Name+"'! (An explicit conversion exists. Are you missing a cast?)")
		print.Error(
			"BINDER",
			print.ExplicitConversionError,
			errorLocation,
			"Cannot convert type \"%s\" to \"%s\"! (An explicit conversion exists. Are you missing a cast?)",
			nameFrom,
			nameTo,
		)
		os.Exit(-1)
		return boundnodes.BoundErrorExpressionNode{}
	}

	if conversionType.IsIdentity {
		return expr
	}

	return boundnodes.CreateBoundConversionExpressionNode(to, expr, expr.Span())
}

func (bin Binder) LookupType(typeClause nodes.TypeClauseNode, canFail bool) (symbols.TypeSymbol, bool) {
	switch typeClause.TypeIdentifier.Value {
	case "void":
		return builtins.Void, true
	case "bool":
		return builtins.Bool, true
	case "byte":
		return builtins.Byte, true
	case "int":
		return builtins.Int, true
	case "long":
		return builtins.Long, true
	case "float":
		return builtins.Float, true
	case "uint":
		return builtins.UInt, true
	case "ulong":
		return builtins.ULong, true
	case "double":
		return builtins.Double, true
	case "string":
		return builtins.String, true
	case "any":
		return builtins.Any, true
	case "array":
		if len(typeClause.SubClauses) != 1 {
			print.Error(
				"BINDER",
				print.InvalidNumberOfSubtypesError,
				typeClause.Span(),
				"Datatype \"%s\" takes in exactly one subtype!",
				typeClause.TypeIdentifier.Value,
			)
			os.Exit(-1)
		}

		baseType, _ := bin.LookupType(typeClause.SubClauses[0], false)
		return symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{baseType}, true, false), true

	case "pointer":
		if len(typeClause.SubClauses) != 1 {
			print.Error(
				"BINDER",
				print.InvalidNumberOfSubtypesError,
				typeClause.Span(),
				"Datatype \"%s\" takes in exactly one subtype!",
				typeClause.TypeIdentifier.Value,
			)
			os.Exit(-1)
		}

		baseType, _ := bin.LookupType(typeClause.SubClauses[0], false)
		return symbols.CreateTypeSymbol("pointer", []symbols.TypeSymbol{baseType}, false, false), true

	case "action":
		if len(typeClause.SubClauses) == 0 {
			print.Error(
				"BINDER",
				print.InvalidNumberOfSubtypesError,
				typeClause.Span(),
				"Datatype \"%s\" takes in at least one subtype!",
				typeClause.TypeIdentifier.Value,
			)
			os.Exit(-1)
		}

		subTypes := make([]symbols.TypeSymbol, 0)

		for _, clause := range typeClause.SubClauses {
			typ, _ := bin.LookupType(clause, false)
			subTypes = append(subTypes, typ)
		}

		return symbols.CreateTypeSymbol("action", subTypes, false, false), true

	default:
		// check if this might be a class
		cls, ok := bin.LookupClass(typeClause.TypeIdentifier.Value, true, typeClause.TypeIdentifier.Span)
		if ok {
			return cls.Type, true
		}

		// check if this might be a struct
		stc, ok := bin.LookupStruct(typeClause.TypeIdentifier.Value, true, typeClause.TypeIdentifier.Span)
		if ok {
			return stc.Type, true
		}

		// check if this binder has been given a pre-initial typeset
		if bin.PreInitialTypeset != nil {

			// if so, use it as a source for type symbols
			// this is done because at this point, no classes are officially registered yet
			// to "kickstart" the type resolving process
			for _, piType := range bin.PreInitialTypeset {
				if piType.Name == typeClause.TypeIdentifier.Value {
					return piType, true
				}
			}
		}

		// otherwise, die()
		if !canFail {
			print.Error(
				"BINDER",
				print.UnknownDataTypeError,
				typeClause.Span(),
				"Couldn't find datatype \"%s\"! Are you sure it exists?",
				typeClause.TypeIdentifier.Value,
			)
			os.Exit(-1)
		}

		return symbols.TypeSymbol{}, false
	}
}

func LookupPrimitiveType(name string, canFail bool, errorLocation print.TextSpan) (symbols.TypeSymbol, bool) {
	switch name {
	case "void":
		return builtins.Void, true
	case "bool":
		return builtins.Bool, true
	case "byte":
		return builtins.Byte, true
	case "int":
		return builtins.Int, true
	case "long":
		return builtins.Long, true
	case "float":
		return builtins.Float, true
	case "uint":
		return builtins.UInt, true
	case "ulong":
		return builtins.ULong, true
	case "double":
		return builtins.Double, true
	case "string":
		return builtins.String, true
	case "any":
		return builtins.Any, true
	default:
		if !canFail {
			//print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
			print.Error(
				"BINDER",
				print.UnknownDataTypeError,
				errorLocation,
				"Couldn't find primitive datatype \"%s\"! Are you sure it exists?",
				name,
			)
			os.Exit(-1)
		}

		return symbols.TypeSymbol{}, false
	}
}

func (bin Binder) LookupClass(name string, canFail bool, errorLocaton print.TextSpan) (symbols.ClassSymbol, bool) {
	cls := bin.ActiveScope.TryLookupSymbol(name)
	if cls == nil {
		return FailClassLookup(name, canFail, errorLocaton)
	}

	if cls.SymbolType() != symbols.Class {
		return FailClassLookup(name, canFail, errorLocaton)
	}

	return cls.(symbols.ClassSymbol), true
}

func (bin Binder) LookupStruct(name string, canFail bool, errorLocaton print.TextSpan) (symbols.StructSymbol, bool) {

	stc := bin.ActiveScope.TryLookupSymbol(name)
	if stc == nil {
		return FailStructLookup(name, canFail, errorLocaton)
	}

	if stc.SymbolType() != symbols.Struct {
		return FailStructLookup(name, canFail, errorLocaton)
	}

	return stc.(symbols.StructSymbol), true
}

func (bin Binder) LookupPackage(name string, canFail bool, errorLocaton print.TextSpan) (symbols.PackageSymbol, bool) {
	pck := bin.ActiveScope.TryLookupSymbol(name)
	if pck == nil {
		return FailPackageLookup(name, canFail, errorLocaton)
	}

	if pck.SymbolType() != symbols.Package {
		return FailPackageLookup(name, canFail, errorLocaton)
	}

	return pck.(symbols.PackageSymbol), true
}

func LookupClassInPackage(name string, pack symbols.PackageSymbol, canFail bool, errorLocation print.TextSpan) (symbols.ClassSymbol, bool) {
	for _, cls := range pack.Classes {
		if cls.Name == name {
			return cls, true
		}
	}

	if !canFail {
		print.Error(
			"BINDER",
			print.UnknownClassError,
			errorLocation,
			"Couldn't find class \"%s\" in package \"%s\"! Are you sure it exists?",
			name,
			pack.Name,
		)
		os.Exit(-1)
	}

	return symbols.ClassSymbol{}, false
}

func LookupFunctionInPackage(name string, pack symbols.PackageSymbol, errorLocation print.TextSpan) symbols.FunctionSymbol {
	for _, fnc := range pack.Functions {
		if fnc.Name == name {
			return fnc
		}
	}

	print.Error(
		"BINDER",
		print.UndefinedFunctionCallError,
		errorLocation,
		"Couldn't find function \"%s\" in package \"%s\"! Are you sure it exists?",
		name,
		pack.Name,
	)
	os.Exit(-1)

	return symbols.FunctionSymbol{}
}

func FailClassLookup(name string, canFail bool, errorLocation print.TextSpan) (symbols.ClassSymbol, bool) {
	if !canFail {
		//print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
		print.Error(
			"BINDER",
			print.UnknownClassError,
			errorLocation,
			"Couldn't find class \"%s\"! Are you sure it exists?",
			name,
		)
		os.Exit(-1)
	}

	return symbols.ClassSymbol{}, false
}

func FailStructLookup(name string, canFail bool, errorLocation print.TextSpan) (symbols.StructSymbol, bool) {
	if !canFail {
		//print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
		print.Error(
			"BINDER",
			print.UnknownClassError,
			errorLocation,
			"Couldn't find struct \"%s\"! Are you sure it exists?",
			name,
		)
		os.Exit(-1)
	}

	return symbols.StructSymbol{}, false
}

func FailPackageLookup(name string, canFail bool, errorLocation print.TextSpan) (symbols.PackageSymbol, bool) {
	if !canFail {
		//print.PrintC(print.Red, "Couldnt find Datatype '"+name+"'!")
		print.Error(
			"BINDER",
			print.UnknownPackageError,
			errorLocation,
			"Couldn't find package \"%s\"! Are you sure it was imported?",
			name,
		)
		os.Exit(-1)
	}

	return symbols.PackageSymbol{}, false
}

// </TYPES> -------------------------------------------------------------------
