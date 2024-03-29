package symbols

import (
	"fmt"
	print2 "github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

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
	Declaration() print2.TextSpan
}

var tempCounter = 0

func GetTempName() string {
	tempCounter++
	return fmt.Sprintf("TMP_%d", tempCounter)
}

var lambdaCounter = 0

func GetLambdaName() string {
	lambdaCounter++
	return fmt.Sprintf("LAMBDA_%d", lambdaCounter)
}

// types of symbols
type SymbolType string

const (
	Function       SymbolType = "FunctionSymbol"
	Class          SymbolType = "ClassSymbol"
	Struct         SymbolType = "StructSymbol"
	Enum           SymbolType = "EnumSymbol"
	GlobalVariable SymbolType = "GlobalVariableSymbol"
	LocalVariable  SymbolType = "LocalVariableSymbol"
	Parameter      SymbolType = "ParameterSymbol"
	Type           SymbolType = "TypeSymbol"
	Package        SymbolType = "PackageSymbol"
)
