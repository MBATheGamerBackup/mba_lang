package parser_test

import (
	"fmt"
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/parser"
)

func checkParserErrors(t *testing.T, p *parser.Parser) {
	var errors = p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf(
		"parser has %d errors.",
		len(errors),
	)

	for _, msg := range errors {
		t.Errorf(
			"parser error: %q",
			msg,
		)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	var integer, ok = il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf(
			"il not *ast.IntegerLiteral. got=%T",
			il,
		)
		return false
	}

	if integer.Value != value {
		t.Errorf(
			"integer.Value not %d. got=%d",
			value,
			integer.Value,
		)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf(
			"integer.TokenLiteral not %d. got=%s",
			value,
			integer.TokenLiteral(),
		)
		return false
	}

	return true
}

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	var identifier, ok = expression.(*ast.Identifier)

	if !ok {
		t.Errorf(
			"expression not *ast.Identifier. got=%T",
			expression,
		)
		return false
	}

	if identifier.Value != value {
		t.Errorf(
			"identifier.Value not %s. got=%s",
			value,
			identifier.Value,
		)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf(
			"identifier.TokenLiteral not %s. got=%s",
			value,
			identifier.TokenLiteral(),
		)
		return false
	}

	return true
}

func testBooleanExpression(t *testing.T, expression ast.Expression, value bool) bool {
	var boolean, ok = expression.(*ast.Boolean)

	if !ok {
		t.Errorf(
			"expression not *ast.Boolean. got=%T",
			expression,
		)
		return false
	}

	if boolean.Value != value {
		t.Errorf(
			"boolean.Value not %t. got=%t",
			value,
			boolean.Value,
		)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf(
			"boolean.TokenLiteral not %t. got=%s",
			value,
			boolean.TokenLiteral(),
		)
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, expression ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	case bool:
		return testBooleanExpression(t, expression, v)
	}

	t.Errorf(
		"type of expression not handled. got=%T",
		expression,
	)

	return false
}

func testInfixExpression(t *testing.T, expression ast.Expression, left interface{}, operator string, right interface{}) bool {
	var opExpression, ok = expression.(*ast.InfixExpression)

	if !ok {
		t.Errorf(
			"expression is not ast.OperatorExpression. got=%T(%s)",
			expression,
			expression,
		)
		return false
	}

	if !testLiteralExpression(t, opExpression.Left, left) {
		return false
	}

	if opExpression.Operator != operator {
		t.Errorf(
			"expression.Operator is not '%s'. got=%q",
			operator,
			opExpression.Operator,
		)
		return false
	}

	if !testLiteralExpression(t, opExpression.Right, right) {
		return false
	}

	return true
}
