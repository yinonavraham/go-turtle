package lexer

import "fmt"

func NewPosition(line int, column int) Position {
	return Position{
		Line:   line,
		Column: column,
	}
}

type Position struct {
	Line   int
	Column int
}

func (p Position) String() string {
	return fmt.Sprintf("line %d, column %d", p.Line, p.Column)
}
