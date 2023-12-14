package evaluator_test

import (
	"github.com/MBATheGamer/mba_lang/evaluator"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/object"
	"github.com/MBATheGamer/mba_lang/parser"
)

func testEval(input string) object.Object {
	var l = lexer.New(input)
	var p = parser.New(l)
	var program = p.ParseProgram()

	return evaluator.Eval(program)
}
