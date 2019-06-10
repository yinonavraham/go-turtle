package lexer

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
