package token

import "fmt"

type TokenType int

const (
	INTEGER TokenType = iota
	PLUS
	MINUS
	DIV
	MUL
	LPAREN
	RPAREN
	EOF
)

type Token struct {
	Type  TokenType
	Value interface{}
}

var StrTok = [...]string{
	"INTEGER",
	"PLUS",
	"MINUS",
	"DIV",
	"MUL",
	"LPAREN",
	"RPAREN",
	"EOF",
}

func (t TokenType) String() string {
	return fmt.Sprintf("%v", StrTok[t])
}
