package parser

import (
	"github.com/riadafridishibly/spi/lexer"
	"github.com/riadafridishibly/spi/token"
	"log"
)

type Parser struct {
	lex          lexer.Lexer
	currentToken token.Token
}

func NewParser(lex lexer.Lexer) Parser {
	return Parser{lex, lex.GetNextToken()}
}

func (par *Parser) Factor() int64 {
	// factor: factor | LPAREN expr RPAREN
	var res int64 = 0
	switch par.currentToken.Type {
	case token.LPAREN:
		par.Consume(token.LPAREN)
		res = par.Expr()
		par.Consume(token.RPAREN)
	case token.INTEGER:
		res = par.currentToken.Value.(int64)
		par.Consume(token.INTEGER)

	default:
		panic("called with unknown token")

	}
	return res
}

func (par *Parser) Consume(toktype token.TokenType) {
	if par.currentToken.Type == toktype {
		par.currentToken = par.lex.GetNextToken()
	} else {
		log.Fatalf("Token mismatch! passed [%+v] current token [%+v]\n",
			token.StrTok[toktype], token.StrTok[par.currentToken.Type])
	}
}

func (par *Parser) Term() int64 {
	// term: factor ((MUL | DIV) factor)*
	result := par.Factor()
	for par.currentToken.Type == token.DIV ||
		par.currentToken.Type == token.MUL {
		switch par.currentToken.Type {
		case token.MUL:
			par.Consume(token.MUL)
			result = result * par.Factor()
		case token.DIV:
			par.Consume(token.DIV)
			result = result / par.Factor()
		}
	}
	return result
}

func (par *Parser) Expr() int64 {
	// expr: term ((PLUS | MINUS) term)*
	result := par.Term()
	for par.currentToken.Type == token.PLUS ||
		par.currentToken.Type == token.MINUS {
		switch par.currentToken.Type {
		case token.PLUS:
			par.Consume(token.PLUS)
			result = result + par.Factor()
		case token.MINUS:
			par.Consume(token.MINUS)
			result = result - par.Factor()
		}
	}
	return result
}
