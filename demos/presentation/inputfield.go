package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

const inputField = `[green]package[white] main

[green]import[white] (
    [red]"strconv"[white]

    [red]tcell "github.com/gdamore/tcell/v2"[white]
    [red][white]
)

[green]func[white] [yellow]main[white]() {
    input := gview.[yellow]NewInputField[white]().
        [yellow]SetLabel[white]([red]"Enter a number: "[white]).
        [yellow]SetAcceptanceFunc[white](
            gview.InputFieldInteger,
        ).[yellow]SetDoneFunc[white]([yellow]func[white](key tcell.Key) {
            text := input.[yellow]GetText[white]()
            n, _ := strconv.[yellow]Atoi[white](text)
            [blue]// We have a number.[white]
        })
    gview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](input, true).
        [yellow]Run[white]()
}`

// InputField demonstrates the InputField.
func InputField(nextSlide func()) (title string, content gview.Primitive) {
	input := gview.NewInputField().
		SetLabel("Enter a number: ").
		SetAcceptanceFunc(gview.InputFieldInteger).SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "Input", Code(input, 30, 1, inputField)
}
