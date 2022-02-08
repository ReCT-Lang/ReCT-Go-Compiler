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
)
