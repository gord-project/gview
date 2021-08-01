package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

const helloWorld = `[green]package[white] main

[green]import[white] (
    [red][white]
)

[green]func[white] [yellow]main[white]() {
    box := gview.[yellow]NewBox[white]().
        [yellow]SetBorder[white](true).
        [yellow]SetTitle[white]([red]"Hello, world!"[white])
    gview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](box, true).
        [yellow]Run[white]()
}`

// HelloWorld shows a simple "Hello world" example.
func HelloWorld(nextSlide func()) (title string, content gview.Primitive) {
	// We use a text view because we want to capture keyboard input.
	textView := gview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	textView.SetBorder(true).SetTitle("Hello, world!")
	return "Hello, world", Code(textView, 30, 10, helloWorld)
}
