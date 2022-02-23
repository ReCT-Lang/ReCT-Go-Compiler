package parser

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/rules"
	"os"
)

// Parser : internal struct for assembling the syntax tree
type Parser struct {
	Tokens []lexer.Token
	Index  int
}

// <HELPERS> ------------------------------------------------------------------

// returns current token
func (prs *Parser) current() lexer.Token {
	return prs.peek(0)
}

// returns current token + a given offset
func (prs *Parser) peek(offset int) lexer.Token {
	// if we are out of bounds -> return EOF token
	if prs.Index+offset < 0 || prs.Index+offset >= len(prs.Tokens) {
		return lexer.Token{
			Kind:  lexer.EOF,
			Value: "",
		}
	}

	// if everything is fine -> great!
	return prs.Tokens[prs.Index+offset]
}

// consume a syntax token is It's what we're looking for if not -> just die for now
func (prs *Parser) consume(expected lexer.TokenKind) lexer.Token {
	if prs.current().Kind != expected {
		// can't tell the user which one because there's no way to get a constants name by its value...
		// so we might need a separate name array for them

		// Switched the TokenKind constants to strings and now added the error message you wanted <3

		// Added this additionalInfo section to help users know why a parser error occurred
		additionalInfo := ""
		if prs.current().Kind == lexer.BadToken {
			additionalInfo = " (may be caused by previous \"UnexpectedCharacterError\" which produces a BadToken)"
		}
		print.Error(
			"PARSER",
			print.UnexpectedTokenError,
			prs.current().Line,
			prs.current().Column,
			5,
			"unexpected Token \"%s\"! Expected \"%s\"!"+additionalInfo,
			prs.current().Kind,
			expected,
		)
		os.Exit(-1) // die(-1);
	}

	// if everything is alright -> step our index and return the token
	prs.Index++
	return prs.peek(-1)
}

func (prs *Parser) rewind(to lexer.Token) {
	for prs.current().String(false) != to.String(false) {
		prs.Index--
	}
}

// </HELPERS> -----------------------------------------------------------------

// Parse parse a compilation (with all its functions, classes, enums, global statements, etc...)
func Parse(tokens []lexer.Token) []nodes.MemberNode {
	parser := Parser{
		Tokens: tokens,
		Index:  0,
	}

	return parser.parseMembers()
}

// parseMembers begin parsing one member at a time until we hit lexer.EOF Token
func (prs *Parser) parseMembers() []nodes.MemberNode {
	members := make([]nodes.MemberNode, 0)

	for prs.current().Kind != lexer.EOF {

		startToken := prs.current()

		// parse all our members
		member := prs.parseMember()
		members = append(members, member)

		// if we got stuck somewhere, just keep moving
		if startToken == prs.current() {
			prs.Index++
		}
	}

	return members
}

// parseMember begins parsing either functions or global statements
func (prs *Parser) parseMember() nodes.MemberNode {
	// functions / classes would go here \/
	if prs.current().Kind == lexer.FunctionKeyword {
		return prs.parseFunctionDeclaration()
	}

	// global statements (any statements outside any functions)
	return prs.parseGlobalStatement()
}

// parseGlobalStatement parse a statement and return as a global statement member
// this becomes a member node and eventually is appended onto our member node list
// then we parse the next member in parseMembers.
func (prs *Parser) parseGlobalStatement() nodes.GlobalStatementMember {
	statement := prs.parseStatement()
	return nodes.CreateGlobalStatementMember(statement)
}

// parseFunctionDeclaration checks for a valid order of Tokens, parses all the statements inside the function
// and then returns it as a function declaration member.
func (prs *Parser) parseFunctionDeclaration() nodes.FunctionDeclarationMember {

	// Example:
	// function functionName(functionArg1 string) string { ... }
	prs.consume(lexer.FunctionKeyword) // We know we are getting a functional already so just consume this

	// consume the "functionName" (which is an identifier token).
	identifier := prs.consume(lexer.IdToken)

	// this is where we parse the parameters (e.g., functionArg1)
	prs.consume(lexer.OpenParenthesisToken)
	params := prs.parseParameterList() // We only need the arguments not the parenthesis tokens UwU
	prs.consume(lexer.CloseParenthesisToken)

	// This is optional, not all functions need a return type as not all functions return a value.
	// Our example returns a string type, so we will need the TypeClause later.
	// if there is no return type the typeClause will be empty.
	typeClause := prs.parseOptionalTypeClause()

	// This is where we parse all our statements in the function
	// the block statement will handle multiple statements inside itself
	body := prs.parseBlockStatement()

	return nodes.CreateFunctionDeclarationMember(identifier, params, typeClause, body)
}

// parseParameterList we parse a list of arguments (usually for a function or functionCall)
// In code this looks like: (arg1 string, arg2 int, arg3 string) ... etc
func (prs *Parser) parseParameterList() []nodes.ParameterNode {
	params := make([]nodes.ParameterNode, 0) // list of our arguments

	for prs.current().Kind != lexer.CloseParenthesisToken &&
		prs.current().Kind != lexer.EOF {

		// Get a single argument and add it to the list
		param := prs.parseParameter()
		params = append(params, param)

		// skip over the commas
		if prs.current().Kind == lexer.CommaToken {
			prs.consume(lexer.CommaToken)
		} else {
			// if it isn't a comma what the hell is going on we need to get out of here!
			break
		}
	}

	return params
}

// parseParameter gets a single parameter and returns it
// example: arg1 string
// arg1 is the identifier, string is the typeClause
func (prs *Parser) parseParameter() nodes.ParameterNode {
	identifier := prs.consume(lexer.IdToken)
	typeClause := prs.parseTypeClause()
	return nodes.CreateParameterNode(identifier, typeClause)
}

// parseOptionalTypeClause we check if a typeClause is there, return an empty one if so, otherwise we return a TypeClauseNode
func (prs *Parser) parseOptionalTypeClause() nodes.TypeClauseNode {
	// if there's no type clause, return an empty one
	if prs.current().Kind != lexer.IdToken && prs.current().Kind != lexer.AccessToken {
		return nodes.TypeClauseNode{}
	}

	return prs.parseTypeClause()
}

// parseTypeClause consumes the datatype and returns it in node form
func (prs *Parser) parseTypeClause() nodes.TypeClauseNode {
	// if theres a '->' token, consume it
	if prs.current().Kind == lexer.AccessToken {
		prs.consume(lexer.AccessToken)
	}

	identifier := prs.consume(lexer.IdToken)
	subTypes := make([]nodes.TypeClauseNode, 0)

	// subtypes
	if prs.current().Kind == lexer.OpenBracketToken {
		prs.consume(lexer.OpenBracketToken)

		// parse all subtypes in the [] recursively
		for true {
			subTypes = append(subTypes, prs.parseTypeClause())

			if prs.current().Kind != lexer.CommaToken {
				break
			} else {
				prs.consume(lexer.CommaToken)
			}
		}

		prs.consume(lexer.CloseBracketToken)
	}

	return nodes.CreateTypeClauseNode(identifier, subTypes)
}

// parseUncertainTypeClause consumes the datatype and returns it in node form
// difference to parseTypeClause is that this one can fail safely if we notice that this isn't a type
func (prs *Parser) parseUncertainTypeClause() (nodes.TypeClauseNode, bool) {
	identifier := prs.consume(lexer.IdToken)
	subTypes := make([]nodes.TypeClauseNode, 0)

	// subtypes
	if prs.current().Kind == lexer.OpenBracketToken {
		prs.consume(lexer.OpenBracketToken)

		// parse all subtypes in the [] recursively
		for true {
			// check if this is an identifier
			if prs.current().Kind != lexer.IdToken {
				return nodes.TypeClauseNode{}, false
			}

			subClause, ok := prs.parseUncertainTypeClause()

			// if the subclause failed -> fail
			if !ok {
				return nodes.TypeClauseNode{}, false
			}

			subTypes = append(subTypes, subClause)

			if prs.current().Kind != lexer.CommaToken {
				break
			} else {
				prs.consume(lexer.CommaToken)
			}
		}

		// if there's no close bracket here -> fail
		if prs.current().Kind != lexer.CloseBracketToken {
			return nodes.TypeClauseNode{}, false
		}

		prs.consume(lexer.CloseBracketToken)
	}

	return nodes.CreateTypeClauseNode(identifier, subTypes), true
}

// <STATEMENTS> ---------------------------------------------------------------

// parseStatement Based off the first keyword it'll parse a statement
func (prs *Parser) parseStatement() nodes.StatementNode {
	var statement nodes.StatementNode = nil
	// nil StatementNode can cause segmentation violation if no correct key is found. (handled in parsePrimaryExpression)

	// select correct parsing function based on kind
	cur := prs.current().Kind
	// { ... }
	if cur == lexer.OpenBraceToken {
		statement = prs.parseBlockStatement()
		// var|set name <- "Jerry";
	} else if cur == lexer.VarKeyword || cur == lexer.SetKeyword {
		statement = prs.parseVariableDeclaration()
		// if ( ... ) { ... }
	} else if cur == lexer.IfKeyword {
		statement = prs.parseIfStatement()
		// return ...
	} else if cur == lexer.ReturnKeyword {
		statement = prs.parseReturnStatement()
		// for (... , ... , ...) { ... }
	} else if cur == lexer.ForKeyword {
		statement = prs.parseForStatement()
		// while ( ... ) { ... }
	} else if cur == lexer.WhileKeyword {
		statement = prs.parseWhileStatement()
		// break; (stops a loop)
	} else if cur == lexer.BreakKeyword {
		statement = prs.parseBreakStatement()
		// continue; (skips to the next iteration of a loop)
	} else if cur == lexer.ContinueKeyword {
		statement = prs.parseContinueStatement()
		// from ( ... ) to ... { ... }
	} else if cur == lexer.FromKeyword {
		statement = prs.parseFromToStatement()

	} else {
		// Lastly we process an expression
		statement = prs.parseExpressionStatement()

		// moved the error message to parsePrimaryExpression()
	}

	// if there's a semicolon -> a b s o r b    i t
	if prs.current().Kind == lexer.Semicolon {
		prs.consume(lexer.Semicolon)
	}

	return statement
}

// parseBlockStatement this statement contains a bunch of other statements while also being a statement itself
// Example if { ... } (the "{ ... }" is the block statement).
func (prs *Parser) parseBlockStatement() nodes.BlockStatementNode {
	// create a list for our statement
	statements := make([]nodes.StatementNode, 0)

	// {
	openBrace := prs.consume(lexer.OpenBraceToken)

	for prs.current().Kind != lexer.EOF &&
		prs.current().Kind != lexer.CloseBraceToken {

		startToken := prs.current()

		// We keep getting more statements until we either hit
		statement := prs.parseStatement()
		statements = append(statements, statement)

		// if we got stuck somewhere, just keep moving
		if startToken == prs.current() {
			prs.Index++
		}
	}

	// }
	closeBrace := prs.consume(lexer.CloseBraceToken)

	return nodes.CreateBlockStatementNode(openBrace, statements, closeBrace)
}

// parseVariableDeclaration parses a variable declaration like var x <- 10;
// Example: var name <- "Jerry";
func (prs *Parser) parseVariableDeclaration() nodes.VariableDeclarationStatementNode {

	// Firstly we absorb the var or set keyword
	// In our case (example) we are absorbing var.
	// We don't need to check the kind because we already do that in parseStatement
	keyword := prs.consume(prs.current().Kind)

	// Next we check if there is a type in the variable declaration like:
	// var x string;
	// We assign typeClause and empty node and check, this means the typeClause will remain empty
	// if there is no type in the declaration.
	typeClause := nodes.TypeClauseNode{}
	if prs.current().Kind == lexer.IdToken &&
		(prs.peek(1).Kind == lexer.IdToken || prs.peek(1).Kind == lexer.OpenBracketToken) {
		typeClause = prs.parseTypeClause()
	}

	// We consume the variable name, the assign token (<-) and then parse the value as an expression.
	identifier := prs.consume(lexer.IdToken)
	assign := prs.consume(lexer.AssignToken)

	// parsing an expression allows variable to be defined as 1*2/5+3 or string + string, instead of a single simple value
	initializer := prs.parseExpression()

	return nodes.CreateVariableDeclarationStatementNode(keyword, typeClause, identifier, assign, initializer)
}

// parseIfStatement as you can probably guess, this function is called when parseStatement gets an IfKeyword Token.
func (prs *Parser) parseIfStatement() nodes.IfStatementNode {
	// Remember if statements?
	// if ( ... ) { ... }

	// Here we get the "if"
	keyword := prs.consume(lexer.IfKeyword)

	// Consuming the ( ... )
	prs.consume(lexer.OpenParenthesisToken)
	condition := prs.parseExpression() // ... is an expression, it should evaluate to true or false (this isn't checked here)
	prs.consume(lexer.CloseParenthesisToken)

	// Finally, the statement which is shown as { ... }
	// Though most people put a block statement after an if, ReCT lets you put a single statement if you want.
	// Remember a blockStatement is a type of statement it's still valid.
	statement := prs.parseStatement()

	// To end it, we parse the "else" of an if statement.
	// this will be an empty elseClauseNode
	elseClause := prs.parseElseClause()

	return nodes.CreateIfStatementNode(keyword, condition, statement, elseClause)
}

// parseElseClause this parses else statements
// Example else { ... } (this usually comes after an if statement like: if ( ... ) { ... } else { ... }
func (prs *Parser) parseElseClause() nodes.ElseClauseNode {

	// if theres no else -> dont parse an else lol - Red
	if prs.current().Kind != lexer.ElseKeyword {
		return nodes.ElseClauseNode{}
	}

	// consume keyword and parse statement
	keyword := prs.consume(lexer.ElseKeyword)

	// Same as if statement this can be a single statement or a block statement
	statement := prs.parseStatement()

	return nodes.CreateElseClauseNode(keyword, statement)
}

// parseReturnStatement handles return statements like: return ...
// happens at the end of a function, if you don't know that you should be reading this tbh
func (prs *Parser) parseReturnStatement() nodes.ReturnStatementNode {

	// a b s o r b return keyword
	keyword := prs.consume(lexer.ReturnKeyword)

	var expression nodes.ExpressionNode = nil
	// if we are at the end of the line (;) there's no return value given
	if prs.current().Kind != lexer.Semicolon {
		// mmm yummy expression
		expression = prs.parseExpression()
	}

	return nodes.CreateReturnStatementNode(keyword, expression)
}

// parseForStatement handles for statements (loops)
// Example: for ( ..., ..., ...) { ... }
func (prs *Parser) parseForStatement() nodes.ForStatementNode {
	// First we consume that for keyword, so we can get to the good parts
	keyword := prs.consume(lexer.ForKeyword)

	// now need to parse each statement ( ... ) in ( ..., ..., ... )
	prs.consume(lexer.OpenParenthesisToken)

	// First is the initializer (a statement declared before the loop begins)
	// this is usually: var i = 0, and "i" is used throughout the loop.
	initialiser := prs.parseVariableDeclaration()
	prs.consume(lexer.Semicolon) // a b s o r b ;

	// Second is the condition (if the condition becomes false the loop ends)
	// Examples of this are: i < 10, i < x->GetLength(), etc.
	condition := prs.parseExpression()
	prs.consume(lexer.Semicolon) // a b s o r b ;

	// Finally, is the updation (name stolen from Java) where the initialiser is modified each iteration (usually)
	updation := prs.parseStatement()
	prs.consume(lexer.CloseParenthesisToken) // we've finished with the loop "parameters" now we should get the "contents"

	// Loop contents can be a single statement or sexy block statement with lots of statements
	statement := prs.parseStatement()

	// pretty sure I write this, glad it didn't require correcting - tokorv
	return nodes.CreateForStatementNode(keyword, initialiser, condition, updation, statement)
}

// parseWhileStatement a while loop
// Example: while ( ... ) { ... }
func (prs *Parser) parseWhileStatement() nodes.WhileStatementNode {

	// We already know WhileKeyword is there because we detected it in parseStatement
	// not we consume it to move to the next token.
	keyword := prs.consume(lexer.WhileKeyword)

	// This is where we handle ( ... )
	prs.consume(lexer.OpenParenthesisToken)  // consume (
	condition := prs.parseExpression()       // conditional expression (if false stop loop)
	prs.consume(lexer.CloseParenthesisToken) // consume )

	// Finally, we get all the statements
	// Usually this is a blockStatement, but it can be a single statement like Print()
	statement := prs.parseStatement()

	return nodes.CreateWhileStatementNode(keyword, condition, statement)
}

// parseFromToStatement a from to loop (quite unique to rect I think)
// this loop takes two arguments, a start, and an end.
// Example: from ( ... ) to ... { ... }
// Code Example: from i <- 0 to 100 { Print(string(i)); }
// the above code ill print all numbers from 0 to 100.
func (prs *Parser) parseFromToStatement() nodes.FromToStatementNode {

	// We're expecting it from parseStatement,
	// we consume it, so we can parse the other tokens
	keyword := prs.consume(lexer.FromKeyword)

	// this is where we handle ( ... ) : ( i <- 0 )
	prs.consume(lexer.OpenParenthesisToken)  // (
	identifier := prs.consume(lexer.IdToken) // and identifier (usually "i", but can be valid identifier really)
	prs.consume(lexer.AssignToken)           // <- (assignToken)
	lowerBound := prs.parseExpression()      // Now we get the initial value of the identifier (in: i <- 0, it's the 0)
	prs.consume(lexer.CloseParenthesisToken) // )

	// We should get a "to" keyword next (see example)
	prs.consume(lexer.ToKeyword)

	// and we get the identifiers limit (if the identifier is >= to the limit the loop will break)
	upperBound := prs.parseExpression()

	// and now we get the statement, same as all other loops, this can be a blockStatement or single statement
	statement := prs.parseStatement()
	return nodes.CreateFromToStatementNode(keyword, identifier, lowerBound, upperBound, statement)
}

// parseBreakStatement processes "break" keyword (honesty nothing special, similar process to continue and return)
func (prs *Parser) parseBreakStatement() nodes.BreakStatementNode {
	keyword := prs.consume(lexer.BreakKeyword)

	return nodes.CreateBreakStatement(keyword)
}

// parseContinueStatement processes the "continue" keyword, not much happening really.
func (prs *Parser) parseContinueStatement() nodes.ContinueStatementNode {
	keyword := prs.consume(lexer.ContinueKeyword)

	return nodes.CreateContinueStatement(keyword)
}

// parseExpressionStatement expressions are 2nd class citizens, they come after statements
// If no statement can be found, we try to process an expression.
func (prs *Parser) parseExpressionStatement() nodes.ExpressionStatementNode {
	expression := prs.parseExpression()
	// We basically parse an expression
	return nodes.CreateExpressionStatementNode(expression)
}

// </STATEMENTS> --------------------------------------------------------------
// <EXPRESSIONS> --------------------------------------------------------------

// parseExpression
func (prs *Parser) parseExpression() nodes.ExpressionNode {
	// check for more "complex" expressions first
	// (these are not allowed in binary expressions)
	// if there's none -> parse normal binary expression

	// variable editors
	if prs.current().Kind == lexer.IdToken &&
		prs.peek(1).Kind == lexer.AssignToken &&
		!prs.peek(1).SpaceAfter &&
		rules.GetBinaryOperatorPrecedence(prs.peek(2)) != 0 {
		// if <-+ <-- <-/ <-* : (for the uneducated these are the equivalents of += -= /= *= in ReCT.
		return prs.parseVariableEditorExpression()
	}

	// single variable editors (++ and --)
	if prs.current().Kind == lexer.IdToken &&
		((prs.peek(1).Kind == lexer.PlusToken && prs.peek(2).Kind == lexer.PlusToken) ||
			(prs.peek(1).Kind == lexer.MinusToken && prs.peek(2).Kind == lexer.MinusToken)) {

		identifier := prs.consume(lexer.IdToken)    // the "i" in "i++"
		operator := prs.consume(prs.current().Kind) // the "+"
		prs.consume(prs.current().Kind)             // another "+", we don't have to check because we do that in the if statement above
		return nodes.CreateVariableEditorExpressionNode(identifier, operator, nil, true)
	}

	// assignment expression
	if prs.current().Kind == lexer.IdToken &&
		prs.peek(1).Kind == lexer.AssignToken {
		return prs.parseAssignmentExpression()
	}

	// array creating
	if prs.current().Kind == lexer.MakeKeyword {
		return prs.parseMakeArrayExpression()
	}

	// if none of the above are what we want, it must be a binary expression!
	return prs.parseBinaryExpression(0)
}

// parseBinaryExpression
func (prs *Parser) parseBinaryExpression(parentPrecedence int) nodes.ExpressionNode {
	var left nodes.ExpressionNode

	// check if this is a unary expression
	unaryPrecedence := rules.GetUnaryOperatorPrecedence(prs.current())
	if unaryPrecedence != 0 && unaryPrecedence > parentPrecedence {
		operator := prs.consume(prs.current().Kind)
		operand := prs.parseBinaryExpression(unaryPrecedence)
		return nodes.CreateUnaryExpressionNode(operator, operand)

		// if not, start by parsing our left expression
	} else {
		left = prs.parsePrimaryExpression()
	}

	// funky while(true) but go-style
	for {
		precedence := rules.GetBinaryOperatorPrecedence(prs.current())

		// if this isn't an operator, or it has less precedence:
		// stop / hand back over to parent
		if precedence == 0 || precedence <= parentPrecedence {
			break
		}

		operator := prs.consume(prs.current().Kind)
		right := prs.parseBinaryExpression(precedence)

		// set left to our current expression and continue
		left = nodes.CreateBinaryExpressionNode(operator, left, right)
	}

	return left
}

// parsePrimaryExpression primary expressions being the simple things
func (prs *Parser) parsePrimaryExpression() nodes.ExpressionNode {
	cur := prs.current().Kind
	if cur == lexer.StringToken {
		return prs.parseStringLiteral()

	} else if cur == lexer.NumberToken {
		return prs.parseNumberLiteral()

	} else if cur == lexer.TrueKeyword || cur == lexer.FalseKeyword {
		return prs.parseBoolLiteral()

	} else if cur == lexer.OpenParenthesisToken {
		return prs.parseParenthesisedExpression()

	} else if prs.current().Kind == lexer.IdToken &&
		prs.peek(1).Kind == lexer.AccessToken {
		return prs.parseTypeCallExpression()

	} else if cur == lexer.IdToken {
		return prs.parseNameOrCallExpression()
	}

	// No proper keyword is found
	// Since ExpressionNode is nil the program will crash anyway, at least we exit safely
	// Added this additionalInfo section to help users know why a parser error occurred
	additionalInfo := ""
	if prs.current().Kind == lexer.BadToken {
		additionalInfo = " (may be caused by previous \"UnexpectedCharacterError\" which produces a BadToken)"
	}
	print.Error(
		"PARSER",
		print.UnexpectedTokenError,
		prs.current().Line,
		prs.current().Column,
		8,
		"unexpected Token \"%s\"!"+additionalInfo,
		prs.current().Kind,
	)
	os.Exit(1)

	return nil
}

// parseAssignmentExpression takes an assignmentExpression and returns a node of the same type
// Example: x <- 100;
// this is when x has already been defined but is now being assigned a new value.
func (prs *Parser) parseAssignmentExpression() nodes.AssignmentExpressionNode {
	identifier := prs.consume(lexer.IdToken) // the variable you're assigning the new value (like "x")
	prs.consume(lexer.AssignToken)           // check and skip past the <- (assignToken)
	value := prs.parseExpression()           // new value of variable (like "x"'s new value is 100).

	return nodes.CreateAssignmentExpressionNode(identifier, value)
}

// parseVariableEditorExpression this parses an expression like i <-+ 1, the variable is reassigned using the
// AssignToken, and an operator instead of an expression
func (prs *Parser) parseVariableEditorExpression() nodes.VariableEditorExpressionNode {
	identifier := prs.consume(lexer.IdToken)    // Get the identifier you want to edit
	prs.consume(lexer.AssignToken)              // a b s o r b assignment token (we don't need it)
	operator := prs.consume(prs.current().Kind) // get the operator (this is important as we want to know if it's an
	// PLUS(+) MINUS(-) DIVIDE(/) or STAR(*))

	// Get new expression value
	expression := prs.parseExpression()

	return nodes.CreateVariableEditorExpressionNode(identifier, operator, expression, false)
}

// parseNameOrCallExpression we're either parsing a NameExpressionNode or a CallExpressionNode
func (prs *Parser) parseNameOrCallExpression() nodes.ExpressionNode {

	// check for parenthesis for function call
	if prs.peek(1).Kind == lexer.OpenParenthesisToken {
		return prs.parseCallExpression()
	}

	// if there's an open bracket, assume it's a cast, if it's not we can just rewind
	if prs.peek(1).Kind == lexer.OpenBracketToken {
		return prs.parseComplexCastExpression()
	}

	// if not, it's a name expression
	return prs.parseNameExpression()
}

// parseCallExpression this for when we're calling a function
// For example: Print("hello");
// we need "Print", and the parameters "hello"
func (prs *Parser) parseCallExpression() nodes.CallExpressionNode {

	// We need the identifier to know which function the program called
	identifier := prs.consume(lexer.IdToken)

	prs.consume(lexer.OpenParenthesisToken)  // (
	args := prs.parseArguments()             // We get the arguments being put into the function
	prs.consume(lexer.CloseParenthesisToken) // )

	return nodes.CreateCallExpressionNode(identifier, args, nodes.TypeClauseNode{})
}

// parseComplexCastExpression this for casting what ive called "complex cast" here
// For example: complexType[string, int](someAny)
// we need the type and expression but have to make sure this actually is a cast and not an array access
func (prs *Parser) parseComplexCastExpression() nodes.ExpressionNode {

	// We store the identifier, so we can rewind in case we need to
	identifier := prs.current()

	// try to parse a type
	typeClause, ok := prs.parseUncertainTypeClause()

	// if it didn't go so well, rewind -> this is actually an array access
	if !ok {
		prs.rewind(identifier)
		return prs.parseArrayAccessExpression()
	}

	// if there's no open parenthesis for the cast -> also rewind (this is an array access)
	if prs.current().Kind != lexer.OpenParenthesisToken {
		prs.rewind(identifier)
		return prs.parseArrayAccessExpression()
	}

	prs.consume(lexer.OpenParenthesisToken)  // (
	expression := prs.parseExpression()      // We get the expression we want to cast
	prs.consume(lexer.CloseParenthesisToken) // )

	// return a call expression as the rest is all managed through it
	return nodes.CreateCallExpressionNode(identifier, []nodes.ExpressionNode{expression}, typeClause)
}

// parseArrayAccessExpression this for accessing arrays!
// For example: someArray[1]
// we need to get the identifier and index we want to access
func (prs *Parser) parseArrayAccessExpression() nodes.ExpressionNode {

	// We need the identifier to know what variable to access
	identifier := prs.consume(lexer.IdToken)

	prs.consume(lexer.OpenBracketToken)  // [
	index := prs.parseExpression()       // We get the index expression
	prs.consume(lexer.CloseBracketToken) // ]

	// if there's an assign token, this is an array assignment
	if prs.current().Kind == lexer.AssignToken {
		prs.consume(lexer.AssignToken) // <-
		value := prs.parseExpression() // the value to store in the array
		// return an array assignment expression
		return nodes.CreateArrayAssignmentExpressionNode(identifier, index, value)
	}

	// return an array access expression
	return nodes.CreateArrayAccessExpressionNode(identifier, index)
}

// parseMakeArrayExpression this for creating arrays
// For example: make string array(10)
// we need to get the type and length of the new array
func (prs *Parser) parseMakeArrayExpression() nodes.MakeArrayExpressionNode {

	prs.consume(lexer.MakeKeyword) // make

	// We need the base type of the array
	baseType := prs.parseTypeClause()

	// now we need to consume the "array" keyword
	// only problem with that is that it doesn't exist (it's an identifier)
	// so for now we just consume an IdToken

	prs.consume(lexer.IdToken) // array

	prs.consume(lexer.OpenBracketToken)  // (
	length := prs.parseExpression()      // We get the length of the new array
	prs.consume(lexer.CloseBracketToken) // )

	// return an array access expression
	return nodes.CreateMakeArrayExpressionNode(baseType, length)
}

// parseTypeCallExpression when we want to call a function attached to a data type (or class in the future)
// Example: string->GetLength()
// the data type is string (though in a program it's going to be a variable with the type string)
// GetLength is the exact function call we need
// Anything in the ( ... ) we need to get as arguments
func (prs *Parser) parseTypeCallExpression() nodes.TypeCallExpressionNode {

	identifier := prs.consume(lexer.IdToken) // the variable we're calling the function on

	prs.consume(lexer.AccessToken)               // ->
	callIdentifier := prs.consume(lexer.IdToken) // now we need the name of the call (like "GetLength()")

	prs.consume(lexer.OpenParenthesisToken)  // (
	args := prs.parseArguments()             // any arguments in the function call itself
	prs.consume(lexer.CloseParenthesisToken) // )

	return nodes.CreateTypeCallExpressionNode(identifier, callIdentifier, args)
}

// parseArguments this is when we want to get a series of arguments in a function call function definition.
// Example: arg1 string, arg2 int, arg3 bool,
func (prs *Parser) parseArguments() []nodes.ExpressionNode {
	args := make([]nodes.ExpressionNode, 0) // our slice of arguments

	// We keep looping to get all the yummy arguments
	for prs.current().Kind != lexer.CloseParenthesisToken &&
		prs.current().Kind != lexer.EOF {

		// arguments are just cooler expressions change my mind.
		expression := prs.parseExpression()
		args = append(args, expression)

		if prs.current().Kind == lexer.CommaToken {
			prs.consume(lexer.CommaToken) // then we re-loop and get the next argument expression
		} else {
			break // if not a comma, we know there are no more arguments left
		}
	}

	return args
}

// parseNameExpression just consumes an identifier
func (prs *Parser) parseNameExpression() nodes.NameExpressionNode {
	// Doesn't get any more simple than this
	identifier := prs.consume(lexer.IdToken)
	return nodes.CreateNameExpressionNode(identifier)
}

// parseParenthesisedExpression this is an expression wrapped in parentheses
// Example: ( 1 + 1 + 1 + 1 + 1 ) or ( "Hello" )
// The above are relatively simple, remember this can be any Expression even complex ones.
func (prs *Parser) parseParenthesisedExpression() nodes.ParenthesisedExpressionNode {

	// It's quite literally just consuming the parentheses and passing the expression as a new Node
	prs.consume(lexer.OpenParenthesisToken)
	expression := prs.parseExpression()
	prs.consume(lexer.CloseParenthesisToken)

	return nodes.CreateParenthesisedExpressionNode(expression)
}

// parseStringLiteral
func (prs *Parser) parseStringLiteral() nodes.LiteralExpressionNode {
	// a string token can be found in the lexer in lexer.getString()
	str := prs.consume(lexer.StringToken)
	return nodes.CreateLiteralExpressionNode(str)
}

// parseNumberLiteral
func (prs *Parser) parseNumberLiteral() nodes.LiteralExpressionNode {
	// a number token can be found in the lexer in lexer.getNumber()
	num := prs.consume(lexer.NumberToken)
	return nodes.CreateLiteralExpressionNode(num)
}

// parseBoolLiteral it's *literally* just a bool (true or false)
// Example: true, false
func (prs *Parser) parseBoolLiteral() nodes.LiteralExpressionNode {
	_bool := prs.consume(prs.current().Kind)
	return nodes.CreateLiteralExpressionNode(_bool)
}

// </EXPRESSIONS> -------------------------------------------------------------
