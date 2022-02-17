# ReCT Lexer/Scanner
This lexer is currently for the 1.1 version of ReCT and handles all the necessary Token types needed for that version.
The source code review goes into more detail about how the lexer works.

# Source Code Review
The lexer is made up of various structures and functions that are used to convert the source code into Tokens.
Tokens are important values in the source code our program needs in order to understand what the source code does.
We translate the source code character by character and check for certain characteristics in each character,
we bundle characters together and assign them TokenKinds/TokenTypes.

## Significant Constructs

### Lexer structure
The lexer structure is a [struct](https://go.dev/tour/moretypes/2) that contains important information about our source code. The structure was designed
to be internal meaning the structure only needs to be used within `lexer.go`, all the necessary information are returned
by the function `Lex()`.
```go
type Lexer struct {
	Code   []rune
	Line   int
	Column int
	Index  int
	Tokens []Token
}
```
The structure stores the source code in an array of [runes](https://go.dev/tour/basics/11) called `Code`, it also stores the current index used for 
getting characters from the `Code`, the `Line` and `Column` are used to keep track of the current position in the
source code, and finally `Tokens` is a Token array which stores the information we need to get out of the Lexer.

### Tokens
Tokens are the product of our lexer, they tell us everything we need to know about the source code.
A single Token is defined by its own struct** shown below. This stores the `Value` of our Token, it's `Kind`,
and the `Line` and `Column` of where our token is in the source code.

```go
type Token struct {
	Value      string
	RealValue  interface{}
	Kind       TokenKind
	Line       int
	Column     int
	SpaceAfter bool
}
```

There are a variety of "constructor" functions in `token.go`; functions used to create tokens, and `TokenKind`
which is an [enum](https://www.sohamkamani.com/golang/enums/) containing all the different types of token the lexer generates.

### Lexical functions
Lexical functions are used to create the token array we see in the lexer structure, these functions are all well 
documented inside the source code itself (just check out `lexer.go`), I'll provide a brief description of what 
each function does below.

- `Lex()` is the start and end point of the lexical analysis, it creates the Lexer instance, loops through each 
character in the source code, and calls the correct function when it encounters a specific kind of character.
- `getId()` is called when `Lex()` finds a letter, it processes an identifier or keyword Token.
- `getNumber()` is called when `Lex()` finds a number, it processes a number token (integers and floats)
- `getString()` is called when `Lex()` finds a ", it processes a string token
- `getComment()` is called when `Lex()` finds a //, it processes a comment (doesn't collect a token)
- `getOperator()` is called when `Lex()` can't find anything else, it tries to process an operator token, but if it can't
find an operator it will display an error and collect a BadToken instead.

## Sideline Constructs

### File handling
In order to process the source code we must open and read the file. This task is simple, but we off load it into its
own function to avoid making `Lex()` messy. This task handled in `handleFileOpen`, this function reads the file or
displays and error if the file failed to open correctly.

### BadToken 
BadTokens are tokens generated in `getOperator()` and represent an unknown character in the source code.
They are generated so the lexical analysis does not stop midway through and continues to lex the entire source code.
The BadTokens are handled later down the line by the parser which does stop the program to allow the user to
correct the mistake.

### Keyword checking
Keyword tokens are collected by `getId`. However, to check they are keywords, `CheckIfKeyword` is used which is a 
simple [switch statement](https://go.dev/tour/flowcontrol/9) that checks if the buffer value is a keyword.

### Error handling
Most errors are handled by `print/error.go`, an example is shown below, this kind of error is present throughout the
compiler source code. 
```go
print.Error(
			"LEXER",
			print.FileVoidError,
			0,
			0,
			5,
			"an unexpected error occurred when reading file \"%s\"!",
			filename,
		)
```


