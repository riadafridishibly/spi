package lexer

import (
	"github.com/riadafridishibly/spi/token"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const NULLCHAR rune = -1

type Lexer struct {
	text        string
	pos         int
	currentChar rune
}

func NewLexer(text string, pos int) Lexer {
	var curr rune = NULLCHAR
	if pos < len(text) {
		curr = rune(text[pos])
	}
	return Lexer{text: text, pos: pos, currentChar: curr}
}

func (lx *Lexer) advance() {
	lx.pos++
	if lx.pos >= len(lx.text) {
		lx.currentChar = NULLCHAR
		// panic("Advance: Lexer.Pos overflows")
	} else {
		lx.currentChar = rune(lx.text[lx.pos])
	}
}

func (lx *Lexer) skipWhiteSpaces() {
	for unicode.IsSpace(rune(lx.currentChar)) {
		lx.advance()
	}
}

func (lx *Lexer) integer() int64 {
	var buf strings.Builder

	for lx.currentChar != NULLCHAR && unicode.IsDigit(lx.currentChar) {
		buf.WriteRune(lx.currentChar)
		lx.advance()
	}

	i, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		log.Printf("%+v\n", lx)
		panic(err)
	}
	return i
}

func (lx *Lexer) GetNextToken() token.Token {
	for lx.currentChar != NULLCHAR && lx.pos < len(lx.text) {
		if unicode.IsSpace(lx.currentChar) {
			lx.skipWhiteSpaces()
			continue
		}

		if unicode.IsDigit(lx.currentChar) {
			ival := lx.integer()
			return token.Token{Type: token.INTEGER, Value: ival}
		}

		if lx.currentChar == '+' {
			lx.advance()
			return token.Token{Type: token.PLUS, Value: '+'}
		}

		if lx.currentChar == '-' {
			lx.advance()
			return token.Token{Type: token.MINUS, Value: '-'}
		}

		if lx.currentChar == '*' {
			lx.advance()
			return token.Token{Type: token.MUL, Value: '*'}
		}

		if lx.currentChar == '/' {
			lx.advance()
			return token.Token{Type: token.DIV, Value: '/'}
		}

		if lx.currentChar == '(' {
			lx.advance()
			return token.Token{Type: token.LPAREN, Value: '('}
		}

		if lx.currentChar == ')' {
			lx.advance()
			return token.Token{Type: token.RPAREN, Value: ')'}
		}

		log.Fatalf("GetNextToken: current char [%v] unknown", string(lx.currentChar))
	}
	return token.Token{Type: token.EOF, Value: nil}
}
