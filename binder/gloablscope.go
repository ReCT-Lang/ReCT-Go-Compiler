package binder

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/builtins"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes/boundnodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

type GlobalScope struct {
	MainFunction symbols.FunctionSymbol

	Functions  []symbols.FunctionSymbol
	Variables  []symbols.VariableSymbol
	Classes    []symbols.ClassSymbol
	Structs    []symbols.StructSymbol
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
	packageAliases := make([]nodes.PackageAliasMember, 0)
	packageUses := make([]nodes.PackageUseMember, 0)
	functionDeclarations := make([]nodes.FunctionDeclarationMember, 0)
	externalFunctionDeclarations := make([]nodes.ExternalFunctionDeclarationMember, 0)
	classDeclarations := make([]nodes.ClassDeclarationMember, 0)
	structDeclarations := make([]nodes.StructDeclarationMember, 0)
	enumDeclarations := make([]nodes.EnumDeclarationMember, 0)
	globalStatements := make([]nodes.GlobalStatementMember, 0)

	// sort all our members into functions and global statements
	for _, member := range members {
		if member.NodeType() == nodes.FunctionDeclaration {
			functionDeclarations = append(functionDeclarations, member.(nodes.FunctionDeclarationMember))
		} else if member.NodeType() == nodes.ExternalFunctionDeclaration {
			externalFunctionDeclarations = append(externalFunctionDeclarations, member.(nodes.ExternalFunctionDeclarationMember))
		} else if member.NodeType() == nodes.ClassDeclaration {
			classDeclarations = append(classDeclarations, member.(nodes.ClassDeclarationMember))
		} else if member.NodeType() == nodes.StructDeclaration {
			structDeclarations = append(structDeclarations, member.(nodes.StructDeclarationMember))
		} else if member.NodeType() == nodes.EnumDeclaration {
			enumDeclarations = append(enumDeclarations, member.(nodes.EnumDeclarationMember))
		} else if member.NodeType() == nodes.PackageReference {
			packageReferences = append(packageReferences, member.(nodes.PackageReferenceMember))
		} else if member.NodeType() == nodes.PackageAlias {
			packageAliases = append(packageAliases, member.(nodes.PackageAliasMember))
		} else if member.NodeType() == nodes.PackageUse {
			packageUses = append(packageUses, member.(nodes.PackageUseMember))
		} else {
			globalStatements = append(globalStatements, member.(nodes.GlobalStatementMember))
		}
	}

	binder := CreateBinder(mainScope, symbols.FunctionSymbol{})

	// first load all enums, this is cool because it doenst require anything else to be set up first
	for _, enm := range enumDeclarations {
		binder.BindEnumDeclaration(enm)
	}

	// this is now the global parent scope node
	MainScope = *binder.ActiveScope

	// Actual cool binding
	// -------------------

	// load all our packages first
	for _, pkg := range packageReferences {
		binder.BindPackageReference(pkg)
	}

	for _, pkg := range packageAliases {
		binder.BindPackageAlias(pkg)
	}

	for _, pkg := range packageUses {
		binder.BindPackageUse(pkg)
	}

	// create a loose list of types so our class and struct members have something to work with
	preInitialTypeset := make([]symbols.TypeSymbol, 0)
	for _, cls := range classDeclarations {
		preInitialTypeset = append(preInitialTypeset, symbols.CreateTypeSymbol(cls.Identifier.Value, make([]symbols.TypeSymbol, 0), true, true, false, symbols.PackageSymbol{}, nil))
	}
	for _, stc := range structDeclarations {
		preInitialTypeset = append(preInitialTypeset, symbols.CreateTypeSymbol(stc.Identifier.Value, make([]symbols.TypeSymbol, 0), false, true, false, symbols.PackageSymbol{}, nil))
	}
	for _, enm := range enumDeclarations {
		preInitialTypeset = append(preInitialTypeset, symbols.CreateTypeSymbol(enm.Identifier.Value, make([]symbols.TypeSymbol, 0), false, false, true, symbols.PackageSymbol{}, nil))
	}

	// declare all our structs and their fields
	for _, stc := range structDeclarations {
		binder.BindStructDeclaration(stc, preInitialTypeset)
	}

	// declare all our classes and their members
	for _, cls := range classDeclarations {
		binder.BindClassDeclaration(cls, preInitialTypeset)
	}

	// declare all our functions
	for _, fnc := range functionDeclarations {
		binder.BindFunctionDeclaration(fnc, false)
	}

	// declare all our external functions
	for _, fnc := range externalFunctionDeclarations {
		binder.BindExternalFunctionDeclaration(fnc, false)
	}

	// this is now the global parent scope node
	MainScope = *binder.ActiveScope

	// bind all our statements
	boundStatements := make([]boundnodes.BoundStatementNode, 0)
	for _, stmt := range globalStatements {
		boundStatements = append(boundStatements, binder.BindStatement(stmt.Statement))
	}

	// this is now the even globaler parent scope node (updated)
	MainScope = *binder.ActiveScope

	return GlobalScope{
		MainFunction: symbols.CreateFunctionSymbol("main", make([]symbols.ParameterSymbol, 0), builtins.Void, nodes.FunctionDeclarationMember{}, true),
		Functions:    binder.ActiveScope.GetAllFunctions(),
		Variables:    binder.ActiveScope.GetAllVariables(),
		Classes:      binder.ActiveScope.GetAllClasses(),
		Structs:      binder.ActiveScope.GetAllStructs(),
		Packages:     binder.ActiveScope.GetAllPackages(),
		Statements:   boundStatements,
	}
}

func BindParentScope(globalScope GlobalScope) Scope {
	parent := BindRootScope()
	workingScope := CreateScope(&parent)

	for _, pck := range globalScope.Packages {
		workingScope.TryDeclareSymbol(pck)
	}

	for _, cls := range globalScope.Classes {
		workingScope.TryDeclareSymbol(cls)
	}

	for _, stc := range globalScope.Structs {
		workingScope.TryDeclareSymbol(stc)
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
