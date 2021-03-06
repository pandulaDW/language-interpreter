package parser

import (
	"fmt"
	"github.com/pandulaDW/language-interpreter/ast"
	"github.com/pandulaDW/language-interpreter/lexer"
	"github.com/pandulaDW/language-interpreter/tokens"
)

// Parser Definition
type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	errors    []string
}

//New is the constructor of the parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Helpers -----------------------------------------------------
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peerError(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peerError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// ParseProgram creates a new AST root node and does the main parsing
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = make([]ast.Statement, 0)

	for p.curToken.Type != tokens.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Parser function -------------------------------
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokens.LET:
		return p.parseLetStatement()
	case tokens.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// Let statement -------------------------------
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Return statement -------------------------------
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
