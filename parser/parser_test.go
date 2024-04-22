package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParser_GetTokens(t *testing.T) {
	t.Run("GetTokens", func(t *testing.T) {
		p := NewParser(strings.NewReader("+-*_;:[](){}?&|/="))

		tokens := p.GetTokens()

		if len(tokens) != 17 {
			t.Errorf("Expected 17 tokens, got %d", len(tokens))
		}

		t.Logf("Tokens: %v", tokens)
	})
}

func TestGrammarParser_Full(t *testing.T) {
	t.Run("Full", func(t *testing.T) {
		input := `statement: assignment | expr | if_statement ;
expr: expr '+' term | expr '-' term | term ;
term: term '*' atom | term '/' atom | atom ;
atom: NAME | NUMBER | '(' expr ')' ;
assignment: target '=' expr ;
target: NAME ;
if_statement: 'if' expr ':' statement ;`
		p := NewGrammarParser(strings.NewReader(input))

		rules := p.Grammar()

		t.Logf("Rules: %v", rules)
	})
}

func TestGrammarParser_Generate_ToyParser(t *testing.T) {
	t.Run("Generate_ToyParser", func(t *testing.T) {
		input := `statement: assignment | expr | if_statement ;
expr: expr '+' term | expr '-' term | term ;
term: term '*' atom | term '/' atom | atom ;
atom: NAME | NUMBER | '(' expr ')' ;
assignment: target '=' expr ;
target: NAME ;
if_statement: 'if' expr ':' statement ;`

		p := NewGrammarParser(strings.NewReader(input))

		rules := p.Grammar()

		for _, rule := range rules {
			fmt.Println()
			fmt.Printf("func (p *ToyParser) %s() {\n", rule.Name)
			fmt.Printf("\tpos := p.Mark()\n")
			for _, alt := range rule.Alts {
				// items := make([]*lexer.Token, 0)
				fmt.Printf("\tif ")
				for _, item := range alt {
					fmt.Printf("%s ", item.Lexeme)
				}
			}
		}

	})
}
