package parser

import (
	"log"

	"github.com/riadafridishibly/spi/ast"
	"github.com/riadafridishibly/spi/lexer"
	"github.com/riadafridishibly/spi/token"
)

type Parser struct {
	lex          lexer.Lexer
	currentToken token.Token
}

func NewParser(lex lexer.Lexer) Parser {
	return Parser{lex, lex.GetNextToken()}
}

func (par *Parser) Factor() ast.Node {
	// factor: (PLUS | MINUS) factor | INTEGER | LPAREN expr RPAREN
	var node ast.Node

	tok := par.currentToken

	switch par.currentToken.Type {
	case token.PLUS:
		par.Consume(token.PLUS)
		node = &ast.UnaryOp{
			Token: tok,
			Op:    tok,
			Expr:  par.Factor(),
		}
	case token.MINUS:
		par.Consume(token.MINUS)
		node = &ast.UnaryOp{
			Token: tok,
			Op:    tok,
			Expr:  par.Factor(),
		}
	case token.LPAREN:
		par.Consume(token.LPAREN)
		node = par.Expr()
		par.Consume(token.RPAREN)
	case token.INTEGER:
		node = &ast.Num{
			Token: par.currentToken,
			Value: par.currentToken.Value.(int64),
		}
		par.Consume(token.INTEGER)

	default:
		panic("called with unknown token")

	}
	return node
}

func (par *Parser) Consume(toktype token.TokenType) {
	if par.currentToken.Type == toktype {
		par.currentToken = par.lex.GetNextToken()
	} else {
		log.Fatalf("Token mismatch! passed [%+v] current token [%+v]\n",
			token.StrTok[toktype], token.StrTok[par.currentToken.Type])
	}
}

func (par *Parser) Term() ast.Node {
	// term: factor ((MUL | DIV) factor)*
	result := par.Factor()
	for par.currentToken.Type == token.DIV ||
		par.currentToken.Type == token.MUL {
		tok := par.currentToken
		switch par.currentToken.Type {
		case token.MUL:
			par.Consume(token.MUL)
		case token.DIV:
			par.Consume(token.DIV)
		}
		result = &ast.BinOp{
			Left:  result,
			Op:    tok,
			Token: tok,
			Right: par.Factor(),
		}

	}
	return result
}

func (par *Parser) Expr() ast.Node {
	// expr: term ((PLUS | MINUS) term)*
	result := par.Term()
	for par.currentToken.Type == token.PLUS ||
		par.currentToken.Type == token.MINUS {
		tok := par.currentToken
		switch par.currentToken.Type {
		case token.PLUS:
			par.Consume(token.PLUS)
		case token.MINUS:
			par.Consume(token.MINUS)
		}
		result = &ast.BinOp{
			Left:  result,
			Op:    tok,
			Token: tok,
			Right: par.Term(),
		}
	}
	return result
}

func (par *Parser) Parse() ast.Node {
	return par.Expr()
}
