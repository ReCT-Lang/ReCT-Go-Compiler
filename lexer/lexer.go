package lexer

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// TODO(tokorv): hey gamer, could you add a field to the token struct for "typed value" (you can prob come up with a better name)
//               its just gonna be the value but already converted into the correct datatype

// lexer : internal struct for Lexical Analysis
type Lexer struct {
	Code   []byte
	Line   int // Not implemented yet
	Column int // Not implemented yet
	Index  int
	Tokens []Token
}

func Lex(filename string) []Token {
	scanner := &Lexer{
		handleFileOpen(filename),
		0,
		0,
		0,
		make([]Token, 0),
	}

	for scanner.Index < len(scanner.Code) {
		c := scanner.Code[scanner.Index]
		if unicode.IsLetter(rune(c)) {
			scanner.getId()
		} else if unicode.IsNumber(rune(c)) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else if c == ' ' || c == '\n' || c == '\t' || c == '\v' {
			scanner.Index++
		} else {
			scanner.getOperator()
		}
	}
	return scanner.Tokens
}

func (lxr *Lexer) getNumber() {
	var buffer string
	buffer = string(lxr.Code[lxr.Index])
	lxr.Index++

	for lxr.Index < len(lxr.Code) && unicode.IsDigit(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}

	lxr.Tokens = append(lxr.Tokens, Token{buffer, NumberToken, lxr.Line, lxr.Column})
}

func (lxr *Lexer) getString() {
	var buffer string
	lxr.Index++

	for lxr.Index < len(lxr.Code) && lxr.Code[lxr.Index] != '"' {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}
	lxr.Index++ // Can't believe I forgot this aaaaaaa
	lxr.Tokens = append(lxr.Tokens, Token{buffer, StringToken, lxr.Line, lxr.Column})
}

func (lxr *Lexer) getId() {
	var buffer string
	buffer = string(lxr.Code[lxr.Index])
	lxr.Index++

	IsLetterOrDigitOrWhatever := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c) || string(c) == "_" || string(c) == "."
	}
	for lxr.Index < len(lxr.Code) && IsLetterOrDigitOrWhatever(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}

	lxr.Tokens = append(lxr.Tokens, Token{buffer, IdToken, lxr.Line, lxr.Column})
}

func (lxr *Lexer) getOperator() {
	var _token TokenKind
	peek := func(offset int) byte {
		if lxr.Index+offset < len(lxr.Code) {
			return lxr.Code[lxr.Index+offset]
		}
		return '\000'
	}
	switch lxr.Code[lxr.Index] {
	case '+':
		_token = PlusToken
	case '-':
		_token = MinusToken
	case '/':
		_token = SlashToken
	case '*':
		_token = StarToken
	case '=':
		_token = EqualsToken
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
	case '<':
		if peek(1) == '-' {
			lxr.Index++
			_token = AssignToken
		} else {
			_token = LessThanToken
		}
	case '>':
		_token = GreaterThanToken
	default:
		fmt.Println("ERROR: Unexpected character somewhere idk where though")
		_token = BadToken
	}
	if _token == AssignToken {
		lxr.Tokens = append(lxr.Tokens, CreateToken(
			string(peek(-1))+string(lxr.Code[lxr.Index]),
			_token,
			0,
			0,
		),
		)
	} else {
		lxr.Tokens = append(lxr.Tokens, CreateToken(string(lxr.Code[lxr.Index]), _token, 0, 0))
	}
	lxr.Index++
}

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
