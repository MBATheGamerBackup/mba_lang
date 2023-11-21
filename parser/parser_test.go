package parser_test

import (
	"testing"

	"github.com/MBATheGamer/mba_lang/ast"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/parser"
)

type ExpectedIdentifier struct {
	expectedIdentifier string
}

func TestLetStatement(t *testing.T) {
	var input = `
let x = 5;
let y = 10;
let foobar = 838383;
	`

	var l = lexer.New(input)
	var p = parser.New(l)

	var program = p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	var tests = []ExpectedIdentifier{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		var statement = program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q",
			s.TokenLiteral(),
		)
		return false
	}

	var letStatement, ok = s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T",
			s,
		)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)
		return false
	}

	return true
}
