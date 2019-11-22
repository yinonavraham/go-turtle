package action

import (
	"fmt"
	"github.com/yinonavraham/go-turtle/pkg/parser/ast"
)

// Action base interface
type Action interface {
	fmt.Stringer
	// The name of the action
	ActionName() string
}

const (
	// PenDown action
	PenDown = simpleAction(ast.CommandPenDown)
	// PenUp action
	PenUp = simpleAction(ast.CommandPenUp)
	// ShowTurtle action
	ShowTurtle = simpleAction(ast.CommandShowTurtle)
	// HideTurtle action
	HideTurtle = simpleAction(ast.CommandHideTurtle)
	// Home action
	Home = simpleAction(ast.CommandHome)
	// Clean action
	Clean = simpleAction(ast.CommandClean)
	// ClearScreen action
	ClearScreen = simpleAction(ast.CommandClearScreen)
)

var _ Action = PenDown
var _ Action = PenUp
var _ Action = ShowTurtle
var _ Action = HideTurtle
var _ Action = Home
var _ Action = Clean
var _ Action = ClearScreen
var _ Action = MoveForward(0)
var _ Action = MoveBack(0)
var _ Action = TurnRight(0)
var _ Action = TurnLeft(0)

type simpleAction string

func (a simpleAction) ActionName() string {
	return string(a)
}

func (a simpleAction) String() string {
	return a.ActionName()
}

// MoveAction interface - for any move action (forward, back)
type MoveAction interface {
	Action
	// The number of steps to move. Positive number moves forward, negative number of steps means move back
	Steps() int
}

// MoveForward creates an Action to move forward a given number of steps
func MoveForward(steps int) MoveAction {
	return move{
		name:  ast.CommandForward,
		steps: steps,
	}
}

// MoveBack creates an Action to move back a given number of steps
func MoveBack(steps int) MoveAction {
	return move{
		name:  ast.CommandBack,
		steps: -steps,
	}
}

type move struct {
	name  string
	steps int
}

func (a move) ActionName() string {
	return a.name
}

func (a move) Steps() int {
	return a.steps
}

func (a move) String() string {
	return fmt.Sprintf("%s %d", a.name, a.steps)
}

// TurnAction interface - for any turn action (right, left)
type TurnAction interface {
	Action
	// The angle (degrees) to turn. Positive number turns right (clockwise), negative number of degrees means turn left
	// (counter clockwise).
	Angle() int
}

// TurnRight creates an Action to turn right in a given angle (degrees)
func TurnRight(angle int) TurnAction {
	return turn{
		name:  ast.CommandRight,
		angle: angle,
	}
}

// TurnLeft creates an Action to turn left in a given angle (degrees)
func TurnLeft(angle int) TurnAction {
	return turn{
		name:  ast.CommandLeft,
		angle: -angle,
	}
}

type turn struct {
	name  string
	angle int
}

func (a turn) ActionName() string {
	return a.name
}

func (a turn) Angle() int {
	return a.angle
}

func (a turn) String() string {
	return fmt.Sprintf("%s %d", a.name, a.angle)
}
