package binder

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
)

type GlobalScope struct {
	MainFunction symbols.FunctionSymbol

	Functions  []symbols.FunctionSymbol
	Variables  []symbols.VariableSymbol
	Statements []boundnodes.BoundStatementNode
}

func (g GlobalScope) Print() {
	print.PrintC(print.Red, ":Main Function")
	g.MainFunction.Print("  ")

	print.PrintC(print.Red, ":Functions")
	for _, fnc := range g.Functions {
		fnc.Print("  ")
	}

	print.PrintC(print.Red, ":Variables")
	for _, variable := range g.Variables {
		variable.Print("  ")
	}

	print.PrintC(print.Red, ":Global Statements")
	for _, stmt := range g.Statements {
		stmt.Print("  ")
	}
}

func BindGlobalScope(members []nodes.MemberNode) GlobalScope {
	rootScope := BindRootScope()
	mainScope := CreateScope(&rootScope)

	functionDeclarations := make([]nodes.FunctionDeclarationMember, 0)
	globalStatements := make([]nodes.GlobalStatementMember, 0)

	// sort all our members into functions and global statements
	for _, member := range members {
		if member.NodeType() == nodes.FunctionDeclaration {
			functionDeclarations = append(functionDeclarations, member.(nodes.FunctionDeclarationMember))
		} else {
			globalStatements = append(globalStatements, member.(nodes.GlobalStatementMember))
		}
	}

	binder := CreateBinder(mainScope, symbols.FunctionSymbol{})

	// declare all our functions
	for _, fnc := range functionDeclarations {
		binder.BindFunctionDeclaration(fnc)
	}

	// bind all our statements
	boundStatements := make([]boundnodes.BoundStatementNode, 0)
	for _, stmt := range globalStatements {
		boundStatements = append(boundStatements, binder.BindStatement(stmt.Statement))
	}

	return GlobalScope{
		MainFunction: symbols.CreateFunctionSymbol("main", make([]symbols.ParameterSymbol, 0), builtins.Void, nodes.FunctionDeclarationMember{}),
		Functions:    binder.ActiveScope.GetAllFunctions(),
		Variables:    binder.ActiveScope.GetAllVariables(),
		Statements:   boundStatements,
	}
}

func BindParentScope(globalScope GlobalScope) Scope {
	parent := BindRootScope()
	workingScope := CreateScope(&parent)

	for _, fnc := range globalScope.Functions {
		workingScope.TryDeclareSymbol(fnc)
	}

	for _, variable := range globalScope.Variables {
		// only bring over global vars
		if !variable.IsGlobal() {
			continue
		}

		workingScope.TryDeclareSymbol(variable)
	}

	return workingScope
}

func BindRootScope() Scope {
	scope := CreateScope(nil)

	scope.TryDeclareSymbol(builtins.Print)
	scope.TryDeclareSymbol(builtins.PrintfI32)
	scope.TryDeclareSymbol(builtins.Write)
	scope.TryDeclareSymbol(builtins.Input)
	scope.TryDeclareSymbol(builtins.InputKey)
	scope.TryDeclareSymbol(builtins.Clear)
	scope.TryDeclareSymbol(builtins.SetCursor)
	scope.TryDeclareSymbol(builtins.GetSizeX)
	scope.TryDeclareSymbol(builtins.GetSizeY)
	scope.TryDeclareSymbol(builtins.SetCursorVisible)
	scope.TryDeclareSymbol(builtins.GetCursorVisible)
	scope.TryDeclareSymbol(builtins.Random)
	scope.TryDeclareSymbol(builtins.Sleep)
	scope.TryDeclareSymbol(builtins.Version)

	return scope
}
