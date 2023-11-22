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

var letTests = map[string]string{
	"first-test": `
let x = 5;
let y = 10;
let foobar = 838383;
	`,
	"second-test": `
let x 5;
let = 10;
let 838383;
	`,
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	var errors = p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %derrors.", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testingLetStatement(input string, t *testing.T) {
	var l = lexer.New(input)
	var p = parser.New(l)

	var program = p.ParseProgram()
	checkParserErrors(t, p)
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

func TestFirstLetStatement(t *testing.T) {
	testingLetStatement(letTests["first-test"], t)
}

func TestSecondLetStatement(t *testing.T) {
	testingLetStatement(letTests["second-test"], t)
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
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	for _, statement := range program.Statements {
		var returnStatement, ok = statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("statement not *ast.ReturnStatement, got=%T",
				statement,
			)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got=%q",
				returnStatement.TokenLiteral(),
			)
		}
	}
}
