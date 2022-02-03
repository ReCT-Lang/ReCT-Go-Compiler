package symbols

import "ReCT-Go-Compiler/nodes"

type FunctionSymbol struct {
	Symbol

	Name        string
	Parameters  []ParameterSymbol
	Type        TypeSymbol
	Declaration nodes.FunctionDeclarationMember
}

// implement the symbol interface
func (FunctionSymbol) SymbolType() SymbolType { return Function }
func (s FunctionSymbol) SymbolName() string   { return s.Name }

// constructor
func CreateFunctionSymbol(name string, params []ParameterSymbol, typeSymbol TypeSymbol, declaration nodes.FunctionDeclarationMember) FunctionSymbol {
	return FunctionSymbol{
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
	}
}
