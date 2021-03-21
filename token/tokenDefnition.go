package token

// Type is defined as a string
type Type string

// Token definition
type Token struct {
	Type    Type
	Literal string
}
