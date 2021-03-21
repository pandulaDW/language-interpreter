package repl

import (
	"bufio"
	"fmt"
	"github.com/pandulaDW/language-interpreter/lexer"
	"github.com/pandulaDW/language-interpreter/tokens"
	"io"
	"log"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprint(out, PROMPT+" ")
		if err != nil {
			log.Fatal(err)
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
