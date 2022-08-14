package symbols

import (
	"ReCT-Go-Compiler/print"
)

type TypeSymbol struct {
	Symbol

	Name          string
	SubTypes      []TypeSymbol
	IsObject      bool
	IsUserDefined bool
}

// implement the interface
func (TypeSymbol) SymbolType() SymbolType { return Type }
func (s TypeSymbol) SymbolName() string   { return s.Name }

func (sym TypeSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ TypeSymbol ["+sym.Fingerprint()+"]")
}

// a unique identifier for each type
func (sym TypeSymbol) Fingerprint() string {
	id := "T_" + sym.Name + "_"

	for _, subtype := range sym.SubTypes {
		id += subtype.Name + ";"
	}

	return id
}

// constructor
func CreateTypeSymbol(name string, subTypes []TypeSymbol, isObject bool, isUserDefined bool) TypeSymbol {
	return TypeSymbol{
		Name:          name,
		SubTypes:      subTypes,
		IsObject:      isObject,
		IsUserDefined: isUserDefined,
	}
}
