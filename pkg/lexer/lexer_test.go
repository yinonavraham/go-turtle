package lexer_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-turtle/pkg/lexer"
	"strings"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		tokens    string
		wantedErr string
	}{
		{
			name:   "empty",
			text:   "",
			tokens: "",
		},
		{
			name:   "only spaces",
			text:   "  \n\t\r\n  ",
			tokens: "",
		},
		{
			name:   "nil",
			text:   "",
			tokens: "",
		},
		{
			name:   "single line, mixed tokens, ends with comment",
			text:   "bla foo 7+ 6 ;some comment",
			tokens: "1,1,bla,identifier~1,5,foo,identifier~1,9,7,literal~1,10,+,operator~1,12,6,literal~1,14,;some comment,comment",
		},
		{
			name:   "line starts with spaces, mixed tokens, ends with spaces",
			text:   "  blabla 4 \t[:foo 56  \n",
			tokens: "1,3,blabla,identifier~1,10,4,literal~1,13,[,separator~1,14,:foo,identifier~1,19,56,literal",
		},
		{
			name:   "multiline",
			text:   ";foo proc\ndo foo :a\n  fd 50\nend",
			tokens: "1,1,;foo proc,comment~2,1,do,identifier~2,4,foo,identifier~2,8,:a,identifier~3,3,fd,identifier~3,6,50,literal~4,1,end,identifier",
		},
		{
			name:   "arithmetic expression",
			text:   "  1+12 /5 - :x * :Y ",
			tokens: "1,3,1,literal~1,4,+,operator~1,5,12,literal~1,8,/,operator~1,9,5,literal~1,11,-,operator~1,13,:x,identifier~1,16,*,operator~1,18,:Y,identifier",
		},
		{
			name:      "error1",
			text:      "  9a",
			wantedErr: "Lexer failed at line 1, column 3: 9a",
		},
		{
			name:      "error2",
			text:      " :foo:bar",
			wantedErr: "Lexer failed at line 1, column 2: :foo:bar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokens, err := lexer.Lex(strings.NewReader(test.text))
			if test.wantedErr == "" {
				assert.NoError(t, err, "unexpected lexer error")
				assert.Equal(t, test.tokens, tokensAsString(tokens), "unexpected tokens result")
			} else if err == nil {
				assert.Fail(t, fmt.Sprintf("expected an error: %s", test.wantedErr))
			} else {
				assert.Equal(t, test.wantedErr, err.Error(), "unexpected error message")
			}
		})
	}
}

func tokensAsString(tokens []lexer.Token) string {
	if tokens == nil {
		return ""
	}
	tokenStrings := make([]string, len(tokens))
	for i, t := range tokens {
		tokenStrings[i] = fmt.Sprintf("%d,%d,%s,%s", t.Line(), t.Column(), t.Value(), t.Type())
	}
	return strings.Join(tokenStrings, "~")
}
