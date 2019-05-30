package lexer

import "fmt"

func NewPosition(line int, column int) Position {
	return position{
		line:   line,
		column: column,
	}
}

type Position interface {
	Line() int
	Column() int
}

type position struct {
	line   int
	column int
}

func (p position) Line() int {
	return p.line
}

func (p position) Column() int {
	return p.column
}

func (p position) String() string {
	return fmt.Sprintf("line %d, column %d", p.line, p.column)
}
