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

	// identifier
	TokenIdentifier

	// basic types / literals
	TokenIntegerType // integer
	TokenStringType  // string

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
)
