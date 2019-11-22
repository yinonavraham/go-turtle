package ast

//###############################################

// Expr interface is a common interface for all AST expressions
type Expr interface {
	Visitable
}

//###############################################

// CommandSequenceExpr is an AST expression for a sequence of commands
type CommandSequenceExpr struct {
	ItemBase
	Commands []Command
}

// Accept the given visitor
func (e CommandSequenceExpr) Accept(v Visitor) {
	v.VisitCommandSequenceExpr(e)
}

//###############################################

// ValueExpr is a common AST expression interface for a value
type ValueExpr interface {
	Expr
}

//###############################################

// IntegerExpr is an AST expression for an integer value
type IntegerExpr struct {
	ItemBase
	Value int
}

// Accept the given visitor
func (e IntegerExpr) Accept(v Visitor) {
	v.VisitIntegerExpr(e)
}

//###############################################

// VariableExpr is an AST expression for a named variable value (placeholder)
type VariableExpr struct {
	ItemBase
	Name string
}

// Accept the given visitor
func (e VariableExpr) Accept(v Visitor) {
	v.VisitVariableExpr(e)
}
