package lexer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/riadafridishibly/spi/token"
)

var RESERVED_KEYWORD map[string]token.Token

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

func (lx *Lexer) peek() rune {
	peekPos := lx.pos + 1
	if peekPos >= len(lx.text) {
		return NULLCHAR
	}
	return rune(lx.text[peekPos])
}

func (lx *Lexer) skipComment() {
	for lx.currentChar != '}' {
		lx.advance()
	}
	lx.advance() // discard the `}`
}

// `id` returns identifier or reserved keywords
func (lx *Lexer) id() token.Token {
	var buf strings.Builder

	isAlphaNum := func(ch rune) bool {
		return (unicode.IsLetter(lx.currentChar) ||
			unicode.IsDigit(lx.currentChar))
	}

	for lx.currentChar != NULLCHAR && isAlphaNum(lx.currentChar) {
		buf.WriteRune(lx.currentChar)
		lx.advance()
	}

	str := buf.String()

	tok, ok := RESERVED_KEYWORD[str]
	if ok {
		return tok
	}
	return token.Token{
		Type:  token.ID,
		Value: str,
	}
}

func (lx *Lexer) skipWhiteSpaces() {
	for unicode.IsSpace(rune(lx.currentChar)) {
		lx.advance()
	}
}

// number parse and returns integer/real token
func (lx *Lexer) number() token.Token {

	var buf strings.Builder
	for lx.currentChar != NULLCHAR &&
		unicode.IsDigit(lx.currentChar) {
		buf.WriteRune(lx.currentChar)
		lx.advance()
	}

	// handle the floating point case
	if lx.currentChar == '.' {
		buf.WriteRune(lx.currentChar)
		lx.advance()

		for lx.currentChar != NULLCHAR &&
			unicode.IsDigit(lx.currentChar) {
			buf.WriteRune(lx.currentChar)
			lx.advance()
		}

		str := buf.String()
		realVal, err := strconv.ParseFloat(buf.String(), 64)

		if err != nil {
			panic(fmt.Sprintf("Couldn't parse float %v", str))
		}

		return token.Token{
			Type:  token.REAL_CONST,
			Value: realVal,
		}
	}

	// we're here that means value is integer
	str := buf.String()
	intVal, err := strconv.ParseInt(buf.String(), 10, 64)

	if err != nil {
		panic(fmt.Sprintf("Couldn't parse int %v", str))
	}
	return token.Token{
		Type:  token.INTEGER_CONST,
		Value: intVal,
	}
}

func (lx *Lexer) GetNextToken() token.Token {
	for lx.currentChar != NULLCHAR && lx.pos < len(lx.text) {

		if unicode.IsSpace(lx.currentChar) {
			lx.skipWhiteSpaces()
			continue
		}

		if lx.currentChar == '{' {
			lx.skipComment()
			continue
		}

		if unicode.IsDigit(lx.currentChar) {
			return lx.number()
		}

		if unicode.IsLetter(lx.currentChar) {
			return lx.id()
		}

		if lx.currentChar == ':' && lx.peek() == '=' {
			lx.advance()
			lx.advance()
			return token.Token{Type: token.ASSIGN, Value: ":="}
		}

		if lx.currentChar == ';' {
			lx.advance()
			return token.Token{Type: token.SEMI, Value: ";"}
		}

		if lx.currentChar == '.' {
			lx.advance()
			return token.Token{Type: token.DOT, Value: "."}
		}

		if lx.currentChar == ':' {
			lx.advance()
			return token.Token{Type: token.COLON, Value: ":"}
		}

		if lx.currentChar == ',' {
			lx.advance()
			return token.Token{Type: token.COMMA, Value: ","}
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
			return token.Token{Type: token.FLOAT_DIV, Value: '/'}
		}

		if lx.currentChar == '(' {
			lx.advance()
			return token.Token{Type: token.LPAREN, Value: '('}
		}

		if lx.currentChar == ')' {
			lx.advance()
			return token.Token{Type: token.RPAREN, Value: ')'}
		}

		panic(fmt.Sprintf("GetNextToken: current char [%v] unknown", string(lx.currentChar)))
	}
	return token.Token{Type: token.EOF, Value: nil}
}

func init() {
	RESERVED_KEYWORD = map[string]token.Token{
		"PROGRAM": token.Token{
			Type:  token.PROGRAM,
			Value: "PROGRAM",
		},
		"VAR": token.Token{
			Type:  token.VAR,
			Value: "VAR",
		},
		"DIV": token.Token{
			Type:  token.INTEGER_DIV,
			Value: "DIV",
		},
		"INTEGER": token.Token{
			Type:  token.INTEGER,
			Value: "INTEGER",
		},
		"REAL": token.Token{
			Type:  token.REAL,
			Value: "REAL",
		},
		"BEGIN": token.Token{
			Type:  token.BEGIN,
			Value: "BEGIN",
		},
		"END": token.Token{
			Type:  token.END,
			Value: "END",
		},
	}
}
