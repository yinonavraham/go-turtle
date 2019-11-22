package action

import "fmt"

func Handle(action Action, handler Handler) error {
	switch action {
	case PenDown:
		return handler.PenDown()
	case PenUp:
		return handler.PenUp()
	case ShowTurtle:
		return handler.ShowTurtle()
	case HideTurtle:
		return handler.HideTurtle()
	case Home:
		return handler.Home()
	case Clean:
		return handler.Clean()
	case ClearScreen:
		return handler.ClearScreen()
	default:
		switch a := action.(type) {
		case MoveAction:
			if a.Steps() < 0 {
				return handler.MoveBack(-a.Steps())
			} else {
				return handler.MoveForward(a.Steps())
			}
		case TurnAction:
			if a.Angle() < 0 {
				return handler.TurnLeft(-a.Angle())
			} else {
				return handler.TurnRight(a.Angle())
			}
		}
	}
	return ErrUnknownAction(fmt.Sprintf("unknown action: %s", action.ActionName()))
}

type Handler interface {
	PenDown() error
	PenUp() error
	ShowTurtle() error
	HideTurtle() error
	Home() error
	Clean() error
	ClearScreen() error
	MoveForward(steps int) error
	MoveBack(steps int) error
	TurnRight(angle int) error
	TurnLeft(angle int) error
}

var _ error = ErrUnknownAction("")

type ErrUnknownAction string

func (e ErrUnknownAction) Error() string {
	return string(e)
}