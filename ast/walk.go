package ast

import (
	"fmt"
	"github.com/riadafridishibly/spi/token"
)

type Visitor interface {
	visitBinOp(b *BinOp) int64
	visitNum(n *Num) int64
}

type nodeVisitor struct{}

func (nv *nodeVisitor) visitBinOp(b *BinOp) int64 {
	fmt.Println(b.Op)
	switch b.Op.Type {
	case token.PLUS:
		return b.Left.Accept(nv) + b.Right.Accept(nv)
	case token.MINUS:
		return b.Left.Accept(nv) - b.Right.Accept(nv)
	case token.MUL:
		return b.Left.Accept(nv) * b.Right.Accept(nv)
	case token.DIV:
		return b.Left.Accept(nv) / b.Right.Accept(nv)
	}
	panic(fmt.Sprintf("visitBinOp: Unknown Operator [%v]", b.Op))
}

func (nv *nodeVisitor) visitNum(n *Num) int64 {
	return n.Value.(int64)
}

func Walk(n Node) int64 {
	nv := nodeVisitor{}
	return n.Accept(&nv)
}
