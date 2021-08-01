package main

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/yellowsink/gord/tview"
)

func main() {
	field := tview.NewInputField()
	field.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlV {
			field.Insert("Fotze")
			return nil
		}

		return event
	})
	app := tview.NewApplication().SetRoot(field, true)
	app.Run()
}
