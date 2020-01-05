package lexer

import (
	"testing"

	"github.com/holyhoppsan/monkeyinterpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{ token.ASSIGN, "=" },
		{ token.PLUS, "+" },
		{ token.LPAREN, "(" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.RBRACE, "}" },
		{ token.COMMA, "," },
		{ token.SEMICOLON, ";" },
		{ token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		nextToken := l.NextToken()

		if nextToken.Type != tt.expectedType {
			t.Fatalf("tests[%d] tokentype wrong. expected=%q, got=%q", i, tt.expectedType, nextToken.Type)
		}

		if nextToken.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, nextToken.Literal)
		}
	}
}