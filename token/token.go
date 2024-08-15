package token

type TokenType string //alias for the existing type string.

type Token struct { 
	Type   TokenType
	Literal string //for debugging purposes
}

// Defining Tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"
	// Identifiers and literals
	IDENT = "IDENT" // add more if needed

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	DOT      = "."
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	BANG     = "!"
	LOE      = "<="
	GOE      = ">="
	EQUAL    = "=="
	NOTEQUAL = "!="
	OR       = "||"
	AND      = "&&"
	// Add more operators if needed

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
							
	// Keywords
	FUNCTION = "FUNCTION"
	VARIABLE = "VARIABLE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var Keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"variable": VARIABLE,
	"true":     TRUE,
	"false":    FALSE,
	"return":   RETURN,
	"if":       IF,
	"else":     ELSE,
}

func LookupIdent(ident string) TokenType {

	if tok, ok := Keywords[ident]; ok { //ok is a booleaen var to determine if the look up was a successful
		return tok
	}
	return IDENT
}
