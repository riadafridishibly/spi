package ast

import (
	"fmt"
	"github.com/riadafridishibly/spi/token"
)

var GLOBAL_SCOPE = map[string]ValueType{}

type Visitor interface {
	visitProgram(p *Program)
	visitBlock(b *Block) ValueType
	visitVarDecl(vardecl *VarDecl) ValueType
	visitType(typ *Type) ValueType
	visitCompound(compound *Compound) ValueType
	visitAssign(assign *Assign) ValueType
	visitVar(vr *Var) ValueType
	visitNoOp(np *NoOp) ValueType
	visitBinOp(b *BinOp) ValueType
	visitNum(n *Num) ValueType
	visitUnaryOp(u *UnaryOp) ValueType
}

type nodeVisitor struct{}

func (nv *nodeVisitor) visitProgram(prog *Program) {
	prog.Block.Accept(nv)
}

func (nv *nodeVisitor) visitBlock(b *Block) ValueType {
	for _, decl := range b.Declarations {
		decl.Accept(nv)
	}
	b.CompoundStatement.Accept(nv)

	return nil
}

func (nv *nodeVisitor) visitVarDecl(vardecl *VarDecl) ValueType {
	return nil
}

func (nv *nodeVisitor) visitType(typ *Type) ValueType {
	return nil
}

func (nv *nodeVisitor) visitCompound(compound *Compound) ValueType {
	for _, child := range compound.Children {
		child.Accept(nv)
	}
	return nil
}

func (nv *nodeVisitor) visitAssign(assign *Assign) ValueType {
	varName := assign.Left.(*Var).Token.Value.(string)
	GLOBAL_SCOPE[varName] = assign.Right.Accept(nv)
	return nil
}

func (nv *nodeVisitor) visitVar(vr *Var) ValueType {
	varName := vr.Token.Value.(string)
	val, ok := GLOBAL_SCOPE[varName]

	if !ok {
		panic("Name Error")
	} else {
		return val
	}
}
func (nv *nodeVisitor) visitNoOp(np *NoOp) ValueType {
	return nil
}

func (nv *nodeVisitor) visitBinOp(b *BinOp) ValueType {
	// tok := b.Token
	switch b.Op.Type {
	case token.PLUS:
		return Plus(b.Left.Accept(nv), b.Right.Accept(nv))
	case token.MINUS:
		return Minus(b.Left.Accept(nv), b.Right.Accept(nv))
	case token.MUL:
		return Multiply(b.Left.Accept(nv), b.Right.Accept(nv))
	case token.INTEGER_DIV:
		return IntegerDiv(b.Left.Accept(nv), b.Right.Accept(nv))
	case token.FLOAT_DIV:
		return FloatDiv(b.Left.Accept(nv), b.Right.Accept(nv))
	}
	panic(fmt.Sprintf("visitBinOp: Unknown Operator [%v]", b.Op))
}
func (nv *nodeVisitor) visitNum(n *Num) ValueType {
	return n.Value
}
func (nv *nodeVisitor) visitUnaryOp(u *UnaryOp) ValueType {
	val := u.Expr.Accept(nv)

	switch v := val.(type) {
	case int64:
		if u.Op.Type == token.PLUS {
			return +v
		} else {
			return -v
		}
	case float64:
		if u.Op.Type == token.PLUS {
			return +v
		} else {
			return -v
		}
	default:
		panic(fmt.Sprintf("visitUnaryOp: Value other than int64/float64 [%v]", val))
	}
}

func Walk(n Node) ValueType {
	nv := nodeVisitor{}
	res := n.Accept(&nv)

	fmt.Println(GLOBAL_SCOPE)
	return res
}
