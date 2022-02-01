package parser

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
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
		fmt.Println("[Parser] ERROR: unexpected Token!")
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
	// if prs.current().Kind == lexer.FunctionKeyword { ... }

	// global statements (any statements outside of any functions)
	return prs.parseGlobalStatement()
}

func (prs *Parser) parseGlobalStatement() nodes.MemberNode {
	statement := prs.parseStatement()
	// return nodes.CreateGlobalStatementMember(statement) // this doesnt work, not entirely sure what i did wrong but ill figure it out later
	return &nodes.GlobalStatementMember{
		Statement: statement,
	}
}

// <STATEMENTS> ---------------------------------------------------------------

func (prs *Parser) parseStatement() nodes.StatementNode {
	// select correct parsing function based on kind
	switch prs.current().Kind {
	case lexer.OpenBraceToken:
		return prs.parseBlockStatement()
	case lexer.VarKeyword:
	case lexer.SetKeyword:
		return prs.parseVariableDeclaration()
	case lexer.IfKeyword:
		return prs.parseIfStatement()
	}

	return nil //prs.parseExpressionStatement()
}

func (prs *Parser) parseBlockStatement() *nodes.BlockStatementNode {
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

	return &nodes.BlockStatementNode{
		OpenBrace:  openBrace,
		Statments:  statements,
		CloseBrace: closeBrace,
	}
}

func (prs *Parser) parseVariableDeclaration() *nodes.VariableDeclarationStatementNode {
	// figure out if we need to consume "set" or "var"
	expecting := lexer.VarKeyword
	if prs.current().Kind == lexer.SetKeyword {
		expecting = lexer.SetKeyword
	}

	keyword := prs.consume(expecting)
	identifier := prs.consume(lexer.IdToken)

	// no explicit type clause, im tired lol

	assign := prs.consume(lexer.AssignToken)
	initializer := prs.parseExpression()

	return &nodes.VariableDeclarationStatementNode{
		Keyword:     keyword,
		Identifier:  identifier,
		AssignToken: assign,
		Initializer: initializer,
	}
}

func (prs *Parser) parseIfStatement() *nodes.IfStatementNode {
	keyword := prs.consume(lexer.IfKeyword)
	condition := prs.parseExpression()
	statement := prs.parseStatement()
	elseClause := prs.parseElseClause()

	return &nodes.IfStatementNode{
		IfKeyword:     keyword,
		Condition:     condition,
		ThenStatement: statement,
		ElseClause:    elseClause,
	}
}

func (prs *Parser) parseElseClause() nodes.ElseClauseNode {
	// if theres no else -> dont parse an else lol
	if prs.current().Kind != lexer.ElseKeyword {
		return nodes.ElseClauseNode{}
	}

	keyword := prs.consume(lexer.ElseKeyword)
	statement := prs.parseStatement()

	return nodes.ElseClauseNode{
		ElseKeyword:   keyword,
		ElseStatement: statement,
	}
}

// </STATEMENTS> --------------------------------------------------------------
// <EXPRESSIONS> --------------------------------------------------------------

func (prs *Parser) parseExpression() nodes.ExpressionNode {
	// we only have literals atm

	switch prs.current().Kind {
	case lexer.StringToken:
		return prs.parseStringLiteral()
	case lexer.NumberToken:
		return prs.parseNumberLiteral()
	}

	return nil
}

func (prs *Parser) parseStringLiteral() *nodes.LiteralExpressionNode {
	str := prs.consume(lexer.StringToken)
	return &nodes.LiteralExpressionNode{
		LiteralToken: str,
		LiteralValue: str.Value,
	}
}

func (prs *Parser) parseNumberLiteral() *nodes.LiteralExpressionNode {
	num := prs.consume(lexer.NumberToken)
	return &nodes.LiteralExpressionNode{
		LiteralToken: num,
		LiteralValue: num.Value,
	}
}

// </EXPRESSIONS> -------------------------------------------------------------
