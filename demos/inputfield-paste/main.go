package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

func main() {
	field := gview.NewInputField()
	field.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlV {
			field.Insert("Fotze")
			return nil
		}

		return event
	})
	app := gview.NewApplication().SetRoot(field, true)
	app.Run()
}
