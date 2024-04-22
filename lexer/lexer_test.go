package lexer

import (
	"strings"
	"testing"
)

func TestLexer_Emit(t *testing.T) {
	l := NewLexer(strings.NewReader("test"))

	l.Emit(NewToken("test", EOF, 1, 1))

	tok := l.NextToken()
	if tok.Lexeme != "test" {
		t.Errorf("Expected 'test', got %s", tok.Lexeme)
	}
	if tok.Type != EOF {
		t.Errorf("Expected EOF, got %d", tok.Type)
	}
	if tok.Line != 1 {
		t.Errorf("Expected 1, got %d", tok.Line)
	}
	if tok.Column != 1 {
		t.Errorf("Expected 1, got %d", tok.Column)
	}
}

func TestLexer_States(t *testing.T) {
	t.Run("lexSpaces", func(t *testing.T) {
		l := NewLexer(strings.NewReader(" \n\t\r"))

		l.s = lexSpaces
		l.s = l.s(l)

		if l.Line != 2 {
			t.Errorf("Expected 2, got %d", l.Line)
		}
		if l.Column != 3 {
			t.Errorf("Expected 3, got %d", l.Column)
		}
	})

	t.Run("lexString", func(t *testing.T) {
		l := NewLexer(strings.NewReader("'hello world!'"))

		l.s = lexString(l)

		token := l.NextToken()
		if token.Lexeme != "hello world!" {
			t.Errorf("Expected 'hello world!', got %s", token.Lexeme)
		}
		if token.Type != STR {
			t.Errorf("Expected STR, got %d", token.Type)
		}
	})

	t.Run("lexSymbol", func(t *testing.T) {
		l := NewLexer(strings.NewReader("+-*_;:[](){}?&|/="))

		l.s = lexSymbol(l)
		token := l.NextToken()
		if token.Lexeme != "+" {
			t.Errorf("Expected '+', got %s", token.Lexeme)
		}
		if token.Type != PLUS {
			t.Errorf("Expected PLUS, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "-" {
			t.Errorf("Expected '-', got %s", token.Lexeme)
		}
		if token.Type != MINUS {
			t.Errorf("Expected MINUS, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "*" {
			t.Errorf("Expected '*', got %s", token.Lexeme)
		}
		if token.Type != STAR {
			t.Errorf("Expected MINUS, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "_" {
			t.Errorf("Expected '_', got %s", token.Lexeme)
		}
		if token.Type != UNDERSCORE {
			t.Errorf("Expected MINUS, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != ";" {
			t.Errorf("Expected ';', got %s", token.Lexeme)
		}
		if token.Type != SEMICOLON {
			t.Errorf("Expected SEMICOLON, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != ":" {
			t.Errorf("Expected ':', got %s", token.Lexeme)
		}
		if token.Type != COLON {
			t.Errorf("Expected SEMICOLON, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "[" {
			t.Errorf("Expected '[', got %s", token.Lexeme)
		}
		if token.Type != LEFT_BRACKET {
			t.Errorf("Expected LEFT_BRACKET, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "]" {
			t.Errorf("Expected ']', got %s", token.Lexeme)
		}
		if token.Type != RIGHT_BRACKET {
			t.Errorf("Expected RIGHT_BRACKET, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "(" {
			t.Errorf("Expected '(', got %s", token.Lexeme)
		}
		if token.Type != LEFT_PAREN {
			t.Errorf("Expected LEFT_PAREN, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != ")" {
			t.Errorf("Expected ')', got %s", token.Lexeme)
		}
		if token.Type != RIGHT_PAREN {
			t.Errorf("Expected RIGHT_PAREN, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "{" {
			t.Errorf("Expected '{', got %s", token.Lexeme)
		}
		if token.Type != LEFT_BRACE {
			t.Errorf("Expected LEFT_BRACE, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "}" {
			t.Errorf("Expected '}', got %s", token.Lexeme)
		}
		if token.Type != RIGHT_BRACE {
			t.Errorf("Expected RIGHT_BRACE, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "?" {
			t.Errorf("Expected '?', got %s", token.Lexeme)
		}
		if token.Type != QUESTION {
			t.Errorf("Expected QUESTION, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "&" {
			t.Errorf("Expected '&', got %s", token.Lexeme)
		}
		if token.Type != AMPERSAND {
			t.Errorf("Expected AMPERSAND, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "|" {
			t.Errorf("Expected '|', got %s", token.Lexeme)
		}
		if token.Type != PIPE {
			t.Errorf("Expected PIPE, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "/" {
			t.Errorf("Expected '/', got %s", token.Lexeme)
		}
		if token.Type != SLASH {
			t.Errorf("Expected SLASH, got %d", token.Type)
		}

		l.s = lexSymbol(l)
		token = l.NextToken()
		if token.Lexeme != "=" {
			t.Errorf("Expected '=', got %s", token.Lexeme)
		}
		if token.Type != EQUALS {
			t.Errorf("Expected EQUALS, got %d", token.Type)
		}
	})

	t.Run("lexSequence", func(t *testing.T) {
		l := NewLexer(strings.NewReader("..."))

		l.s = lexSequence(l)

		token := l.NextToken()
		if token.Lexeme != "..." {
			t.Errorf("Expected '...', got %s", token.Lexeme)
		}
		if token.Type != SEQUENCE {
			t.Errorf("Expected SEQUENCE, got %d", token.Type)
		}
	})

	t.Run("lexWord", func(t *testing.T) {
		l := NewLexer(strings.NewReader("hello"))

		l.s = lexWord(l)

		token := l.NextToken()
		if token.Lexeme != "hello" {
			t.Errorf("Expected 'hello', got %s", token.Lexeme)
		}
		if token.Type != WORD {
			t.Errorf("Expected WORD, got %d", token.Type)
		}
	})

	t.Run("lexNumber", func(t *testing.T) {
		l := NewLexer(strings.NewReader("123"))

		l.s = lexNumber(l)

		token := l.NextToken()
		if token.Lexeme != "123" {
			t.Errorf("Expected '123', got %s", token.Lexeme)
		}
		if token.Type != NUMBER {
			t.Errorf("Expected NUMBER, got %d", token.Type)
		}
	})
}

func TestLexer_Full(t *testing.T) {
	input := `Input = Space Sum !_ ;
Sum = Number (Plus Number)* ;
Number = Digits Space ;
Plus = "+" Space ;
Digits = [0-9]+ ;
Space = " "* ;`

	l := NewLexer(strings.NewReader(input))

	// Check correctness visually
	for token := l.NextToken(); token != nil && token.Type != EOF; token = l.NextToken() {
		t.Logf("Token: %s, Type: %s\n", token.Lexeme, token.Type)
	}
}

func TestLexer_Full_Gvanrossum(t *testing.T) {
	input := `statement: assignment | expr | if_statement ;
expr: expr '+' term | expr '-' term | term ;
term: term '*' atom | term '/' atom | atom ;
atom: NAME | NUMBER | '(' expr ')' ;
assignment: target '=' expr ;
target: NAME ;
if_statement: 'if' expr ':' statement ;`

	l := NewLexer(strings.NewReader(input))

	// Check correctness visually
	for token := l.NextToken(); token != nil && token.Type != EOF; token = l.NextToken() {
		t.Logf("Token: %s, Type: %s\n", token.Lexeme, token.Type)
	}
}
