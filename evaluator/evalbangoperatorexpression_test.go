package evaluator_test

import "testing"

func TestBangOperator(t *testing.T) {
	var tests = []TestEvalBoolean{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}
