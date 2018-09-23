package glambda

import (
	"fmt"
	"regexp"
	"strings"
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
)

type token struct {
	tokenType tokenType
	value     string
}

func (t token) String() string {
	return t.value
}

type lexer struct {
	input  string
	tokens chan token
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
}

func (l *lexer) lexOneToken() {
	for _, tr := range tokenRegexes {
		if value := tr.regex.FindString(l.input); value != "" {
			l.emit(token{tr.tokenType, value})
			l.input = l.input[len(value):]
			return
		}
	}
	l.errorf("invalid token at:\n%s", l.input)
}

func (l *lexer) run() {
	for len(l.input) > 0 {
		l.input = strings.Trim(l.input, " \n\t")
		l.lexOneToken()
	}
	close(l.tokens)
}
