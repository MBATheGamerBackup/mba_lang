package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestIdentifierExpression(t *testing.T) {
	var input = "foobar;"

	var l = lexer.New(input)
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
			"program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0],
		)
	}

	var identifier *ast.Identifier
	identifier, ok = statement.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf(
			"expression is not *ast.Identifier. got=%T",
			statement.Expression,
		)
	}

	if identifier.Value != "foobar" {
		t.Errorf(
			"identifier.Value is not %s. got=%s",
			"foobar",
			identifier.Value,
		)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf(
			"identifier.TokenLiteral() is not %s. got=%s",
			"foobar",
			identifier.TokenLiteral(),
		)
	}
}
