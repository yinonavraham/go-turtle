package ast

type Script interface {
	Commands() []Command
}

type script struct {
	commands []Command
}

func (s script) Commands() []Command {
	return s.commands
}
