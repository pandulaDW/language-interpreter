package ast

import (
	"github.com/pandulaDW/language-interpreter/tokens"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name:  &Identifier{Token: tokens.Token{Type: tokens.IDENT, Literal: "myVar"}},
				Value: &Identifier{Token: tokens.Token{Type: tokens.IDENT, Literal: "anotherVar"}},
			}}}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
