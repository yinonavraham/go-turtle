package lexer

type Token interface {
	Value() string
	Type() TokenType
	Line() int
	Column() int
}

type TokenType string

const (
	//identifier: names, built-in and custom (variables, procedures, ...)
	TokenTypeIdentifier = "identifier"
	//comment: line, block (";this is a comment")
	TokenTypeComment = "comment"
	//separator (also known as punctuators): punctuation characters and paired-delimiters ("[", "]", ..)
	TokenTypeSeparator = "separator"
	//operator: symbols that operate on arguments and produce results ("+", "-", "*", "/", ...)
	TokenTypeOperator = "operator"
	//literal: numeric, logical, textual, reference literals
	TokenTypeLiteral = "literal"
)

type token struct {
	value     string
	tokenType TokenType
	line      int
	column    int
}

func (t token) Value() string {
	return t.value
}

func (t token) Type() TokenType {
	return t.tokenType
}

func (t token) Line() int {
	return t.line
}

func (t token) Column() int {
	return t.column
}

func (t token) String() string {
	return t.value
}
