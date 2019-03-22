package parser

import (
	"github.com/riyanshkarani011235/meme/lexer"
)

type Parser struct {
	l *lexer.Lexer
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
}
