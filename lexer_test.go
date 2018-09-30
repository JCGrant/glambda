package glambda

import "testing"

func TestLex(t *testing.T) {
	tests := []struct {
		input    string
		expected []token
	}{
		{
			``,
			[]token{
				token{tokenEOF, ``},
			},
		},
		{
			`\`,
			[]token{
				token{tokenLambda, `\`},
				token{tokenEOF, ``},
			},
		},
		{
			`0 = \ f x . f (f x)`,
			[]token{
				token{tokenIdentifier, `0`},
				token{tokenEquals, `=`},
				token{tokenLambda, `\`},
				token{tokenIdentifier, `f`},
				token{tokenIdentifier, `x`},
				token{tokenDot, `.`},
				token{tokenIdentifier, `f`},
				token{tokenLeftParen, `(`},
				token{tokenIdentifier, `f`},
				token{tokenIdentifier, `x`},
				token{tokenRightParen, `)`},
				token{tokenEOF, ``},
			},
		},
		{
			`-- hey this is a test`,
			[]token{
				token{tokenComment, "-- hey this is a test"},
				token{tokenEOF, ""},
			},
		},
		{
			`
			 -- there are spaces to the left of this comment

			0 = \ f      x .      f (f     x)


			`,
			[]token{
				token{tokenNewLine, "\n"},
				token{tokenComment, "-- there are spaces to the left of this comment"},
				token{tokenNewLine, "\n\n"},
				token{tokenIdentifier, `0`},
				token{tokenEquals, `=`},
				token{tokenLambda, `\`},
				token{tokenIdentifier, `f`},
				token{tokenIdentifier, `x`},
				token{tokenDot, `.`},
				token{tokenIdentifier, `f`},
				token{tokenLeftParen, `(`},
				token{tokenIdentifier, `f`},
				token{tokenIdentifier, `x`},
				token{tokenRightParen, `)`},
				token{tokenNewLine, "\n\n\n"},
				token{tokenEOF, ``},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			lexer := lex(test.input)
			var tokens []token
			for token := range lexer.tokens {
				tokens = append(tokens, token)
			}
			assertEqual(t, test.expected, tokens)
		})
	}
}
