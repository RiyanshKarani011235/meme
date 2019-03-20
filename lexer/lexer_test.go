package lexer_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/riyanshkarani011235/meme/lexer"
	"github.com/riyanshkarani011235/meme/token"
)

type testStruct struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var _ = Describe("NextToken", func() {
	Context("Testing Braces and Parens", func() {
		It("Should generate correct tokens", func() {
			testString := "(){}[]<>"
			lex := lexer.NewLexer(testString)
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

			testNextToken(testString, lex, testOutput)

		})
	})

	Context("Testing oneof", func() {
		It("Should generate correct Tokens", func() {
			testString := "oneof"
			testOutput := []*testStruct{
				&testStruct{token.TokenOneOf, "oneof"},
				&testStruct{token.TokenEOF, "EOF"},
			}
			lex := lexer.NewLexer(testString)

			testNextToken(testString, lex, testOutput)
		})
	})

	Context("Testing Keywords and Identifiers", func() {
		It("Should generate correct Tokens", func() {
			testString := `concept Hello<T> extends World {
				required foo [oneof(Concept, Relation)]
				optional bar [anyof(Baz, Box)]
				}`
			lex := lexer.NewLexer(testString)
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
				&testStruct{token.TokenRightBrace, "}"},
				&testStruct{token.TokenEOF, "EOF"},
			}

			testNextToken(testString, lex, testOutput)
		})
	})
})

func testNextToken(input string, l *lexer.Lexer, testInput []*testStruct) {
	tokens := l.Tokenize()

	// verify each token
	for i, expectedOutput := range testInput {
		t := tokens[i]

		Expect(t.Literal).To(Equal(expectedOutput.expectedLiteral))
		Expect(t.Type).To(Equal(expectedOutput.expectedType))
		// fmt.Printf("%v\n", t)
	}
}
