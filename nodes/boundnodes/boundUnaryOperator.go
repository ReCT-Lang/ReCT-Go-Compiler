package boundnodes

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/symbols"
)

type BoundUnaryOperatorType string

const (
	Identity        BoundUnaryOperatorType = "Identity"
	Negation        BoundUnaryOperatorType = "Negation"
	LogicalNegation BoundUnaryOperatorType = "LogicalNegation"
)

type BoundUnaryOperator struct {
	Exists bool

	TokenKind    lexer.TokenKind
	OperatorKind BoundUnaryOperatorType
	OperandType  symbols.TypeSymbol
	ResultType   symbols.TypeSymbol
}

// constructor
func CreateBoundUnaryOperator(tok lexer.TokenKind, kind BoundUnaryOperatorType, operand symbols.TypeSymbol, result symbols.TypeSymbol) BoundUnaryOperator {
	return BoundUnaryOperator{
		Exists:       true,
		TokenKind:    tok,
		OperatorKind: kind,
		OperandType:  operand,
		ResultType:   result,
	}
}

// allowed operations
var UnaryOperators []BoundUnaryOperator = []BoundUnaryOperator{
	// int operations
	/* + */ CreateBoundUnaryOperator(lexer.PlusToken, Identity, builtins.Int, builtins.Int),
	/* - */ CreateBoundUnaryOperator(lexer.MinusToken, Negation, builtins.Int, builtins.Int),

	// float operations
	/* + */ CreateBoundUnaryOperator(lexer.PlusToken, Identity, builtins.Float, builtins.Float),
	/* - */ CreateBoundUnaryOperator(lexer.MinusToken, Negation, builtins.Float, builtins.Float),

	// bool operations
	/* ! */ CreateBoundUnaryOperator(lexer.NotToken, LogicalNegation, builtins.Bool, builtins.Bool),
}

func BindUnaryOperator(tokenKind lexer.TokenKind, operandType symbols.TypeSymbol) BoundUnaryOperator {
	for _, op := range UnaryOperators {
		if op.TokenKind == tokenKind &&
			op.OperandType.Fingerprint() == operandType.Fingerprint() {
			return op
		}
	}

	return BoundUnaryOperator{}
}
