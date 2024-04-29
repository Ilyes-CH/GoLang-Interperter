package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string //for debugging purposes
}


// Defining Tokens
const (
	ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    INT = "INT"
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
)

var Keywords = map[string]TokenType{
   "fn": "FUNCTION",
   "variable": "VARIABLE",
}

func LookupIdent(ident string) TokenType {

    if tok, ok := Keywords[ident]; ok {
        return tok
    }
    return IDENT
}