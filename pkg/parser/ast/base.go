package ast

import "github.com/yinonavraham/go-turtle/pkg/lexer"

type Script struct {
	Commands []Command
}

func (s Script) Accept(v Visitor) {
	v.VisitScript(s)
}

//###############################################

type Command interface {
	Visitable
	CommandName() string
}

//###############################################

type ItemBase struct {
	StartPosition lexer.Position
	EndPosition   lexer.Position
}
