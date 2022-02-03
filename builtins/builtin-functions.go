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

	InputKey = symbols.CreateFunctionSymbol(
		"InputKey",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)

	Clear = symbols.CreateFunctionSymbol(
		"Clear",
		[]symbols.ParameterSymbol{},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	SetCursor = symbols.CreateFunctionSymbol(
		"SetCursor",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("x", 0, Int),
			symbols.CreateParameterSymbol("y", 1, Int),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	GetSizeX = symbols.CreateFunctionSymbol(
		"GetSizeX",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	GetSizeY = symbols.CreateFunctionSymbol(
		"GetSizeY",
		[]symbols.ParameterSymbol{},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	SetSize = symbols.CreateFunctionSymbol(
		"GetSize",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("width", 0, Int),
			symbols.CreateParameterSymbol("height", 1, Int),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	SetCursorVisible = symbols.CreateFunctionSymbol(
		"SetCursorVisible",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("state", 0, Bool),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	GetCursorVisible = symbols.CreateFunctionSymbol(
		"GetCursorVisible",
		[]symbols.ParameterSymbol{},
		Bool,
		nodes.FunctionDeclarationMember{},
	)

	Random = symbols.CreateFunctionSymbol(
		"Random",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("maxNum", 0, Int),
		},
		Int,
		nodes.FunctionDeclarationMember{},
	)

	Sleep = symbols.CreateFunctionSymbol(
		"Sleep",
		[]symbols.ParameterSymbol{
			symbols.CreateParameterSymbol("milliseconds", 0, Int),
		},
		Void,
		nodes.FunctionDeclarationMember{},
	)

	Version = symbols.CreateFunctionSymbol(
		"Version",
		[]symbols.ParameterSymbol{},
		String,
		nodes.FunctionDeclarationMember{},
	)
)
