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
	Classes    []symbols.ClassSymbol
	Packages   []symbols.PackageSymbol
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

	packageReferences := make([]nodes.PackageReferenceMember, 0)
	functionDeclarations := make([]nodes.FunctionDeclarationMember, 0)
	classDeclarations := make([]nodes.ClassDeclarationMember, 0)
	globalStatements := make([]nodes.GlobalStatementMember, 0)

	// sort all our members into functions and global statements
	for _, member := range members {
		if member.NodeType() == nodes.FunctionDeclaration {
			functionDeclarations = append(functionDeclarations, member.(nodes.FunctionDeclarationMember))
		} else if member.NodeType() == nodes.ClassDeclaration {
			classDeclarations = append(classDeclarations, member.(nodes.ClassDeclarationMember))
		} else if member.NodeType() == nodes.PackageReference {
			packageReferences = append(packageReferences, member.(nodes.PackageReferenceMember))
		} else {
			globalStatements = append(globalStatements, member.(nodes.GlobalStatementMember))
		}
	}

	binder := CreateBinder(mainScope, symbols.FunctionSymbol{})

	// load all our packages first
	for _, pkg := range packageReferences {
		binder.BindPackageReference(pkg)
	}

	// create a loose list of types so our class members have something to work with
	preInitialTypeset := make([]symbols.TypeSymbol, 0)
	for _, cls := range classDeclarations {
		preInitialTypeset = append(preInitialTypeset, symbols.CreateTypeSymbol(cls.Identifier.Value, make([]symbols.TypeSymbol, 0), true, true))
	}

	// declare all our classes and their members
	for _, cls := range classDeclarations {
		binder.BindClassDeclaration(cls, preInitialTypeset)
	}

	// declare all our functions
	for _, fnc := range functionDeclarations {
		binder.BindFunctionDeclaration(fnc, false)
	}

	// bind all our statements
	boundStatements := make([]boundnodes.BoundStatementNode, 0)
	for _, stmt := range globalStatements {
		boundStatements = append(boundStatements, binder.BindStatement(stmt.Statement))
	}

	return GlobalScope{
		MainFunction: symbols.CreateFunctionSymbol("main", make([]symbols.ParameterSymbol, 0), builtins.Void, nodes.FunctionDeclarationMember{}, true),
		Functions:    binder.ActiveScope.GetAllFunctions(),
		Variables:    binder.ActiveScope.GetAllVariables(),
		Classes:      binder.ActiveScope.GetAllClasses(),
		Packages:     binder.ActiveScope.GetAllPackages(),
		Statements:   boundStatements,
	}
}

func BindParentScope(globalScope GlobalScope) Scope {
	parent := BindRootScope()
	workingScope := CreateScope(&parent)

	for _, cls := range globalScope.Classes {
		workingScope.TryDeclareSymbol(cls)
	}

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

	// no more builtins which could go here...

	return scope
}
