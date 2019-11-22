package action

import (
	"fmt"
	"github.com/yinonavraham/go-turtle/pkg/parser/ast"
)

type Action interface {
	fmt.Stringer
	ActionName() string
}

const (
	PenDown     = simpleAction(ast.CommandPenDown)
	PenUp       = simpleAction(ast.CommandPenUp)
	ShowTurtle  = simpleAction(ast.CommandShowTurtle)
	HideTurtle  = simpleAction(ast.CommandHideTurtle)
	Home        = simpleAction(ast.CommandHome)
	Clean       = simpleAction(ast.CommandClean)
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

type MoveAction interface {
	Action
	Steps() int
}

func MoveForward(steps int) MoveAction {
	return move{
		name:  ast.CommandForward,
		steps: steps,
	}
}

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

type TurnAction interface {
	Action
	Angle() int
}

func TurnRight(angle int) TurnAction {
	return turn{
		name:  ast.CommandRight,
		angle: angle,
	}
}

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