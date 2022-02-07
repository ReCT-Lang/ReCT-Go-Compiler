package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
)

type FunctionSymbol struct {
	Symbol
	Exists bool

	Name        string
	Parameters  []ParameterSymbol
	Type        TypeSymbol
	Declaration nodes.FunctionDeclarationMember
}

// implement the symbol interface
func (FunctionSymbol) SymbolType() SymbolType { return Function }
func (s FunctionSymbol) SymbolName() string   { return s.Name }

func (sym FunctionSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ FunctionSymbol ["+sym.Name+"]")
}

// constructor
func CreateFunctionSymbol(name string, params []ParameterSymbol, typeSymbol TypeSymbol, declaration nodes.FunctionDeclarationMember) FunctionSymbol {
	return FunctionSymbol{
		Exists:      true,
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
	}
}
