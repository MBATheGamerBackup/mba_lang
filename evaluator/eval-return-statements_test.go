package evaluator_test

import "testing"

func TestReturnStatements(t *testing.T) {
	var tests = []TestEvalInterger{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`
if (10 > 1) {
	if (10 > 1) {
		return 10;
	}
	return 10;
}
			`,
			10,
		},
	}

	for _, tt := range tests {
		var evaluated = testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
