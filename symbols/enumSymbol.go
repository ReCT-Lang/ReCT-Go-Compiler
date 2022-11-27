package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
)

type EnumSymbol struct {
	Symbol

	Exists bool

	Type TypeSymbol

	Name        string
	Declaration nodes.EnumDeclarationMember
	Fields      map[string]int
}

// implement the symbol interface
func (EnumSymbol) SymbolType() SymbolType { return Enum }
func (s EnumSymbol) SymbolName() string   { return s.Name }

func (sym EnumSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ EnumSymbol ["+sym.Name+"]")
}

func (s EnumSymbol) Fingerprint() string {
	id := "E_" + s.Name + "_"
	return id
}

// constructor
func CreateEnumSymbol(name string, declaration nodes.EnumDeclarationMember, fields map[string]int) EnumSymbol {
	sym := EnumSymbol{
		Exists:      true,
		Name:        name,
		Declaration: declaration,
		Fields:      fields,
	}

	sym.Type = CreateTypeSymbol(name, make([]TypeSymbol, 0), false, false, true, PackageSymbol{}, sym)
	return sym
}
