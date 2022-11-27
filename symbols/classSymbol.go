package symbols

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"

	"github.com/llir/llvm/ir/types"
)

type ClassSymbol struct {
	Symbol

	Exists bool

	Type   TypeSymbol
	IRType types.Type

	Package PackageSymbol

	Name        string
	Declaration nodes.ClassDeclarationMember
	Functions   []FunctionSymbol
	Fields      []VariableSymbol
}

// implement the symbol interface
func (ClassSymbol) SymbolType() SymbolType { return Class }
func (s ClassSymbol) SymbolName() string   { return s.Name }

func (sym ClassSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ ClassSymbol ["+sym.Name+"]")
}

func (s ClassSymbol) Fingerprint() string {
	id := "C_" + s.Name + "_"
	return id
}

// constructor
func CreateClassSymbol(name string, declration nodes.ClassDeclarationMember, functions []FunctionSymbol, fields []VariableSymbol, pck PackageSymbol) ClassSymbol {
	sym := ClassSymbol{
		Exists:      true,
		Name:        name,
		Declaration: declration,
		Functions:   functions,
		Fields:      fields,
		Package:     pck,
	}

	sym.Type = CreateTypeSymbol(name, make([]TypeSymbol, 0), true, true, false, pck, sym)
	return sym
}
