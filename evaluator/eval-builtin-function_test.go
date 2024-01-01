package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/object"
)

type TestFunction struct {
	input    string
	expected interface{}
}

func TestBuiltinFunction(t *testing.T) {
	var tests = []TestFunction{
		{
			`len("")`,
			0,
		},
		{
			`len("four")`,
			4,
		},
		{
			`len("Hello, World!")`,
			13,
		},
		{
			`len(1)`,
			"argument to `len` not supported, got INTEGER",
		},
		{
			`len("one", "two")`,
			"wrong number of arguments. got=2, want=1",
		},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			var errObj, ok = evaluated.(*object.Error)

			if !ok {
				t.Errorf(
					"object is not Error. got=%T (%+v)",
					evaluated,
					evaluated,
				)
				continue
			}

			if errObj.Message != expected {
				t.Errorf(
					"wrong error message. expected=%q, got=%q",
					expected,
					errObj.Message,
				)
			}
		}
	}
}
