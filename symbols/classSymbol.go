package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
)

type ClassSymbol struct {
	Symbol

	Exists bool

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
func CreateClassSymbol(name string, declration nodes.ClassDeclarationMember, functions []FunctionSymbol, fields []VariableSymbol) ClassSymbol {
	return ClassSymbol{
		Exists:      true,
		Name:        name,
		Declaration: declration,
		Functions:   functions,
		Fields:      fields,
	}
}