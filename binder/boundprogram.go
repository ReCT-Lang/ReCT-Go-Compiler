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
	GlobalScope  *GlobalScope
	MainFunction symbols.FunctionSymbol
	Functions    []BoundFunction
	Classes      []BoundClass
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
	classes := make([]BoundClass, 0)

	mainBody := boundnodes.CreateBoundBlockStatementNode(globalScope.Statements)
	loweredMainBody := lowerer.Lower(globalScope.MainFunction, mainBody)
	functionBodies = append(functionBodies, BoundFunction{
		Symbol: globalScope.MainFunction,
		Body:   loweredMainBody,
	})

	for _, fnc := range globalScope.Functions {
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

		for _, fnc := range cls.Functions {
			binder := CreateBinder(classScope, fnc)
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
		GlobalScope:  &globalScope,
		MainFunction: globalScope.MainFunction,
		Functions:    functionBodies,
		Classes:      classes,
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
