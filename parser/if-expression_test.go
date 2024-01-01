package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestIfExpression(t *testing.T) {
	var input = `if (x < y) { x }`

	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Bdoy does not contain %d statements. ogt=%d\n",
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

	var expression *ast.IfExpression
	expression, ok = statement.Expression.(*ast.IfExpression)

	if !ok {
		t.Fatalf(
			"statement.Expressiom is not ast.IfExpression. got=%T",
			statement.Expression,
		)
	}

	if !testInfixExpression(t, expression.Condition, "x", "<", "y") {
		return
	}

	if len(expression.Consequence.Statements) != 1 {
		t.Errorf(
			"consequence is not statements> got=%d\n",
			len(expression.Consequence.Statements),
		)
	}

	var consequence *ast.ExpressionStatement
	consequence, ok = expression.Consequence.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf(
			"State,emts[0] is not ast.ExpressionStatement. got=%T",
			expression.Consequence.Statements[0],
		)
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if expression.Alternative != nil {
		t.Errorf(
			"expression.Alternative.Statements was not nil. got=%+v",
			expression.Alternative,
		)
	}
}
