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

	GetArrayLength = symbols.CreateBuiltInTypeFunctionSymbol(
		"GetLength",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
		Array,
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

	Push = symbols.CreateBuiltInTypeFunctionSymbol(
		"Push",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("object", 0, Any),
		},
		Void,
		nodes.FunctionDeclarationMember{},
		Array,
	)

	PPush = symbols.CreateBuiltInTypeFunctionSymbol(
		"PPush",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("element", 0, Identity),
		},
		Void,
		nodes.FunctionDeclarationMember{},
		Array,
	)

	Start = symbols.CreateBuiltInTypeFunctionSymbol(
		"Start",
		[]symbols.ParameterSymbol{},
		Void,
		nodes.FunctionDeclarationMember{},
		Thread,
	)

	Join = symbols.CreateBuiltInTypeFunctionSymbol(
		"Join",
		[]symbols.ParameterSymbol{},
		Void,
		nodes.FunctionDeclarationMember{},
		Thread,
	)

	Kill = symbols.CreateBuiltInTypeFunctionSymbol(
		"Kill",
		[]symbols.ParameterSymbol{},
		Void,
		nodes.FunctionDeclarationMember{},
		Thread,
	)
)
