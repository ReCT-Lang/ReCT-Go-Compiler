package parser

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/rules"
	"fmt"
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

// consume a syntax token is its what we're looking for if not -> just die for now
func (prs *Parser) consume(expected lexer.TokenKind) lexer.Token {
	if prs.current().Kind != expected {
		// cant tell the user which one because theres no way to get a constants name by its value...
		// so we might need a separate name array for them

		// Switched the TokenKind constants to strings and now added the error message you wanted <3
		fmt.Printf(
			"ERROR(%d, %d): Unexpected Token \"%s\"! Expected \"%s\".\n",
			prs.current().Line,
			prs.current().Column,
			prs.current().Kind,
			expected,
		)
		os.Exit(-1) // die(-1);
	}

	// if everything is alright -> step our index and return the token
	prs.Index++
	return prs.peek(-1)
}

// </HELPERS> -----------------------------------------------------------------

// parse a compilation (with all its functions, classes, enums, global statements, etc...)
func Parse(tokens []lexer.Token) []nodes.MemberNode {
	parser := Parser{
		Tokens: tokens,
		Index:  0,
	}

	return parser.parseMembers()
}

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

func (prs *Parser) parseMember() nodes.MemberNode {
	// functions / classes would go here \/
	if prs.current().Kind == lexer.FunctionKeyword {
		return prs.parseFunctionDeclaration()
	}

	// global statements (any statements outside of any functions)
	return prs.parseGlobalStatement()
}

func (prs *Parser) parseGlobalStatement() nodes.GlobalStatementMember {
	statement := prs.parseStatement()
	return nodes.CreateGlobalStatementMember(statement)
}

func (prs *Parser) parseFunctionDeclaration() nodes.FunctionDeclarationMember {
	prs.consume(lexer.FunctionKeyword)
	identifier := prs.consume(lexer.IdToken)

	prs.consume(lexer.OpenParenthesisToken)
	params := prs.parseParameterList()
	prs.consume(lexer.CloseParenthesisToken)

	typeClause := prs.parseOptionalTypeClause()

	body := prs.parseBlockStatement()

	return nodes.CreateFunctionDeclarationMember(identifier, params, typeClause, body)
}

func (prs *Parser) parseParameterList() []nodes.ParameterNode {
	params := make([]nodes.ParameterNode, 0)

	for prs.current().Kind != lexer.CloseParenthesisToken &&
		prs.current().Kind != lexer.EOF {

		param := prs.parseParameter()
		params = append(params, param)

		if prs.current().Kind == lexer.CommaToken {
			prs.consume(lexer.CommaToken)
		} else {
			break
		}
	}

	return params
}

func (prs *Parser) parseParameter() nodes.ParameterNode {
	identifier := prs.consume(lexer.IdToken)
	typeClause := prs.parseTypeClause()
	return nodes.CreateParameterNode(identifier, typeClause)
}

func (prs *Parser) parseOptionalTypeClause() nodes.TypeClauseNode {
	// if theres no type clause, return an empty one
	if prs.current().Kind != lexer.IdToken && prs.current().Kind != lexer.AccessToken {
		return nodes.TypeClauseNode{}
	}

	return prs.parseTypeClause()
}

func (prs *Parser) parseTypeClause() nodes.TypeClauseNode {
	// if theres a '->' token, consume it
	if prs.current().Kind == lexer.AccessToken {
		prs.consume(lexer.AccessToken)
	}

	identifier := prs.consume(lexer.IdToken)
	return nodes.CreateTypeClauseNode(identifier)
}

// <STATEMENTS> ---------------------------------------------------------------

func (prs *Parser) parseStatement() nodes.StatementNode {
	var statement nodes.StatementNode = nil
	// nil StatementNode can cause segmentation violation if no correct key is found, I've added an error handler to counter this.

	// select correct parsing function based on kind
	cur := prs.current().Kind
	if cur == lexer.OpenBraceToken {
		statement = prs.parseBlockStatement()

	} else if cur == lexer.VarKeyword || cur == lexer.SetKeyword {
		statement = prs.parseVariableDeclaration()

	} else if cur == lexer.IfKeyword {
		statement = prs.parseIfStatement()

	} else if cur == lexer.ReturnKeyword {
		statement = prs.parseReturnStatement()

	} else if cur == lexer.ForKeyword {
		statement = prs.parseForStatement()

	} else if cur == lexer.WhileKeyword {
		statement = prs.parseWhileStatement()

	} else if cur == lexer.BreakKeyword {
		statement = prs.parseBreakStatement()

	} else if cur == lexer.ContinueKeyword {
		statement = prs.parseContinueStatement()

	} else if cur == lexer.FromKeyword {
		statement = prs.parseFromToStatement()

	} else {
		statement = prs.parseExpressionStatement()

		// moved the error message to parsePrimaryExpression()
	}

	// if theres a semicolon -> a b s o r b    i t
	if prs.current().Kind == lexer.Semicolon {
		prs.consume(lexer.Semicolon)
	}

	return statement
}

func (prs *Parser) parseBlockStatement() nodes.BlockStatementNode {
	// create a list for our statement
	statements := make([]nodes.StatementNode, 0)

	// {
	openBrace := prs.consume(lexer.OpenBraceToken)

	for prs.current().Kind != lexer.EOF &&
		prs.current().Kind != lexer.CloseBraceToken {

		startToken := prs.current()

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

func (prs *Parser) parseVariableDeclaration() nodes.VariableDeclarationStatementNode {

	// We already check for var/set in parseStatement(), this code just repeats the process.
	// Replaced expecting with prs.current().kind when defining keyword - Tokorv
	// smort  - RedCube

	keyword := prs.consume(prs.current().Kind) // Replaced expecting

	typeClause := nodes.TypeClauseNode{}
	if prs.current().Kind == lexer.IdToken && prs.peek(1).Kind == lexer.IdToken {
		typeClause = prs.parseTypeClause()
	}

	identifier := prs.consume(lexer.IdToken)
	assign := prs.consume(lexer.AssignToken)
	initializer := prs.parseExpression()

	return nodes.CreateVariableDeclarationStatementNode(keyword, typeClause, identifier, assign, initializer)
}

func (prs *Parser) parseIfStatement() nodes.IfStatementNode {
	keyword := prs.consume(lexer.IfKeyword)

	prs.consume(lexer.OpenParenthesisToken)
	condition := prs.parseExpression()
	prs.consume(lexer.CloseParenthesisToken)

	statement := prs.parseStatement()
	elseClause := prs.parseElseClause()

	return nodes.CreateIfStatementNode(keyword, condition, statement, elseClause)
}

func (prs *Parser) parseElseClause() nodes.ElseClauseNode {
	// if theres no else -> dont parse an else lol
	if prs.current().Kind != lexer.ElseKeyword {
		return nodes.ElseClauseNode{}
	}

	keyword := prs.consume(lexer.ElseKeyword)
	statement := prs.parseStatement()

	return nodes.CreateElseClauseNode(keyword, statement)
}

func (prs *Parser) parseReturnStatement() nodes.ReturnStatementNode {
	keyword := prs.consume(lexer.ReturnKeyword)

	var expression nodes.ExpressionNode = nil
	// if we are at the end of the line (;) theres no return value given
	if prs.current().Kind != lexer.Semicolon {
		expression = prs.parseExpression()
	}

	return nodes.CreateReturnStatementNode(keyword, expression)
}

func (prs *Parser) parseForStatement() nodes.ForStatementNode {
	keyword := prs.consume(lexer.ForKeyword)

	// For ( S; E; E) S
	prs.consume(lexer.OpenParenthesisToken)
	initialiser := prs.parseVariableDeclaration()
	prs.consume(lexer.Semicolon)

	condition := prs.parseExpression()
	prs.consume(lexer.Semicolon)

	updation := prs.parseStatement()
	prs.consume(lexer.CloseParenthesisToken)

	statement := prs.parseStatement()

	return nodes.CreateForStatementNode(keyword, initialiser, condition, updation, statement)
}

func (prs *Parser) parseWhileStatement() nodes.WhileStatementNode {
	keyword := prs.consume(lexer.WhileKeyword)

	prs.consume(lexer.OpenParenthesisToken)
	condition := prs.parseExpression()
	prs.consume(lexer.CloseParenthesisToken)

	statement := prs.parseStatement()

	return nodes.CreateWhileStatementNode(keyword, condition, statement)
}

func (prs *Parser) parseFromToStatement() nodes.FromToStatementNode {
	keyword := prs.consume(lexer.FromKeyword)

	prs.consume(lexer.OpenParenthesisToken)
	identifier := prs.consume(lexer.IdToken)
	prs.consume(lexer.AssignToken)
	lowerBound := prs.parseExpression()
	prs.consume(lexer.CloseParenthesisToken)

	prs.consume(lexer.ToKeyword)
	upperBound := prs.parseExpression()

	statement := prs.parseStatement()
	return nodes.CreateFromToStatementNode(keyword, identifier, lowerBound, upperBound, statement)
}

func (prs *Parser) parseBreakStatement() nodes.BreakStatementNode {
	keyword := prs.consume(lexer.BreakKeyword)

	return nodes.CreateBreakStatement(keyword)
}

func (prs *Parser) parseContinueStatement() nodes.ContinueStatementNode {
	keyword := prs.consume(lexer.ContinueKeyword)

	return nodes.CreateContinueStatement(keyword)
}

func (prs *Parser) parseExpressionStatement() nodes.ExpressionStatementNode {
	expression := prs.parseExpression()
	return nodes.CreateExpressionStatementNode(expression)
}

// </STATEMENTS> --------------------------------------------------------------
// <EXPRESSIONS> --------------------------------------------------------------

func (prs *Parser) parseExpression() nodes.ExpressionNode {
	// check for more "complex" expressions first
	// (these are not allowed in binary expressions)
	// if theres none -> parse normal binary expression

	// assignment expression
	if prs.current().Kind == lexer.IdToken &&
		prs.peek(1).Kind == lexer.AssignToken {
		return prs.parseAssignmentExpression()
	}

	return prs.parseBinaryExpression(0)
}

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

		// if this isnt an operator or it has less precedence:
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

// primary expressions being the simple things
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
	} else if cur == lexer.IdToken {
		return prs.parseNameOrCallExpression()
	}

	// No proper keyword is found
	// Since ExpressionNode is nil the program will crash anyway, at least we exit safely
	fmt.Printf("ERROR: Unexpected Token \"%s\" found! ExpressionNode nil -> forced exit.", prs.current().Kind)
	os.Exit(1)

	return nil
}

func (prs *Parser) parseAssignmentExpression() nodes.AssignmentExpressionNode {
	identifier := prs.consume(lexer.IdToken)
	prs.consume(lexer.AssignToken)
	value := prs.parseExpression()

	return nodes.CreateAssignmentExpressionNode(identifier, value)
}

func (prs *Parser) parseNameOrCallExpression() nodes.ExpressionNode {

	// check for parenthesis for function call
	if prs.peek(1).Kind == lexer.OpenParenthesisToken {
		return prs.parseCallExpression()
	}

	// if not its a name expression
	return prs.parseNameExpression()
}

func (prs *Parser) parseCallExpression() nodes.CallExpressionNode {
	identifier := prs.consume(lexer.IdToken)
	prs.consume(lexer.OpenParenthesisToken)
	args := prs.parseArguments()
	prs.consume(lexer.CloseParenthesisToken)

	return nodes.CreateCallExpressionNode(identifier, args)
}

func (prs *Parser) parseArguments() []nodes.ExpressionNode {
	args := make([]nodes.ExpressionNode, 0)

	for prs.current().Kind != lexer.CloseParenthesisToken &&
		prs.current().Kind != lexer.EOF {

		expression := prs.parseExpression()
		args = append(args, expression)

		if prs.current().Kind == lexer.CommaToken {
			prs.consume(lexer.CommaToken)
		} else {
			break
		}
	}

	return args
}

func (prs *Parser) parseNameExpression() nodes.NameExpressionNode {
	identifier := prs.consume(lexer.IdToken)
	return nodes.CreateNameExpressionNode(identifier)
}

func (prs *Parser) parseParenthesisedExpression() nodes.ParenthesisedExpressionNode {

	prs.consume(lexer.OpenParenthesisToken)
	expression := prs.parseExpression()
	prs.consume(lexer.CloseParenthesisToken)

	return nodes.CreateParenthesisedExpressionNode(expression)
}

func (prs *Parser) parseStringLiteral() nodes.LiteralExpressionNode {
	str := prs.consume(lexer.StringToken)
	return nodes.CreateLiteralExpressionNode(str)
}

func (prs *Parser) parseNumberLiteral() nodes.LiteralExpressionNode {
	num := prs.consume(lexer.NumberToken)
	return nodes.CreateLiteralExpressionNode(num)
}

func (prs *Parser) parseBoolLiteral() nodes.LiteralExpressionNode {
	_bool := prs.consume(prs.current().Kind)
	return nodes.CreateLiteralExpressionNode(_bool)
}

// </EXPRESSIONS> -------------------------------------------------------------
