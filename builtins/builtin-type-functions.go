package builtins

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/symbols"
)

var (
	GetLength = symbols.CreateBuiltInTypeFunctionSymbol(
		"GetLength",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
		String,
	)

	Substring = symbols.CreateBuiltInTypeFunctionSymbol(
		"Substring",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("startingIndex", 0, Int),
			symbols.CreateParameterSymbol("length", 1, Int),
		},
		String,
		nodes.FunctionDeclarationMember{},
		String,
	)
)
