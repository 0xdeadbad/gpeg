package parser

import (
	"fmt"
	"gpeg/lexer"
)

type Rule struct {
	Name *lexer.Token
	Alts [][]*lexer.Token
}

func NewRule(name *lexer.Token) *Rule {
	return &Rule{
		Name: name,
		Alts: make([][]*lexer.Token, 0),
	}
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s: %v", r.Name.Lexeme, r.Alts)
}
