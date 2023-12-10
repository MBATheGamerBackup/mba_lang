package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type LetTest struct {
	input              string
	expectedIdentifier string
	expectedValue      interface{}
}

func TestLetStatement(t *testing.T) {
	var tests = []LetTest{
		{
			input:              "let x = 5;",
			expectedIdentifier: "x",
			expectedValue:      5,
		},
		{
			input:              "let y = true;",
			expectedIdentifier: "y",
			expectedValue:      true,
		},
		{
			input:              "let foobar = y;",
			expectedIdentifier: "foobar",
			expectedValue:      "y",
		},
	}

	for _, tt := range tests {
		var l = lexer.New(tt.input)
		var p = parser.New(l)
		var program = p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain 1 statements. got=%d",
				len(program.Statements),
			)
		}

		var statement = program.Statements[0]

		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}

		var value = statement.(*ast.LetStatement).Value

		if !testLiteralExpression(t, value, tt.expectedValue) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf(
			"s.TokenLiteral not 'let'. got=%q",
			s.TokenLiteral(),
		)
		return false
	}

	var letStatement, ok = s.(*ast.LetStatement)

	if !ok {
		t.Errorf(
			"s not *ast.LetStatement. got=%T",
			s,
		)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf(
			"letStatement.Name.Value not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf(
			"s.Name not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)
		return false
	}

	return true
}
