package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestCallExpression(t *testing.T) {
	var input = "add(1, 2 * 3, 4 + 5);"

	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	var statement, ok = program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf(
			"statement is not ast.ExpressionStatement. got=%T",
			program.Statements[0],
		)
	}

	var expression *ast.CallExpression
	expression, ok = statement.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf(
			"statement.Expression is not ast.CallExpression. got=%T",
			statement.Expression,
		)
	}

	if testIdentifier(t, expression.Function, "add") {
		return
	}

	if len(expression.Arguments) != 3 {
		t.Fatalf(
			"wrong length of arguments. got=%d",
			len(expression.Arguments),
		)
	}

	testLiteralExpression(t, expression.Arguments[0], 1)
	testInfixExpression(t, expression.Arguments[1], 2, "*", 3)
	testInfixExpression(t, expression.Arguments[2], 4, "+", 5)
}
