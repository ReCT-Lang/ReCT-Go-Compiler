package builtins

import "ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0), false, false, false)
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0), false, false, false)
	Byte   = symbols.CreateTypeSymbol("byte", make([]symbols.TypeSymbol, 0), false, false, false)
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0), false, false, false)
	UInt   = symbols.CreateTypeSymbol("uint", make([]symbols.TypeSymbol, 0), false, false, false)
	Long   = symbols.CreateTypeSymbol("long", make([]symbols.TypeSymbol, 0), false, false, false)
	ULong  = symbols.CreateTypeSymbol("ulong", make([]symbols.TypeSymbol, 0), false, false, false)
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0), false, false, false)
	Double = symbols.CreateTypeSymbol("double", make([]symbols.TypeSymbol, 0), false, false, false)
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0), true, false, false)
	Any    = symbols.CreateTypeSymbol("any", make([]symbols.TypeSymbol, 0), true, false, false)

	// lambda/functionExpression/action/etc
	Action = symbols.CreateTypeSymbol("action", make([]symbols.TypeSymbol, 0), false, false, false)

	// threads
	Thread = symbols.CreateTypeSymbol("thread", make([]symbols.TypeSymbol, 0), true, false, false)

	// generic array types so the emitter has something to work with
	Array  = symbols.CreateTypeSymbol("array", make([]symbols.TypeSymbol, 0), true, false, false)
	PArray = symbols.CreateTypeSymbol("parray", make([]symbols.TypeSymbol, 0), true, false, false)

	// lazy shortcut
	AnyArr = symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{Any}, true, false, false)

	// placeholder
	Enum = symbols.CreateTypeSymbol("enum", make([]symbols.TypeSymbol, 0), false, false, true)

	// spoopy
	Pointer = symbols.CreateTypeSymbol("pointer", make([]symbols.TypeSymbol, 0), false, false, false)

	// the cursed ones
	Error    = symbols.CreateTypeSymbol("?", make([]symbols.TypeSymbol, 0), false, false, false)
	Identity = symbols.CreateTypeSymbol("¯\\_(ツ)_/¯", make([]symbols.TypeSymbol, 0), false, false, false)

	Types = []symbols.TypeSymbol{
		Void, Bool, Byte, Int, Long, UInt, ULong, Float, Double, String, Any, Action, Array, PArray, Pointer, Thread, Enum,
	}
)
