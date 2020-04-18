package ast

import "github.com/riadafridishibly/spi/token"

type ValueType interface{}

type Node interface {
	Accept(v Visitor) ValueType
}

type Program struct {
	Name  string
	Block Node
}

type Block struct {
	Declarations      []Node
	CompoundStatement Node
}

type VarDecl struct {
	VarNode  Node
	TypeNode Node
}

type Type struct {
	Token token.Token
}

type Compound struct {
	Children []Node
}

type Assign struct {
	Left  Node
	Op    token.Token
	Right Node
}

type Var struct {
	Token token.Token
}

type NoOp struct{}

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
	Value interface{}
}

func (prog *Program) Accept(v Visitor) ValueType {
	v.visitProgram(prog)
	return nil
}

func (block *Block) Accept(v Visitor) ValueType {
	return v.visitBlock(block)
}

func (vr *Var) Accept(v Visitor) ValueType {
	return v.visitVar(vr)
}

func (nop *NoOp) Accept(v Visitor) ValueType {
	return v.visitNoOp(nop)
}
func (b *BinOp) Accept(v Visitor) ValueType {
	return v.visitBinOp(b)
}

func (n *Num) Accept(v Visitor) ValueType {
	return v.visitNum(n)
}

func (u *UnaryOp) Accept(v Visitor) ValueType {
	return v.visitUnaryOp(u)
}

func (varDcl *VarDecl) Accept(v Visitor) ValueType {
	return v.visitVarDecl(varDcl)
}

func (typ *Type) Accept(v Visitor) ValueType {
	return v.visitType(typ)
}

func (compound *Compound) Accept(v Visitor) ValueType {
	return v.visitCompound(compound)
}

func (assign *Assign) Accept(v Visitor) ValueType {
	return v.visitAssign(assign)
}
