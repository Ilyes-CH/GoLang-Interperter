package lexer

import (
	"ilymar/token"
	"testing"
)


func TestLogicalOp(t *testing.T) {
	input := `= <> 
	
		if 5 < 10 {
			return false
		}
		10 == 10
		10 >= 9
		1 <= 9
		10 != 1
	`

	tests := []struct { //this is a slice of structs
		expectedType token.TokenType  // we define an anonymous struct type. This struct has two fields
		expectedLiteral string
	}{
		{token.ASSIGN,"="},
		{token.LT,"<"},
		{token.GT,">"},
	
		{token.IF,"if"},
		{token.INT,"5"},
		{token.LT,"<"},
		{token.INT,"10"},
		{token.LBRACE,"{"},
		{token.RETURN,"return"},
		{token.FALSE,"false"},
		{token.RBRACE,"}"},
	
		{token.INT,"10"},
		{token.EQUAL,"=="}, 
		{token.INT,"10"},
	
		{token.INT,"10"},
		{token.GOE,">="},
		{token.INT,"9"},
	
		{token.INT,"1"},
		{token.LOE,"<="},
		{token.INT,"9"},
		{token.INT,"10"},
		{token.NOTEQUAL,"!="}, 
		{token.INT,"1"},

	}

	l :=New(string(input))

	for i, tt :=  range tests{
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] wrong Type - expected %q . got %q",i,tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] wrong Literal - expected %q . got %q",i,tt.expectedLiteral, tok.Literal)
		}
	}
}



func TestBlock(t *testing.T) {
	input := `
	variable n = 5;
	variable m = 10;
	if
	variable add = fn(x, y){
		x + y
	}
	variable result = add(n,m);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VARIABLE, "variable"},
		{token.IDENT, "n"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.VARIABLE, "variable"},
		{token.IDENT, "m"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.VARIABLE, "variable"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},

		{token.VARIABLE, "variable"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "n"},
		{token.COMMA, ","},
		{token.IDENT, "m"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
