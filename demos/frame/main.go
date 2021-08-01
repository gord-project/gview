// Demo code for the Frame primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

func main() {
	app := gview.NewApplication()
	frame := gview.NewFrame(gview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("Header left", true, gview.AlignLeft, tcell.ColorWhite).
		AddText("Header middle", true, gview.AlignCenter, tcell.ColorWhite).
		AddText("Header right", true, gview.AlignRight, tcell.ColorWhite).
		AddText("Header second middle", true, gview.AlignCenter, tcell.ColorRed).
		AddText("Footer middle", false, gview.AlignCenter, tcell.ColorGreen).
		AddText("Footer second middle", false, gview.AlignCenter, tcell.ColorGreen)
	if err := app.SetRoot(frame, true).Run(); err != nil {
		panic(err)
	}
}
