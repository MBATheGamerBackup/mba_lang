package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MBATheGamer/mba_lang/evaluator"
	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/object"
	"github.com/MBATheGamer/mba_lang/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)
	var env = object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		var scanned = scanner.Scan()

		if !scanned {
			return
		}

		var line = scanner.Text()

		var l = lexer.New(line)
		var p = parser.New(l)
		var program = p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		var evaluated = evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! I think we have a problem(s) to fix here!\n")
	io.WriteString(out, "  parser error:\n")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
}
