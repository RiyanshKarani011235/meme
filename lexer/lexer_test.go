package lexer_test

import (
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
	Describe("Testing Single Character Tokens", func() {
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
			}

			testNextToken(testString, lex, testOutput)

		})
	})
})

func testNextToken(input string, lex *lexer.Lexer, testInput []*testStruct) {
	for _, expectedOutput := range testInput {
		nextToken := lex.NextToken()

		Expect(nextToken.Type).To(Equal(expectedOutput.expectedType))
		Expect(nextToken.Literal).To(Equal(expectedOutput.expectedLiteral))
	}
}
