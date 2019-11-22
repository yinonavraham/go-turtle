package lexer

// NewPosition creates a new Position instance with the provided coordinates
func NewPosition(line int, column int) Position {
	return Position{
		Line:   line,
		Column: column,
	}
}

// Position struct holds a position in a text
type Position struct {
	Line   int
	Column int
}
