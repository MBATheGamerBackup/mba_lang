package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/evaluator"
	"github.com/MBATheGamer/mba_lang/object"
)

type TestIfElse struct {
	input    string
	expected interface{}
}

func TestIfElseExpression(t *testing.T) {
	var tests = []TestIfElse{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)
		var integer, ok = tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != evaluator.NULL {
		t.Errorf(
			"object is not NULL. got=%T (%+v)",
			obj,
			obj,
		)
	}

	return true
}
