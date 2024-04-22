package lexer

import (
	"bufio"
	"io"
)

type Lexer struct {
	r *bufio.Reader
	c chan *Token
	s State

	Line   int
	Column int
}

func NewLexer(reader io.Reader) *Lexer {
	r := bufio.NewReader(reader)
	c := make(chan *Token, 2)

	return &Lexer{
		r:      r,
		c:      c,
		s:      lex,
		Line:   1,
		Column: 1,
	}
}

func (l *Lexer) NextToken() *Token {
	for {
		select {
		case t, _ := <-l.c:
			return t
		default:
			l.s = l.s(l)
			if l.s == nil {
				l.Emit(NewToken("EOF", EOF, l.Line, l.Column))
				close(l.c)
			}
		}
	}
}

func (l *Lexer) Emit(t *Token) {
	l.c <- t
}

func (l *Lexer) readRune() (r rune, err error) {
	r, _, err = l.r.ReadRune()
	// Update line and column after reading a rune
	defer func() {
		if err != nil && err != io.EOF {
			return
		}
		if r == '\n' {
			l.Line++
			l.Column = 1
		} else {
			l.Column++
		}
	}()
	return
}

func (l *Lexer) peekRune() (r rune, err error) {
	r, _, err = l.r.ReadRune()
	if err != nil {
		return
	}
	err = l.r.UnreadRune()
	return
}
