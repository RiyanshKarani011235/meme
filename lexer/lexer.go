package lexer

import (
	"github.com/riyanshkarani011235/meme/token"
)

type Lexer interface {
	NextToken() token.Token
}
