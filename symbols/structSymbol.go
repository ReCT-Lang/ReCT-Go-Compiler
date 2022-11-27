package symbols

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/nodes"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"

	"github.com/llir/llvm/ir/types"
)

type StructSymbol struct {
	Symbol

	Exists bool

	Type   TypeSymbol
	IRType types.Type

	Name        string
	Declaration nodes.StructDeclarationMember
	Fields      []VariableSymbol
}

// implement the symbol interface
func (StructSymbol) SymbolType() SymbolType { return Struct }
func (s StructSymbol) SymbolName() string   { return s.Name }

func (sym StructSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ StructSymbol ["+sym.Name+"]")
}

func (s StructSymbol) Fingerprint() string {
	id := "S_" + s.Name + "_"
	return id
}

// constructor
func CreateStructSymbol(name string, declaration nodes.StructDeclarationMember, fields []VariableSymbol) StructSymbol {
	sym := StructSymbol{
		Exists:      true,
		Name:        name,
		Declaration: declaration,
		Fields:      fields,
	}

	sym.Type = CreateTypeSymbol(name, make([]TypeSymbol, 0), false, true, false, PackageSymbol{}, sym)

	return sym
}
