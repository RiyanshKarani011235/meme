package lexer

import (
	"fmt"
	"github.com/riyanshkarani011235/meme/token"
)

const (
	eof = '\r'
)

type Lexer struct {
	input      string
	lineNumber int              // the line number in the input string the lexer is at
	startPos   int              // starting position of this item
	currentPos int              // current position in the input
	tokens     chan token.Token // the channel at which tokens are emitted
}

func NewLexer(input string) (l *Lexer, c chan token.Token) {
	l = &Lexer{
		input:      input,
		lineNumber: 0,
		startPos:   0,
		currentPos: 0,
		tokens:     make(chan token.Token),
	}

	go l.run()

	go func() {
		for message := range l.tokens {
			fmt.Printf("%v\n", message)
		}
	}()
	return l, l.tokens
}

// run the lexer until nil state is reached
func (l *Lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}

	close(l.tokens)
}

// emits a token to the tokens channel of the lexer
func (l *Lexer) emit(t token.Token) {
	l.tokens <- t
	l.startPos = l.currentPos
}

// reads one character from the input and returns it
func (l *Lexer) next() (character byte) {
	if l.currentPos >= len(l.input) {
		character = eof
		l.currentPos += 1
		return
	}

	character = l.input[l.currentPos]
	l.currentPos += 1
	return
}

// skips over the pending input before this point
func (l *Lexer) ignore() {
	l.startPos = l.currentPos
}

// backs up in the input by one character
func (l *Lexer) backup() {
	l.currentPos -= 1

	// @todo panic instead?
	if l.startPos > l.currentPos {
		l.startPos = l.currentPos
	}
}

// ----------------
// helper functions
// ----------------

func isAlphabet(character byte) bool {
	return ('a' <= character && character <= 'z') || ('A' <= character && character <= 'Z')
}

func isNumeric(character byte) bool {
	return '0' <= character && character <= '9'
}

func isAlphaNumeric(character byte) bool {
	return isAlphabet(character) || isNumeric(character)
}

func isAlphaNumericOrUnderscore(character byte) bool {
	return isAlphaNumeric(character) || character == '_'
}

func isWhiteSpace(character byte) bool {
	switch character {
	case ' ', '\t', '\n':
		return true
	default:
		return false
	}
}

// ---------------
// State functions
// ---------------
type stateFn func(*Lexer) stateFn

func startState(l *Lexer) stateFn {
	l.startPos = 0
	l.currentPos = 0
	return tokenizeText
}

func eatWhiteSpace(l *Lexer) stateFn {
	switch l.next() {
	case ' ', '\t':
		// keep eating white space
		return eatWhiteSpace
	case '\n':
		// increment the line number
		l.lineNumber += 1
		// continue eating white space
		return eatWhiteSpace
	default:
		// no longer a white space, ignore the white space read
		l.backup()
		l.ignore()
		// start reading the text
		return tokenizeText
	}
}

func tokenizeText(l *Lexer) stateFn {
	c := l.next()

	// whitespace
	if isWhiteSpace(c) {
		l.backup()
		return eatWhiteSpace
	}

	// keyword or identifier
	if isAlphabet(c) {
		l.backup()
		return tokenizeKeywordOrIdentifier
	}

	switch c {
	// parens or braces
	case '{', '}', '(', ')', '[', ']', '<', '>':
		l.backup()
		return tokenizeParensAndBraces

	// end of file
	case eof:
		l.backup()
		return tokenizeEndOfFile

	// comma
	case ',':
		l.backup()
		return tokenizeDelimiter

	// anything else, must be a syntax error
	default:
		return syntaxError

	}

}

func tokenizeDelimiter(l *Lexer) stateFn {
	c := l.next()
	delimiterTokenTypeMap := map[string]token.TokenType{
		",": token.TokenComma,
	}

	// valid delimiter, emit token
	if tokenType, ok := delimiterTokenTypeMap[string(c)]; ok {
		l.emit(token.Token{
			Type:       tokenType,
			Literal:    string(c),
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})

		return tokenizeText

	}

	// invalid delimiter, syntax error
	return syntaxError
}

func tokenizeParensAndBraces(l *Lexer) stateFn {
	c := l.next()

	t := token.Token{
		Literal:    string(c),
		LineNumber: l.lineNumber,
		// @todo FileInfo
		ColumnNumberStart: l.currentPos,
		ColumnNumberEnd:   l.currentPos,
	}

	switch c {
	case '{':
		t.Type = token.TokenLeftBrace
		break
	case '}':
		t.Type = token.TokenRightBrace
		break
	case '(':
		t.Type = token.TokenLeftParen
		break
	case ')':
		t.Type = token.TokenRightParen
		break
	case '[':
		t.Type = token.TokenLeftSquareBrace
		break
	case ']':
		t.Type = token.TokenRightSquareBrace
		break
	case '<':
		t.Type = token.TokenLeftAngleBrace
		break
	case '>':
		t.Type = token.TokenRightAngleBrace
		break
	default:
		panic(fmt.Sprintf("invalid paren/brace character %v", c))
	}

	// emit the token
	l.emit(t)

	// start reading text again
	return tokenizeText
}

func tokenizeKeywordOrIdentifier(l *Lexer) stateFn {
	keywordsTokenTypeMap := map[string]token.TokenType{
		"concept":  token.TokenConcept,
		"relation": token.TokenRelation,
		"required": token.TokenRequired,
		"optional": token.TokenOptional,
		"integer":  token.TokenIntegerType,
		"string":   token.TokenStringType,
		"oneof":    token.TokenOneOf,
		"anyof":    token.TokenAnyOf,
	}

	c := l.next()

	// alphanumeric or underscore
	if isAlphaNumericOrUnderscore(c) {
		return tokenizeKeywordOrIdentifier
	}

	// anything else
	l.backup()
	literal := l.input[l.startPos:l.currentPos]

	if tokenType, ok := keywordsTokenTypeMap[literal]; ok {
		// keyword
		l.emit(token.Token{
			Type:       tokenType,
			Literal:    literal,
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})
	} else {
		// identifier
		l.emit(token.Token{
			Type:       token.TokenIdentifier,
			Literal:    literal,
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})
	}

	return tokenizeText
}

func tokenizeEndOfFile(l *Lexer) stateFn {
	c := l.next()

	switch c {
	case eof:
		l.emit(token.Token{
			Type:       token.TokenEOF,
			Literal:    string(c),
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})
		return nil
	default:
		// should be an EOF, but is not. Panic
		panic(fmt.Sprintf("character %v is not EOF", c))
	}
}

func generateSyntaxError(errorString string) func(*Lexer) stateFn {
	return func(l *Lexer) stateFn {
		l.emit(token.Token{
			Type:       token.TokenError,
			Literal:    errorString,
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})
		return nil
	}
}

func syntaxError(l *Lexer) stateFn {
	l.emit(token.Token{
		Type:       token.TokenError,
		Literal:    "Syntax Error",
		LineNumber: l.lineNumber,
		// @todo FileInfo
		ColumnNumberStart: l.startPos,
		ColumnNumberEnd:   l.currentPos,
	})
	return nil
}

/*
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '{':
		tok = newToken(token.TokenLeftBrace, l.ch)
	case '}':
		tok = newToken(token.TokenRightBrace, l.ch)
	case '(':
		tok = newToken(token.TokenLeftParen, l.ch)
	case ')':
		tok = newToken(token.TokenRightParen, l.ch)
	case '[':
		tok = newToken(token.TokenLeftSquareBrace, l.ch)
	case ']':
		tok = newToken(token.TokenRightSquareBrace, l.ch)
	case '<':
		tok = newToken(token.TokenLeftAngleBrace, l.ch)
	case '>':
		tok = newToken(token.TokenRightAngleBrace, l.ch)
	}

	// read the next character
	l.readChar()

	// return the read token
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII code for NUL
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(type_ token.TokenType, literal byte) token.Token {
	return token.Token{Type: type_, Literal: string(literal)}
}
*/
