package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
)

type TypeFunctionSymbol struct {
	Symbol

	Exists  bool
	BuiltIn bool

	Name        string
	Parameters  []ParameterSymbol
	Type        TypeSymbol
	Declaration nodes.FunctionDeclarationMember
	OriginType  TypeSymbol // This is the data type of function access call like string->Split()
}

// implement the symbol interface
func (TypeFunctionSymbol) SymbolType() SymbolType { return Function }
func (sym TypeFunctionSymbol) SymbolName() string { return sym.Name }

func (sym TypeFunctionSymbol) Print(indent string) {
	if sym.BuiltIn {
		print.PrintC(print.Cyan, indent+"└ FunctionSymbol ["+sym.Name+"]")
	} else {
		print.PrintC(print.Magenta, indent+"└ FunctionSymbol ["+sym.Name+"]")
	}
}

func (sym TypeFunctionSymbol) Fingerprint() string {
	id := "TF_" + sym.OriginType.Fingerprint() + "_" + sym.Name + "_"

	for _, param := range sym.Parameters {
		id += "[" + param.Type.Fingerprint() + "]"
	}

	id += sym.Type.Name
	return id
}

// constructor
func CreateTypeFunctionSymbol(
	name string,
	params []ParameterSymbol,
	typeSymbol TypeSymbol,
	declaration nodes.FunctionDeclarationMember,
	origin TypeSymbol,
) TypeFunctionSymbol {
	return TypeFunctionSymbol{
		Exists:      true,
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
	}
}

func CreateBuiltInTypeFunctionSymbol(
	name string,
	params []ParameterSymbol,
	typeSymbol TypeSymbol,
	declaration nodes.FunctionDeclarationMember,
	origin TypeSymbol,
) TypeFunctionSymbol {
	return TypeFunctionSymbol{
		Exists:      true,
		BuiltIn:     true,
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
		OriginType:  origin,
	}
}
