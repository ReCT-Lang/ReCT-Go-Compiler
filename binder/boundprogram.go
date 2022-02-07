package binder

import (
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
}

type BoundFunction struct {
	Symbol symbols.FunctionSymbol
	Body   boundnodes.BoundBlockStatementNode
}

func BindProgram(members []nodes.MemberNode) BoundProgram {
	globalScope := BindGlobalScope(members)
	parentScope := BindParentScope(globalScope)
	functionBodies := make([]BoundFunction, 0)

	mainBody := boundnodes.CreateBoundBlockStatementNode(globalScope.Statements)
	functionBodies = append(functionBodies, BoundFunction{
		Symbol: globalScope.MainFunction,
		Body:   mainBody,
	})

	for _, fnc := range globalScope.Functions {
		binder := CreateBinder(parentScope, fnc)
		body := binder.BindBlockStatement(fnc.Declaration.Body)

		functionBodies = append(functionBodies, BoundFunction{
			Symbol: fnc,
			Body:   body,
		})
	}

	return BoundProgram{
		GlobalScope:  &globalScope,
		MainFunction: globalScope.MainFunction,
		Functions:    functionBodies,
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
}
