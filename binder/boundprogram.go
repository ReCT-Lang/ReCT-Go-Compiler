package binder

import (
	"ReCT-Go-Compiler/lowerer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundProgram struct {
	GlobalScope       *GlobalScope
	MainFunction      symbols.FunctionSymbol
	Functions         []BoundFunction
	ExternalFunctions []symbols.FunctionSymbol
	Classes           []BoundClass
	Packages          []symbols.PackageSymbol
}

type BoundFunction struct {
	Symbol symbols.FunctionSymbol
	Body   boundnodes.BoundBlockStatementNode
}

type BoundClass struct {
	Symbol    symbols.ClassSymbol
	Functions []BoundFunction
}

func BindProgram(members []nodes.MemberNode) BoundProgram {
	globalScope := BindGlobalScope(members)
	parentScope := BindParentScope(globalScope)
	functionBodies := make([]BoundFunction, 0)
	functionReferences := make([]symbols.FunctionSymbol, 0)
	classes := make([]BoundClass, 0)

	mainBody := boundnodes.CreateBoundBlockStatementNode(globalScope.Statements, print.TextSpan{})
	loweredMainBody := lowerer.Lower(globalScope.MainFunction, mainBody)
	functionBodies = append(functionBodies, BoundFunction{
		Symbol: globalScope.MainFunction,
		Body:   loweredMainBody,
	})

	for _, fnc := range globalScope.Functions {
		// skip external functions
		if fnc.External {
			functionReferences = append(functionReferences, fnc)
			continue
		}

		binder := CreateBinder(parentScope, fnc)
		body := binder.BindBlockStatement(fnc.Declaration.Body)
		loweredBody := lowerer.Lower(fnc, body)

		functionBodies = append(functionBodies, BoundFunction{
			Symbol: fnc,
			Body:   loweredBody,
		})
	}

	for _, cls := range globalScope.Classes {
		classScope := BindParentScope(globalScope)
		classScope.InsertVariableSymbols(cls.Fields)
		classScope.InsertFunctionSymbols(cls.Functions)

		classFunctionBodies := make([]BoundFunction, 0)

		constructorNeedsInjection := false
		constructorInjection := make([]nodes.StatementNode, 0)

		// assemble an injection for the constructor to initialize fields
		for _, mem := range cls.Declaration.Members {
			if mem.NodeType() != nodes.GlobalStatement {
				continue
			}

			// all of these should be variable declarations
			// anything else should have been filtered out by the binder
			declaration := mem.(nodes.GlobalStatementMember).Statement.(nodes.VariableDeclarationStatementNode)
			if declaration.Initializer != nil {
				constructorNeedsInjection = true

				// find the field associated with this declaration
				for _, fld := range cls.Fields {

					// when found, fabricate an assignment expression
					if fld.SymbolName() == declaration.Identifier.Value {
						constructorInjection = append(constructorInjection,
							nodes.CreateExpressionStatementNode(
								nodes.CreateAssignmentExpressionNode(declaration.Identifier, declaration.Initializer),
							),
						)
					}
				}

			}
		}

		for _, fnc := range cls.Functions {
			// if this class has fields with assignments -> put those in the constructor
			if constructorNeedsInjection && fnc.Name == "Constructor" {
				fnc.Declaration.Body.Statements = append(constructorInjection, fnc.Declaration.Body.Statements...)
			}

			binder := CreateBinder(classScope, fnc)
			binder.InClass = true
			body := binder.BindBlockStatement(fnc.Declaration.Body)
			loweredBody := lowerer.Lower(fnc, body)

			classFunctionBodies = append(classFunctionBodies, BoundFunction{
				Symbol: fnc,
				Body:   loweredBody,
			})
		}

		classes = append(classes, BoundClass{
			Symbol:    cls,
			Functions: classFunctionBodies,
		})
	}

	return BoundProgram{
		GlobalScope:       &globalScope,
		MainFunction:      globalScope.MainFunction,
		Functions:         functionBodies,
		ExternalFunctions: functionReferences,
		Classes:           classes,
		Packages:          globalScope.Packages,
	}
}

func (b *BoundProgram) Print() {
	print.PrintC(print.Red, ":Main Function")
	b.MainFunction.Print("  ")

	print.PrintC(print.Red, ":Functions")
	for _, fnc := range b.Functions {
		fnc.Symbol.Print("  ")
		if !fnc.Symbol.BuiltIn {
			fmt.Println("  └ Function Body:")
			fnc.Body.Print("    ")
		}
	}

	print.PrintC(print.Red, ":Classes")
	for _, cls := range b.Classes {
		cls.Symbol.Print("  ")

		print.PrintC(print.Red, "  :Functions")
		for _, fnc := range cls.Functions {
			if !fnc.Symbol.BuiltIn {
				fnc.Symbol.Print("    ")
				fmt.Println("      └ Function Body:")
				fnc.Body.Print("      ")
			}
		}
	}
}

func (b *BoundProgram) PrintStatements() {
	print.PrintC(print.Red, ":Main Function")
	b.MainFunction.Print("  ")

	print.PrintC(print.Red, ":Functions")
	for _, fnc := range b.Functions {
		fnc.Symbol.Print("  ")
		if !fnc.Symbol.BuiltIn {
			fmt.Println("  └ Function Body:")

			for _, stmt := range fnc.Body.Statements {
				if stmt.NodeType() == boundnodes.BoundGarbageCollectionStatement {
					stmt.Print("    ")
				} else {
					fmt.Println("    └ " + stmt.NodeType())
				}
			}
		}
	}
}
