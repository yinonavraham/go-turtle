package ast

import "github.com/yinonavraham/go-turtle/pkg/lexer"

// Script is the root LOGO AST element. It contains all the commands and other elements in a LOGO script.
type Script struct {
	Commands []Command
}

// Accept the given visitor
func (s Script) Accept(v Visitor) {
	v.VisitScript(s)
}

//###############################################

// Command is the common interface for all LOGO AST command elements
type Command interface {
	Visitable
	// CommandName returns the name of the command
	CommandName() string
}

//###############################################

// ItemBase is a base struct for all AST items. It mainly holds the location in the original script from which the item
// was parsed.
type ItemBase struct {
	StartPosition lexer.Position
	EndPosition   lexer.Position
}
