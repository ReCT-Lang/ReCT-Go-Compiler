package lexer

import (
	"ReCT-Go-Compiler/print"
	"errors"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Lexer : Lexer struct for lexing :GentlemenSphere:
type Lexer struct {
	Code   []rune
	Line   int
	Column int
	Index  int
	Tokens []Token
}

// Lex takes a filename and converts it into it's respective lexical tokens
func Lex(filename string) []Token {
	// Opens the file and returns its contents as a byte array
	// It then creates a lexer pointer using the byte array and a few default values.
	scanner := &Lexer{handleFileOpen(filename), 1, 1, 0, make([]Token, 0)}

	// Scanning for all the juicy tokens
	for scanner.Index < len(scanner.Code) {
		c := scanner.Code[scanner.Index]

		peek := func(offset int) rune {
			if scanner.Index+offset < len(scanner.Code) {
				return scanner.Code[scanner.Index+offset]
			}
			return '\000'
		}

		if unicode.IsLetter(c) {
			scanner.getId()
		} else if unicode.IsNumber(c) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else if c == '/' && peek(1) == '/' {
			scanner.getComment()
		} else if c != ' ' && c != '\n' && c != '\t' && c != '\v' {
			scanner.getOperator()
		} else {
			scanner.Increment()
		}
	}
	// Finally, adding an End of File token to help detect the end of the file in the parser (syntax)
	scanner.Tokens = append(scanner.Tokens, CreateToken("\000", EOF, scanner.Line, scanner.Column))
	return scanner.Tokens
}

// getNumber keeps getting bytes until it finds a non-number
// then it generates an integer (or a float) token and slaps it back to the lexer.
func (lxr *Lexer) getNumber() {
	buffer := string(lxr.Code[lxr.Index])
	lxr.Increment()

	// Checks if rune value is a digit, dot or underscore
	// Simplifies code
	isDigitOrDotOrUnderScore := func(c rune) bool {
		return unicode.IsDigit(c) || c == '.' || c == '_'
	}

	// Underscores are now allowed (don't tell Red, it can be our little secret!)
	// Integers and floats can now contain underscores to increase readability
	// Example: 10_000_000_000 (instead of 10000000000)
	// Underscores are removed at lexing, and aren't parsed or processed further than this.
	// - Added without difficulty, your favourite Duck, Tokorv. <3
	for lxr.Index < len(lxr.Code) && isDigitOrDotOrUnderScore(lxr.Code[lxr.Index]) {

		// Here we check for, and remove underscores.
		// Underscores are removed by simply not appending them to the buffer
		if lxr.Code[lxr.Index] != '_' {
			buffer += string(lxr.Code[lxr.Index])
		}
		lxr.Increment()
	}

	// Checking if number is actually an imposter... float... sus
	if strings.Contains(buffer, ".") {
		// float real value
		realValueBuffer, err := strconv.ParseFloat(buffer, 32)
		if err != nil {
			print.Error(
				"LEXER",
				print.RealValueConversionError,
				lxr.Line,
				lxr.Column,
				0,
				"value \"%s\" could not be converted to real value [float] (NumberToken)!",
				buffer,
			)
		}
		lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, float32(realValueBuffer), NumberToken, lxr.Line, lxr.Column))

	} else {
		// int real value
		realValueBuffer, err := strconv.Atoi(buffer)
		if err != nil {
			print.Error(
				"LEXER",
				print.RealValueConversionError,
				lxr.Line,
				lxr.Column,
				0,
				"value \"%s\" could not be converted to real value [int] (NumberToken)!",
				buffer,
			)
		}
		lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, realValueBuffer, NumberToken, lxr.Line, lxr.Column))
	}
}

// getString once it finds an " it'll keep getting bytes until it finds another "
// Basically it's a string detector, string tokens are given back to the lexer (via Tokens []Token).
func (lxr *Lexer) getString() {
	var buffer string
	lxr.Increment()

	for lxr.Index < len(lxr.Code) && lxr.Code[lxr.Index] != '"' {
		if lxr.Code[lxr.Index] == '\\' {
			if lxr.Code[lxr.Index+1] == 'n' {
				lxr.Increment()
				lxr.Increment()
				buffer += "\n"
				continue
			}
		}

		buffer += string(lxr.Code[lxr.Index])
		lxr.Increment()
	}
	lxr.Increment()
	lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, buffer, StringToken, lxr.Line, lxr.Column))
}

// getComment we don't want to add comments to the Tokens because they have nothing of value
// for us to process (at least for now), instead we just keep incrementing through them to increase
// the Lexer Column and Line until we find a new line.
func (lxr *Lexer) getComment() {

	// just increment until we're at the end of file or and of a line
	for lxr.Index < len(lxr.Code) && lxr.Code[lxr.Index] != '\n' {
		lxr.Increment()
	}
	lxr.Increment()
}

// Increment increases the scanner's Index, Column, and Line (if needed).
// This will also check if the index is out of range (End Of File) but leaves
// Error handling to the parent function.
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

// getId gets and identifier Token and appends it to the Lexer Tokens
// A identifier token is just a series of alphanumerical characters like name29, FunctionCall, a38rja
// ReCT identifiers can't start with a number nor an underscore
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

	// checks if identifier is actually a keyword
	kwType := CheckIfKeyword(buffer)

	// Mmm converting true/false into actual boolean values
	if kwType == TrueKeyword || kwType == FalseKeyword {
		lxr.Tokens = append(lxr.Tokens, CreateTokenReal(buffer, kwType == TrueKeyword, kwType, lxr.Line, lxr.Column))
	} else {
		lxr.Tokens = append(lxr.Tokens, CreateToken(buffer, kwType, lxr.Line, lxr.Column))
	}
}

// getOperator checks for plus/minus/assign/etc tokens
// this functional also handles unexpected character!
func (lxr *Lexer) getOperator() {
	var _token TokenKind

	peek := func(offset int) rune {
		if lxr.Index+offset < len(lxr.Code) {
			return lxr.Code[lxr.Index+offset]
		}
		return '\000'
	}

	// save our current index for later
	startIndex := lxr.Index

	// save the line and column, so we're always using the first of possibly many characters
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
	case '[':
		_token = OpenBracketToken
	case ']':
		_token = CloseBracketToken
	case ';':
		_token = Semicolon
	case '?':
		_token = QuestionMarkToken
	case ':':
		_token = ColonToken
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
		print.Error(
			"LEXER",
			print.UnexpectedCharacterError,
			lxr.Line,
			lxr.Column,
			5,
			"an unexpected character was found \"%s\"! Lexer is unable to process this character! (BadToken)",
			string(lxr.Code[lxr.Index]),
		)
		_token = BadToken
	}
	// AssignToken is 2 characters long while every other operator is 1 character.
	// (that is why they are separated).

	// Generalised this a little because we now got a few multi-char operators - Red, thanks - Tokorv xD
	lxr.Tokens = append(
		lxr.Tokens,
		CreateTokenSpaced(
			string(lxr.Code)[startIndex:lxr.Index+1],
			_token,
			line,
			column,
			peek(1) == ' ',
		),
	)

	lxr.Increment()
}

// handleFileOpen reads the file and returns a byte array ([]byte) // nah fam we usin runes
// only handles NotExist and Permission error btw
func handleFileOpen(filename string) []rune {
	contents, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		print.Error(
			"LEXER",
			print.FileDoesNotExistError,
			0,
			0,
			0,
			"file \"%s\" does not exit! Maybe you spelt it wrong?!",
			filename,
		)
		os.Exit(1)
	} else if errors.Is(err, os.ErrPermission) {
		print.Error(
			"LEXER",
			print.FilePermissionError,
			0,
			0,
			0,
			"do not have permissions to open file \"%s\"!",
			filename,
		)
		os.Exit(1)
	} else if err != nil {
		print.Error(
			"LEXER",
			print.FileVoidError,
			0,
			0,
			5,
			"an unexpected error occurred when reading file \"%s\"!",
			filename,
		)
		os.Exit(1)
	}
	// Offload a copy of contents for error handling
	// Also split at new lines because that makes referencing easier
	print.CodeReference = make([]string, 0)
	print.CodeReference = strings.Split(string(contents), "\n")
	return []rune(string(contents))
}

// CheckIfKeyword used by Lexer.getId to convert an identifier Token to a keyword Token
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
	case "class":
		return ClassKeyword
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
	case "make":
		return MakeKeyword
	case "Thread":
		return ThreadKeyword
	default:
		return IdToken
	}
}
