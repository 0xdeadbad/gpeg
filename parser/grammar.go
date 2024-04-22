package parser

import (
	"gpeg/lexer"
	"io"
)

type Fn func()

type GrammarParser struct {
	*Parser
	memo map[string][]*lexer.Token
}

func NewGrammarParser(r io.Reader) *GrammarParser {
	return &GrammarParser{
		Parser: &Parser{
			l:      lexer.NewLexer(r),
			tokens: make([]*lexer.Token, 0),
			pos:    0,
		},
	}
}

func (p *GrammarParser) Grammar() []*Rule {
	pos := p.Mark()
	if rule := p.Rule(); rule != nil {
		rules := make([]*Rule, 0)
		rules = append(rules, rule)

		for rule = p.Rule(); rule != nil; rule = p.Rule() {
			rules = append(rules, rule)
		}

		if p.ExpectType(lexer.EOF) != nil {
			return rules
		}
	}
	p.Reset(pos)
	return nil
}

func (p *GrammarParser) Rule() *Rule {
	p.GetTokens()
	pos := p.Mark()
	if name := p.ExpectType(lexer.WORD); name != nil {
		if p.ExpectLexeme(":") != nil {
			if alt := p.Alternative(); alt != nil {
				alts := make([][]*lexer.Token, 0)
				alts = append(alts, alt)
				apos := p.Mark()

				expected := p.ExpectLexeme("|")
				alt = p.Alternative()
				for expected != nil && alt != nil {
					alts = append(alts, alt)
					apos = p.Mark()
					expected = p.ExpectLexeme("|")
					alt = p.Alternative()
				}
				p.Reset(apos)
				if p.ExpectLexeme(";") != nil {
					return &Rule{
						Name: name,
						Alts: alts,
					}
				}
			}
		}
	}
	p.Reset(pos)
	return nil
}

func (p *GrammarParser) Alternative() []*lexer.Token {
	items := make([]*lexer.Token, 0)
	for item := p.Item(); item != nil; item = p.Item() {
		items = append(items, item)
	}
	return items
}

func (p *GrammarParser) Item() *lexer.Token {
	if item := p.ExpectType(lexer.WORD); item != nil {
		return item
	}
	if item := p.ExpectType(lexer.STR); item != nil {
		return item
	}
	return nil
}
