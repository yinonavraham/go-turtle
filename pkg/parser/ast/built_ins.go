package ast

const (
	CommandForward     string = "FORWARD"
	CommandBack        string = "BACK"
	CommandRight       string = "RIGHT"
	CommandLeft        string = "LEFT"
	CommandHome        string = "HOME"
	CommandRepeat      string = "REPEAT"
	CommandPenUp       string = "PENUP"
	CommandPenDown     string = "PENDOWN"
	CommandShowTurtle  string = "SHOWTURTLE"
	CommandHideTurtle  string = "HIDETURTLE"
	CommandClean       string = "CLEAN"
	CommandClearScreen string = "CLEARSCREEN"
)

//###############################################

type Forward struct {
	ItemBase
	Steps ValueExpr
}

func (c Forward) CommandName() string {
	return CommandForward
}

func (c Forward) Accept(v Visitor) {
	v.VisitForward(c)
}

//###############################################

type Back struct {
	ItemBase
	Steps ValueExpr
}

func (c Back) CommandName() string {
	return CommandBack
}

func (c Back) Accept(v Visitor) {
	v.VisitBack(c)
}

//###############################################

type Right struct {
	ItemBase
	Angle ValueExpr
}

func (c Right) CommandName() string {
	return CommandRight
}

func (c Right) Accept(v Visitor) {
	v.VisitRight(c)
}

//###############################################

type Left struct {
	ItemBase
	Angle ValueExpr
}

func (c Left) CommandName() string {
	return CommandLeft
}

func (c Left) Accept(v Visitor) {
	v.VisitLeft(c)
}

//###############################################

type Home struct {
	ItemBase
}

func (c Home) CommandName() string {
	return CommandHome
}

func (c Home) Accept(v Visitor) {
	v.VisitHome(c)
}

//###############################################

type Repeat struct {
	ItemBase
	Times    ValueExpr
	Commands CommandSequenceExpr
}

func (c Repeat) CommandName() string {
	return CommandRepeat
}

func (c Repeat) Accept(v Visitor) {
	v.VisitRepeat(c)
}

//###############################################

type PenUp struct {
	ItemBase
}

func (c PenUp) CommandName() string {
	return CommandPenUp
}

func (c PenUp) Accept(v Visitor) {
	v.VisitPenUp(c)
}

//###############################################

type PenDown struct {
	ItemBase
}

func (c PenDown) CommandName() string {
	return CommandPenDown
}

func (c PenDown) Accept(v Visitor) {
	v.VisitPenDown(c)
}

//###############################################

type ShowTurtle struct {
	ItemBase
}

func (c ShowTurtle) CommandName() string {
	return CommandShowTurtle
}

func (c ShowTurtle) Accept(v Visitor) {
	v.VisitShowTurtle(c)
}

//###############################################

type HideTurtle struct {
	ItemBase
}

func (c HideTurtle) CommandName() string {
	return CommandHideTurtle
}

func (c HideTurtle) Accept(v Visitor) {
	v.VisitHideTurtle(c)
}

//###############################################

type Clean struct {
	ItemBase
}

func (c Clean) CommandName() string {
	return CommandClean
}

func (c Clean) Accept(v Visitor) {
	v.VisitClean(c)
}

//###############################################

type ClearScreen struct {
	ItemBase
}

func (c ClearScreen) CommandName() string {
	return CommandClearScreen
}

func (c ClearScreen) Accept(v Visitor) {
	v.VisitClearScreen(c)
}
