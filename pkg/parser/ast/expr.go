package ast

//###############################################

type Expr interface {
	Visitable
}

//###############################################

type CommandSequenceExpr struct {
	ItemBase
	Commands []Command
}

func (e CommandSequenceExpr) Accept(v Visitor) {
	v.VisitCommandSequenceExpr(e)
}

//###############################################

type ValueExpr interface {
	Expr
}

//###############################################

type IntegerExpr struct {
	ItemBase
	Value int
}

func (e IntegerExpr) Accept(v Visitor) {
	v.VisitIntegerExpr(e)
}

//###############################################

type VariableExpr struct {
	ItemBase
	Name string
}

func (e VariableExpr) Accept(v Visitor) {
	v.VisitVariableExpr(e)
}
