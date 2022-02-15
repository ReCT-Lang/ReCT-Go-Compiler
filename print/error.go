package print

import (
	_ "encoding/json" // I know JSON is a data interchange but imma use it for storing the error lookup data anyway - tokorv :)))
	"fmt"
	"strings"
)

/* Error messages!
 * ---------------
 * Format: [AREA] ERROR_TYPE Error(Line, Column): Message \n Error look up code: ERROR_CODE
 * AREA = lexer, parser, binder, etc
 * ERROR_TYPE = "UnexpectedCharacter", etc
 * Line/column = self explanatory
 * Message = message of what went wrong
 * ERROR_CODE = an integer representing the error
 *
 * Error messages will also show the line, highlight the token where the error is, and
 * do that underline thing from C ( ^^^^^^^^^^^^^^^^^ ) under the error.
 *
 * Types of error messages:
 * ------------------------
 * DEV: // For in-progress code (when no error message is available)
 * NotImplemented
 * IDK
 *
 * LEXER:
 * UnexpectedCharacter
 * FileDoesNotExit
 * FilePermission
 * FileVoid
 * RealValueConversion
 *
 * PARSER:
 * UnexpectedToken
 *
 * BINDER:
 * DuplicateParameter
 * DuplicateFunction
 * DuplicateVariableDeclaration
 * UndefinedVariableReference
 * TypeFunctionDoesNotExist
 * Conversion
 * ExplicitConversion
 * UnexpectedExpressionStatement
 * OutsideReturn
 * VoidReturn
 * OutsideBreak
 * UnexpectedNonIntegerValue
 * OutsideContinue
 * BinaryOperatorType
 * IncorrectTypeFunctionCall
 * BadNumberOfParameters
 * UndefinedFunctionCall
 * UnaryOperatorType
 * UnknownDataType
 *
 *
 * Printing error messages:
 * ------------------------
 * Example:
 * var myName <-@ "Jerry";
 *            ^^^^^
 * [LEXER] UnexpectedCharacter Error(1, 13): character "@" was not expected!
 * [> Error look up code: 101 (use: rgoc lookup 101, for more information)]
 *
 * "Look up" error messages:
 * -------------------------
 * Cool idea I remember from Rust, is where users could do something along the lines of
 * rgoc error 103 or rgoc lookup ERROR_CODE
 * and it'll give more detailed information of the error than the error message itself.
 * This would require a code because "UnexpectedCharacterError" is a bit of a mouthful to type.
 * I can also put out some documentation on the ERROR_CODEs for people to refer to outside
 * the command line (on Discord for example).
 * this would require a place to work the lookup documentation, I was just going to use JSON because,
 * it's easy to format and "lookup" stuff (since it's a hashmap).
 *
 * Lookup documentation would be split up into: "header", "explanation", "example", "why"
 * Header: Would include the error code, name of error, where abouts it occurs
 * Explanation: why does the error happen? and common causes?
 * Example: A practical example of the error happening.
 * Why: Why does this cause the compiler to stop/output an error?
 *
 * Example of a "lookup" error documentation:
 *
 * -- UnexpectedCharacterError --
 * Name: UnexpectedCharacter
 * Code: 101
 * Area: Lexer
 *
 * An UnexpectedCharacter Error occurs when the Lexer/scanner of the compiler encounters a character that the compiler
 * does not know how to process. Since the compiler does not know how to process this character, it cannot proceed and
 * instead outputs an UnexpectedCharacter Error so the developer of the program can correct the issues and either remove
 * or replace the unexpected character.
 *
 * Example:
 *
 *  1 | var Je@rryNameVariable <- "Jerry";
 *          ^^^^^
 * [LEXER] UnexpectedCharacter Error(1, 7): character "@" was not expected! Compiler does not know how to process this character!
 * [> Error look up code: 101 (use: rgoc lookup 101, for more information)]
 *
 * The compiler does not know how to process this character and therefore can no longer proceed though the code until the
 * bad character is removed or replaced.
 */

// CodeReference stores code for both error lookups and compiler-time error messages.
// It stores code for error lookups, when compiling it is overwritten with the code to compile.
var CodeReference []string = []string{
	"&dyvar &wjerr&r@&wy &w<- &g\"Hello, World\"&g;",
}

// When no data can be found for line, length or column

// Error prints custom error message and code snippet to terminal/console
// Uses old colour formatting method, will switch to Format() later
func Error(area string, _type ErrorType, line, column, length int, message string, fargs ...interface{}) {
	PrintCodeSnippet(line, column, length)
	WriteCF(Cyan, "[%s] ", strings.ToUpper(area))
	WriteC(DarkCyan, string(_type))
	WriteCF(Red, " Error(%d, %d): ", line, column)
	WriteCF(DarkYellow, message, fargs...)
	code := ErrorTypeToCode(_type)
	WriteC(DarkYellow, "\n[> Error look up code: ")
	WriteCF(Cyan, "%d", code)
	WriteC(DarkYellow, " (use: ")
	WriteC(Yellow, "rgoc -lookup ")
	WriteCF(Cyan, "%d", code)
	PrintC(DarkYellow, ", for more information)]\n")
}

// ErrorS basically Error but returns a string instead of printing
func ErrorS(area string, _type ErrorType, line, column, length int, message string, a ...interface{}) string {
	output := PrintCodeSnippetS(line, column, length)
	output += Format(
		fmt.Sprintf(
			"\n&dc[&c%s&dc] %s &rError(&dr%d&r, &dr%d&r): &dy%s\n[> &rError&dy look up code: &c%d&dy (use: &yrgoc -lookup=&c%d&dy, for more information!)]\n",
			strings.ToUpper(area),
			string(_type),
			line,
			column,
			fmt.Sprintf(message, a...),
			ErrorTypeToCode(_type),
			ErrorTypeToCode(_type),
		),
		Gray,
	)
	return output
}

// PrintCodeSnippet does what it says on the label, it prints a snippet of the code in CodeReference.
func PrintCodeSnippet(line, column, length int) {
	if line <= 0 || column <= 0 || length <= 0 {
		return
	}
	PrintCF(White, "\n%d |  %s", line, CodeReference[line-1])
	if column > 3 {
		WriteC(Gray, strings.Repeat(" ", (column)+len(fmt.Sprintf("%d", line))))
		PrintC(Red, strings.Repeat("^", length))
	} else {
		PrintC(Red, strings.Repeat("^", length+column))
	}
}

// PrintCodeSnippetS returns a code snippet string instead of returning it like PrintCodeSnippet.
func PrintCodeSnippetS(line, column, length int) string {
	output := ""
	if line <= 0 || column <= 0 || length <= 0 {
		return output
	}
	output += Format(
		fmt.Sprintf("\n%d | %s", line, CodeReference[line-1]),
		Gray,
	)
	if column > 3 {
		output += "\n" + strings.Repeat(" ", (column)+len(fmt.Sprintf("%d", line)))
		output += Format(
			strings.Repeat("^", length),
			Red,
		)
	} else {
		output += strings.Repeat(" ", (column)+len(fmt.Sprintf("%d", line)))
	}
	return output + "\n"
}

// Still working wonky
func LookUp(code ErrorCode) {
	// Relative path is not good, maybe it's time for the compiler to have some kind of config?
	if data := errorData[code]; data != nil {
		fmt.Println(Format("ReCT-Go-Compiler v&c%s&c - Error Look up &c%d", White, "1.1", int(code)))
		fmt.Println(Format("Name: &m%s", Gray, data["name"]))
		fmt.Println(Format("Area: &b%s\n", Gray, data["area"]))
		fmt.Println(Format(data["explanation"]+"\n", Gray))
		if data["example"] != "" {
			fmt.Println(Format(data["example"]+"\n", Gray))
		}
		if data["additional"] != "" {
			fmt.Println(Format(data["additional"]+"\n", Gray))
		}
		fmt.Println(Format("Still having troubles? Get help on the offical Discord server: &db%s!", Gray, "https://discord.gg/kk9MsnABdF"))
	} else {
		fmt.Println(Format("The error code &c%d&c is &drinvalid&r!", Gray, int(code)))
	}
}

// ErrorCode the numerical representation of an Error, this allows it to be "looked up"
// using the Error lookup system. Each ErrorCode increments by 1 per declaration.
// ErrorCodes are also given additional values depending on the section they come from.
// For example, Errors from the binder have a default value of 3000, this value it added
// onto the iota increment to produce the specific ErrorCode.
type ErrorCode int

const (
	// Developer ErrorCodes (start at 9000) (Why? Fuck logic that's why).
	NotImplementedErrorCode ErrorCode = iota + 9000
	IDKErrorCode                      = iota + 9000 // Depreciated
	NULLErrorCode                     = -1 + 9000

	// Lexer ErrorCodes (start at 1000)
	UnexpectedCharacterErrorCode = iota + 1000 // 1003
	FileDoesNotExistErrorCode    = iota + 1000 // 1004
	FilePermissionErrorCode      = iota + 1000 // 1005
	FileVoidErrorCode            = iota + 1000 // 1006
	RealValueConversionErrorCode = iota + 1000 // 1007

	// Parser ErrorCodes (start at 2000)
	UnexpectedTokenErrorCode = iota + 2000

	// Binder ErrorCodes (start at 3000) (Chonk warning)
	DuplicateParameterErrorCode            = iota + 3000 // 3009
	DuplicateFunctionErrorCode             = iota + 3000
	DuplicateVariableDeclarationErrorCode  = iota + 3000
	UndefinedVariableReferenceErrorCode    = iota + 3000
	TypeFunctionDoesNotExistErrorCode      = iota + 3000
	ConversionErrorCode                    = iota + 3000
	ExplicitConversionErrorCode            = iota + 3000
	UnexpectedExpressionStatementErrorCode = iota + 3000
	OutsideReturnErrorCode                 = iota + 3000
	VoidReturnErrorCode                    = iota + 3000
	OutsideBreakErrorCode                  = iota + 3000
	UnexpectedNonIntegerValueErrorCode     = iota + 3000
	OutsideContinueErrorCode               = iota + 3000
	BinaryOperatorTypeErrorCode            = iota + 3000
	IncorrectTypeFunctionCallErrorCode     = iota + 3000
	BadNumberOfParametersErrorCode         = iota + 3000
	UndefinedFunctionCallErrorCode         = iota + 3000
	UnaryOperatorTypeErrorCode             = iota + 3000
	UnknownDataTypeErrorCode               = iota + 3000
	UnknownStatementErrorCode              = iota + 3000 // 3028
)

type ErrorType string

const (
	// Developer Error
	NotImplementedError ErrorType = "NotImplemented"
	IDK                           = "IDK(cringe)" // Depreciated (because it's kind of dumb)

	// Lexer Errors
	UnexpectedCharacterError = "UnexpectedCharacter"
	FileDoesNotExistError    = "FileDoesNotExist"
	FilePermissionError      = "FilePermission"
	FileVoidError            = "FileVoid"
	RealValueConversionError = "RealValueConversion"

	// Parser Errors
	UnexpectedTokenError = "UnexpectedToken"

	// Binder Errors
	DuplicateParameterError            = "DuplicateParameter"
	DuplicateFunctionError             = "DuplicateFunction"
	DuplicateVariableDeclarationError  = "DuplicateVariableDeclaration"
	UndefinedVariableReferenceError    = "UndefinedVariableReference"
	TypeFunctionDoesNotExistError      = "TypeFunctionDoesNotExist"
	ConversionError                    = "Conversion"
	ExplicitConversionError            = "ExplicitConversion"
	UnexpectedExpressionStatementError = "UnexpectedExpressionStatement"
	OutsideReturnError                 = "OutsideReturn"
	VoidReturnError                    = "VoidReturn"
	OutsideBreakError                  = "OutsideBreak"
	UnexpectedNonIntegerValueError     = "UnexpectedNonIntegerValue"
	OutsideContinueError               = "OutsideContinue"
	BinaryOperatorTypeError            = "BinaryOperatorType"
	IncorrectTypeFunctionCallError     = "IncorrectTypeFunctionCall"
	BadNumberOfParametersError         = "BadNumberOfParameters"
	UndefinedFunctionCallError         = "UndefinedFunctionCall"
	UnaryOperatorTypeError             = "UnaryOperatorType"
	UnknownDataTypeError               = "UnknownDataType"
	UnknownStatementError              = "UnknownStatement"
)

func ErrorTypeToCode(e ErrorType) ErrorCode {
	switch e {
	case UnexpectedCharacterError:
		return UnexpectedCharacterErrorCode
	case DuplicateFunctionError:
		return DuplicateFunctionErrorCode
	case DuplicateVariableDeclarationError:
		return DuplicateVariableDeclarationErrorCode
	case UndefinedVariableReferenceError:
		return UndefinedVariableReferenceErrorCode
	case DuplicateParameterError:
		return DuplicateParameterErrorCode
	case TypeFunctionDoesNotExistError:
		return TypeFunctionDoesNotExistErrorCode
	case ConversionError:
		return ConversionErrorCode
	case ExplicitConversionError:
		return ExplicitConversionErrorCode
	case UnexpectedExpressionStatementError:
		return UnexpectedExpressionStatementErrorCode
	case OutsideReturnError:
		return OutsideReturnErrorCode
	case VoidReturnError:
		return VoidReturnErrorCode
	case OutsideBreakError:
		return OutsideBreakErrorCode
	case UnexpectedNonIntegerValueError:
		return UnexpectedNonIntegerValueErrorCode
	case OutsideContinueError:
		return OutsideContinueErrorCode
	case BinaryOperatorTypeError:
		return BinaryOperatorTypeErrorCode
	case IncorrectTypeFunctionCallError:
		return IncorrectTypeFunctionCallErrorCode
	case BadNumberOfParametersError:
		return BadNumberOfParametersErrorCode
	case UndefinedFunctionCallError:
		return UndefinedFunctionCallErrorCode
	case UnaryOperatorTypeError:
		return UnaryOperatorTypeErrorCode
	case UnknownDataTypeError:
		return UnknownDataTypeErrorCode
	case NotImplementedError:
		return NotImplementedErrorCode
	case IDK:
		return IDKErrorCode
	case UnexpectedTokenError:
		return UnexpectedTokenErrorCode
	case FileDoesNotExistError:
		return FileDoesNotExistErrorCode
	case FilePermissionError:
		return FilePermissionErrorCode
	case FileVoidError:
		return FileVoidErrorCode
	case RealValueConversionError:
		return RealValueConversionErrorCode
	case UnknownStatementError:
		return UnknownStatementErrorCode
	default:
		return NULLErrorCode
	}
}

// errorData stores all the lookup errorData
var errorData = map[ErrorCode]map[string]string{
	NotImplementedErrorCode: {
		"name": "NotImplemented",
		"area": "Developer",
		"code": string(rune(NotImplementedErrorCode)),
		"explanation": `This error is used as a &wplace marker&w for features that are &wnot fully developed&w yet. 
Since the feature it not fully developed, it &rwill not&r have a &wspecific error code&w or type for you to check out.`,
		"example":    "",
		"additional": "If you think a &mNotImplemented&m error is a mistake, please contact one of the main contributors of the project, or contribute the new error yourself.",
	},
	IDKErrorCode: {
		"name":        "IDK(cringe)",
		"area":        "Developer",
		"code":        string(rune(IDKErrorCode)),
		"explanation": `This error is &drdepreciated&dr. It may be used as an alternative for a &mNotImplemented&m Error, please use: &dyrgoc -lookup &c9000&c, for more information.`,
		"example":     "",
		"additional":  "",
	},
	NULLErrorCode: {
		"name":        "NULL",
		"area":        "Developer",
		"code":        string(rune(NULLErrorCode)),
		"explanation": "This error is &mNULL&m!",
		"example":     "",
		"additional":  "",
	},
	UnexpectedCharacterErrorCode: {
		"name": "UnexpectedCharacter",
		"area": "Lexer",
		"code": string(rune(UnexpectedCharacterErrorCode)),
		"explanation": `An &mUnexpectedCharacter&m Error occurs when the &bLexer/scanner&b of the compiler encounters a &wcharacter&w that the &wcompiler&w &rdoes not&r know how to &wprocess&w. 
Since the compiler does not know how to process this character, it &drcannot proceed&dr and instead outputs an &mUnexpectedCharacter&m Error so the developer 
of the program can correct the issues and &weither remove or replace&w the &wunexpected character&w.`,
		"example": ErrorS(
			"Lexer",
			UnexpectedCharacterError,
			1,
			10,
			5,
			"an unexpected character was found \"%s\"! Lexer is unable to process this character! (BadToken)",
			string(CodeReference[0][15]),
		),
		"additional": "This error can cause a &mBadToken&m error in the &bParser&b later on.",
	},
	FileDoesNotExistErrorCode: {
		"name": "FileDoesNotExist",
		"area": "Lexer",
		"code": string(rune(FileDoesNotExistErrorCode)),
		"explanation": `The &wcompiler will check if your file exists&w, and &wif&w it does &drnot&dr the compiler will output this error.
Usually the cause of this error is entering the &rwrong path&r to the file or a &rtypo&r in the file's name.'`,
		"example":    "",
		"additional": "",
	},
	FilePermissionErrorCode: {
		"name": "FilePermission",
		"area": "Lexer",
		"code": string(rune(FilePermissionErrorCode)),
		"explanation": `The compiler will &wtry to open your file&w, and if it cannot it will make a &wseries of checks&w to see &rwhy it can't open your file&r.
In this case, the &wcompiler found your file&w but the compiler doesn't have the &cpermissions to open the file&c.
You may need to &wrun the compiler as administrator&w, &wmove the file&w into a different directory (which can update permissions), 
or directly &wmodify the file's write/read permissions&w.'`,
		"example":    "",
		"additional": "",
	},
	FileVoidErrorCode: {
		"name": "FileVoid",
		"area": "Lexer",
		"code": string(rune(FileVoidErrorCode)),
		"explanation": `This error occurs when the &wcompiler is trying to open your file&w. If opening your &rfile fails&r, the compiler will take a 
series of &wsteps to identify the problem&w, often this leads to a &mFilePermission&m error or a &mFileDoesNotExist&m error.
However, if the compiler &wcannot diagnose the problem&w, it will output a &mFileVoid&m error. 
Put simply, &wsomething is wrong with the file&w, and the &ccompiler doesn't know what&c.`,
		"example":    "",
		"additional": "",
	},
	RealValueConversionErrorCode: {
		"name": "RealValueConversion",
		"area": "Lexer",
		"code": string(rune(RealValueConversionErrorCode)),
		"explanation": `The compiler will try to convert some values like &dyint&dy and &dyfloat&dy into their true values to help down the line.
This can &wissues if the conversion fails&w. &wYou should check your float and int values for oddities.&w
The most likely cause of this error is &drmultiple points in float literals&dr.`,
		"example":    "",
		"additional": "",
	},
	UnexpectedTokenErrorCode: {
		"name": "UnexpectedToken",
		"area": "Parser",
		"code": string(rune(UnexpectedTokenErrorCode)),
		"explanation": `An UnexpectedToken error occurs when the compiler is expecting a different value, identifier, keyword, or operator 
than what was provided. A common cause of this error is the previous occurrence of a &mUnexpectedCharacter&m error. This is because
&unexpectedCharacter&m errors produce a &mBadToken&m which is then processed by the parser to produce an &mUnexpectedToken&m error.`,
		"example": "",
		"additional": `Another common cause of an &mUnexpectedToken&m error is a value, identifier, keyword, or operator appearing
where it shouldn't.'`,
	},
}

/*`1 | &mPrint&m(&g"Here's an example: "&g);
2 | &dyvar&dy Je&r@&rrryNameVariable <- &g"Jerry"&g;
        &dr^^^^^
&c[LEXER] &dcUnexpectedCharacter &rError(&dr2&r, &dr7&r): &dycharacter &g"@"&dy was not expected! Compiler does not know how to process this character!
[> Error look up code: &c1003&dy (use: &brgoc -lookup &c1003&dy, for more information)]`*/
