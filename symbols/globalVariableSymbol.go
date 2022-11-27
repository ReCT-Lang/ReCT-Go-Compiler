package symbols

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

type GlobalVariableSymbol struct {
	VariableSymbol

	Name     string
	ReadOnly bool
	Type     TypeSymbol
	UniqueID int
}

// implement the symbol interface
func (GlobalVariableSymbol) SymbolType() SymbolType { return GlobalVariable }
func (s GlobalVariableSymbol) SymbolName() string   { return s.Name }

func (sym GlobalVariableSymbol) Print(indent string) {
	print.PrintC(print.Magenta, indent+"└ GlobalVariableSymbol ["+sym.Name+"]")
}

// implement the var interface
func (GlobalVariableSymbol) IsGlobal() bool        { return true }
func (s GlobalVariableSymbol) IsReadOnly() bool    { return s.ReadOnly }
func (s GlobalVariableSymbol) VarType() TypeSymbol { return s.Type }
func (s GlobalVariableSymbol) Fingerprint() string { return fmt.Sprintf("VG_%d", s.UniqueID) }

// constructor
func CreateGlobalVariableSymbol(name string, readonly bool, typeSymbol TypeSymbol) GlobalVariableSymbol {
	variableCounter++
	return GlobalVariableSymbol{
		Name:     name,
		ReadOnly: readonly,
		Type:     typeSymbol,
		UniqueID: variableCounter,
	}
}
