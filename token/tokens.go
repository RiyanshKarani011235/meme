package token

type TokenType int

const (
	TokenError TokenType = iota // error occurred
	// value is the text of error

	// special types
	TokenIllegal // illegal
	TokenEOF     // end of file

	// keywords
	TokenConcept  // concept
	TokenRelation // relation
	TokenRequired // required
	TokenOptional // optional
	TokenExtends  // extends

	// identifier
	TokenIdentifier

	// basic types / literals
	TokenIntegerType // integer
	TokenStringType  // string
	TokenBooleanType // boolean

	// composite type constructors
	TokenOneOf // oneof
	TokenAnyOf // anyof

	// parens and braces
	TokenLeftParen        // (
	TokenRightParen       // )
	TokenLeftBrace        // {
	TokenRightBrace       // }
	TokenLeftSquareBrace  // [
	TokenRightSquareBrace // ]
	TokenLeftAngleBrace   // <
	TokenRightAngleBrace  // >

	// delimiters
	TokenComma // ,

	// comments
	TokenSingleLineComment
	TokenMultiLineComment
)

var TokenTypeLookupMap = map[string]TokenType{
	// keywords
	"concept":  TokenConcept,
	"relation": TokenRelation,
	"required": TokenRequired,
	"optional": TokenOptional,
	"extends":  TokenExtends,

	// basic types / literals
	"integer": TokenIntegerType,
	"string":  TokenStringType,
	"boolean": TokenBooleanType,

	// composite type constructors
	"oneof": TokenOneOf,
	"anyof": TokenAnyOf,

	// parens and braces
	"(": TokenLeftParen,
	")": TokenRightParen,
	"{": TokenLeftBrace,
	"}": TokenRightBrace,
	"[": TokenLeftSquareBrace,
	"]": TokenRightSquareBrace,
	"<": TokenLeftAngleBrace,
	">": TokenRightAngleBrace,

	// delimiters
	",": TokenComma,
}

var tokenString = map[TokenType]string{
	TokenError:             "ERROR",
	TokenIllegal:           "ILLEGAL",
	TokenEOF:               "EOF",
	TokenConcept:           "CONCEPT",
	TokenRelation:          "RELATION",
	TokenRequired:          "REQUIRED",
	TokenOptional:          "OPTIONAL",
	TokenExtends:           "EXTENDS",
	TokenIdentifier:        "IDENTIFIER",
	TokenIntegerType:       "INTEGER",
	TokenStringType:        "STRING",
	TokenBooleanType:       "BOOLEAN",
	TokenOneOf:             "ONEOF",
	TokenAnyOf:             "ANYOF",
	TokenLeftParen:         "LEFT_PAREN",
	TokenRightParen:        "RIGHT_PAREN",
	TokenLeftBrace:         "LEFT_BRACE",
	TokenRightBrace:        "RIGHT_BRACE",
	TokenLeftSquareBrace:   "LEFT_SQUARE_BRACE",
	TokenRightSquareBrace:  "RIGHT_SQUARE_BRACE",
	TokenLeftAngleBrace:    "LEFT_ANGLE_BRACE",
	TokenRightAngleBrace:   "RIGHT_ANGLE_BRACE",
	TokenComma:             "COMMA",
	TokenSingleLineComment: "SINGLE_LINE_COMMENT",
	TokenMultiLineComment:  "MULTI_LINE_COMMENT",
}

func (tokenType TokenType) String() string {
	t, ok := tokenString[tokenType]
	if !ok {
		panic("undefined tokenType")
	}

	return t
}
