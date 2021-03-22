package ast

import (
	"github.com/pandulaDW/language-interpreter/tokens"
	"go/token"
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

// Let Statement --------------------------------
type LetStatement struct {
	Token token.Token // The tokens.LET token
	Name  *Identifier
	Value Expression
}

func (l LetStatement) TokenLiteral() string {
	panic("implement me")
}

func (l LetStatement) statementNode() {
	panic("implement me")
}
