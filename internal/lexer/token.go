package lexer

type TokenType uint8

// + - * / %
// < >
// & ^ | ! ~ =

// -= += /= *= %=
// <= >=
// &= ^= |= != ~= ==
// && ||

const (
	TokenType_INVALID TokenType = iota
	TokenType_EOF
	TokenType_PLUS
	TokenType_MINUS
	TokenType_STAR
	TokenType_DIVIDE
	TokenType_MODULO
	TokenType_LESS_THAN
	TokenType_GREATER_THAN
	TokenType_BITWISE_AND
	TokenType_BITWISE_XOR
	TokenType_BITWISE_OR
	TokenType_BITWISE_NOT
	TokenType_LOGICAL_AND
	TokenType_LOGICAL_XOR
	TokenType_LOGICAL_OR
	TokenType_LOGICAL_NOT

	TokenType_ASSIGN
	TokenType_PLUS_ASSIGN
	TokenType_MINUS_ASSIGN
	TokenType_STAR_ASSIGN
	TokenType_DIVIDE_ASSIGN
	TokenType_MODULO_ASSIGN
	TokenType_LESS_THAN_ASSIGN
	TokenType_GREATER_THAN_ASSIGN
	TokenType_BITWISE_AND_ASSIGN
	TokenType_BITWISE_XOR_ASSIGN
	TokenType_BITWISE_OR_ASSIGN
	TokenType_BITWISE_NOT_ASSIGN
	TokenType_EQUALS
	TokenType_NOT_EQUALS

	TokenType_FUNCTION
	TokenType_RBRACE
	TokenType_LBRACE
	TokenType_RBRACKET
	TokenType_LBRACKET
	TokenType_RPAREN
	TokenType_LPAREN

	TokenType_PERIOD

	TokenType_IDENTIFIER
	TokenType_STRING
)

type Token struct {
	tContents string
	tType     TokenType
}

func (t *Token) GetType() TokenType {
	return t.tType
}

func (t *Token) GetContents() string {
	return t.tContents
}
