package token

import "fmt"

type TokenType int

const (
	PROGRAM TokenType = iota
	SEMI
	DOT
	VAR
	ID
	COMMA
	COLON
	INTEGER
	REAL
	BEGIN
	END
	ASSIGN
	INTEGER_DIV
	FLOAT_DIV
	INTEGER_CONST
	REAL_CONST
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
	"PROGRAM",
	"SEMI",
	"DOT",
	"VAR",
	"ID",
	"COMMA",
	"COLON",
	"INTEGER",
	"REAL",
	"BEGIN",
	"END",
	"ASSIGN",
	"INTEGER_DIV",
	"FLOAT_DIV",
	"INTEGER_CONST",
	"REAL_CONST",
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
