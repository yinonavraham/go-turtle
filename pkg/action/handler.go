package action

import "fmt"

// Handle a given action using the provided Handler.
//
// Returns an error in case handling failed. If the given action name is unknown (not supported), the error is of type
// ErrUnknownAction.
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
			}
			return handler.MoveForward(a.Steps())
		case TurnAction:
			if a.Angle() < 0 {
				return handler.TurnLeft(-a.Angle())
			}
			return handler.TurnRight(a.Angle())
		}
	}
	return ErrUnknownAction(fmt.Sprintf("unknown action: %s", action.ActionName()))
}

// Handler interface, used for handling actions.
//
// See: Handle(Action, Handler) error
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

// ErrUnknownAction returned by the Handle function in case the given action name is not supported.
type ErrUnknownAction string

func (e ErrUnknownAction) Error() string {
	return string(e)
}
