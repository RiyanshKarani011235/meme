package lexer

import (
	"fmt"

	"github.com/riyanshkarani011235/meme/token"
)

const (
	eof = '\r'
)

type stateFn func(*Lexer) stateFn

type Lexer struct {
	input      string
	lineNumber int              // the line number in the input string the lexer is at
	startPos   int              // starting position of this item
	currentPos int              // current position in the input
	tokens     chan token.Token // the channel at which tokens are emitted
	state      stateFn
}

func NewLexer(input string) (l *Lexer) {
	l = &Lexer{
		input:      input,
		lineNumber: 0,
		startPos:   0,
		currentPos: 0,
		tokens:     make(chan token.Token, 1),
		state:      startState,
	}
	return
}

// Tokenizes the entire input and returns
// a slice of token.Token instances. If using
// the lexer in conjunction with a parser,
// consider using the NextToken method to
// tokenize and parse the input incrementally.
// That, lexing can be done lazily and stopped
// prematurely in case of a Parsing/Lexing error.
func (l *Lexer) tokenize() []token.Token {
	go l.run()
	tokens := make([]token.Token, 0)
	for t := range l.tokens {
		tokens = append(tokens, t)
	}

	return tokens
}

// run the lexer until nil state is reached
func (l *Lexer) run() {
	for l.state != nil {
		l.state = l.state(l)
	}

	close(l.tokens)
}

// reads the input and returns the next token.
// ok is true if the lexer has more tokens to emit.
// ok is false if the lexer has no more tokens to emit.
func (l *Lexer) NextToken() (t token.Token, ok bool) {
	if l.state == nil {
		return token.Token{}, false
	}

	// generate the next state while at least
	// one token is emitted, then stop. This should
	// always work because if the previous state is
	// not nil and the current state becomes nil,
	// either an EOF token or a SynatxError token is
	// emitted.
	for len(l.tokens) == 0 {
		l.state = l.state(l)
	}

	return <-l.tokens, true
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

// peeks the next character without incrementing
// currentPos
func (l *Lexer) peek() (character byte) {
	character = l.next()
	l.backup()
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
		return tokenizeSpecialCharacters

	// end of file
	case eof:
		l.backup()
		return tokenizeEndOfFile

	// delimiters
	case ',':
		l.backup()
		return tokenizeSpecialCharacters

	// comments
	case '/':
		nextCharacter := l.next()

		// do not backup here, because the comment content
		// starts after "//" or "/*"

		if nextCharacter == '/' {
			return tokenizeSingleLineComment
		} else if nextCharacter == '*' {
			return tokenizeMultiLineComment
		} else {
			return syntaxError
		}

	// anything else, must be a syntax error
	default:
		return syntaxError
	}
}

func tokenizeSingleLineComment(l *Lexer) stateFn {
	switch l.next() {
	case '\n', eof:
		// end of comment
		l.backup()
		l.emit(token.Token{
			Type:       token.TokenSingleLineComment,
			Literal:    l.input[l.startPos:l.currentPos],
			LineNumber: l.lineNumber,
			// @todo FileInfo
			ColumnNumberStart: l.startPos,
			ColumnNumberEnd:   l.currentPos,
		})
		return tokenizeText
	default:
		// continue reading comment
		return tokenizeSingleLineComment
	}
}

func tokenizeMultiLineComment(l *Lexer) stateFn {
	switch l.next() {
	case '*':
		if l.peek() == '/' {
			// end of comment
			l.next()
			l.emit(token.Token{
				Type:       token.TokenMultiLineComment,
				Literal:    l.input[l.startPos:l.currentPos],
				LineNumber: l.lineNumber,
				// @todo FileInfo
				ColumnNumberStart: l.startPos,
				ColumnNumberEnd:   l.currentPos,
			})
			fmt.Println("returning to tokenizeText")
			return tokenizeText
		} else {
			// continue reading comment
			return tokenizeMultiLineComment
		}
	case eof:
		// premature end of file, before end of multi line comment
		return syntaxError
	default:
		// continue reading the comment
		return tokenizeMultiLineComment
	}

}

func tokenizeSpecialCharacters(l *Lexer) stateFn {
	c := l.next()
	tokenType, ok := token.TokenTypeLookupMap[string(c)]
	if !ok {
		panic(fmt.Sprintf("unrecognized special character %v", c))
	}

	l.emit(token.Token{
		Type:       tokenType,
		Literal:    string(c),
		LineNumber: l.lineNumber,
		// @todo FileInfo
		ColumnNumberStart: l.currentPos,
		ColumnNumberEnd:   l.currentPos,
	})

	return tokenizeText
}

func tokenizeKeywordOrIdentifier(l *Lexer) stateFn {
	c := l.next()

	// alphanumeric or underscore
	if isAlphaNumericOrUnderscore(c) {
		return tokenizeKeywordOrIdentifier
	}

	// anything else
	l.backup()
	literal := l.input[l.startPos:l.currentPos]

	// we now have a literal, build a token out of it
	t := token.Token{
		Literal:    literal,
		LineNumber: l.lineNumber,
		// @todo FileInfo
		ColumnNumberStart: l.startPos,
		ColumnNumberEnd:   l.currentPos,
	}

	if tokenType, ok := token.TokenTypeLookupMap[literal]; ok {
		// is a keyword
		t.Type = tokenType
	} else {
		// is an identifier
		t.Type = token.TokenIdentifier
	}

	l.emit(t)
	return tokenizeText
}

func tokenizeEndOfFile(l *Lexer) stateFn {
	c := l.next()

	switch c {
	case eof:
		l.emit(token.Token{
			Type:       token.TokenEOF,
			Literal:    "EOF",
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
