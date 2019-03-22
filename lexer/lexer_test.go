package lexer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/riyanshkarani011235/meme/lexer"
	"github.com/riyanshkarani011235/meme/token"
)

type testStruct struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var _ = Describe("tokenize and NextToken", func() {
	Context("Testing Braces and Parens", func() {
		testString := "(){}[]<>"
		testOutput := []*testStruct{
			&testStruct{token.TokenLeftParen, "("},
			&testStruct{token.TokenRightParen, ")"},
			&testStruct{token.TokenLeftBrace, "{"},
			&testStruct{token.TokenRightBrace, "}"},
			&testStruct{token.TokenLeftSquareBrace, "["},
			&testStruct{token.TokenRightSquareBrace, "]"},
			&testStruct{token.TokenLeftAngleBrace, "<"},
			&testStruct{token.TokenRightAngleBrace, ">"},
			&testStruct{token.TokenEOF, "EOF"},
		}

		It("Tokenize Should generate correct tokens", func() {
			l := NewLexer(testString)
			testTokenize(l, testOutput)
		})

		It("NextToken Should generate correct tokens", func() {
			l := NewLexer(testString)
			testNextToken(l, testOutput)
		})
	})

	Context("Testing oneof", func() {
		testString := "oneof"
		testOutput := []*testStruct{
			&testStruct{token.TokenOneOf, "oneof"},
			&testStruct{token.TokenEOF, "EOF"},
		}

		It("Tokenize should generate correct Tokens", func() {
			l := NewLexer(testString)
			testTokenize(l, testOutput)
		})

		It("NextToken Should generate correct tokens", func() {
			l := NewLexer(testString)
			testNextToken(l, testOutput)
		})
	})

	Context("Testing Keywords and Identifiers", func() {
		testString := `concept Hello<T> extends World {
				required foo [oneof(Concept, Relation)]
				optional bar [anyof(Baz, Box)]
				required baz [anyof(boolean, integer, string)]
				}`
		testOutput := []*testStruct{
			&testStruct{token.TokenConcept, "concept"},
			&testStruct{token.TokenIdentifier, "Hello"},
			&testStruct{token.TokenLeftAngleBrace, "<"},
			&testStruct{token.TokenIdentifier, "T"},
			&testStruct{token.TokenRightAngleBrace, ">"},
			&testStruct{token.TokenExtends, "extends"},
			&testStruct{token.TokenIdentifier, "World"},
			&testStruct{token.TokenLeftBrace, "{"},
			&testStruct{token.TokenRequired, "required"},
			&testStruct{token.TokenIdentifier, "foo"},
			&testStruct{token.TokenLeftSquareBrace, "["},
			&testStruct{token.TokenOneOf, "oneof"},
			&testStruct{token.TokenLeftParen, "("},
			&testStruct{token.TokenIdentifier, "Concept"},
			&testStruct{token.TokenComma, ","},
			&testStruct{token.TokenIdentifier, "Relation"},
			&testStruct{token.TokenRightParen, ")"},
			&testStruct{token.TokenRightSquareBrace, "]"},
			&testStruct{token.TokenOptional, "optional"},
			&testStruct{token.TokenIdentifier, "bar"},
			&testStruct{token.TokenLeftSquareBrace, "["},
			&testStruct{token.TokenAnyOf, "anyof"},
			&testStruct{token.TokenLeftParen, "("},
			&testStruct{token.TokenIdentifier, "Baz"},
			&testStruct{token.TokenComma, ","},
			&testStruct{token.TokenIdentifier, "Box"},
			&testStruct{token.TokenRightParen, ")"},
			&testStruct{token.TokenRightSquareBrace, "]"},
			&testStruct{token.TokenRequired, "required"},
			&testStruct{token.TokenIdentifier, "baz"},
			&testStruct{token.TokenLeftSquareBrace, "["},
			&testStruct{token.TokenAnyOf, "anyof"},
			&testStruct{token.TokenLeftParen, "("},
			&testStruct{token.TokenBooleanType, "boolean"},
			&testStruct{token.TokenComma, ","},
			&testStruct{token.TokenIntegerType, "integer"},
			&testStruct{token.TokenComma, ","},
			&testStruct{token.TokenStringType, "string"},
			&testStruct{token.TokenRightParen, ")"},
			&testStruct{token.TokenRightSquareBrace, "]"},
			&testStruct{token.TokenRightBrace, "}"},
			&testStruct{token.TokenEOF, "EOF"},
		}

		It("Tokenize should generate correct Tokens", func() {
			l := NewLexer(testString)
			testTokenize(l, testOutput)
		})

		It("NextToken Should generate correct tokens", func() {
			l := NewLexer(testString)
			testNextToken(l, testOutput)
		})
	})
})

func testTokensList(tokens []token.Token, testInput []*testStruct) {
	// verify each token
	for i, expectedOutput := range testInput {
		t := tokens[i]
		// fmt.Println(expectedOutput)
		// fmt.Println(t)

		Expect(t.Literal).To(Equal(expectedOutput.expectedLiteral))
		Expect(t.Type).To(Equal(expectedOutput.expectedType))
	}
}

func testTokenize(l *Lexer, testInput []*testStruct) {
	testTokensList(l.tokenize(), testInput)
}

func testNextToken(l *Lexer, testInput []*testStruct) {
	tokens := make([]token.Token, 0)

	i := 0
	for t, ok := l.NextToken(); ok; t, ok = l.NextToken() {
		i += 1
		tokens = append(tokens, t)
	}

	testTokensList(tokens, testInput)
}
