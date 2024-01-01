package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestFunctionLiteralParsing(t *testing.T) {
	var input = `fn(x, y) { x + y; }`

	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Body does not contain %d statements. got=%d\n",
			1,
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

	var function *ast.FunctionLiteral

	function, ok = statement.Expression.(*ast.FunctionLiteral)

	if !ok {
		t.Fatalf(
			"statement.Expression is not ast.FunctionLiteral. got=%T",
			statement.Expression,
		)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf(
			"function literal parameters wrong. want 2, got=%d\n",
			len(function.Parameters),
		)
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf(
			"function.Body.Statements has not 1 statements. got=%d\n",
			len(function.Body.Statements),
		)
	}

	var bodyStatement *ast.ExpressionStatement
	bodyStatement, ok = function.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf(
			"function.Body.Statements[0] is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0],
		)
	}

	testInfixExpression(t, bodyStatement.Expression, "x", "+", "y")
}
