package lexer

import (
	"ilymar/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer { //constructor function

	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0 //set len of character to zero
	} else {
		l.ch = l.input[l.readPosition] //set ca=haracter equal to the position of the cursor in the input string
	}
	l.position = l.readPosition //new position == to read position
	l.readPosition += 1         //moving accross position
}

func newToken(tokenType token.TokenType, ch byte) token.Token { //constructing the token
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string { //obtain the substring in input from start to end
	position := l.position

	for isLetter(l.ch) {
		l.readChar() //moving poiter until condition is false
	}
	// res := fmt.Sprintf("pos: %d ... endPos: %d", position, l.position)
	// fmt.Println(res)
	return l.input[position:l.position] //slice
}

func (l *Lexer) NextToken() token.Token {

	var tok  token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		}else{
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '<':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LOE, Literal: literal}
		}else{
			tok = newToken(token.LT, l.ch)
		}

	case '>':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GOE, Literal: literal}
		}else{
			tok = newToken(token.GT, l.ch)
		}

	case '!':
		if l.peekChar() == '='{
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOTEQUAL, Literal: literal}
		}else{
			tok = newToken(token.BANG, l.ch)
		}

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '/':
		tok = newToken(token.SLASH, l.ch)

	case '.':
		tok = newToken(token.DOT, l.ch)

	case ':':
		tok = newToken(token.COLON, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}




func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input){
		return 0
	}
	return l.input[l.readPosition]  // to take a peek ahead at the ch we are at
}
