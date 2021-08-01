// Demo code for the Checkbox primitive.
package main

import "github.com/gord-project/gview"

func main() {
	app := gview.NewApplication()
	checkbox := gview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).Run(); err != nil {
		panic(err)
	}
}
