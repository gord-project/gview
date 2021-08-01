// Demo code for the Button primitive.
package main

import "github.com/gord-project/gview"

func main() {
	app := gview.NewApplication()
	button := gview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(true).SetRect(0, 0, 22, 3)
	if err := app.SetRoot(button, false).Run(); err != nil {
		panic(err)
	}
}
