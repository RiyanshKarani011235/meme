package token

import (
	"fmt"
)

type FileInfo struct {
	FileName string
	FilePath string
}

type Token struct {
	Type              TokenType
	Literal           string
	FileInfo          *FileInfo
	LineNumber        int
	ColumnNumberStart int
	ColumnNumberEnd   int
}

func (token Token) String() string {
	// handle special cases
	switch token.Type {
	case TokenError:
		return token.Literal
	case TokenIdentifier, TokenSingleLineComment, TokenMultiLineComment:
		// Identifiers / comments, print at most 10 characters
		if len(token.Literal) > 10 {
			return fmt.Sprintf("%v(%.10v...)", token.Type, token.Literal)
		} else {
			return fmt.Sprintf("%v(%v)", token.Type, token.Literal)
		}
	default:
		// anything else
		return fmt.Sprintf("%v", token.Type)
	}
}
