package rules

import "ReCT-Go-Compiler/lexer"

// rules: this package is for holding basic language rules like operator order
// ---------------------------------------------------------------------------

func GetUnaryOperatorPrecedence(tok lexer.Token) int {
	switch tok.Kind {

	case lexer.PlusToken:
	case lexer.MinusToken:
	case lexer.NotToken:
		return 6 // always one higher than the highest binary operator

	default:
		return 0
	}

	return 0
}

func GetBinaryOperatorPrecedence(tok lexer.Token) int {
	switch tok.Kind {

	case lexer.StarToken:
	case lexer.SlashToken:
		return 5

	case lexer.PlusToken:
	case lexer.MinusToken:
		return 4

	case lexer.EqualsToken:
	case lexer.NotEqualsToken:
	case lexer.LessThanToken:
	case lexer.GreaterThanToken:
	case lexer.LessEqualsToken:
	case lexer.GreaterEqualsToken:
		return 3

	case lexer.AmpersandsToken:
		return 2

	case lexer.PipesToken:
	case lexer.HatToken:
		return 1

	default:
		return 0
	}

	return 0
}
