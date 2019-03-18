package lexer

import (
	"meme/token"
)

type Lexer interface {
	NextToken() token.Token
}
