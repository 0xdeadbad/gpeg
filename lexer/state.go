package lexer

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type State func(*Lexer) State

func lex(l *Lexer) (s State) {
	s = nil

	// Always take into account that the rune is peeked for the next state
	r, err := l.peekRune()
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}

	switch r {
	case '\t', '\r', ' ', '\n':
		s = lexSpaces
	case '\'', '"':
		s = lexString
	case '+', '-', '*', '_', ';', ':', '(', ')', '{', '}', '[', ']', '!', '?', '&', '|', '/', '=':
		s = lexSymbol
	case '.':
		s = lexSequence
	default:
		if unicode.IsLetter(r) {
			s = lexWord
		}
		if unicode.IsDigit(r) {
			s = lexNumber
		}
	}

	return
}

func lexSpaces(l *Lexer) (s State) {
	s = lex

	readRuneWhile(l, func(r rune) bool {
		return r == '\t' || r == '\r' || r == ' ' || r == '\n'
	})

	return
}

func lexString(l *Lexer) (s State) {
	s = lex

	var e rune

	sb := new(strings.Builder)

	e = justReadRune(l)
	readRuneWhile(l, func(r rune) (b bool) {
		if r != e {
			b = true
			_, err := sb.WriteRune(r)
			if err != nil {
				panic(err)
			}
		}
		return
	})
	_ = expectRune(l, e)

	l.Emit(NewToken(sb.String(), STR, l.Line, l.Column))

	return
}

func lexSymbol(l *Lexer) (s State) {
	s = lex

	var tokenType TokenType

	r := justReadRune(l)
	switch r {
	case '+':
		tokenType = PLUS
	case '-':
		tokenType = MINUS
	case '*':
		tokenType = STAR
	case '_':
		tokenType = UNDERSCORE
	case ';':
		tokenType = SEMICOLON
	case ':':
		tokenType = COLON
	case '(':
		tokenType = LEFT_PAREN
	case ')':
		tokenType = RIGHT_PAREN
	case '{':
		tokenType = LEFT_BRACE
	case '}':
		tokenType = RIGHT_BRACE
	case '[':
		tokenType = LEFT_BRACKET
	case ']':
		tokenType = RIGHT_BRACKET
	case '!':
		tokenType = BANG
	case '?':
		tokenType = QUESTION
	case '&':
		tokenType = AMPERSAND
	case '|':
		tokenType = PIPE
	case '/':
		tokenType = SLASH
	case '=':
		tokenType = EQUALS
	}

	l.Emit(NewToken(string(r), tokenType, l.Line, l.Column))

	return
}

func lexSequence(l *Lexer) (s State) {
	s = lex

	_ = expectRune(l, '.')
	_ = expectRune(l, '.')
	_ = expectRune(l, '.')

	l.Emit(NewToken("...", SEQUENCE, l.Line, l.Column))

	return

}

func lexWord(l *Lexer) (s State) {
	s = lex

	sb := new(strings.Builder)

	readRuneWhile(l, func(r rune) (b bool) {
		if unicode.IsLetter(r) || r == '_' {
			b = true
			_, err := sb.WriteRune(r)
			if err != nil {
				panic(err)
			}
		}
		return
	})

	l.Emit(NewToken(sb.String(), WORD, l.Line, l.Column))

	return

}

func lexNumber(l *Lexer) (s State) {
	s = lex

	sb := new(strings.Builder)

	readRuneWhile(l, func(r rune) (b bool) {
		if unicode.IsDigit(r) {
			b = true
			_, err := sb.WriteRune(r)
			if err != nil {
				panic(err)
			}
		}
		return
	})

	l.Emit(NewToken(sb.String(), NUMBER, l.Line, l.Column))

	return
}

func panicNotEOF(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

func justReadRune(l *Lexer) (r rune) {
	var err error
	r, err = l.readRune()
	panicNotEOF(err)
	return
}

func justPeekRune(l *Lexer) (r rune) {
	var err error
	r, err = l.peekRune()
	panicNotEOF(err)
	return
}

func readRuneWhile(l *Lexer, f func(r rune) bool) {
	for r := justPeekRune(l); f(r); r = justPeekRune(l) {
		r = justReadRune(l)
	}
}

func expectRune(l *Lexer, r rune) (ret rune) {
	if ret = justReadRune(l); ret != r {
		panic(fmt.Sprintf("Expected rune %c, got %c", r, ret))
	}
	return
}
