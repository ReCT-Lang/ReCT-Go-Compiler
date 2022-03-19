package builtins

import "ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0), false)
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0), false)
	Byte   = symbols.CreateTypeSymbol("byte", make([]symbols.TypeSymbol, 0), false)
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0), false)
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0), false)
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0), true)
	Any    = symbols.CreateTypeSymbol("any", make([]symbols.TypeSymbol, 0), true)

	// lambda/functionExpression/action/etc
	Action = symbols.CreateTypeSymbol("action", make([]symbols.TypeSymbol, 0), true)

	// threads
	Thread = symbols.CreateTypeSymbol("thread", make([]symbols.TypeSymbol, 0), true)

	// generic array types so the emitter has something to work with
	Array  = symbols.CreateTypeSymbol("array", make([]symbols.TypeSymbol, 0), true)
	PArray = symbols.CreateTypeSymbol("parray", make([]symbols.TypeSymbol, 0), true)

	// the cursed ones
	Error    = symbols.CreateTypeSymbol("?", make([]symbols.TypeSymbol, 0), false)
	Identity = symbols.CreateTypeSymbol("¯\\_(ツ)_/¯", make([]symbols.TypeSymbol, 0), false)

	Types = []symbols.TypeSymbol{
		Void, Bool, Int, Float, String, Any, Action, Array, PArray, Thread,
	}
)
