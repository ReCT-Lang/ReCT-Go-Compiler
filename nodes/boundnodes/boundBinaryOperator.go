package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/symbols"
)

type BoundBinaryOperatorType string

const (
	Addition        BoundBinaryOperatorType = "Addition"
	Subtraction     BoundBinaryOperatorType = "Subtraction"
	Multiplication  BoundBinaryOperatorType = "Multiplication"
	Division        BoundBinaryOperatorType = "Division"
	Modulus         BoundBinaryOperatorType = "Modulus"
	LogicalAnd      BoundBinaryOperatorType = "LogicalAnd"
	LogicalOr       BoundBinaryOperatorType = "LogicalOr"
	BitwiseAnd      BoundBinaryOperatorType = "BitwiseAnd"
	BitwiseOr       BoundBinaryOperatorType = "BitwiseOr"
	BitwiseXor      BoundBinaryOperatorType = "BitwiseXor"
	Equals          BoundBinaryOperatorType = "Equals"
	NotEquals       BoundBinaryOperatorType = "NotEquals"
	Less            BoundBinaryOperatorType = "LessThan"
	LessOrEquals    BoundBinaryOperatorType = "LessOrEquals"
	Greater         BoundBinaryOperatorType = "GreaterThan"
	GreaterOrEquals BoundBinaryOperatorType = "GreaterOrEquals"
)

type BoundBinaryOperator struct {
	Exists bool

	TokenKind    lexer.TokenKind
	OperatorKind BoundBinaryOperatorType
	LeftType     symbols.TypeSymbol
	RightType    symbols.TypeSymbol
	ResultType   symbols.TypeSymbol
}

// constructors
func CreateBoundBinaryOperator(tok lexer.TokenKind, kind BoundBinaryOperatorType, left symbols.TypeSymbol, right symbols.TypeSymbol, result symbols.TypeSymbol) BoundBinaryOperator {
	return BoundBinaryOperator{
		Exists:       true,
		TokenKind:    tok,
		OperatorKind: kind,
		LeftType:     left,
		RightType:    right,
		ResultType:   result,
	}
}

func CreateBoundBinaryOperatorSameInputs(tok lexer.TokenKind, kind BoundBinaryOperatorType, input symbols.TypeSymbol, result symbols.TypeSymbol) BoundBinaryOperator {
	return BoundBinaryOperator{
		Exists:       true,
		TokenKind:    tok,
		OperatorKind: kind,
		LeftType:     input,
		RightType:    input,
		ResultType:   result,
	}
}

func CreateBoundBinaryOperatorAllSame(tok lexer.TokenKind, kind BoundBinaryOperatorType, datatype symbols.TypeSymbol) BoundBinaryOperator {
	return BoundBinaryOperator{
		Exists:       true,
		TokenKind:    tok,
		OperatorKind: kind,
		LeftType:     datatype,
		RightType:    datatype,
		ResultType:   datatype,
	}
}

// allowed operations
var BinaryOperators []BoundBinaryOperator = []BoundBinaryOperator{
	// integer operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Int),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Int),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Int),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Int),
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.Int),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Int),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.Int),
	/* ^  */ CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.Int),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Int, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Int, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Int, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Int, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Int, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Int, builtins.Bool),

	// byte operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Byte),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Byte),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Byte),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Byte),
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.Byte),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Byte),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.Byte),
	/* ^  */ CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.Byte),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Byte, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Byte, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Byte, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Byte, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Byte, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Byte, builtins.Bool),

	// long operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Long),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Long),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Long),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Long),
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.Long),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Long),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.Long),
	/* ^  */ CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.Long),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Long, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Long, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Long, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Long, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Long, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Long, builtins.Bool),

	// float operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Float),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Float),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Float),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Float),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Float),
	/* &  */ //CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.Float),
	/* |  */ //CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.Float),
	/* ^  */ //CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.Float),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Float, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Float, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Float, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Float, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Float, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Float, builtins.Bool),

	// uint operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.UInt),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.UInt),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.UInt),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.UInt),
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.UInt),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.UInt),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.UInt),
	/* ^  */ CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.UInt),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.UInt, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.UInt, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.UInt, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.UInt, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.UInt, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.UInt, builtins.Bool),

	// ulong operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.ULong),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.ULong),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.ULong),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.ULong),
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.ULong),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.ULong),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.ULong),
	/* ^  */ CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.ULong),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.ULong, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.ULong, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.ULong, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.ULong, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.ULong, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.ULong, builtins.Bool),

	// double operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Double),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Double),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Double),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Double),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Double),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Double, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Double, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Double, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Double, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Double, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Double, builtins.Bool),

	// pointer operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.Pointer),
	/* -  */ CreateBoundBinaryOperatorAllSame(lexer.MinusToken, Subtraction, builtins.Pointer),
	/* *  */ CreateBoundBinaryOperatorAllSame(lexer.StarToken, Multiplication, builtins.Pointer),
	/* /  */ CreateBoundBinaryOperatorAllSame(lexer.SlashToken, Division, builtins.Pointer),
	/* %  */ CreateBoundBinaryOperatorAllSame(lexer.ModulusToken, Modulus, builtins.Pointer),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.Pointer, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.Pointer, builtins.Bool),
	/* <  */ CreateBoundBinaryOperatorSameInputs(lexer.LessThanToken, Less, builtins.Pointer, builtins.Bool),
	/* <= */ CreateBoundBinaryOperatorSameInputs(lexer.LessEqualsToken, LessOrEquals, builtins.Pointer, builtins.Bool),
	/* >  */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterThanToken, Greater, builtins.Pointer, builtins.Bool),
	/* >= */ CreateBoundBinaryOperatorSameInputs(lexer.GreaterEqualsToken, GreaterOrEquals, builtins.Pointer, builtins.Bool),

	// bool operations
	/* &  */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandToken, BitwiseAnd, builtins.Bool),
	/* && */ CreateBoundBinaryOperatorAllSame(lexer.AmpersandsToken, LogicalAnd, builtins.Bool),
	/* |  */ CreateBoundBinaryOperatorAllSame(lexer.PipeToken, BitwiseOr, builtins.Bool),
	/* || */ CreateBoundBinaryOperatorAllSame(lexer.PipesToken, LogicalOr, builtins.Bool),
	/* ^  */ //CreateBoundBinaryOperatorAllSame(lexer.HatToken, BitwiseXor, builtins.Bool),
	/* =  */ CreateBoundBinaryOperatorAllSame(lexer.EqualsToken, Equals, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorAllSame(lexer.NotEqualsToken, NotEquals, builtins.Bool),

	// string operations
	/* +  */ CreateBoundBinaryOperatorAllSame(lexer.PlusToken, Addition, builtins.String),
	/* =  */ CreateBoundBinaryOperatorSameInputs(lexer.EqualsToken, Equals, builtins.String, builtins.Bool),
	/* != */ CreateBoundBinaryOperatorSameInputs(lexer.NotEqualsToken, NotEquals, builtins.String, builtins.Bool),
}

func BindBinaryOperator(tokenKind lexer.TokenKind, leftType symbols.TypeSymbol, rightType symbols.TypeSymbol) BoundBinaryOperator {
	for _, op := range BinaryOperators {
		if op.TokenKind == tokenKind &&
			op.LeftType.Fingerprint() == leftType.Fingerprint() &&
			op.RightType.Fingerprint() == rightType.Fingerprint() {
			return op
		}
		if op.TokenKind == tokenKind &&
			op.LeftType.Name == leftType.Name && leftType.Name == "pointer" &&
			op.RightType.Name == rightType.Name && rightType.Name == "pointer" {
			return op
		}
	}

	return BoundBinaryOperator{}
}
