package langserverinterface

type TokenMeaningType string

const (
	LocalVariableMeaning  TokenMeaningType = "LocalVariable"
	GlobalVariableMeaning TokenMeaningType = "GlobalVariable"
	ParameterMeaning      TokenMeaningType = "Parameter"
	FunctionMeaning       TokenMeaningType = "Function"
	TypeFunctionMeaning   TokenMeaningType = "TypeFunction"
	ClassMeaning          TokenMeaningType = "Class"
	StructMeaning         TokenMeaningType = "Struct"
	EnumMeaning           TokenMeaningType = "Enum"
	EnumFieldMeaning      TokenMeaningType = "EnumField"
	PackageMeaning        TokenMeaningType = "Package"
	TypeMeaning           TokenMeaningType = "Type"
)
