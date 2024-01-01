package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/object"
)

func TestFunctionObject(t *testing.T) {
	var input = "fn(x) { x + 2; };"

	var evaluated = testEval(input)

	var fn, ok = evaluated.(*object.Function)

	if !ok {
		t.Fatalf(
			"object is not Function. got=%T (%+v)",
			evaluated,
			evaluated,
		)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf(
			"function has wrong parameters. Parameters=%+v",
			fn.Parameters,
		)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf(
			"parameter is not 'x'. got=%q",
			fn.Parameters[0],
		)
	}

	var expectedBody = "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf(
			"body is not %q. got=%q",
			expectedBody,
			fn.Body.String(),
		)
	}
}
