package parser_test

import (
	"fmt"
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type ExpectedIdentifier struct {
	expectedIdentifier string
}

type PrefixTest struct {
	input        string
	operator     string
	integerValue int64
}

type InfixTest struct {
	input      string
	leftValue  int64
	operator   string
	rightValue int64
}

type OperatorTest struct {
	input    string
	expected string
}

type BooleanTest struct {
	input    string
	expected bool
}

var letTests = map[string]string{
	"first-test": `
let x = 5;
let y = 10;
let foobar = 838383;
	`,
	"second-test": `
let x 5;
let = 10;
let 838383;
	`,
}

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

func testingLetStatement(input string, t *testing.T) {
	var l = lexer.New(input)
	var p = parser.New(l)

	var program = p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	var tests = []ExpectedIdentifier{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		var statement = program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestFirstLetStatement(t *testing.T) {
	testingLetStatement(letTests["first-test"], t)
}

func TestSecondLetStatement(t *testing.T) {
	testingLetStatement(letTests["second-test"], t)
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

func TestReturnStatemnt(t *testing.T) {
	var input = `
return 5;
return 10;
return 993322;
	`

	var l = lexer.New(input)
	var p = parser.New(l)

	var program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	for _, statement := range program.Statements {
		var returnStatement, ok = statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf(
				"statement not *ast.ReturnStatement, got=%T",
				statement,
			)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf(
				"returnStatement.TokenLiteral not 'return', got=%q",
				returnStatement.TokenLiteral(),
			)
		}
	}
}

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

func TestIntegerLiteralExpression(t *testing.T) {
	var input = "5;"

	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements has not enough statements. got=%d",
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

	var literal *ast.IntegerLiteral
	literal, ok = statement.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf(
			"expression not *ast.IntegerLiteral. got=%T",
			statement.Expression,
		)
	}

	if literal.Value != 5 {
		t.Errorf(
			"literal.Value not %d. got=%d",
			5,
			literal.Value,
		)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf(
			"literal.TokenLiteral not %s. got=%s",
			"5",
			literal.TokenLiteral(),
		)
	}
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

func TestPrefixExpression(t *testing.T) {
	var prefixTests = []PrefixTest{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
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

		if !testIntegerLiteral(t, expression.Right, tt.integerValue) {
			return
		}
	}
}

func TestParseInfixExpression(t *testing.T) {
	var infixTests = []InfixTest{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		var l = lexer.New(tt.input)
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
				"program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0],
			)
		}

		var expression *ast.InfixExpression
		expression, ok = statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf(
				"expression is not ast.InfixExpression. got=%T",
				statement.Expression,
			)
		}

		if expression.Operator != tt.operator {
			t.Fatalf(
				"expression.Operator is not '%s'. got=%s",
				tt.operator,
				expression.Operator,
			)
		}

		if !testIntegerLiteral(t, expression.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	var tests = []OperatorTest{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 > 4 != 3 < 4", "((5 > 4) != (3 < 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
	}

	for _, tt := range tests {
		var l = lexer.New(tt.input)
		var p = parser.New(l)
		var program = p.ParseProgram()
		checkParserErrors(t, p)

		var actual = program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
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

func testLiteralExpression(t *testing.T, expression ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	}

	t.Errorf(
		"type of expression not handled. got=%T",
		expression,
	)

	return false
}

func testInfixExpression(
	t *testing.T,
	expression ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
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
