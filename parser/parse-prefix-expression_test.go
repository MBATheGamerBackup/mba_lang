package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type PrefixTest struct {
	input      string
	operator   string
	rightValue interface{}
}

func TestParsePrefixExpression(t *testing.T) {
	var prefixTests = []PrefixTest{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		var l = lexer.New(tt.input)
		var p = parser.New(l)
		var program = p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain %d statements. got=%d",
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

		var expression *ast.PrefixExpression
		expression, ok = statement.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf(
				"expression is not *ast.PrefixExpression. got=%T",
				statement.Expression,
			)
		}

		if expression.Operator != tt.operator {
			t.Errorf(
				"expression.Operator is not '%s'. got=%s",
				tt.operator,
				expression.Operator,
			)
		}

		if !testLiteralExpression(t, expression.Right, tt.rightValue) {
			return
		}
	}
}
