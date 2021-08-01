package gview

import "github.com/gdamore/tcell/v2"

// MouseSupport defines wether a component supports accepting mouse events
type MouseSupport interface {
	MouseHandler() func(event *tcell.EventMouse) bool
}
