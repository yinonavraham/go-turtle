package lexer

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

func Lex(r io.Reader) (tokens []Token, err error) {
	tokens = make([]Token, 0)
	scanner := bufio.NewScanner(r)
	line := 0
	for scanner.Scan() {
		line++
		column := 1
		text, trimmed := trimLeadingSpaces(scanner.Text())
		for len(text) > 0 {
			column += trimmed
			var t Token
			if t, err = nextToken(text, line, column); err != nil {
				return nil, err
			}
			tokens = append(tokens, t)
			column = t.Position().Column() + len(t.Value())
			text, trimmed = trimLeadingSpaces(text[len(t.Value()):])
		}
	}
	return tokens, nil
}

func nextToken(text string, line int, column int) (Token, error) {
	for _, tokenType := range tokenTypesByOrder {
		if values := tokenTypeToRegex[tokenType].FindStringSubmatch(text); len(values) >= 2 {
			value := values[1]
			return token{
				value:     value,
				tokenType: tokenType,
				position:  NewPosition(line, column),
			}, nil
		}
	}
	return nil, LexError{Text: text, Line: line, Column: column}
}

type LexError struct {
	Text         string
	Line, Column int
}

func (e LexError) Error() string {
	return fmt.Sprintf("Lexer failed at line %d, column %d: %s", e.Line, e.Column, e.Text)
}

var (
	leadingSpacesRegex = regexp.MustCompile("^\\s+")
	tokenTypesByOrder  = []TokenType{
		TokenTypeComment,
		TokenTypeSeparator,
		TokenTypeOperator,
		TokenTypeLiteral,
		TokenTypeIdentifier,
	}
	tokenTypeToRegex = map[TokenType]*regexp.Regexp{
		TokenTypeComment:    regexp.MustCompile("^(;.+)"),
		TokenTypeSeparator:  regexp.MustCompile("^([[\\]])"),
		TokenTypeOperator:   regexp.MustCompile("^([+*/-])"),
		TokenTypeLiteral:    regexp.MustCompile("^([0-9]+)(\\s|$|[+*/-])"),
		TokenTypeIdentifier: regexp.MustCompile("^(:?[a-zA-Z]+[a-zA-Z0-9_]*)(\\s|$|[+*/-])"),
	}
)

func trimLeadingSpaces(s string) (string, int) {
	leadingSpaces := leadingSpacesRegex.FindString(s)
	l := len(leadingSpaces)
	return s[l:], l
}
