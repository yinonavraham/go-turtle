package parser_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-turtle/pkg/parser"
	"github.com/yinonavraham/go-turtle/pkg/parser/ast"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	script, err := parser.Parse(strings.NewReader("FORWARD 7 BACK 12 FD 32 BK 99"))
	assert.NoError(t, err, "unexpected parse error")

	out := strings.Builder{}
	_, err = ast.Print(&out, script)
	assert.NoError(t, err, "unexpected print error")
	expected := `FORWARD 7
BACK 12
FORWARD 32
BACK 99`
	assert.Equal(t, expected, out.String())
}
