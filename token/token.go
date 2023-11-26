package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Illegal tokens + end of file token
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// Operators
	ASSIGN   TokenType = "ASSIGN"
	PLUS     TokenType = "+"
	BANG     TokenType = "!"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// Delimiters
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	LBRACE TokenType = "LBRACE"
	RBRACE TokenType = "RBRACE"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	RETURN   TokenType = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
