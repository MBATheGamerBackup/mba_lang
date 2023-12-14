package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/object"
)

type TestEvalBoolean struct {
	input    string
	expected bool
}

func TestEvalBooleanExpression(t *testing.T) {
	var tests = []TestEvalBoolean{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	var result, ok = obj.(*object.Boolean)

	if !ok {
		t.Errorf(
			"object is not Boolean. got=%T (%+v)",
			obj,
			obj,
		)
		return false
	}

	if result.Value != expected {
		t.Errorf(
			"object has wrong value> got=%t, want=%t",
			result.Value,
			expected,
		)
	}

	return true
}
