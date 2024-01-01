package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type BooleanTest struct {
	input    string
	expected bool
}

func TestBooleanExpression(t *testing.T) {
	var tests = []BooleanTest{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		var l = lexer.New(tt.input)
		var p = parser.New(l)
		var program = p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program has not enough statements. got=%d",
				len(program.Statements),
			)
		}

		var statement, ok = program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf(
				"program.Statement[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0],
			)
		}

		var boolean *ast.Boolean
		boolean, ok = statement.Expression.(*ast.Boolean)

		if !ok {
			t.Fatalf(
				"expression not *ast.Boolean. got=%T",
				statement.Expression,
			)
		}

		if boolean.Value != tt.expected {
			t.Errorf(
				"boolean.Value not %t. got=%t",
				tt.expected,
				boolean.Value,
			)
		}
	}
}
