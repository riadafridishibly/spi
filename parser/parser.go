package parser

import (
	"fmt"
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

func (p *Parser) Program() ast.Node {
	// program: PROGRAM variable SEMI block DOT
	p.Consume(token.PROGRAM)

	varNode := p.Variable()
	// Variable method returns *ast.Var boxed in ast.Node interface
	// so it's guaranteed the following type conversion will always work
	progName := varNode.(*ast.Var).Token.Value.(string)
	p.Consume(token.SEMI)

	blockNode := p.Block()

	programNode := &ast.Program{
		Name:  progName,
		Block: blockNode,
	}
	p.Consume(token.DOT)
	return programNode
}

func (p *Parser) Block() ast.Node {
	// block: declarations compound_statement
	declarationNodes := p.Declarations()
	compoundStmtNode := p.CompoundStatement()
	return &ast.Block{
		Declarations:      declarationNodes,
		CompoundStatement: compoundStmtNode,
	}
}

func (p *Parser) Declarations() []ast.Node {
	// declarations: VAR (variable_declaration SEMI)+ | empty

	var declarations []ast.Node

	if p.currentToken.Type == token.VAR {
		p.Consume(token.VAR)

		for p.currentToken.Type == token.ID {
			varDecls := p.VariableDeclarations()
			declarations = append(declarations, varDecls...)
			p.Consume(token.SEMI)
		}
	}

	return declarations
}

func (p *Parser) VariableDeclarations() []ast.Node {
	// variable_declaration: ID (COMMA ID)* COLON type_spec
	var varNodes []*ast.Var
	varNodes = append(varNodes, &ast.Var{Token: p.currentToken})
	p.Consume(token.ID)

	// a, b, c : INTEGER;
	//  ^  <- currentToken is here now
	for p.currentToken.Type == token.COMMA {
		p.Consume(token.COMMA)
		varNodes = append(varNodes, &ast.Var{Token: p.currentToken})
		p.Consume(token.ID)
	}

	// a, b, c : INTEGER;
	//         ^  <- currentToken is here now
	p.Consume(token.COLON)
	typeNode := p.TypeSpec()

	var varDeclNodes []ast.Node

	for _, nodePtr := range varNodes {
		varDeclNodes = append(varDeclNodes, &ast.VarDecl{
			VarNode:  nodePtr,
			TypeNode: typeNode,
		})
	}

	return varDeclNodes
}

func (p *Parser) TypeSpec() ast.Node {
	// type_spec: INTEGER | REAL
	currTok := p.currentToken

	if p.currentToken.Type == token.INTEGER {
		p.Consume(token.INTEGER)
	} else if p.currentToken.Type == token.REAL {
		p.Consume(token.REAL)
	} else {
		panic(fmt.Sprintf("TypeSpec: Unknown type %v", currTok))
	}

	return &ast.Type{Token: currTok}
}

func (p *Parser) CompoundStatement() ast.Node {
	// compound_statement: BEGIN statement_list END
	p.Consume(token.BEGIN)
	nodes := p.StatementList()
	p.Consume(token.END)

	root := &ast.Compound{}

	for _, node := range nodes {
		root.Children = append(root.Children, node)
	}

	return root
}

func (p *Parser) StatementList() []ast.Node {
	node := p.Statement()
	var result []ast.Node
	result = append(result, node)

	for p.currentToken.Type == token.SEMI {
		p.Consume(token.SEMI)
		result = append(result, p.Statement())
	}
	return result
}

func (p *Parser) Statement() ast.Node {
	// statement: compound_statement
	//          | assignment_statement
	//          | empty

	var node ast.Node

	if p.currentToken.Type == token.BEGIN {
		node = p.CompoundStatement()
	} else if p.currentToken.Type == token.ID {
		node = p.AssignmentStatement()
	} else {
		node = p.Empty()
	}

	return node
}

func (p *Parser) AssignmentStatement() ast.Node {
	// assignment_statement: variable ASSIGN expr
	left := p.Variable()
	tok := p.currentToken

	p.Consume(token.ASSIGN)

	right := p.Expr()

	assignNode := &ast.Assign{
		Left:  left,
		Op:    tok,
		Right: right,
	}
	return assignNode
}

func (p *Parser) Variable() ast.Node {
	// variable: ID
	node := &ast.Var{
		Token: p.currentToken,
	}
	p.Consume(token.ID)
	return node
}

func (p *Parser) Empty() ast.Node {
	// empty:
	return &ast.NoOp{}
}

func (par *Parser) Factor() ast.Node {
	// factor: PLUS factor
	//       | MINUS factor
	//       | INTEGER_CONST
	//       | REAL_CONST
	//       | LPAREN expr RPAREN
	//       | variable

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
	case token.INTEGER_CONST:
		par.Consume(token.INTEGER_CONST)
		node = &ast.Num{
			Token: tok,
			Value: tok.Value,
		}
	case token.REAL_CONST:
		par.Consume(token.REAL_CONST)
		node = &ast.Num{
			Token: tok,
			Value: tok.Value,
		}
	case token.LPAREN:
		par.Consume(token.LPAREN)
		node = par.Expr()
		par.Consume(token.RPAREN)
	case token.ID:
		node = par.Variable()

	default:
		panic("called with unknown token")

	}
	if node == nil {
		panic("Factor returns nil ast.Node")
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
	// term: factor ((MUL | INTEGER_DIV | FLOAT_DIV) factor)*

	result := par.Factor()
	for par.currentToken.Type == token.INTEGER_DIV ||
		par.currentToken.Type == token.FLOAT_DIV ||
		par.currentToken.Type == token.MUL {

		tok := par.currentToken

		switch par.currentToken.Type {
		case token.MUL:
			par.Consume(token.MUL)
		case token.INTEGER_DIV:
			par.Consume(token.INTEGER_DIV)
		case token.FLOAT_DIV:
			par.Consume(token.FLOAT_DIV)
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

func (p *Parser) Parse() ast.Node {
	node := p.Program()
	if p.currentToken.Type != token.EOF {
		panic(fmt.Sprintf("Couldn't parse the full program, last token [%v]", p.currentToken))
	}
	return node
}
