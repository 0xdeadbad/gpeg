package parser

import (
	"fmt"
	"gpeg/lexer"
	"io"
)

type Parser struct {
	l      *lexer.Lexer
	tokens []*lexer.Token
	pos    int
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		l:      lexer.NewLexer(r),
		tokens: make([]*lexer.Token, 0),
		pos:    0,
	}
}

func (p *Parser) GetTokens() []*lexer.Token {
	for t := p.l.NextToken(); t != nil; t = p.l.NextToken() {
		p.tokens = append(p.tokens, t)
	}

	return p.tokens
}

func (p *Parser) Mark() int {
	return p.pos
}

func (p *Parser) Reset(pos int) {
	p.pos = pos
}

func (p *Parser) Peek() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}

	return p.tokens[p.pos]
}

func (p *Parser) Next() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}

	defer func() { p.pos++ }()
	return p.tokens[p.pos]
}

func (p *Parser) ExpectLexeme(lexemes ...string) *lexer.Token {
	for _, lexeme := range lexemes {
		if pt := p.Peek(); pt != nil && pt.Lexeme == lexeme {
			defer func() { p.pos++ }()
			return p.tokens[p.pos]
		}
	}
	return nil
}

func (p *Parser) ExpectType(typ ...lexer.TokenType) *lexer.Token {
	for _, t := range typ {
		if pt := p.Peek(); pt != nil && pt.Type == t {
			defer func() { p.pos++ }()
			return p.tokens[p.pos]
		}
	}
	return nil
}

func OutputToyParser(rules []*Rule) {
	fmt.Printf("%s\n", "class ToyParser(Parser):")
}
