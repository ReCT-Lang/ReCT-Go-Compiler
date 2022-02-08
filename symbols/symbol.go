package symbols

type Symbol interface {
	SymbolType() SymbolType
	SymbolName() string
	Print(indent string)
}

var variableCounter = 0

type VariableSymbol interface {
	Symbol

	IsReadOnly() bool
	IsGlobal() bool
	VarType() TypeSymbol
	Fingerprint() string
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
