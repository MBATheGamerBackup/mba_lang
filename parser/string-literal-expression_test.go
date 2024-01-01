package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestStringLiteralExpression(t *testing.T) {
	var input = `"Hello, World!"`

	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	var statement = program.Statements[0].(*ast.ExpressionStatement)
	var literal, ok = statement.Expression.(*ast.StringLiteral)

	if !ok {
		t.Fatalf(
			"expression not *ast.StringLiteral. got=%T",
			statement.Expression,
		)
	}

	if literal.Value != "Hello, World!" {
		t.Errorf(
			"literal.Value not %q. got=%q",
			"Hello, World!",
			literal.Value,
		)
	}
}
