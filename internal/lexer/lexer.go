package lexer

type Lexer struct {
	ptr      int32
	contents []byte
}

func NewLexer(contents string) *Lexer {
	return &Lexer{ptr: 0, contents: []byte(contents)}
}

func (l *Lexer) getChr() byte {
	if l.ptr >= int32(len(l.contents)) {
		return 0
	}
	return l.contents[l.ptr]
}

func (l *Lexer) isEnd() bool {
	return l.ptr >= int32(len(l.contents))
}

func (l *Lexer) isWhitespace() bool {
	return l.getChr() == '\n' || l.getChr() == '\r' || l.getChr() == '\t' || l.getChr() == ' '
}

var TOKEN_EOF = newToken(TokenType_EOF, "EOF")

func (l *Lexer) NextToken() *Token {
	for l.isWhitespace() {
		l.ptr++
	}

	switch l.getChr() {
	case '=':
		l.ptr++
		if l.getChr() == '=' {
			l.ptr++
			return newToken(TokenType_EQUALS, "==")
		}
		return newToken(TokenType_ASSIGN, "=")
	case '+':
		l.ptr++
		if l.getChr() == '=' {
			l.ptr++
			return newToken(TokenType_PLUS_ASSIGN, "+=")
		}
		return newToken(TokenType_PLUS, "+")
	case '-':
		l.ptr++
		if l.getChr() == '=' {
			l.ptr++
			return newToken(TokenType_MINUS_ASSIGN, "-=")
		}
		return newToken(TokenType_MINUS, "-")
	case '/':
		l.ptr++
		if l.getChr() == '=' {
			l.ptr++
			return newToken(TokenType_DIVIDE_ASSIGN, "/=")
		}
		return newToken(TokenType_DIVIDE, "/")
	case '*':
		l.ptr++
		if l.getChr() == '=' {
			l.ptr++
			return newToken(TokenType_STAR_ASSIGN, "*=")
		}
		return newToken(TokenType_STAR, "*")
	}

	return TOKEN_EOF
}

func newToken(tt TokenType, value string) *Token {
	return &Token{tContents: value, tType: tt}
}
