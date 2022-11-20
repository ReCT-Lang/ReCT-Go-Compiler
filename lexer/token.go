package lexer

import (
	print2 "ReCT-Go-Compiler/print"
	"fmt"
)

// TokenKind basically an enum containing all token types.
// TokenKind has been changed from int to string for better debugging.
type TokenKind string

// seems like we will have to set the type for every single one because if not go will think they are just strings...
const (
	// Keywords
	VarKeyword       TokenKind = "var (Keyword)"
	SetKeyword       TokenKind = "set (Keyword)"
	ToKeyword        TokenKind = "to (Keyword)"
	IfKeyword        TokenKind = "if (Keyword)"
	ElseKeyword      TokenKind = "else (Keyword)"
	TrueKeyword      TokenKind = "true (Keyword)"
	FalseKeyword     TokenKind = "false (Keyword)"
	FunctionKeyword  TokenKind = "function (Keyword)"
	ClassKeyword     TokenKind = "class (Keyword)"
	FromKeyword      TokenKind = "from (Keyword)"
	ForKeyword       TokenKind = "for (Keyword)"
	ReturnKeyword    TokenKind = "return (Keyword)"
	WhileKeyword     TokenKind = "while (Keyword)"
	ContinueKeyword  TokenKind = "continue (keyword)"
	BreakKeyword     TokenKind = "break (Keyword)"
	MakeKeyword      TokenKind = "make (Keyword)"
	PackageKeyword   TokenKind = "package (keyword)"
	ExternalKeyword  TokenKind = "external (keyword)"
	CVariadicKeyword TokenKind = "c_variadic (keyword)"
	CAdaptedKeyword  TokenKind = "c_adapted (keyword)"
	RefKeyword       TokenKind = "ref (keyword)"
	DerefKeyword     TokenKind = "deref (keyword)"
	StructKeyword    TokenKind = "struct (keyword)"
	LambdaKeyword    TokenKind = "lambda (keyword)"
	ThisKeyword      TokenKind = "this (keyword)"
	MainKeyword      TokenKind = "main (keyword)"

	// Tokens
	EOF               TokenKind = "EndOfFile"
	IdToken           TokenKind = "Identifier"
	StringToken       TokenKind = "String"
	NativeStringToken TokenKind = "NativeString"
	NumberToken       TokenKind = "Number"

	// Symbol Tokens
	PlusToken          TokenKind = "Plus '+'"
	ModulusToken       TokenKind = "Modulus '%'"
	MinusToken         TokenKind = "Minus '-'"
	StarToken          TokenKind = "Star '*'"
	SlashToken         TokenKind = "Slash '/'"
	EqualsToken        TokenKind = "Equals '='"
	NotToken           TokenKind = "Not '!'"
	NotEqualsToken     TokenKind = "Not Equals '!='"
	CommaToken         TokenKind = "Comma ','"
	GreaterThanToken   TokenKind = "GreaterThanToken '>'"
	LessThanToken      TokenKind = "LessThanToken '<'"
	GreaterEqualsToken TokenKind = "GreaterEqualsToken '>='"
	LessEqualsToken    TokenKind = "LessEqualsToken '<='"
	AmpersandToken     TokenKind = "AmpersandToken '&'"
	AmpersandsToken    TokenKind = "AmpersandsToken '&&'"
	PipeToken          TokenKind = "PipeToken '|'"
	PipesToken         TokenKind = "PipesToken '||'"
	HatToken           TokenKind = "HatToken '^'"
	AssignToken        TokenKind = "AssignToken '<-'"
	AccessToken        TokenKind = "AccessToken '->'"
	ShiftLeftToken     TokenKind = "ShiftLeftToken '<<'"
	ShiftRightToken    TokenKind = "ShiftRightToken '>>'"

	OpenBraceToken        TokenKind = "OpenBrace '{'"
	CloseBraceToken       TokenKind = "Closebrace '}'"
	OpenBracketToken      TokenKind = "OpenBracket '['"
	CloseBracketToken     TokenKind = "CloseBracket ']'"
	OpenParenthesisToken  TokenKind = "OpenParenthesis '('"
	CloseParenthesisToken TokenKind = "CloseParenthesis ')'"

	QuestionMarkToken TokenKind = "QuestionMark '?'"
	ColonToken        TokenKind = "Colon ':'"

	PackageToken TokenKind = "Package '::'"

	HashtagToken TokenKind = "Hashtag '#'"

	BadToken TokenKind = "Token Error (BadToken)" // Naughty ;)

	Semicolon TokenKind = "Semicolon ';'" // Used to separate statements (for now... )
)

// Token stores information about lexical structures in the text
type Token struct {
	Value      string
	RealValue  interface{}
	Kind       TokenKind
	Span       print2.TextSpan
	SpaceAfter bool
}

// CreateToken returns a Token created from the arguments provided
func CreateToken(value string, kind TokenKind, span print2.TextSpan) Token {
	return Token{
		Value:     value,
		RealValue: nil,
		Kind:      kind,
		Span:      span,
	}
}

// CreateTokenSpaced just another constructor to not have to include the spaced bool every time
func CreateTokenSpaced(value string, kind TokenKind, span print2.TextSpan, spaced bool) Token {
	return Token{
		Value:      value,
		RealValue:  nil,
		Kind:       kind,
		Span:       span,
		SpaceAfter: spaced,
	}
}

// CreateTokenReal the majority of the code base uses CreateToken, however, the Token struct has
// a "RealValue" which should store the true value of a Token. For example, NumberToken (TokenKind) created using
// CreateToken will only store its string value and not its real number value. CreateTokenReal will store the
// converted type (so NumberToken actually stores a number).
func CreateTokenReal(buffer string, real interface{}, kind TokenKind, span print2.TextSpan) Token {
	return Token{
		Value:     buffer,
		RealValue: real,
		Kind:      kind,
		Span:      span,
	}
}

// String easy representation of a Token
// You can also make it *pretty* - like we ever used that lmao
func (t Token) String(pretty bool) string {
	if !pretty {
		return fmt.Sprintf("Token { value: %s, kind: %s, position: (%d, %d), real: %v}", t.Value, t.Kind, t.Span.StartLine, t.Span.StartColumn, t.RealValue)
	} else {
		return fmt.Sprintf("Token { \n\tvalue: %s, \n\tkind: %s, \n\tposition: (L%d, SC%d, EC%d, Len %d)\n}", t.Value, t.Kind, t.Span.StartLine, t.Span.StartColumn, t.Span.EndColumn, t.Span.EndIndex-t.Span.StartIndex)
	}
}
