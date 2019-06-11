package ast_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-turtle/pkg/lexer"
	"github.com/yinonavraham/go-turtle/pkg/parser/ast"
	"strings"
	"testing"
)

func TestBuiltInsAndValueExpressions(t *testing.T) {
	//   0         1         2         3         4
	//    1234567890123456789012345678901234567890
	// 1: FORWARD 7 BACK 16  |         |         |
	// 2: RIGHT 90 LEFT 270  |         |         |
	// 3: HOME     |         |         |         |
	// 4: REPEAT :T [FD :X RT :Y]      |         |
	// 5: PENUP PENDOWN SHOWTURTLE HIDETURTLE    |
	// 6: CLEAN CLEARSCREEN  |         |         |
	script := ast.Script{
		Commands: []ast.Command{
			// 1: FORWARD 7 BACK 16  |         |         |
			ast.Forward{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 1, Column: 1}, EndPosition: lexer.Position{Line: 1, Column: 10}},
				Steps:    ast.IntegerExpr{Value: 7},
			},
			ast.Back{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 1, Column: 11}, EndPosition: lexer.Position{Line: 1, Column: 18}},
				Steps:    ast.IntegerExpr{Value: 16},
			},
			// 2: RIGHT 90 LEFT 270  |         |         |
			ast.Right{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 2, Column: 1}, EndPosition: lexer.Position{Line: 2, Column: 9}},
				Angle:    ast.IntegerExpr{Value: 90},
			},
			ast.Left{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 2, Column: 10}, EndPosition: lexer.Position{Line: 2, Column: 18}},
				Angle:    ast.IntegerExpr{Value: 270},
			},
			// 3: HOME     |         |         |         |
			ast.Home{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 3, Column: 1}, EndPosition: lexer.Position{Line: 3, Column: 5}},
			},
			// 4: REPEAT :T [FD :X RT :Y]      |         |
			ast.Repeat{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 4, Column: 1}, EndPosition: lexer.Position{Line: 4, Column: 24}},
				Times:    ast.VariableExpr{Name: ":T"},
				Commands: ast.CommandSequenceExpr{
					Commands: []ast.Command{
						ast.Forward{
							ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 4, Column: 12}, EndPosition: lexer.Position{Line: 4, Column: 17}},
							Steps:    ast.VariableExpr{Name: ":X"},
						},
						ast.Back{
							ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 4, Column: 18}, EndPosition: lexer.Position{Line: 4, Column: 23}},
							Steps:    ast.VariableExpr{Name: ":Y"},
						},
					},
				},
			},
			// 5: PENUP PENDOWN SHOWTURTLE HIDETURTLE    |
			ast.PenUp{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 5, Column: 1}, EndPosition: lexer.Position{Line: 5, Column: 6}},
			},
			ast.PenDown{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 5, Column: 7}, EndPosition: lexer.Position{Line: 5, Column: 14}},
			},
			ast.ShowTurtle{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 5, Column: 15}, EndPosition: lexer.Position{Line: 5, Column: 25}},
			},
			ast.HideTurtle{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 5, Column: 26}, EndPosition: lexer.Position{Line: 5, Column: 36}},
			},
			// 6: CLEAN CLEARSCREEN  |         |         |
			ast.Clean{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 6, Column: 1}, EndPosition: lexer.Position{Line: 6, Column: 6}},
			},
			ast.ClearScreen{
				ItemBase: ast.ItemBase{StartPosition: lexer.Position{Line: 6, Column: 7}, EndPosition: lexer.Position{Line: 6, Column: 18}},
			},
		},
	}
	expected := `FORWARD 7
BACK 16
RIGHT 90
LEFT 270
HOME
REPEAT :T [FORWARD :X BACK :Y]
PENUP
PENDOWN
SHOWTURTLE
HIDETURTLE
CLEAN
CLEARSCREEN`
	b := strings.Builder{}
	c, err := ast.Print(&b, script)
	assert.NoError(t, err)
	assert.Equal(t, 125, c)
	assert.Equal(t, expected, b.String())
}
