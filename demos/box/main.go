// Demo code for the Box primitive.
package main

import "github.com/gord-project/gview"

func main() {
	box := gview.NewBox().
		SetBorder(true).
		SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:] [::bu]title")
	if err := gview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
