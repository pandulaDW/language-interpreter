package ast

import (
	"github.com/pandulaDW/language-interpreter/tokens"
)

// Interfaces ----------------------------------
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// AST Program ---------------------------------
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Identifier Statement ------------------------
type Identifier struct {
	Token tokens.Token // the tokens.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement --------------------------------
type LetStatement struct {
	Token tokens.Token // The tokens.LET token
	Name  *Identifier
	Value Expression
}

func (l LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l LetStatement) statementNode() {
	panic("implement me")
}

// ReturnStatement ------------------------------
type ReturnStatement struct {
	Token       tokens.Token // The tokens.RETURN token
	ReturnValue Expression
}

func (r ReturnStatement) statementNode() {}

func (r ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
