package symbols

import "ReCT-Go-Compiler/print"

type ParameterSymbol struct {
	VariableSymbol

	Name    string
	Ordinal int
	Type    TypeSymbol
}

// implement the symbol interface
func (ParameterSymbol) SymbolType() SymbolType { return Parameter }
func (s ParameterSymbol) SymbolName() string   { return s.Name }

func (sym ParameterSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ ParameterSymbol ["+sym.Name+"]")
}

// implement the var interface
func (ParameterSymbol) IsGlobal() bool        { return false }
func (s ParameterSymbol) IsReadOnly() bool    { return true }
func (s ParameterSymbol) VarType() TypeSymbol { return s.Type }

// constructor
func CreateParameterSymbol(name string, ordinal int, typeSymbol TypeSymbol) ParameterSymbol {
	return ParameterSymbol{
		Name:    name,
		Ordinal: ordinal,
		Type:    typeSymbol,
	}
}
