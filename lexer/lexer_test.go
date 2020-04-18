package lexer

import (
	// "github.com/riadafridishibly/spi/token"
	"fmt"
	"testing"
)

func TestSkipWhiteSpaces(t *testing.T) {
	// test if it sets the Lexer.CurrentChar as expected!
	// if the position goes beyond the string length, then
	// Lexer.CurrentChar is set to Zero (0)
	testcases := []struct {
		text string
		want rune
	}{
		{"     ", NULLCHAR},
		{"  2  ", '2'},
		{"12345", '1'},
		{"\t\t", NULLCHAR},
	}

	for _, tt := range testcases {
		lx := NewLexer(tt.text, 0)
		lx.skipWhiteSpaces()

		if lx.currentChar != tt.want {
			t.Errorf("%+v, %v != %v\n", tt, lx.currentChar, tt.want)
		}
	}
}

func TestInteger(t *testing.T) {
	testcases := []struct {
		text string
		want int64
	}{
		{"0", 0},
		{"12345", 12345},
		{"9223372036854775807", 9223372036854775807},
	}

	for _, tt := range testcases {
		lx := NewLexer(tt.text, 0)

		i := lx.number()

		if i.Value.(int64) != tt.want {
			t.Errorf("%+v, %v != %v\n", tt, i, tt.want)
		}
	}
}

func TestGetNextToken(t *testing.T) {
	text := `
		PROGRAM Part10;
		VAR
		number     : INTEGER;
		a, b, c, x : INTEGER;
		y          : REAL;

		BEGIN {Part10}
		BEGIN
			number := 2;
			a := number;
			b := 10 * a + 10 * number DIV 4;
			c := a - - b
		END;
		x := 11;
		y := 20 / 7 + 3.14;
		{ writeln('a = ', a); }
		{ writeln('b = ', b); }
		{ writeln('c = ', c); }
		{ writeln('number = ', number); }
		{ writeln('x = ', x); }
		{ writeln('y = ', y); }
		END.  {Part10}
	`
	lx := NewLexer(text, 0)

	for lx.currentChar != NULLCHAR {
		tok := lx.GetNextToken()
		fmt.Println(tok)
		t.Log(tok)
	}
}
