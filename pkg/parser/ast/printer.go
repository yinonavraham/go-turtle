package ast

import (
	"fmt"
	"io"
)

type printer struct {
	w     io.Writer
	count int
	err   error
}

// Print the provided Visitable argument to the given Writer.
//
// This will use a script printer to write the script text
func Print(w io.Writer, v Visitable) (int, error) {
	p := &printer{w: w}
	v.Accept(p)
	if p.err != nil {
		return 0, p.err
	}
	return p.count, nil
}

func (p *printer) VisitScript(script Script) {
	p.visitCommands(script.Commands, "\n")
}

func (p *printer) VisitForward(forward Forward) {
	p.visitCommand(forward, forward.Steps)
}

func (p *printer) VisitBack(back Back) {
	p.visitCommand(back, back.Steps)
}

func (p *printer) VisitRight(right Right) {
	p.visitCommand(right, right.Angle)
}

func (p *printer) VisitLeft(left Left) {
	p.visitCommand(left, left.Angle)
}

func (p *printer) VisitHome(home Home) {
	p.visitCommand(home)
}

func (p *printer) VisitRepeat(repeat Repeat) {
	p.visitCommand(repeat, repeat.Times, repeat.Commands)
}

func (p *printer) VisitPenUp(penUp PenUp) {
	p.visitCommand(penUp)
}

func (p *printer) VisitPenDown(penDown PenDown) {
	p.visitCommand(penDown)
}

func (p *printer) VisitShowTurtle(showTurtle ShowTurtle) {
	p.visitCommand(showTurtle)
}

func (p *printer) VisitHideTurtle(hideTurtle HideTurtle) {
	p.visitCommand(hideTurtle)
}

func (p *printer) VisitClean(clean Clean) {
	p.visitCommand(clean)
}

func (p *printer) VisitClearScreen(clearScreen ClearScreen) {
	p.visitCommand(clearScreen)
}

func (p *printer) VisitIntegerExpr(expr IntegerExpr) {
	p.writeString(fmt.Sprintf("%d", expr.Value))
}

func (p *printer) VisitVariableExpr(expr VariableExpr) {
	p.writeString(expr.Name)
}

func (p *printer) VisitCommandSequenceExpr(expr CommandSequenceExpr) {
	_ = !p.writeString("[") || !p.visitCommands(expr.Commands, " ") || !p.writeString("]")
}

func (p *printer) visitCommands(commands []Command, separator string) bool {
	for i, cmd := range commands {
		if p.err != nil {
			return false
		}
		if i > 0 {
			if !p.writeString(separator) {
				return false
			}
		}
		cmd.Accept(p)
	}
	return p.err == nil
}

func (p *printer) visitCommand(cmd Command, args ...Expr) bool {
	if !p.writeString(cmd.CommandName()) {
		return false
	}
	for _, arg := range args {
		if !p.writeString(" ") {
			return false
		}
		arg.Accept(p)
	}
	return p.err == nil
}

func (p *printer) writeString(s string) bool {
	c, err := p.w.Write([]byte(s))
	if err != nil {
		p.err = err
		return false
	}
	p.count += c
	return true
}
