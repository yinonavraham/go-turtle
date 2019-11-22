package ast

// Visitable is a common interface for all items which the Visitor can visit
type Visitable interface {
	Accept(visitor Visitor)
}

// Visitor is an interface for defining all the possible item types a visitor can visit
type Visitor interface {
	VisitScript(script Script)
	VisitForward(forward Forward)
	VisitBack(back Back)
	VisitRight(right Right)
	VisitLeft(left Left)
	VisitHome(home Home)
	VisitRepeat(repeat Repeat)
	VisitPenUp(penUp PenUp)
	VisitPenDown(penDown PenDown)
	VisitShowTurtle(showTurtle ShowTurtle)
	VisitHideTurtle(hideTurtle HideTurtle)
	VisitClean(clean Clean)
	VisitClearScreen(clearScreen ClearScreen)
	VisitIntegerExpr(expr IntegerExpr)
	VisitVariableExpr(expr VariableExpr)
	VisitCommandSequenceExpr(expr CommandSequenceExpr)
}
