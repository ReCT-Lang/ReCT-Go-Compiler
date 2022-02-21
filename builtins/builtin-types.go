package builtins

import "ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0), false)
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0), false)
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0), false)
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0), false)
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0), true)
	Any    = symbols.CreateTypeSymbol("any", make([]symbols.TypeSymbol, 0), true)

	// the cursed one
	Error = symbols.CreateTypeSymbol("?", make([]symbols.TypeSymbol, 0), false)

	Types = []symbols.TypeSymbol{
		Void, Bool, Int, Float, String, Any,
	}
)
