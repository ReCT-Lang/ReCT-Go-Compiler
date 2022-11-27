package builtins

import "github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Byte   = symbols.CreateTypeSymbol("byte", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	UInt   = symbols.CreateTypeSymbol("uint", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Long   = symbols.CreateTypeSymbol("long", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	ULong  = symbols.CreateTypeSymbol("ulong", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Double = symbols.CreateTypeSymbol("double", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0), true, false, false, symbols.PackageSymbol{}, nil)
	Any    = symbols.CreateTypeSymbol("any", make([]symbols.TypeSymbol, 0), true, false, false, symbols.PackageSymbol{}, nil)

	// lambda/functionExpression/action/etc
	Action = symbols.CreateTypeSymbol("action", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)

	// threads
	Thread = symbols.CreateTypeSymbol("thread", make([]symbols.TypeSymbol, 0), true, false, false, symbols.PackageSymbol{}, nil)

	// generic array types so the emitter has something to work with
	Array  = symbols.CreateTypeSymbol("array", make([]symbols.TypeSymbol, 0), true, false, false, symbols.PackageSymbol{}, nil)
	PArray = symbols.CreateTypeSymbol("parray", make([]symbols.TypeSymbol, 0), true, false, false, symbols.PackageSymbol{}, nil)

	// lazy shortcut
	AnyArr = symbols.CreateTypeSymbol("array", []symbols.TypeSymbol{Any}, true, false, false, symbols.PackageSymbol{}, nil)

	// placeholder
	Enum = symbols.CreateTypeSymbol("enum", make([]symbols.TypeSymbol, 0), false, false, true, symbols.PackageSymbol{}, nil)

	// spoopy
	Pointer = symbols.CreateTypeSymbol("pointer", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)

	// the cursed ones
	Error    = symbols.CreateTypeSymbol("?", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)
	Identity = symbols.CreateTypeSymbol("¯\\_(ツ)_/¯", make([]symbols.TypeSymbol, 0), false, false, false, symbols.PackageSymbol{}, nil)

	Types = []symbols.TypeSymbol{
		Void, Bool, Byte, Int, Long, UInt, ULong, Float, Double, String, Any, Action, Array, PArray, Pointer, Thread, Enum,
	}
)
