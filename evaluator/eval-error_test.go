package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/object"
)

type TestError struct {
	input    string
	expected string
}

func TestErrorHandling(t *testing.T) {
	var tests = []TestError{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"false + true;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; false + true; 5;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { false + true; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
	if (10 > 1) {
		return true + false;
	}
	return 1;
}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)

		var errObj, ok = evaluated.(*object.Error)

		if !ok {
			t.Errorf(
				"no error object returned. got=%T(%+v)",
				evaluated,
				evaluated,
			)
			continue
		}

		if errObj.Message != tt.expected {
			t.Errorf(
				"wrong error message. expected=%q, got=%q",
				tt.expected,
				errObj.Message,
			)
		}
	}
}
