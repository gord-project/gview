// Demo code for the Modal primitive.
package main

import (
	"github.com/gord-project/gview"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func main() {
	app := gview.NewApplication()

	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	modal := func(p gview.Primitive, width, height int) gview.Primitive {
		return gview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(gview.NewFlex().SetDirection(gview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, false).
				AddItem(nil, 0, 1, false), width, 1, false).
			AddItem(nil, 0, 1, false)
	}

	background := gview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetText(strings.Repeat("background ", 1000))

	box := gview.NewBox().
		SetBorder(true).
		SetTitle("Centered Box")

	pages := gview.NewPages().
		AddPage("background", background, true, true).
		AddPage("modal", modal(box, 40, 10), true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
