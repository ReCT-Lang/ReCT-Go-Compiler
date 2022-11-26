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

	GetBuffer = symbols.CreateBuiltInTypeFunctionSymbol(
		"GetBuffer",
		[]symbols.ParameterSymbol{},
		symbols.CreateTypeSymbol("pointer", []symbols.TypeSymbol{Byte}, false, false, false),
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

	Run = symbols.CreateBuiltInTypeFunctionSymbol(
		"Run",
		[]symbols.ParameterSymbol{}, // ---> these two get filled in on a case by case basis by the binder
		Void,                        // -/
		nodes.FunctionDeclarationMember{},
		Action,
	)

	RunThread = symbols.CreateBuiltInTypeFunctionSymbol(
		"RunThread",
		[]symbols.ParameterSymbol{}, // ---> gets filled in on a case by case basis by the binder
		Thread,
		nodes.FunctionDeclarationMember{},
		Action,
	)
)
