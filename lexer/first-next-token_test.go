package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/token"
)

func TestFirstNextToken(t *testing.T) {
	const input = "=+(){},;"

	var tests = []TestToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testInput(input, tests, t)
}
