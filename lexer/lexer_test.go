package lexer

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/token"
)

type TestToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var inputs = map[string]string{
	"first-test": `=+(){},;`,
}

func TestNextToken(t *testing.T) {
	var input = inputs["first-test"]

	var tests = []TestToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	var l = New(input)

	for i, tt := range tests {
		var tok = l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
