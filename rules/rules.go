package rules

import "github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"

// rules: this package is for holding basic language rules like operator order
// ---------------------------------------------------------------------------

func GetUnaryOperatorPrecedence(tok lexer.Token) int {
	switch tok.Kind {

	case lexer.PlusToken,
		lexer.MinusToken,
		lexer.NotToken:
		return 6 // always one higher than the highest binary operator

	default:
		return 0
	}
}

func GetBinaryOperatorPrecedence(tok lexer.Token) int {
	switch tok.Kind {

	case lexer.StarToken, lexer.SlashToken, lexer.ModulusToken:
		return 5

	case lexer.PlusToken, lexer.MinusToken:
		return 4

	case lexer.EqualsToken,
		lexer.NotEqualsToken,
		lexer.LessThanToken,
		lexer.GreaterThanToken,
		lexer.LessEqualsToken,
		lexer.GreaterEqualsToken,
		lexer.ShiftLeftToken,
		lexer.ShiftRightToken:
		return 3

	case lexer.AmpersandsToken:
		return 2

	case lexer.PipesToken,
		lexer.HatToken:
		return 1

	default:
		return 0
	}
}
