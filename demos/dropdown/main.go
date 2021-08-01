// Demo code for the DropDown primitive.
package main

import "github.com/gord-project/gview"

func main() {
	app := gview.NewApplication()
	dropdown := gview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
	if err := app.SetRoot(dropdown, true).Run(); err != nil {
		panic(err)
	}
}
