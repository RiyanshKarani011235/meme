package ast

import (
	"github.com/riyanshkarani011235/meme/lexer"
	"github.com/riyanshkarani011235/meme/token"
)

type Node interface {
	Token() *token.Token
}

type Statement interface {
	Node
	statementNode()
}
