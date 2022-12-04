package print

import (
	_ "encoding/json" // I know JSON is a data interchange but imma use it for storing the error lookup data anyway - tokorv :)))
	"fmt"
	"os"
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

// ErrorReport is a structure to hold all info about an error
type ErrorReport struct {
	Area        string
	ErrType     ErrorType
	Span        TextSpan
	Message     string
	MessageArgs []interface{}
}

var ErrorList = make([]ErrorReport, 0)

// CodeReference stores code for both error lookups and compiler-time error messages.
// It stores code for error lookups, when compiling it is overwritten with the code to compile.
var CodeReference []string = []string{
	"&dyvar &wjerr&r@&wy &w<- &g\"Hello, World\"&g;",
}

var OutputErrorMessages = true

var SourceFiles = make(map[string]string)

// When no data can be found for line, length or column

// Error prints custom error message and code snippet to terminal/console
// Uses old colour formatting method, will switch to Format() later
func Error(area string, _type ErrorType, span TextSpan, message string, fargs ...interface{}) {
	if OutputErrorMessages {
		PrintCodeSnippet(span)
		WriteCF(Cyan, "[%s] ", strings.ToUpper(area))
		WriteC(DarkCyan, string(_type))
		WriteCF(Red, " Error(%d, %d, %s): ", span.StartLine, span.StartColumn, span.File)
		WriteCF(DarkYellow, message, fargs...)
		code := ErrorTypeToCode(_type)
		WriteC(DarkYellow, "\n[> Error look up code: ")
		WriteCF(Cyan, "%d", code)
		WriteC(DarkYellow, " (use: ")
		WriteC(Yellow, "rgoc -lookup ")
		WriteCF(Cyan, "%d", code)
		PrintC(DarkYellow, ", for more information)]\n")
	}

	// remember this error
	ErrorList = append(ErrorList, ErrorReport{area, _type, span, message, fargs})
}

func CrashIfErrorsFound() {
	if len(ErrorList) > 0 {
		os.Exit(-1)
	}
}

// ErrorS basically Error but returns a string instead of printing
func ErrorS(area string, _type ErrorType, span TextSpan, message string, a ...interface{}) string {
	output := PrintCodeSnippetS(span)
	output += Format(
		fmt.Sprintf(
			"\n&dc[&c%s&dc] %s &rError(&dr%d&r, &dr%d&r): &dy%s\n[> &rError&dy look up code: &c%d&dy (use: &yrgoc -lookup=&c%d&dy, for more information!)]\n",
			strings.ToUpper(area),
			string(_type),
			span.StartLine,
			span.StartColumn,
			fmt.Sprintf(message, a...),
			ErrorTypeToCode(_type),
			ErrorTypeToCode(_type),
		),
		Gray,
	)
	return output
}

// Warning prints custom warning message and code snippet to terminal/console
func Warning(area string, _type ErrorType, span TextSpan, message string, fargs ...interface{}) {
	PrintCodeSnippet(span)
	WriteCF(Cyan, "[%s] ", strings.ToUpper(area))
	WriteC(DarkCyan, string(_type))
	WriteCF(DarkYellow, " Warning(%d, %d, %s): ", span.StartLine, span.StartColumn, span.File)
	WriteCF(Gray, message, fargs...)
	code := ErrorTypeToCode(_type)
	WriteC(DarkYellow, "\n[> Error look up code: ")
	WriteCF(Cyan, "%d", code)
	WriteC(DarkYellow, " (use: ")
	WriteC(Yellow, "rgoc -lookup ")
	WriteCF(Cyan, "%d", code)
	PrintC(DarkYellow, ", for more information)]\n")
}

// PrintCodeSnippet does what it says on the label, it prints a snippet of the code in CodeReference.
func PrintCodeSnippet(span TextSpan) {
	// no file? tough luck, you won't get a snippet
	if span.File == "" {
		return
	}

	errorFile := SourceFiles[span.File]
	errorLines := strings.Split(errorFile, "\n")

	// is the error contained on a single line?
	if span.StartLine == span.EndLine {
		line := errorLines[span.StartLine-1] // lines are 1-indexed

		// output the line
		offset := PrintLineOfCode(span.StartLine, line)
		fmt.Printf("\n")

		// spacer to the start of the error
		WriteC(Gray, strings.Repeat(" ", span.StartColumn+offset-1))

		// add the error marker
		PrintC(Red, strings.Repeat("^", span.EndColumn-span.StartColumn))

		// we don
		return
	}

	// is the error not contained on a single line we'll show all affected lines with a start and end marker

	offset := GetOffset(span.StartLine)

	// spacer to the start of the error
	WriteC(Gray, strings.Repeat(" ", span.StartColumn+offset-1))

	// add the error marker for the start of the borblem
	WriteC(Red, "v")

	// loins
	WriteC(Red, strings.Repeat("-", len(errorLines[span.StartLine])-span.StartColumn+1))

	// print the lines
	for i := span.StartLine; i < span.EndLine; i++ {
		line := errorLines[i-1] // lines are 1-indexed

		// output the line
		PrintLineOfCode(i, line)
	}
	fmt.Printf("\n")

	// spacer to the start of the error
	WriteC(Red, strings.Repeat("-", span.EndColumn+GetOffset(span.EndLine)+1))

	// add the error marker
	WriteC(Red, "^")

	fmt.Printf("\n")
}

func PrintLineOfCode(nr int, line string) int {
	prefix := fmt.Sprintf("%d |  ", nr)
	WriteCF(White, "\n%s%s", prefix, strings.Replace(line, "\t", " ", -1))

	return len(prefix)
}

func GetOffset(nr int) int {
	return len(fmt.Sprintf("%d |  ", nr))
}

// commented out and replaced with new method -Red
//func PrintCodeSnippet(span TextSpan) {
//	if span.StartLine <= 0 || span.StartColumn <= 0 || span.EndIndex-span.StartIndex <= 0 {
//		return
//	}
//	PrintCF(White, "\n%d |  %s", span.StartLine, CodeReference[span.StartLine-1])
//	if span.StartColumn > 3 {
//		WriteC(Gray, strings.Repeat(" ", (span.StartColumn)+len(fmt.Sprintf("%d", span.StartLine))))
//		PrintC(Red, strings.Repeat("^", span.EndIndex-span.EndIndex))
//	} else {
//		PrintC(Red, strings.Repeat("^", span.EndIndex-span.EndIndex+span.StartColumn))
//	}
//}

// PrintCodeSnippetS returns a code snippet string instead of returning it like PrintCodeSnippet.
func PrintCodeSnippetS(span TextSpan) string {
	output := ""
	if span.StartLine <= 0 || span.StartColumn <= 0 || span.EndIndex-span.StartIndex <= 0 {
		return output
	}
	output += Format(
		fmt.Sprintf("\n%d | %s", span.StartLine, CodeReference[span.StartLine-1]),
		Gray,
	)
	if span.StartColumn > 3 {
		output += "\n" + strings.Repeat(" ", (span.StartColumn)+len(fmt.Sprintf("%d", span.StartLine)))
		output += Format(
			strings.Repeat("^", span.EndIndex-span.EndIndex),
			Red,
		)
	} else {
		output += strings.Repeat(" ", (span.StartColumn)+len(fmt.Sprintf("%d", span.StartLine)))
	}
	return output + "\n"
}

// LookUp prints an explanation of an error message
func LookUp(code ErrorCode) {
	// Relative path is not good, maybe it's time for the compiler to have some kind of config?
	if data := errorData[code]; data != nil {
		fmt.Println(Format("ReCT-Go-Compiler v&c%s&c - Error Look up &c%d", White, "1.1", int(code)))
		fmt.Println(Format("Name: &m%s", Gray, data["name"]))
		fmt.Println(Format("Area: &b%s", Gray, data["area"]))
		fmt.Println(Format("Code: &c%d\n", Gray, code))
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

// ErrorType stores the string value of errors, it's helpful to have all the error in one place.
type ErrorType string

const (
	// Developer Error
	NotImplementedError ErrorType = "NotImplemented"
	IDK                           = "IDK(cringe)" // Depreciated (because it's kind of dumb)

	// Preprocessor Errors
	FileAlreadyInSourcesWarning = "FileAlreadyInSources Warning"

	// Lexer Errors
	UnexpectedCharacterError = "UnexpectedCharacter"
	FileDoesNotExistError    = "FileDoesNotExist"
	FilePermissionError      = "FilePermission"
	FileVoidError            = "FileVoid"
	RealValueConversionError = "RealValueConversion"

	// Parser Errors
	UnexpectedTokenError = "UnexpectedToken"

	// Binder Errors
	DuplicateParameterError               = "DuplicateParameter"
	DuplicateFunctionError                = "DuplicateFunction"
	DuplicateVariableDeclarationError     = "DuplicateVariableDeclaration"
	DuplicatePackageImportError           = "DuplicatePackageImportError"
	UndefinedVariableReferenceError       = "UndefinedVariableReference"
	TypeFunctionDoesNotExistError         = "TypeFunctionDoesNotExist"
	ConversionError                       = "Conversion"
	ExplicitConversionError               = "ExplicitConversion"
	UnexpectedExpressionStatementError    = "UnexpectedExpressionStatement"
	OutsideReturnError                    = "OutsideReturn"
	VoidReturnError                       = "VoidReturn"
	OutsideBreakError                     = "OutsideBreak"
	UnexpectedNonIntegerValueError        = "UnexpectedNonIntegerValue"
	OutsideContinueError                  = "OutsideContinue"
	BinaryOperatorTypeError               = "BinaryOperatorType"
	IncorrectTypeFunctionCallError        = "IncorrectTypeFunctionCall"
	BadNumberOfParametersError            = "BadNumberOfParameters"
	UndefinedFunctionCallError            = "UndefinedFunctionCall"
	UnaryOperatorTypeError                = "UnaryOperatorType"
	UnknownDataTypeError                  = "UnknownDataType"
	UnknownStatementError                 = "UnknownStatement"
	IllegalVariableDeclarationError       = "IllegalVariableDeclarationError"
	IllegalFunctionSignatureError         = "IllegalFunctionSignatureError"
	IllegalNestedClassesError             = "IllegalNestedClassesError"
	InvalidStatementPlacementError        = "InvalidStatementPlacementError"
	OutsideConstructorCallError           = "OutsideConstructorCallError"
	InvalidClassAccessError               = "InvalidClassAccessError"
	IllegalConstructorCallError           = "IllegalConstructorCallError"
	TernaryOperatorTypeError              = "TernaryOperatorTypeError"
	UnknownClassError                     = "UnknownClassError"
	UnknownStructError                    = "UnknownStructError"
	FunctionAccessViolationError          = "FunctionAccessViolationError"
	UnknownFieldError                     = "UnknownFieldError"
	InvalidNumberOfSubtypesError          = "InvalidNumberOfSubtypesError"
	UnknownPackageError                   = "UnknownPackageError"
	UnexpectedNonArrayValueError          = "UnexpectedNonArrayValueError"
	InvalidExternalFunctionPlacementError = "InvalidExternalFunctionPlacementError"
	UnexpectedNonPointerValueError        = "UnexpectedNonPointerValueError"
	TooManyStructParametersError          = "TooManyStructParametersError"
	OutsideThisError                      = "OutsideThisError"

	// Emitter Errors
	UnknownVTableError       = "UnknownVTableError"
	UnknownConstructorError  = "UnknownConsructorError"
	CAdapterCompilationError = "CAdapterCompilationError"
	ExternalCAdapterWarning  = "ExternalCAdapterWarning"

	// Packager Errors
	UnknownPackageModuleFileError     = "UnknownPackageModuleFileError"
	IllegalBoxedTypeError             = "IllegalBoxedTypeError"
	IllegalUnspecificArrayTypeError   = "IllegalUnspecificArrayTypeError"
	MonkeError                        = "MonkeError"
	InvalidNonPointerReferenceError   = "InvalidNonPointerReferenceError"
	UnparsableFingerprintError        = "UnparsableFingerprintError"
	ImpossibleFunctionProcessingError = "ImpossibleFunctionProcessingError"
	ImpossibleFieldProcessingError    = "ImpossibleFieldProcessingError"
)

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

	// Preprocessor ErrorCodes
	FileAlreadyInSourcesWarningCode = iota

	// Lexer ErrorCodes (start at 1000)
	UnexpectedCharacterErrorCode = iota + 1000 // 1003
	FileDoesNotExistErrorCode    = iota + 1000 // 1004
	FilePermissionErrorCode      = iota + 1000 // 1005
	FileVoidErrorCode            = iota + 1000 // 1006
	RealValueConversionErrorCode = iota + 1000 // 1007

	// Parser ErrorCodes (start at 2000)
	UnexpectedTokenErrorCode = iota + 2000

	// Binder ErrorCodes (start at 3000) (Chonk warning)
	DuplicateParameterErrorCode               = iota + 3000 // 3009
	DuplicateFunctionErrorCode                = iota + 3000
	DuplicateVariableDeclarationErrorCode     = iota + 3000
	DuplicatePackageImportErrorCode           = iota + 3000
	UndefinedVariableReferenceErrorCode       = iota + 3000
	TypeFunctionDoesNotExistErrorCode         = iota + 3000
	ConversionErrorCode                       = iota + 3000
	ExplicitConversionErrorCode               = iota + 3000
	UnexpectedExpressionStatementErrorCode    = iota + 3000
	OutsideReturnErrorCode                    = iota + 3000
	VoidReturnErrorCode                       = iota + 3000
	OutsideBreakErrorCode                     = iota + 3000
	UnexpectedNonIntegerValueErrorCode        = iota + 3000
	OutsideContinueErrorCode                  = iota + 3000
	BinaryOperatorTypeErrorCode               = iota + 3000
	IncorrectTypeFunctionCallErrorCode        = iota + 3000
	BadNumberOfParametersErrorCode            = iota + 3000
	UndefinedFunctionCallErrorCode            = iota + 3000
	UnaryOperatorTypeErrorCode                = iota + 3000
	UnknownDataTypeErrorCode                  = iota + 3000 // NOTE: this error does not exist, and I have no idea why!
	UnknownStatementErrorCode                 = iota + 3000
	IllegalVariableDeclarationErrorCode       = iota + 3000
	IllegalFunctionSignatureErrorCode         = iota + 3000
	IllegalNestedClassesErrorCode             = iota + 3000
	InvalidStatementPlacementErrorCode        = iota + 3000
	OutsideConstructorCallErrorCode           = iota + 3000
	InvalidClassAccessErrorCode               = iota + 3000
	IllegalConstructorCallErrorCode           = iota + 3000
	TernaryOperatorTypeErrorCode              = iota + 3000
	UnknownClassErrorCode                     = iota + 3000
	FunctionAccessViolationErrorCode          = iota + 3000
	UnknownFieldErrorCode                     = iota + 3000
	InvalidNumberOfSubtypesErrorCode          = iota + 3000
	UnknownPackageErrorCode                   = iota + 3000
	UnexpectedNonArrayValueErrorCode          = iota + 3000
	InvalidExternalFunctionPlacementErrorCode = iota + 3000
	UnexpectedNonPointerValueErrorCode        = iota + 3000
	UnknownStructErrorCode                    = iota + 3000
	TooManyStructParametersErrorCode          = iota + 3000
	OutsideThisErrorCode                      = iota + 3000

	// Emitter ErrorCodes
	UnknownVTableErrorCode       = iota + 4000
	UnknownConstructorErrorCode  = iota + 4000
	CAdapterCompilationErrorCode = iota + 4000
	ExternalCAdapterWarningCode  = iota + 4000

	// Packager ErrorCodes
	UnknownPackageModuleFileErrorCode     = iota + 5000
	IllegalBoxedTypeErrorCode             = iota + 5000
	IllegalUnspecificArrayTypeErrorCode   = iota + 5000
	MonkeErrorCode                        = iota + 5000
	InvalidNonPointerReferenceErrorCode   = iota + 5000
	UnparsableFingerprintErrorCode        = iota + 5000
	ImpossibleFunctionProcessingErrorCode = iota + 5000
	ImpossibleFieldProcessingErrorCode    = iota + 5000
)

var ErrorTypeCodeRelations = map[ErrorType]ErrorCode{
	UnexpectedCharacterError:              UnexpectedCharacterErrorCode,
	NotImplementedError:                   NotImplementedErrorCode,
	FileDoesNotExistError:                 FileDoesNotExistErrorCode,
	FilePermissionError:                   FilePermissionErrorCode,
	FileVoidError:                         FileVoidErrorCode,
	RealValueConversionError:              RealValueConversionErrorCode,
	UnexpectedTokenError:                  UnexpectedTokenErrorCode,
	DuplicateParameterError:               DuplicateParameterErrorCode,
	DuplicateFunctionError:                DuplicateFunctionErrorCode,
	DuplicateVariableDeclarationError:     DuplicateVariableDeclarationErrorCode,
	UndefinedVariableReferenceError:       UndefinedVariableReferenceErrorCode,
	TypeFunctionDoesNotExistError:         TypeFunctionDoesNotExistErrorCode,
	ConversionError:                       ConversionErrorCode,
	ExplicitConversionError:               ExplicitConversionErrorCode,
	UnexpectedExpressionStatementError:    UnexpectedExpressionStatementErrorCode,
	OutsideReturnError:                    OutsideReturnErrorCode,
	VoidReturnError:                       VoidReturnErrorCode,
	OutsideBreakError:                     OutsideBreakErrorCode,
	UnexpectedNonIntegerValueError:        UnexpectedNonIntegerValueErrorCode,
	OutsideContinueError:                  OutsideContinueErrorCode,
	BinaryOperatorTypeError:               BinaryOperatorTypeErrorCode,
	IncorrectTypeFunctionCallError:        IncorrectTypeFunctionCallErrorCode,
	BadNumberOfParametersError:            BadNumberOfParametersErrorCode,
	UndefinedFunctionCallError:            UndefinedFunctionCallErrorCode,
	UnaryOperatorTypeError:                UnaryOperatorTypeErrorCode,
	UnknownDataTypeError:                  UnknownDataTypeErrorCode,
	UnknownStatementError:                 UnknownStatementErrorCode,
	IllegalVariableDeclarationError:       IllegalVariableDeclarationErrorCode,
	IllegalFunctionSignatureError:         IllegalFunctionSignatureErrorCode,
	IllegalNestedClassesError:             IllegalNestedClassesErrorCode,
	InvalidStatementPlacementError:        InvalidStatementPlacementErrorCode,
	DuplicatePackageImportError:           DuplicatePackageImportErrorCode,
	OutsideConstructorCallError:           OutsideConstructorCallErrorCode,
	InvalidClassAccessError:               InvalidClassAccessErrorCode,
	IllegalConstructorCallError:           IllegalConstructorCallErrorCode,
	TernaryOperatorTypeError:              TernaryOperatorTypeErrorCode,
	UnknownClassError:                     UnknownClassErrorCode,
	FunctionAccessViolationError:          FunctionAccessViolationErrorCode,
	UnknownFieldError:                     UnknownFieldErrorCode,
	InvalidNumberOfSubtypesError:          InvalidNumberOfSubtypesErrorCode,
	UnknownPackageError:                   UnknownPackageErrorCode,
	UnknownVTableError:                    UnknownVTableErrorCode,
	UnknownConstructorError:               UnknownConstructorErrorCode,
	UnknownPackageModuleFileError:         UnknownPackageModuleFileErrorCode,
	IllegalBoxedTypeError:                 IllegalBoxedTypeErrorCode,
	IllegalUnspecificArrayTypeError:       IllegalUnspecificArrayTypeErrorCode,
	MonkeError:                            MonkeErrorCode,
	InvalidNonPointerReferenceError:       InvalidNonPointerReferenceErrorCode,
	UnparsableFingerprintError:            UnparsableFingerprintErrorCode,
	UnexpectedNonArrayValueError:          UnexpectedNonArrayValueErrorCode,
	ImpossibleFunctionProcessingError:     ImpossibleFunctionProcessingErrorCode,
	ImpossibleFieldProcessingError:        ImpossibleFieldProcessingErrorCode,
	FileAlreadyInSourcesWarning:           FileAlreadyInSourcesWarningCode,
	InvalidExternalFunctionPlacementError: InvalidExternalFunctionPlacementErrorCode,
	UnexpectedNonPointerValueError:        UnexpectedNonPointerValueErrorCode,
	UnknownStructError:                    UnknownStructErrorCode,
	TooManyStructParametersError:          TooManyStructParametersErrorCode,
	CAdapterCompilationError:              CAdapterCompilationErrorCode,
	ExternalCAdapterWarning:               ExternalCAdapterWarningCode,
	OutsideThisError:                      OutsideThisErrorCode,
}

// ErrorTypeToCode https://discord.com/channels/751171532398788720/937451421702455306/943557950260269179
func ErrorTypeToCode(e ErrorType) ErrorCode {
	code, ok := ErrorTypeCodeRelations[e]
	if !ok {
		return NULLErrorCode
	} else {
		return code
	}
}

// errorData stores all the lookup errorData
// This is super long so let me give a more in-detail explanation!
// Each error there values:
//   - "name" 		-> Stores the name of the error
//   - "area" 		-> Stores where in the compiler the error is called
//   - "code" 		-> An integer user can use to "lookup" the error
//   - "explanation" 	-> An explanation of what the error is/why it occurs.
//   - "example" 		-> This is generated using ErrorS() and is an example of the error in practice.
//     -> Some error messages will not have an example as they may not be caused by the code itself.
//     -> CodeReference stores all example code in a string array. This is wiped at compile time.
//   - "additional" 	-> This is for additional information, sometimes relating to side effects or further explanation of
//     -> the example code.
var errorData = map[ErrorCode]map[string]string{
	NotImplementedErrorCode: {
		"name": "NotImplemented",
		"area": "Developer",
		"explanation": `This error is used as a &wplace marker&w for features that are &wnot fully developed&w yet. 
Since the feature it not fully developed, it &rwill not&r have a &wspecific error code&w or type for you to check out.`,
		"example":    "",
		"additional": "If you think a &mNotImplemented&m error is a mistake, please contact one of the main contributors of the project, or contribute the new error yourself.",
	},
	IDKErrorCode: {
		"name":        "IDK(cringe)",
		"area":        "Developer",
		"explanation": `This error is &drdepreciated&dr. It may be used as an alternative for a &mNotImplemented&m Error, please use: &dyrgoc -lookup &c9000&c, for more information.`,
		"example":     "",
		"additional":  "",
	},
	NULLErrorCode: {
		"name":        "NULL",
		"area":        "Developer",
		"explanation": "This error is &mNULL&m!",
		"example":     "",
		"additional":  "",
	},
	UnexpectedCharacterErrorCode: {
		"name": "UnexpectedCharacter",
		"area": "Lexer",
		"explanation": `An &mUnexpectedCharacter&m Error occurs when the &bLexer/scanner&b of the compiler encounters a &wcharacter&w that the &wcompiler&w &rdoes not&r know how to &wprocess&w. 
Since the compiler does not know how to process this character, it &drcannot proceed&dr and instead outputs an &mUnexpectedCharacter&m Error so the developer 
of the program can correct the issues and &weither remove or replace&w the &wunexpected character&w.`,
		//"example": ErrorS(
		//	"Lexer",
		//	UnexpectedCharacterError,
		//	1,
		//	10,
		//	5,
		//	"an unexpected character was found \"%s\"! Lexer is unable to process this character! (BadToken)",
		//	string(CodeReference[0][15]),
		//),
		"additional": "This error can cause a &mBadToken&m error in the &bParser&b later on.",
	},
	FileDoesNotExistErrorCode: {
		"name": "FileDoesNotExist",
		"area": "Lexer",
		"explanation": `The &wcompiler will check if your file exists&w, and &wif&w it does &drnot&dr the compiler will output this error.
Usually the cause of this error is entering the &rwrong path&r to the file or a &rtypo&r in the file's name.'`,
		"example":    "",
		"additional": "",
	},
	FilePermissionErrorCode: {
		"name": "FilePermission",
		"area": "Lexer",
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
		"explanation": `The compiler will try to convert some values like &dyint&dy and &dyfloat&dy into their true values to help down the line.
This can &wissues if the conversion fails&w. &wYou should check your float and int values for oddities.&w
The most likely cause of this error is &drmultiple points in float literals&dr.`,
		"example":    "",
		"additional": "",
	},
	UnexpectedTokenErrorCode: {
		"name": "UnexpectedToken",
		"area": "Parser",
		"explanation": `An UnexpectedToken error occurs when the compiler is expecting a different value, identifier, keyword, or operator 
than what was provided. A common cause of this error is the previous occurrence of a &mUnexpectedCharacter&m error. This is because
&unexpectedCharacter&m errors produce a &mBadToken&m which is then processed by the parser to produce an &mUnexpectedToken&m error.`,
		"example": "",
		"additional": `Another common cause of an &mUnexpectedToken&m error is a value, identifier, keyword, or operator appearing
where it shouldn't.'`,
	},
	DuplicateParameterErrorCode: {
		"name": "DuplicateParameter",
		"area": "Binder",
		"explanation": `This error is caused by multiple of the same parameter declared in a single function declaration.
In order to fix this, you need to remove the duplicate parameter or rename one of the parameters so they
appear to be different.`,
		"example":    "",
		"additional": "",
	},
	DuplicateFunctionErrorCode: {
		"name": "DuplicateFunction",
		"area": "Binder",
		"explanation": `This error occurs when the compiler detects multiple functions of the same name are being defined.
The compiler will always detect the second declaration as it can only check the function with previously processed
function symbols. In order to fix this issue, the user needs to change the name of one of the functions.`,
		"example":    "",
		"additional": "",
	},
	DuplicateVariableDeclarationErrorCode: {
		"name": "DuplicateVariableDeclaration",
		"area": "Binder",
		"explanation": `Similar to &mDuplicantFunction&m and &mDuplicantParameter&m errors, &DuplicantVariableDeclaration&m
error occurs when two variables of the same name are defined within the same or parent-to-child scope.
In order to fix this, the user needs to change the name or remove one of the variable declarations.`,
		"example":    "",
		"additional": "",
	},
	UndefinedVariableReferenceErrorCode: {
		"name": "UndefinedVariableReference",
		"area": "Binder",
		"explanation": `The compiler must have previous record of a variable's existence to ensure that the variable
the user is referencing exists. If the variable does not exist the compile will produce this error.
The variable you are trying to reference may have a typo in the name or be declared in a scope that compiler is not considering.`,
		"example":    "",
		"additional": "",
	},
	TypeFunctionDoesNotExistErrorCode: {
		"name": "TypeFunctionDoesNotExist",
		"area": "Binder",
		"explanation": `A type function that was referenced in the code doesn't exist!
A type function is a function accessed through a variable of a specific type such as GetLength() on a string variable.
This error occurs when the function in question doesn't exist for that datatype. `,
		"example":    "",
		"additional": "",
	},
	ConversionErrorCode: {
		"name": "Conversion",
		"area": "Binder",
		"explanation": `This error occurs when a program attempts to convert from one type to another type but the conversion
doesn't exist. This means the compiler doesn't know how to convert between the types and therefore returns and error.`,
		"example":    "",
		"additional": "",
	},
	ExplicitConversionErrorCode: {
		"name": "ExplicitConversion",
		"area": "Binder",
		"explanation": `Similar to &mConversion&m error, this error occurs when the program tries to convert from one type
to another but does not know how. However, in this case, you can write an explicit type cast which allows the compiler to 
understand which type to convert to.`,
		"example":    "",
		"additional": "",
	},
	UnexpectedExpressionStatementErrorCode: {
		"name": "UnexpectedExpressionStatement",
		"area": "Binder",
		"explanation": `&mUnexpectedExpressionStatement&m error occurs when an expression other than call or assignment
is used as a statement. Only specific expressions are allowed to be used as statements such as 
function calls and variable assignments.`,
		"example":    "",
		"additional": "",
	},
	OutsideReturnErrorCode: {
		"name": "OutsideReturn",
		"area": "Binder",
		"explanation": `This error occurs when a return statement is used outside of a function.
to fix this, you will need to remove the return statement. Maybe you put it in the wrong scope?`,
		"example":    "",
		"additional": "",
	},
	VoidReturnErrorCode: {
		"name": "VoidReturn",
		"area": "Binder",
		"explanation": `This error occurs when a return statement is used inside of a void function.
A void function cannot return any value and therefore a return statement is not allowed.
Similar to &mOutsideReturn&m error, you will need to remove the return statement.`,
		"example":    "",
		"additional": "",
	},
	OutsideBreakErrorCode: {
		"name": "OutsideBreak",
		"area": "Binder",
		"explanation": `This error occurs when a break statement is used outside of a loop.
A break statement cannot be used outside of a loop; the compiler does not know how to 
manage a break statement if it is not inside a loop.
You must remove the break statement. Maybe it is in the wrong scope?'`,
		"example":    "",
		"additional": "",
	},
	UnexpectedNonIntegerValueErrorCode: {
		"name": "UnexpectedNonIntegerValue",
		"area": "Binder",
		"explanation": `This error occurs when the compiler is expecting an integer value (literal, or expression),
but instead finds a different type. To fix this, you must remove the non-integer value and replace it with
and integer value.`,
		"example":    "",
		"additional": "",
	},
	OutsideContinueErrorCode: {
		"name": "OutsideContinue",
		"area": "Binder",
		"explanation": `This error occurs when a continue statement is used outside of a loop.
A continue statement cannot be used outside of a loop as it's functionality is to do with the 
loop it is contained within. To fix this, you must remove the continue statement or place it inside a loop.`,
		"example":    "",
		"additional": "",
	},
	BinaryOperatorTypeErrorCode: {
		"name": "BinaryOperatorType",
		"area": "Binder",
		"explanation": `This error occurs when the user attempts to use a binary operator between two types the
compiler does not know how to use the binary operator with.`,
		"example":    "",
		"additional": "",
	},
	IncorrectTypeFunctionCallErrorCode: {
		"name": "IncorrectTypeFunctionCall",
		"area": "Binder",
		"explanation": `A type function is a function that can be accessed and applied to it's type.
An example of a type function would be GetLength on a string datatype.
This error occurs when a type function is used on a datatype that doesn't have access to that 
type function. The compiler does not know how to use the type function on that datatype and 
therefore, the compiler error.`,
		"example":    "",
		"additional": "",
	},
	BadNumberOfParametersErrorCode: {
		"name": "BadNumberOfParameters",
		"area": "Binder",
		"explanation": `A function call expects a certain number of arguments but too many or too little
arguments are provided. The function can only be ran if it has the correct number of argument it expects.`,
		"example":    "",
		"additional": "",
	},
	UndefinedFunctionCallErrorCode: {
		"name": "UndefinedFunctionCall",
		"area": "Binder",
		"explanation": `This error occurs when the user tries to use a function call that does not exist.
This error is similar to &mTypeFunctionDoesNotExit&m and &mUndefinedVariableReference&m as in both cases the 
user is trying to access language constructs that don't exist.
Usually this error is caused by a typo in the function call name.'`,
		"example":    "",
		"additional": "",
	},
	UnaryOperatorTypeErrorCode: {
		"name": "UnaryOperatorType",
		"area": "Binder",
		"explanation": `Similar to &mBinaryOperatorType&m error, &mUnaryOperatorType&m error occurs when 
the user tries to use the unary operator with two types that the compiler does not now how to process with
a particular unary operator.`,
		"example":    "",
		"additional": "",
	},
	UnknownStatementErrorCode: {
		"name": "UnknownStatement",
		"area": "Binder",
		"explanation": `This error occurs when a statement is found that should not exist. The compiler checks through all
the possible statement types (like ifStatement, whileStatement, etc) and the one it found does not match any that exist.`,
		"example":    "",
		"additional": "",
	},
}
