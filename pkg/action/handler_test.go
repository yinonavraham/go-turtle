package action_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-turtle/pkg/action"
	"strconv"
	"strings"
	"testing"
)

var testData = []struct {
	action action.Action
	text   string
}{
	{action: action.Home, text: "Home"},
	{action: action.PenDown, text: "Pen Down"},
	{action: action.ShowTurtle, text: "Show"},
	{action: action.MoveForward(10), text: "Forward 10"},
	{action: action.TurnRight(90), text: "Right 90"},
	{action: action.MoveBack(20), text: "Back 20"},
	{action: action.TurnLeft(270), text: "Left 270"},
	{action: action.HideTurtle, text: "Hide"},
	{action: action.PenUp, text: "Pen UP"},
	{action: action.Clean, text: "Clean"},
	{action: action.ClearScreen, text: "Clear Screen"},
}

func TestHandler(t *testing.T) {
	printer := printerHandler{}
	expected := strings.Builder{}
	for _, td := range testData {
		assert.NoError(t, action.Handle(td.action, &printer), td.action.String())
		_, _ = printer.WriteRune('\n')
		expected.WriteString(td.text)
		expected.WriteRune('\n')
	}
	assert.Equal(t, expected.String(), printer.String())
}

func TestHandler_Error(t *testing.T) {
	printer := printerHandler{
		err: true,
	}
	for _, td := range testData {
		err := action.Handle(td.action, &printer)
		assert.EqualError(t, err, td.text)
	}
}

func TestHandler_ErrUnknownAction(t *testing.T) {
	printer := printerHandler{}
	a := dummyAction("foo")
	err := action.Handle(a, &printer)
	assert.EqualError(t, err, "unknown action: foo")
	assert.IsType(t, action.ErrUnknownAction(""), err)
}

var _ action.Action = dummyAction("")

type dummyAction string

func (a dummyAction) ActionName() string { return string(a) }
func (a dummyAction) String() string     { return string(a) }

var _ action.Handler = &printerHandler{}

type printerHandler struct {
	strings.Builder
	err bool
}

func (h *printerHandler) writeString(s string) error {
	if h.err {
		return fmt.Errorf(s)
	}
	_, err := h.WriteString(s)
	return err
}

func (h *printerHandler) PenUp() error {
	return h.writeString("Pen UP")
}

func (h *printerHandler) PenDown() error {
	return h.writeString("Pen Down")
}

func (h *printerHandler) ShowTurtle() error {
	return h.writeString("Show")
}

func (h *printerHandler) HideTurtle() error {
	return h.writeString("Hide")
}

func (h *printerHandler) Home() error {
	return h.writeString("Home")
}

func (h *printerHandler) Clean() error {
	return h.writeString("Clean")
}

func (h *printerHandler) ClearScreen() error {
	return h.writeString("Clear Screen")
}

func (h *printerHandler) MoveForward(steps int) error {
	return h.writeString("Forward " + strconv.Itoa(steps))
}

func (h *printerHandler) MoveBack(steps int) error {
	return h.writeString("Back " + strconv.Itoa(steps))
}

func (h *printerHandler) TurnRight(angle int) error {
	return h.writeString("Right " + strconv.Itoa(angle))
}

func (h *printerHandler) TurnLeft(angle int) error {
	return h.writeString("Left " + strconv.Itoa(angle))
}
