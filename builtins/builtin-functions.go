package builtins

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/symbols"
)

var (
	Print = symbols.CreateBuiltInFunctionSymbol(
		"Print",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("text", 0, String),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Write = symbols.CreateBuiltInFunctionSymbol(
		"Write",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("text", 0, String),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Input = symbols.CreateBuiltInFunctionSymbol(
		"Input",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)

	InputKey = symbols.CreateBuiltInFunctionSymbol(
		"InputKey",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)

	Clear = symbols.CreateBuiltInFunctionSymbol(
		"Clear",
		[]symbols.ParameterSymbol{},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	SetCursor = symbols.CreateBuiltInFunctionSymbol(
		"SetCursor",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("x", 0, Int),
			symbols.CreateParameterSymbol("y", 1, Int),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	GetSizeX = symbols.CreateBuiltInFunctionSymbol(
		"GetSizeX",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	GetSizeY = symbols.CreateBuiltInFunctionSymbol(
		"GetSizeY",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	SetCursorVisible = symbols.CreateBuiltInFunctionSymbol(
		"SetCursorVisible",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("state", 0, Bool),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	GetCursorVisible = symbols.CreateBuiltInFunctionSymbol(
		"GetCursorVisible",
		[]symbols.ParameterSymbol{},
		Bool,
		nodes.FunctionDeclarationMember{},
	)

	Random = symbols.CreateBuiltInFunctionSymbol(
		"Random",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("maxNum", 0, Int),
		},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	Sleep = symbols.CreateBuiltInFunctionSymbol(
		"Sleep",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("milliseconds", 0, Int),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Version = symbols.CreateBuiltInFunctionSymbol(
		"Version",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)

	Sqrt = symbols.CreateBuiltInFunctionSymbol(
		"Sqrt",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("num", 0, Int),
		},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	Now = symbols.CreateBuiltInFunctionSymbol(
		"Now",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	Char = symbols.CreateBuiltInFunctionSymbol(
		"Char",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("index", 0, Int),
		},
		String,
		nodes.FunctionDeclarationMember{},
	)

	Functions = []symbols.FunctionSymbol{
		Print, Write, Input, InputKey, Clear, SetCursor, GetSizeX, GetSizeY, SetCursorVisible, GetCursorVisible, Random, Sleep, Version, Sqrt, Now, Char,
	}
)
