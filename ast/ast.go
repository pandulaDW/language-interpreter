package ast

import (
	"bytes"
	"github.com/pandulaDW/language-interpreter/tokens"
)

// Interfaces ----------------------------------
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
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

func (l LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()
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

func (r ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

//ExpressionStatement -----------------------------
type ExpressionStatement struct {
	Token      tokens.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
