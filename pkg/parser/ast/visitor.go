package ast

type Visitable interface {
	Accept(visitor Visitor)
}

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
