package lexer

// TokenType represents the different types of tokens the lexer identifies
type TokenType string

const (
	// TokenTypeIdentifier : names, built-in and custom (variables, procedures, ...)
	TokenTypeIdentifier = "identifier"
	// TokenTypeComment : line, block (";this is a comment")
	TokenTypeComment = "comment"
	// TokenTypeSeparator (also known as punctuators): punctuation characters and paired-delimiters ("[", "]", ..)
	TokenTypeSeparator = "separator"
	// TokenTypeOperator : symbols that operate on arguments and produce results ("+", "-", "*", "/", ...)
	TokenTypeOperator = "operator"
	// TokenTypeLiteral : numeric, logical, textual, reference literals
	TokenTypeLiteral = "literal"
)

// Token is a single token identified in the text by the lexer
type Token struct {
	// Value is the token string value
	Value string
	// Type is the type of the token
	Type TokenType
	// Position is the start position of the token in the text
	Position Position
}

func (t Token) String() string {
	return t.Value
}
