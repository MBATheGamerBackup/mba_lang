package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type FunctionParametersTest struct {
	input          string
	expectedParams []string
}

func TestFunctionParametersParsing(t *testing.T) {
	var tests = []FunctionParametersTest{
		{
			input:          "fn() {};",
			expectedParams: []string{},
		},
		{
			input:          "fn(x) {};",
			expectedParams: []string{"x"},
		},
		{
			input: "fn(x, y, z) {};",
			expectedParams: []string{
				"x",
				"y",
				"z",
			},
		},
	}

	for _, tt := range tests {
		var l = lexer.New(tt.input)
		var p = parser.New(l)
		var program = p.ParseProgram()
		checkParserErrors(t, p)

		var statement = program.Statements[0].(*ast.ExpressionStatement)
		var function = statement.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf(
				"length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams),
				len(function.Parameters),
			)
		}

		for i, identifier := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], identifier)
		}
	}
}
