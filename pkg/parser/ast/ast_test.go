package ast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-turtle/pkg/lexer"
	"io"
	"strings"
	"testing"
)

func TestBuiltInsAndValueExpressions(t *testing.T) {
	//    0         1         2         3         4
	//    01234567890123456789012345678901234567890
	// 1: FORWARD 7 BACK 16   |         |         |
	// 2: RIGHT 90 LEFT 270   |         |         |
	// 3: HOME      |         |         |         |
	// 4: REPEAT :T [FD :X RT :Y]       |         |
	// 5: PENUP PENDOWN SHOWTURTLE HIDETURTLE     |
	// 6: CLEAN CLEARSCREEN   |         |         |
	script := Script{
		Commands: []Command{
			// 1: FORWARD 7 BACK 16   |         |         |
			Forward{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 1, Column: 0}, EndPosition: lexer.Position{Line: 1, Column: 8}},
				Steps:    IntegerExpr{Value: 7},
			},
			Back{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 1, Column: 10}, EndPosition: lexer.Position{Line: 1, Column: 16}},
				Steps:    IntegerExpr{Value: 16},
			},
			// 2: RIGHT 90 LEFT 270   |         |         |
			Right{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 2, Column: 0}, EndPosition: lexer.Position{Line: 2, Column: 7}},
				Angle:    IntegerExpr{Value: 90},
			},
			Left{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 2, Column: 0}, EndPosition: lexer.Position{Line: 2, Column: 16}},
				Angle:    IntegerExpr{Value: 270},
			},
			// 3: HOME      |         |         |         |
			Home{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 3, Column: 0}, EndPosition: lexer.Position{Line: 3, Column: 3}},
			},
			// 4: REPEAT :T [FD :X RT :Y]       |         |
			Repeat{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 4, Column: 0}, EndPosition: lexer.Position{Line: 4, Column: 22}},
				Times:    VariableExpr{Name: ":T"},
				Commands: CommandSequenceExpr{
					Commands: []Command{
						Forward{
							ItemBase: ItemBase{StartPosition: lexer.Position{Line: 4, Column: 11}, EndPosition: lexer.Position{Line: 4, Column: 15}},
							Steps:    VariableExpr{Name: ":X"},
						},
						Back{
							ItemBase: ItemBase{StartPosition: lexer.Position{Line: 4, Column: 17}, EndPosition: lexer.Position{Line: 4, Column: 21}},
							Steps:    VariableExpr{Name: ":Y"},
						},
					},
				},
			},
			// 5: PENUP PENDOWN SHOWTURTLE HIDETURTLE     |
			PenUp{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 5, Column: 0}, EndPosition: lexer.Position{Line: 5, Column: 4}},
			},
			PenDown{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 5, Column: 6}, EndPosition: lexer.Position{Line: 5, Column: 12}},
			},
			ShowTurtle{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 5, Column: 14}, EndPosition: lexer.Position{Line: 5, Column: 23}},
			},
			HideTurtle{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 5, Column: 25}, EndPosition: lexer.Position{Line: 5, Column: 34}},
			},
			// 6: CLEAN CLEARSCREEN   |         |         |
			Clean{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 6, Column: 0}, EndPosition: lexer.Position{Line: 6, Column: 4}},
			},
			ClearScreen{
				ItemBase: ItemBase{StartPosition: lexer.Position{Line: 6, Column: 6}, EndPosition: lexer.Position{Line: 6, Column: 16}},
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
	c, err := writeAst(&b, script)
	assert.NoError(t, err)
	assert.Equal(t, 125, c)
	assert.Equal(t, expected, b.String())
}

type astWriter struct {
	w     io.Writer
	count int
	err   error
}

func writeAst(w io.Writer, v Visitable) (int, error) {
	astw := &astWriter{w: w}
	v.Accept(astw)
	if astw.err != nil {
		return 0, astw.err
	}
	return astw.count, nil
}

func (w *astWriter) VisitScript(script Script) {
	w.visitCommands(script.Commands, "\n")
}

func (w *astWriter) VisitForward(forward Forward) {
	w.visitCommand(forward, forward.Steps)
}

func (w *astWriter) VisitBack(back Back) {
	w.visitCommand(back, back.Steps)
}

func (w *astWriter) VisitRight(right Right) {
	w.visitCommand(right, right.Angle)
}

func (w *astWriter) VisitLeft(left Left) {
	w.visitCommand(left, left.Angle)
}

func (w *astWriter) VisitHome(home Home) {
	w.visitCommand(home)
}

func (w *astWriter) VisitRepeat(repeat Repeat) {
	w.visitCommand(repeat, repeat.Times, repeat.Commands)
}

func (w *astWriter) VisitPenUp(penUp PenUp) {
	w.visitCommand(penUp)
}

func (w *astWriter) VisitPenDown(penDown PenDown) {
	w.visitCommand(penDown)
}

func (w *astWriter) VisitShowTurtle(showTurtle ShowTurtle) {
	w.visitCommand(showTurtle)
}

func (w *astWriter) VisitHideTurtle(hideTurtle HideTurtle) {
	w.visitCommand(hideTurtle)
}

func (w *astWriter) VisitClean(clean Clean) {
	w.visitCommand(clean)
}

func (w *astWriter) VisitClearScreen(clearScreen ClearScreen) {
	w.visitCommand(clearScreen)
}

func (w *astWriter) VisitIntegerExpr(expr IntegerExpr) {
	w.writeString(fmt.Sprintf("%d", expr.Value))
}

func (w *astWriter) VisitVariableExpr(expr VariableExpr) {
	w.writeString(expr.Name)
}

func (w *astWriter) VisitCommandSequenceExpr(expr CommandSequenceExpr) {
	_ = !w.writeString("[") || !w.visitCommands(expr.Commands, " ") || !w.writeString("]")
}

func (w *astWriter) visitCommands(commands []Command, separator string) bool {
	for i, cmd := range commands {
		if w.err != nil {
			return false
		}
		if i > 0 {
			if !w.writeString(separator) {
				return false
			}
		}
		cmd.Accept(w)
	}
	return w.err == nil
}

func (w *astWriter) visitCommand(cmd Command, args ...Expr) bool {
	if !w.writeString(cmd.CommandName()) {
		return false
	}
	for _, arg := range args {
		if !w.writeString(" ") {
			return false
		}
		arg.Accept(w)
	}
	return w.err == nil
}

func (w *astWriter) writeString(s string) bool {
	c, err := w.w.Write([]byte(s))
	if err != nil {
		w.err = err
		return false
	}
	w.count += c
	return true
}
