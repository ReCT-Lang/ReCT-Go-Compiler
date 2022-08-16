package symbols

import (
	"ReCT-Go-Compiler/print"
)

type PackageSymbol struct {
	Symbol

	Exists bool

	Name      string
	Functions []FunctionSymbol
	Classes   []ClassSymbol
}

// implement the symbol interface
func (PackageSymbol) SymbolType() SymbolType { return Package }
func (s PackageSymbol) SymbolName() string   { return s.Name }

func (sym PackageSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ PackageSymbol ["+sym.Name+"]")
}

func (s PackageSymbol) Fingerprint() string {
	id := "P_" + s.Name + "_"
	return id
}

// constructor
func CreatePackageSymbol(name string, functions []FunctionSymbol, classes []ClassSymbol) PackageSymbol {
	return PackageSymbol{
		Exists:    true,
		Name:      name,
		Functions: functions,
		Classes:   classes,
	}
}
