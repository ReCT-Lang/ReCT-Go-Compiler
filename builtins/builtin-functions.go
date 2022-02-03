package builtins

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/symbols"
)

var (
	Print = symbols.CreateFunctionSymbol(
		"Print",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("text", 0, String),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Write = symbols.CreateFunctionSymbol(
		"Write",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("text", 0, String),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Input = symbols.CreateFunctionSymbol(
		"Input",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)
)
