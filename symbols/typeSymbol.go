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
	IsEnum        bool

	SourceSymbol Symbol
	Package      PackageSymbol
}

// implement the interface
func (TypeSymbol) SymbolType() SymbolType { return Type }
func (s TypeSymbol) SymbolName() string   { return s.Name }

func (sym TypeSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ TypeSymbol ["+sym.Fingerprint()+"]")
}

// a unique identifier for each type
func (sym TypeSymbol) Fingerprint() string {
	id := "T"

	if sym.IsObject {
		id += "O"
	}

	id += "_" + sym.Name + "_["

	for _, subtype := range sym.SubTypes {
		if subtype.Name == "array" {
			id += subtype.Fingerprint() + ";"
		} else {
			id += subtype.Name + ";"
		}
	}

	id += "]"

	return id
}

// constructor
func CreateTypeSymbol(name string, subTypes []TypeSymbol, isObject bool, isUserDefined bool, isEnum bool, pck PackageSymbol, src Symbol) TypeSymbol {
	return TypeSymbol{
		Name:          name,
		SubTypes:      subTypes,
		IsObject:      isObject,
		IsUserDefined: isUserDefined,
		IsEnum:        isEnum,
		Package:       pck,
		SourceSymbol:  src,
	}
}
