package langserverinterface

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
)

// Map of meanings
var TokenMapping map[lexer.Token]TokenMeaning

// TokenMeaning (base class) holds data on what a token actually is
type TokenMeaning interface {
	Type() TokenMeaningType
}

// VariableTokenMeaning holds data about variable tokens
type VariableTokenMeaning struct {
	TokenMeaning

	Variable symbols.VariableSymbol // wat dis variabl?
}

func (m VariableTokenMeaning) Type() TokenMeaningType {
	if m.Variable.SymbolType() == symbols.LocalVariable {
		return LocalVariableMeaning
	} else if m.Variable.SymbolType() == symbols.GlobalVariable {
		return GlobalVariableMeaning
	} else if m.Variable.SymbolType() == symbols.Parameter {
		return ParameterMeaning
	}

	return ""
}

// FunctionTokenMeaning holds data about function ID tokens
type FunctionTokenMeaning struct {
	TokenMeaning

	Function symbols.FunctionSymbol // wat dis functn?
}

func (m FunctionTokenMeaning) Type() TokenMeaningType {
	return FunctionMeaning
}

// FunctionTokenMeaning holds data about function ID tokens
type TypeFunctionTokenMeaning struct {
	TokenMeaning

	TypeFunction symbols.TypeFunctionSymbol // wat dis functn?
}

func (m TypeFunctionTokenMeaning) Type() TokenMeaningType {
	return TypeFunctionMeaning
}

// ClassTokenMeaning holds data about class ID tokens
type ClassTokenMeaning struct {
	TokenMeaning

	Class symbols.ClassSymbol // wat dis class?
}

func (m ClassTokenMeaning) Type() TokenMeaningType {
	return ClassMeaning
}

// StructTokenMeaning holds data about struct ID tokens
type StructTokenMeaning struct {
	TokenMeaning

	Struct symbols.StructSymbol // wat dis stct?
}

func (m StructTokenMeaning) Type() TokenMeaningType {
	return StructMeaning
}

// EnumTokenMeaning holds data about enum ID tokens
type EnumTokenMeaning struct {
	TokenMeaning

	Enum symbols.EnumSymbol // wat dis enm?
}

func (m EnumTokenMeaning) Type() TokenMeaningType {
	return EnumMeaning
}

// EnumFieldTokenMeaning holds data about enum ID tokens
type EnumFieldTokenMeaning struct {
	TokenMeaning

	Value int
}

func (m EnumFieldTokenMeaning) Type() TokenMeaningType {
	return EnumFieldMeaning
}

// PackageTokenMeaning holds data about package ID tokens
type PackageTokenMeaning struct {
	TokenMeaning

	Package symbols.PackageSymbol // wat dis enm?
}

func (m PackageTokenMeaning) Type() TokenMeaningType {
	return PackageMeaning
}

// TypeTokenMeaning holds data about type ID tokens
type TypeTokenMeaning struct {
	TokenMeaning

	TypeSym symbols.TypeSymbol // wat dis enm?
}

func (m TypeTokenMeaning) Type() TokenMeaningType {
	return TypeMeaning
}
