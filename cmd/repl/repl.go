package main

import (
	"bufio"
	"fmt"
	"github.com/riyanshkarani011235/meme/lexer"
	"github.com/riyanshkarani011235/meme/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
	}
}
