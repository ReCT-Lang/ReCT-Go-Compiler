package print

import (
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

// Global variable :gentleman:
var CodeReference []string

func Error(area string, _type ErrorType, line int, column int, message string, fargs ...interface{}) {
	PrintCodeSnippet(line, column)
	WriteCF(Cyan, "\n[%s] ", strings.ToUpper(area))
	WriteC(DarkCyan, string(_type))
	WriteCF(Red, " Error(%d, %d): ", line, column)
	WriteCF(DarkYellow, message, fargs...)
	code := ErrorTypeToCode(_type)
	WriteC(DarkYellow, "[> Error look up code: ")
	WriteCF(Cyan, "%d", code)
	WriteC(DarkYellow, " (use: ")
	WriteC(Yellow, "rgoc lookup ")
	WriteCF(Cyan, "%d", code)
	PrintC(DarkYellow, ", for more information)]\n")
}

func PrintCodeSnippet(line int, column int) {
	PrintCF(White, "\n%d |  %s", line, CodeReference[line-1])
	if column > 3 {
		WriteC(Gray, strings.Repeat(" ", (column)+len(fmt.Sprintf("%d", line))))
		PrintC(Red, "^^^^^^^")
	} else {
		PrintC(Red, "^^^^^^^")
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
	IDKErrorCode                      = iota + 9000
	NULLErrorCode                     = -1 + 9000

	// Lexer ErrorCodes (start at 1000)
	UnexpectedCharacterErrorCode = iota + 1000
	FileDoesNotExitErrorCode     = iota + 1000
	FilePermissionErrorCode      = iota + 1000
	FileVoidErrorCode            = iota + 1000
	RealValueConversionErrorCode = iota + 1000

	// Parser ErrorCodes (start at 2000)
	UnexpectedTokenErrorCode = iota + 2000

	// Binder ErrorCodes (start at 3000) (Chonk warning)
	DuplicateParameterErrorCode            = iota + 3000
	DuplicateFunctionErrorCode             = iota + 3000
	DuplicateVariableDeclarationErrorCode  = iota + 3000
	UndefinedVariableReferenceErrorCode    = iota + 3000
	TypeFunctionDoesNotExitErrorCode       = iota + 3000
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
)

type ErrorType string

const (
	// Developer Error
	NotImplementedError ErrorType = "NotImplemented"
	IDK                           = "IDK(cringe)"

	// Lexer Errors
	UnexpectedCharacterError = "UnexpectedCharacter"
	FileDoesNotExitError     = "FileDoesNotExit"
	FilePermissionError      = "FilePermission"
	FileVoidError            = "FileVoid"
	RealValueConversionError = "RealValueConversionError"

	// Parser Errors
	UnexpectedTokenError = "UnexpectedTokenError"

	// Binder Errors
	DuplicateParameterError            = "DuplicateParameter"
	DuplicateFunctionError             = "DuplicateFunction"
	DuplicateVariableDeclarationError  = "DuplicateVariableDeclaration"
	UndefinedVariableReferenceError    = "UndefinedVariableReference"
	TypeFunctionDoesNotExitError       = "TypeFunctionDoesNotExist"
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
	case TypeFunctionDoesNotExitError:
		return TypeFunctionDoesNotExitErrorCode
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
	case FileDoesNotExitError:
		return FileDoesNotExitErrorCode
	case FilePermissionError:
		return FilePermissionErrorCode
	case FileVoidError:
		return FileVoidErrorCode
	case RealValueConversionError:
		return RealValueConversionErrorCode
	default:
		return NULLErrorCode
	}
}
