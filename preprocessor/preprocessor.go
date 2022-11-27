package preprocessor

import (
	"errors"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"os"
	"strings"
)

type Preprocessor struct {
	Code       string
	Filename   string
	Index      int
	Tokens     []lexer.Token
	Statements []PreProcStatement

	// has any statement in here changed this files code?
	ChangedFile bool

	// reference to the source file list
	Sources *[]string

	// reference to the list of arguments
	Args *[]string
}

type PreProcStatement struct {
	Keyword string
	Content string
	Span    print.TextSpan
}

// <HELPERS> ------------------------------------------------------------------

// returns current token
func (ppc *Preprocessor) current() lexer.Token {
	return ppc.peek(0)
}

// returns current token + a given offset
func (ppc *Preprocessor) peek(offset int) lexer.Token {
	// if we are out of bounds -> return EOF token
	if ppc.Index+offset < 0 || ppc.Index+offset >= len(ppc.Tokens) {
		return lexer.Token{
			Kind:  lexer.EOF,
			Value: "",
		}
	}

	// if everything is fine -> great!
	return ppc.Tokens[ppc.Index+offset]
}

// </HELPERS> -----------------------------------------------------------------

func Preprocess(filename string, sources *[]string, arguments *[]string) string {
	// create a preprocessor object
	preproc := Preprocessor{Sources: sources, Args: arguments}

	// read the files contents
	code := ReadFile(filename, print.TextSpan{})
	preproc.Code = string(code)
	preproc.Filename = filename

	runPreprocessor := true
	for runPreprocessor {
		preproc.PreprocessString()
		runPreprocessor = preproc.ChangedFile
	}

	return preproc.Code
}

func (ppc *Preprocessor) PreprocessString() {
	// reset the changed flag
	ppc.ChangedFile = false

	// lex the file into tokens
	ppc.Tokens = lexer.LexInternal([]rune(ppc.Code), ppc.Filename, false)

	// go through all tokens
	for ppc.current().Kind != lexer.EOF {
		// is this a hashtag token?
		if ppc.current().Kind == lexer.HashtagToken {
			// is this a valid preprocessor statement?
			// following the pattern:
			// #nameOfStatement("argument string");
			if ppc.peek(1).Kind == lexer.IdToken &&
				ppc.peek(2).Kind == lexer.OpenParenthesisToken &&
				ppc.peek(3).Kind == lexer.StringToken &&
				ppc.peek(4).Kind == lexer.CloseParenthesisToken {

				span := ppc.current().Span.SpanBetween(ppc.peek(4).Span)

				// did some buffoon put a semicolon on the end of this??
				// (I like to put semicolons at the end of my preproc statements B) -Red)
				if ppc.peek(5).Kind == lexer.Semicolon {
					span = span.SpanBetween(ppc.peek(5).Span)
				}

				// create a new preprocessor statement
				ppc.Statements = append(ppc.Statements, PreProcStatement{
					Keyword: ppc.peek(1).Value,
					Content: ppc.peek(3).RealValue.(string),
					Span:    span,
				})
			}
		}

		// step the pointer
		ppc.Index++
	}

	// we now have a list of preprocessor statementz!!!!!!!! (mmmmm jes)
	for _, statement := range ppc.Statements {
		ppc.ProcessStatement(statement)
	}

	// ok we done with the statements, bye bye now
	ppc.Statements = make([]PreProcStatement, 0)
	ppc.Index = 0
}

func (ppc *Preprocessor) ProcessStatement(stmt PreProcStatement) {
	switch stmt.Keyword {
	case "attach":
		ppc.ProcessAttachStatement(stmt)
		break

	case "source":
		ppc.ProcessSourceStatement(stmt)
		break

	case "arg":
		ppc.ProcessArgStatement(stmt)
		break
	}
}

func (ppc *Preprocessor) ProcessAttachStatement(stmt PreProcStatement) {
	// load the attached file (if it can be found!)
	fileContents := ReadFile(stmt.Content, stmt.Span)

	// replace the preproc statement with the files content
	ppc.ReplaceSpan(string(fileContents), stmt.Span)

	// we mead chang
	ppc.ChangedFile = true
}

func (ppc *Preprocessor) ProcessSourceStatement(stmt PreProcStatement) {
	// check if the given path is already added to the source list
	for _, s := range *ppc.Sources {
		// this gamer is already in the list (cringe)
		if s == stmt.Content {
			print.Warning(
				"PREPROCESSOR",
				print.FileAlreadyInSourcesWarning,
				stmt.Span,
				"The source file '%s' has already been added to the sources list!",
				s)
			return
		}
	}

	appendedSources := append(*ppc.Sources, stmt.Content)
	*ppc.Sources = appendedSources
}

func (ppc *Preprocessor) ProcessArgStatement(stmt PreProcStatement) {
	appendedArgs := append(*ppc.Args, stmt.Content)
	*ppc.Args = appendedArgs
}

func (ppc *Preprocessor) ReplaceSpan(text string, span print.TextSpan) {
	// replace the span with the given text
	pre := ppc.Code[:span.StartIndex]
	post := ppc.Code[span.EndIndex:]

	//fmt.Println("==> " + ppc.Code[span.StartIndex:span.EndIndex])

	replacement := pre + text + post
	ppc.Code = replacement
}

// ReadFile reads the file and returns a byte array ([]byte) // nah fam we usin runes
// only handles NotExist and Permission error btw
func ReadFile(filename string, errorLocation print.TextSpan) []rune {
	contents, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		print.Error(
			"LEXER",
			print.FileDoesNotExistError,
			errorLocation,
			"file \"%s\" does not exist! Maybe you spelt it wrong?!",
			filename,
		)
		os.Exit(1)
	} else if errors.Is(err, os.ErrPermission) {
		print.Error(
			"LEXER",
			print.FilePermissionError,
			errorLocation,
			"do not have permissions to open file \"%s\"!",
			filename,
		)
		os.Exit(1)
	} else if err != nil {
		print.Error(
			"LEXER",
			print.FileVoidError,
			errorLocation,
			"an unexpected error occurred when reading file \"%s\"!",
			filename,
		)
		os.Exit(1)
	}
	// destroy all CR in the file
	contents = []byte(strings.Replace(string(contents), "\r", "", -1))

	// Offload a copy of contents for error handling
	// Also split at new lines because that makes referencing easier
	print.CodeReference = make([]string, 0)
	print.CodeReference = strings.Split(string(contents), "\n")
	print.SourceFiles[filename] = string(contents)
	return []rune(string(contents))
}
