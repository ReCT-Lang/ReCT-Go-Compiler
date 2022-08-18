package symbols

import (
	"ReCT-Go-Compiler/print"

	"github.com/llir/llvm/ir"
)

type PackageSymbol struct {
	Symbol

	Exists bool

	Name      string
	Functions []FunctionSymbol
	Classes   []ClassSymbol

	Module *ir.Module
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
func CreatePackageSymbol(name string, functions []FunctionSymbol, classes []ClassSymbol, module *ir.Module) PackageSymbol {
	return PackageSymbol{
		Exists:    true,
		Name:      name,
		Functions: functions,
		Classes:   classes,
		Module:    module,
	}
}
