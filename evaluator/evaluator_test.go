package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/evaluator"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/object"
	"github.com/MBATheGamer/mba_lang/parser"
)

type TestEvalBoolean struct {
	input    string
	expected bool
}

type TestEvalInterger struct {
	input    string
	expected int64
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

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	var result, ok = obj.(*object.Integer)

	if !ok {
		t.Errorf(
			"object is not Integer. got=%T (%+v)",
			obj,
			obj,
		)
		return false
	}

	if result.Value != expected {
		t.Errorf(
			"object has wrong value. got=%d,, want=%d",
			result.Value,
			expected,
		)
		return false
	}

	return true
}

func testEval(input string) object.Object {
	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()
	var env = object.NewEnvironment()

	return evaluator.Eval(program, env)
}
