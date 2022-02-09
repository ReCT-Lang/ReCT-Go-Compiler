package print

import "strings"

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

/* Everything is in place, I'll write the rest tonight, I'm going out to eat - Tokorv */
// Also I write a lot of docs lol

type ErrorCode int

const (
	// Developer ErrorCodes (start at 9000)
	NotImplementedErrorCode ErrorCode = 9000
	IDKErrorCode                      = 9001

	// Lexer ErrorCodes (start at 1000)
	UnexpectedCharacterErrorCode = 1000
	FileDoesNotExitErrorCode     = 1001
	FilePermissionErrorCode      = 1002
	FileVoidErrorCode            = 1003
	RealValueConversion          = 1004
)

type ErrorType string

const (
	UnexpectedCharacterError ErrorType = "UnexpectedCharacter"
	FileDoesNotExitError               = "FileDoesNotExit"
	FilePermissionError                = "FilePermission"
	FileVoidError                      = "FileVoid"
	RealValueConversionError           = "RealValueConversionError"
)

func Error(area string, _type ErrorType, line int, column int, message string, fargs ...interface{}) {
	WriteCF(Cyan, "[%s] ", strings.ToUpper(area))
	WriteC(DarkCyan, string(_type))
	WriteCF(Red, " Error(%d, %d): ", line, column)
	WriteCF(DarkYellow, message, fargs...)
	code := ErrorTypeToCode(_type)
	WriteCF(Green, "[> Error look up code: %d (use: rgoc lookup %d, for more information)])", code, code)
}

func ErrorTypeToCode(e ErrorType) ErrorCode {
	switch e {
	case UnexpectedCharacterError:
		return UnexpectedCharacterErrorCode
	default:
		return -1
	}
}
