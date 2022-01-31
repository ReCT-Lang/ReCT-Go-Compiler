package lexer

// TokenKind : Basically an enum
type TokenKind int

const (
	// Keywords
	VarKeyword TokenKind = 0
	SetKeyword
	ToKeyword
	IfKeyword
	ElseKeyword
	TrueKeyword
	FalseKeyword

	// Tokens
	EOF
	StringToken
	NumberToken
	PlusToken
	MinusToken
	StarToken
	SlashToken
	EqualsToken
	OpenBraceToken
	CloseBraceToken
	OpenParenthesisToken
	CloseParenthesisToken
	AssignToken
)

type Token struct {
	Value  string
	Kind   TokenKind
	Line   int
	Column int
}

func CreateToken(value string, kind TokenKind, line int, column int) Token {
	return Token{
		value, kind, line, column,
	}
}
