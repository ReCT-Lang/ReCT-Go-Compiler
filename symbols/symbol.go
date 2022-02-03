package symbols

type Symbol interface {
	SymbolType() SymbolType
	SymbolName() string
}

type VariableSymbol interface {
	Symbol

	IsReadOnly() bool
	IsGlobal() bool
	VarType() TypeSymbol
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
