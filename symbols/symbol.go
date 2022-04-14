package symbols

import "fmt"

type Symbol interface {
	SymbolType() SymbolType
	SymbolName() string
	Print(indent string)
	Fingerprint() string
}

var variableCounter = 0

type VariableSymbol interface {
	Symbol

	IsReadOnly() bool
	IsGlobal() bool
	VarType() TypeSymbol
}

var tempCounter = 0

func GetTempName() string {
	tempCounter++
	return fmt.Sprintf("TMP_%d", tempCounter)
}

// types of symbols
type SymbolType string

const (
	Function       SymbolType = "FunctionSymbol"
	GlobalVariable SymbolType = "GlobalVariableSymbol"
	LocalVariable  SymbolType = "LocalVariableSymbol"
	Parameter      SymbolType = "ParameterSymbol"
	Type           SymbolType = "TypeSymbol"
)
