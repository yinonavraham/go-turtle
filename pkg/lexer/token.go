package lexer

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

type Token struct {
	Value     string
	TokenType TokenType
	Position  Position
}

func (t Token) String() string {
	return t.Value
}
