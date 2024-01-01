package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/token"
)

type TestToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testInput(input string, tests []TestToken, t *testing.T) {
	var l = lexer.New(input)

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
