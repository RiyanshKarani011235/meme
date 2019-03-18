package token

type TokenType string

const (
	// special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// keywords
	CONCEPT  = "concept"
	RELATION = "relation"
	REQUIRED = "required"
	OPTIONAL = "optional"

	// composite type constructors
	ONEOF = "oneof"
	ANYOF = "anyof"

	// parens and braces
	L_BRACE        = "{"
	R_BRACE        = "}"
	L_PAREN        = "("
	R_PAREN        = ")"
	L_SQUARE_PAREN = "["
	R_SQUARE_PAREN = "]"
	L_ANGLE_PAREN  = "<"
	R_ANGLE_BRACE  = ">"

	// delimiters
	COMMA = ","
)
