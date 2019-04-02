package main

import (
	"bufio"
	"fmt"
	"github.com/riyanshkarani011235/meme/lexer"
	"github.com/riyanshkarani011235/meme/token"
	"io"
	"os"
)

const PROMPT = "> "

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
		tokens := make([]token.Token, 0)

		for {
			t, ok := l.NextToken()
			if !ok {
				break
			}

			tokens = append(tokens, t)
		}

		for _, token := range tokens {
			fmt.Printf("%v\n", token)
		}
	}
}

func main() {
	fmt.Printf("The meme programming language v0.1\n")
	Start(os.Stdin, os.Stdout)
}
