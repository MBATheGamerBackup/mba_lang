package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MBATheGamer/mba_lang/lexer"
	"github.com/MBATheGamer/mba_lang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)

	for {
		fmt.Printf("%s", PROMPT)

		var scan = scanner.Scan()

		if !scan {
			return
		}

		var line = scanner.Text()

		var l = lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
