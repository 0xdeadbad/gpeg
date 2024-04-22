package lexer

type Token struct {
	Lexeme string
	Type   TokenType
	Line   int
	Column int
}

func NewToken(lexeme string, t TokenType, line int, column int) *Token {
	return &Token{Lexeme: lexeme, Type: t, Line: line, Column: column}
}

func (t *Token) String() string {
	return t.Lexeme
}

func (t *Token) IsType(tokenType TokenType) bool {
	return t.Type == tokenType
}

type TokenType int

const (
	EOF TokenType = iota
	STR
	PLUS
	MINUS
	STAR
	UNDERSCORE
	SEMICOLON
	COLON
	RIGHT_PAREN
	LEFT_PAREN
	RIGHT_BRACE
	LEFT_BRACE
	RIGHT_BRACKET
	LEFT_BRACKET
	BANG
	QUESTION
	AMPERSAND
	PIPE
	SLASH
	SEQUENCE
	EQUALS
	WORD
	NUMBER
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case STR:
		return "STR"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case UNDERSCORE:
		return "UNDERSCORE"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACKET:
		return "RIGHT_BRACKET"
	case LEFT_BRACKET:
		return "LEFT_BRACKET"
	case BANG:
		return "BANG"
	case QUESTION:
		return "QUESTION"
	case AMPERSAND:
		return "AMPERSAND"
	case PIPE:
		return "PIPE"
	case SLASH:
		return "SLASH"
	case SEQUENCE:
		return "SEQUENCE"
	case EQUALS:
		return "EQUALS"
	case WORD:
		return "WORD"
	case NUMBER:
		return "NUMBER"
	}
	return "UNKNOWN"
}
