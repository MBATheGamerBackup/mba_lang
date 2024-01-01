package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

func TestReturnStatemnt(t *testing.T) {
	var input = `
		return 5;
		return 10;
		return 993322;
	`

	var l = lexer.New(input)
	var p = parser.New(l)

	var program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	for _, statement := range program.Statements {
		var returnStatement, ok = statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf(
				"statement not *ast.ReturnStatement, got=%T",
				statement,
			)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf(
				"returnStatement.TokenLiteral not 'return', got=%q",
				returnStatement.TokenLiteral(),
			)
		}
	}
}
