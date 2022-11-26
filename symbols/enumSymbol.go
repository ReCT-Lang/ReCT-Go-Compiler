package symbols

import (
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"

	"github.com/llir/llvm/ir/types"
)

type StructSymbol struct {
	Symbol

	Exists bool

	Type   TypeSymbol

	Name        string
	Declaration nodes.EnumDeclarationMember
	Fields      map[string]int
}

// implement the symbol interface
func (StructSymbol) SymbolType() SymbolType { return Enum }
func (s StructSymbol) SymbolName() string   { return s.Name }

func (sym StructSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ EnumSymbol ["+sym.Name+"]")
}

func (s StructSymbol) Fingerprint() string {
	id := "E_" + s.Name + "_"
	return id
}

// constructor
func CreateEnumSymbol(name string, declaration nodes.EnumDeclarationMember, fields map[string]int) EnumSymbol {
	return EnumSymbol{
		Exists:      true,
		Name:        name,
		Declaration: declaration,
		Fields:      fields,
		Type:        CreateTypeSymbol(name, make([]TypeSymbol, 0), false, false),
	}
}
