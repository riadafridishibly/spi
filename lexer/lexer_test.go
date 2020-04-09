package lexer

import (
	"github.com/riadafridishibly/spi/token"
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

		i := lx.integer()

		if i != tt.want {
			t.Errorf("%+v, %v != %v\n", tt, i, tt.want)
		}
	}
}

func TestGetNextToken(t *testing.T) {
	text := "12 + 3 / (8 - 5) * 2"
	tok := []token.Token{
		{Type: token.INTEGER, Value: int64(12)},
		{Type: token.PLUS, Value: '+'},
		{Type: token.INTEGER, Value: int64(3)},
		{Type: token.DIV, Value: '/'},
		{Type: token.LPAREN, Value: '('},
		{Type: token.INTEGER, Value: int64(8)},
		{Type: token.MINUS, Value: '-'},
		{Type: token.INTEGER, Value: int64(5)},
		{Type: token.RPAREN, Value: ')'},
		{Type: token.MUL, Value: '*'},
		{Type: token.INTEGER, Value: int64(2)},
	}

	tokenEqual := func(got, want token.Token) bool {

		if want.Type != got.Type {
			return false
		}

		gv := got.Value
		wv := want.Value

		switch got.Type {
		case token.INTEGER:
			return gv.(int64) == wv.(int64)
		case token.PLUS, token.MINUS, token.DIV,
			token.MUL, token.LPAREN, token.RPAREN:
			return gv.(int32) == wv.(int32)
		default:
			return false
		}
	}

	lx := NewLexer(text, 0)

	index := 0
	for lx.currentChar != NULLCHAR {
		curr := lx.GetNextToken()
		if !tokenEqual(curr, tok[index]) {
			t.Errorf("Expected %+v but got%+v\n", tok[index], curr)
		}
		index++
	}
}
