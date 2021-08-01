// Demo code for the InputField primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

func main() {
	app := gview.NewApplication()
	inputField := gview.NewInputField().
		SetLabel("Enter a number: ").
		SetPlaceholder("E.g. 1234").
		SetFieldWidth(10).
		SetAcceptanceFunc(gview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).Run(); err != nil {
		panic(err)
	}
}
