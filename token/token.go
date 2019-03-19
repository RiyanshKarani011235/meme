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
	case TokenEOF:
		return "EOF"
	case TokenError:
		return token.Literal
	}

	// general case, print at most 10 characters
	if len(token.Literal) > 10 {
		return fmt.Sprintf("%.10q...", token.Literal)
	}
	return fmt.Sprintf("%q", token.Literal)
}
