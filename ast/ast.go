package ast

import "github.com/riadafridishibly/spi/token"

type Node interface {
	Accept(v Visitor) int64
}

type BinOp struct {
	Left  Node
	Right Node
	Token token.Token
	Op    token.Token // Do I need these two?
}

type UnaryOp struct {
	Token token.Token
	Op    token.Token
	Expr  Node
}

type Num struct {
	Token token.Token
	Value interface{} // This is basically `int64`
}

func (b *BinOp) Accept(v Visitor) int64 {
	return v.visitBinOp(b)
}

func (n *Num) Accept(v Visitor) int64 {
	return v.visitNum(n)
}

func (u *UnaryOp) Accept(v Visitor) int64 {
	return v.visitUnaryOp(u)
}
