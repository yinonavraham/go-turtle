package ast

const (
	// CommandForward - command long name
	CommandForward string = "FORWARD"
	// CommandBack - command long name
	CommandBack string = "BACK"
	// CommandRight - command long name
	CommandRight string = "RIGHT"
	// CommandLeft - command long name
	CommandLeft string = "LEFT"
	// CommandHome - command long name
	CommandHome string = "HOME"
	// CommandRepeat - command long name
	CommandRepeat string = "REPEAT"
	// CommandPenUp - command long name
	CommandPenUp string = "PENUP"
	// CommandPenDown - command long name
	CommandPenDown string = "PENDOWN"
	// CommandShowTurtle - command long name
	CommandShowTurtle string = "SHOWTURTLE"
	// CommandHideTurtle - command long name
	CommandHideTurtle string = "HIDETURTLE"
	// CommandClean - command long name
	CommandClean string = "CLEAN"
	// CommandClearScreen - command long name
	CommandClearScreen string = "CLEARSCREEN"
)

//###############################################

// Forward is an AST item that represents the FORWARD command
type Forward struct {
	ItemBase
	Steps ValueExpr
}

// CommandName returns the name of the command
func (c Forward) CommandName() string {
	return CommandForward
}

// Accept the given visitor
func (c Forward) Accept(v Visitor) {
	v.VisitForward(c)
}

//###############################################

// Back is an AST item that represents the BACK command
type Back struct {
	ItemBase
	Steps ValueExpr
}

// CommandName returns the name of the command
func (c Back) CommandName() string {
	return CommandBack
}

// Accept the given visitor
func (c Back) Accept(v Visitor) {
	v.VisitBack(c)
}

//###############################################

// Right is an AST item that represents the RIGHT command
type Right struct {
	ItemBase
	Angle ValueExpr
}

// CommandName returns the name of the command
func (c Right) CommandName() string {
	return CommandRight
}

// Accept the given visitor
func (c Right) Accept(v Visitor) {
	v.VisitRight(c)
}

//###############################################

// Left is an AST item that represents the LEFT command
type Left struct {
	ItemBase
	Angle ValueExpr
}

// CommandName returns the name of the command
func (c Left) CommandName() string {
	return CommandLeft
}

// Accept the given visitor
func (c Left) Accept(v Visitor) {
	v.VisitLeft(c)
}

//###############################################

// Home is an AST item that represents the HOME command
type Home struct {
	ItemBase
}

// CommandName returns the name of the command
func (c Home) CommandName() string {
	return CommandHome
}

// Accept the given visitor
func (c Home) Accept(v Visitor) {
	v.VisitHome(c)
}

//###############################################

// Repeat is an AST item that represents the REPEAT command
type Repeat struct {
	ItemBase
	Times    ValueExpr
	Commands CommandSequenceExpr
}

// CommandName returns the name of the command
func (c Repeat) CommandName() string {
	return CommandRepeat
}

// Accept the given visitor
func (c Repeat) Accept(v Visitor) {
	v.VisitRepeat(c)
}

//###############################################

// PenUp is an AST item that represents the PENUP command
type PenUp struct {
	ItemBase
}

// CommandName returns the name of the command
func (c PenUp) CommandName() string {
	return CommandPenUp
}

// Accept the given visitor
func (c PenUp) Accept(v Visitor) {
	v.VisitPenUp(c)
}

//###############################################

// PenDown is an AST item that represents the PENDOWN command
type PenDown struct {
	ItemBase
}

// CommandName returns the name of the command
func (c PenDown) CommandName() string {
	return CommandPenDown
}

// Accept the given visitor
func (c PenDown) Accept(v Visitor) {
	v.VisitPenDown(c)
}

//###############################################

// ShowTurtle is an AST item that represents the SHOWTURTLE command
type ShowTurtle struct {
	ItemBase
}

// CommandName returns the name of the command
func (c ShowTurtle) CommandName() string {
	return CommandShowTurtle
}

// Accept the given visitor
func (c ShowTurtle) Accept(v Visitor) {
	v.VisitShowTurtle(c)
}

//###############################################

// HideTurtle is an AST item that represents the HIDETURTLE command
type HideTurtle struct {
	ItemBase
}

// CommandName returns the name of the command
func (c HideTurtle) CommandName() string {
	return CommandHideTurtle
}

// Accept the given visitor
func (c HideTurtle) Accept(v Visitor) {
	v.VisitHideTurtle(c)
}

//###############################################

// Clean is an AST item that represents the CLEAN command
type Clean struct {
	ItemBase
}

// CommandName returns the name of the command
func (c Clean) CommandName() string {
	return CommandClean
}

// Accept the given visitor
func (c Clean) Accept(v Visitor) {
	v.VisitClean(c)
}

//###############################################

// ClearScreen is an AST item that represents the CLEARSCREEN command
type ClearScreen struct {
	ItemBase
}

// CommandName returns the name of the command
func (c ClearScreen) CommandName() string {
	return CommandClearScreen
}

// Accept the given visitor
func (c ClearScreen) Accept(v Visitor) {
	v.VisitClearScreen(c)
}
