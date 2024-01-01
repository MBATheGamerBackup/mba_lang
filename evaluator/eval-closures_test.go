package evaluator_test

import "testing"

func TestClosures(t *testing.T) {
	var input = `
let newAdder = fn(x) {
	fn(y) { x + y };
};

let addTwo = newAdder(2);
addTwo(2);
	`

	testIntegerObject(t, testEval(input), 4)
}
