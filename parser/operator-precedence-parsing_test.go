package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type OperatorTest struct {
	input    string
	expected string
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	var tests = []OperatorTest{
		{
			input:    "-a * b",
			expected: "((-a) * b)",
		},
		{
			input:    "!-a",
			expected: "(!(-a))",
		},
		{
			input:    "a + b + c",
			expected: "((a + b) + c)",
		},
		{
			input:    "a + b - c",
			expected: "((a + b) - c)",
		},
		{
			input:    "a * b * c",
			expected: "((a * b) * c)",
		},
		{
			input:    "a * b / c",
			expected: "((a * b) / c)",
		},
		{
			input:    "a + b / c",
			expected: "(a + (b / c))",
		},
		{
			input:    "a + b * c + d / e - f",
			expected: "(((a + (b * c)) + (d / e)) - f)",
		},
		{
			input:    "3 + 4; -5 * 5",
			expected: "(3 + 4)((-5) * 5)",
		},
		{
			input:    "5 > 4 == 3 < 4",
			expected: "((5 > 4) == (3 < 4))",
		},
		{
			input:    "5 > 4 != 3 < 4",
			expected: "((5 > 4) != (3 < 4))",
		},
		{
			input:    "3 + 4 * 5 == 3 * 1 + 4 * 5",
			expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			input:    "true",
			expected: "true",
		},
		{
			input:    "false",
			expected: "false",
		},
		{
			input:    "3 > 5 == false",
			expected: "((3 > 5) == false)",
		},
		{
			input:    "3 < 5 == true",
			expected: "((3 < 5) == true)",
		},
		{
			input:    "1 + (2 + 3) + 4",
			expected: "((1 + (2 + 3)) + 4)",
		},
		{
			input:    "(5 + 5) * 2",
			expected: "((5 + 5) * 2)",
		},
		{
			input:    "2 / (5 + 5)",
			expected: "(2 / (5 + 5))",
		},
		{
			input:    "-(5 + 5)",
			expected: "(-(5 + 5))",
		},
		{
			input:    "!(true == true)",
			expected: "(!(true == true))",
		},
		{
			input:    "a + add(b * c) + d",
			expected: "((a + add((b * c))) + d)",
		},
		{
			input:    "add(a, b, 1, 2 * 3, 4 * 5, add(6, 7 * 8))",
			expected: "add(a, b, 1, (2 * 3), (4 * 5), add(6, (7 * 8)))",
		},
		{
			input:    "add(a + b + c * d / f + g)",
			expected: "add((((a + b) + ((c * d) / f)) + g))",
		},
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
