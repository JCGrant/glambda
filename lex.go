package glambda

import (
	"fmt"
	"regexp"
)

type tokenType string

const (
	tokenErr        tokenType = `error`
	tokenLambda               = `\`
	tokenDot                  = `.`
	tokenLeftParen            = `(`
	tokenRightParen           = `)`
	tokenEquals               = `=`
	tokenIdentifier           = `identifier`
	tokenWhitespace           = `whitespace`
)

type token struct {
	tokenType tokenType
	value     string
}

func (t token) String() string {
	return t.value
}

type lexer struct {
	input    string
	position int
	tokens   chan token
}

func lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token),
	}
	go l.run()
	return l
}

func (l *lexer) emit(token token) {
	l.tokens <- token
}

func (l *lexer) nextItem() token {
	return <-l.tokens
}

func (l *lexer) errorf(format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	l.emit(token{tokenErr, value})
}

var tokenRegexes = []struct {
	tokenType tokenType
	regex     *regexp.Regexp
}{
	{tokenLambda, regexp.MustCompile(`\A(\\)`)},
	{tokenDot, regexp.MustCompile(`\A(\.)`)},
	{tokenLeftParen, regexp.MustCompile(`\A(\()`)},
	{tokenRightParen, regexp.MustCompile(`\A(\))`)},
	{tokenEquals, regexp.MustCompile(`\A(=)`)},
	{tokenIdentifier, regexp.MustCompile(`\A(\b[a-zA-Z0-9]+\b)`)},
	{tokenWhitespace, regexp.MustCompile(`\s`)},
}

func (l *lexer) lexOneToken() {
	for _, tr := range tokenRegexes {
		input := l.input[l.position:]
		if value := tr.regex.FindString(input); value != "" {
			// Ignore whitespace tokens but still want to track position
			if tr.tokenType != tokenWhitespace {
				l.emit(token{tr.tokenType, value})
			}
			l.position += len(value)
			return
		}
	}
	l.errorf("invalid token at:\n%s", l.input)
}

func (l *lexer) run() {
	for len(l.input) > 0 {
		l.lexOneToken()
	}
	close(l.tokens)
}
