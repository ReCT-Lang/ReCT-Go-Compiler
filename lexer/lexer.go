package lexer

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Lexer : internal struct for Lexical Analysis
type Lexer struct {
	Code   []byte
	Line   int
	Column int
	Index  int
	Tokens []Token
}

// Lex
func Lex(filename string) []Token {
	scanner := &Lexer{handleFileOpen(filename), 1, 1, 0, make([]Token, 0)}

	for scanner.Index < len(scanner.Code) {
		c := scanner.Code[scanner.Index]

		if unicode.IsLetter(rune(c)) {
			scanner.getId()
		} else if unicode.IsNumber(rune(c)) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else if c != ' ' && c != '\n' && c != '\t' && c != '\v' {
			scanner.getOperator()
		} else {
			scanner.Increment()
		}
	}
	scanner.Tokens = append(scanner.Tokens, CreateToken("\000", EOF, scanner.Line, scanner.Column))
	return scanner.Tokens
}

// getNumber
func (lxr *Lexer) getNumber() {
	buffer := string(lxr.Code[lxr.Index])
	lxr.Increment()

	for lxr.Index < len(lxr.Code) && unicode.IsDigit(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Increment()
	}

	realValueBuffer, err := strconv.Atoi(buffer)
	if err != nil {
		fmt.Printf("ERROR: value \"%s\" could not be converted to real value (NumberToken)!", buffer)
	}
	lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, realValueBuffer, NumberToken, lxr.Line, lxr.Column))
}

// getString
func (lxr *Lexer) getString() {
	var buffer string
	lxr.Increment()

	for lxr.Index < len(lxr.Code) && lxr.Code[lxr.Index] != '"' {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Increment()
	}
	lxr.Increment()
	lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, buffer, StringToken, lxr.Line, lxr.Column))
}

// Increment increases the scanner's Index, Column, and Lin (if needed).
func (lxr *Lexer) Increment() {
	lxr.Index++
	lxr.Column++

	if lxr.Index >= len(lxr.Code) {
		return
	} else if lxr.Code[lxr.Index] == '\n' {
		lxr.Line++
		lxr.Column = 1
	}
}

// getId
func (lxr *Lexer) getId() {
	buffer := string(lxr.Code[lxr.Index])
	lxr.Increment()

	IsLetterOrDigitOrWhatever := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c) || string(c) == "_" || string(c) == "."
	}

	for lxr.Index < len(lxr.Code) && IsLetterOrDigitOrWhatever(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Increment()
	}

	kwType := CheckIfKeyword(buffer)

	if kwType == TrueKeyword || kwType == FalseKeyword {
		lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, kwType == TrueKeyword, kwType, lxr.Line, lxr.Column))
	} else {
		lxr.Tokens = append(lxr.Tokens, CreateToken(buffer, kwType, lxr.Line, lxr.Column))
	}
}

// getOperator
func (lxr *Lexer) getOperator() {
	var _token TokenKind

	peek := func(offset int) byte {
		if lxr.Index+offset < len(lxr.Code) {
			return lxr.Code[lxr.Index+offset]
		}
		return '\000'
	}

	// save our current index for later
	startIndex := lxr.Index

	// save the line and column so we're always using the first of possibly many characters
	line := lxr.Line
	column := lxr.Column

	switch lxr.Code[lxr.Index] {
	case '+':
		_token = PlusToken
  case '%':
    _token = ModulusToken
	case '-':
		if peek(1) == '>' {
			lxr.Increment()
			_token = AccessToken
		} else {
			_token = MinusToken
		}
	case '/':
		_token = SlashToken
	case '*':
		_token = StarToken
	case '=':
		_token = EqualsToken
	case '!':
		if peek(1) == '=' {
			lxr.Increment()
			_token = NotEqualsToken
		} else {
			_token = NotToken
		}
	case '(':
		_token = OpenParenthesisToken
	case ')':
		_token = CloseParenthesisToken
	case '{':
		_token = OpenBraceToken
	case '}':
		_token = CloseBraceToken
	case ';':
		_token = Semicolon
	case ',':
		_token = CommaToken
	case '<':
		if peek(1) == '-' {
			lxr.Increment()
			_token = AssignToken
		} else if peek(1) == '=' {
			lxr.Increment()
			_token = LessEqualsToken
		} else {
			_token = LessThanToken
		}
	case '>':
		if peek(1) == '=' {
			lxr.Increment()
			_token = GreaterEqualsToken
		} else {
			_token = GreaterThanToken
		}
	case '^':
		_token = HatToken
	case '&':
		if peek(1) == '&' {
			lxr.Increment()
			_token = AmpersandsToken
		} else {
			_token = AmpersandToken
		}
	case '|':
		if peek(1) == '|' {
			lxr.Increment()
			_token = PipesToken
		} else {
			_token = PipeToken
		}

	default:
		fmt.Printf(
			"ERROR(%d, %d): Unexpected character \"%s\"!\n",
			lxr.Line,
			lxr.Column,
			string(lxr.Code[lxr.Index]),
		)
		_token = BadToken
	}
	// AssignToken is 2 characters long while every other operator is 1 character.
	// (that is why they are separated).

	// Generalised this a litte because we now got a few multi-char operators - Red
	lxr.Tokens = append(
		lxr.Tokens,
		CreateToken(
			string(lxr.Code)[startIndex:lxr.Index+1],
			_token,
			line,
			column,
		),
	)

	lxr.Increment()
}

// handleFileOpen
func handleFileOpen(filename string) []byte {
	contents, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("ERROR: file \"%s\" does not exist!\n", filename)
		os.Exit(1)
	} else if errors.Is(err, os.ErrPermission) {
		fmt.Printf("ERROR: do not have permission to open file \"%s\"!\n", filename)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("ERROR: unable to open file \"%s\" for unknown reasons!", filename)
		os.Exit(1)
	}
	return contents
}

// CheckIfKeyword forgot about this till I started reading the parser code lol
func CheckIfKeyword(buffer string) TokenKind {
	switch buffer {
	case "var":
		return VarKeyword
	case "set":
		return SetKeyword
	case "to":
		return ToKeyword
	case "if":
		return IfKeyword
	case "else":
		return ElseKeyword
	case "true":
		return TrueKeyword
	case "false":
		return FalseKeyword
	case "function":
		return FunctionKeyword
	case "from":
		return FromKeyword
	case "for":
		return ForKeyword
	case "return":
		return ReturnKeyword
	case "while":
		return WhileKeyword
	case "break":
		return BreakKeyword
	case "continue":
		return ContinueKeyword
	default:
		return IdToken
	}
}
