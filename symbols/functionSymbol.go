package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"

	"github.com/llir/llvm/ir"
)

type FunctionSymbol struct {
	Symbol

	Exists   bool
	BuiltIn  bool
	Public   bool
	External bool
	Variadic bool

	IRFunction *ir.Func

	Name        string
	Parameters  []ParameterSymbol
	Type        TypeSymbol
	Declaration nodes.FunctionDeclarationMember
}

// implement the symbol interface
func (FunctionSymbol) SymbolType() SymbolType { return Function }
func (s FunctionSymbol) SymbolName() string   { return s.Name }

func (sym FunctionSymbol) Print(indent string) {
	if sym.BuiltIn {
		print.PrintC(print.Cyan, indent+"└ FunctionSymbol ["+sym.Name+"]")
	} else {
		print.PrintC(print.Magenta, indent+"└ FunctionSymbol ["+sym.Name+"]")
	}
}

func (s FunctionSymbol) Fingerprint() string {
	id := "F_" + s.Name + "_"

	for _, param := range s.Parameters {
		id += "[" + param.Type.Fingerprint() + "]"
	}

	id += s.Type.Name
	return id
}

// constructor
func CreateFunctionSymbol(name string, params []ParameterSymbol, typeSymbol TypeSymbol, declaration nodes.FunctionDeclarationMember, public bool) FunctionSymbol {
	return FunctionSymbol{
		Exists:      true,
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
		Public:      public,
	}
}

func CreateExternalFunctionSymbol(name string, params []ParameterSymbol, typeSymbol TypeSymbol, variadic bool) FunctionSymbol {
	return FunctionSymbol{
		Exists:     true,
		Name:       name,
		Parameters: params,
		Type:       typeSymbol,
		External:   true,
		Variadic:   variadic,
		Public:     true,
	}
}

func CreateBuiltInFunctionSymbol(name string, params []ParameterSymbol, typeSymbol TypeSymbol, declaration nodes.FunctionDeclarationMember) FunctionSymbol {
	return FunctionSymbol{
		Exists:      true,
		BuiltIn:     true,
		Name:        name,
		Parameters:  params,
		Type:        typeSymbol,
		Declaration: declaration,
		Public:      true,
	}
}
