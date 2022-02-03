package builtins

import "ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0))
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0))
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0))
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0))
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0))
)
