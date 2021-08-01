package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

// Flex demonstrates flexbox layout.
func Flex(nextSlide func()) (title string, content gview.Primitive) {
	modalShown := false
	pages := gview.NewPages()
	textView := gview.NewTextView().
		SetDoneFunc(func(key tcell.Key) {
			if modalShown {
				nextSlide()
				modalShown = false
			} else {
				pages.ShowPage("modal")
				modalShown = true
			}
		})
	textView.SetBorder(true).SetTitle("Flexible width, twice of middle column")
	flex := gview.NewFlex().
		AddItem(textView, 0, 2, true).
		AddItem(gview.NewFlex().
			SetDirection(gview.FlexRow).
			AddItem(gview.NewBox().SetBorder(true).SetTitle("Flexible width"), 0, 1, false).
			AddItem(gview.NewBox().SetBorder(true).SetTitle("Fixed height"), 15, 1, false).
			AddItem(gview.NewBox().SetBorder(true).SetTitle("Flexible height"), 0, 1, false), 0, 1, false).
		AddItem(gview.NewBox().SetBorder(true).SetTitle("Fixed width"), 30, 1, false)
	modal := gview.NewModal().
		SetText("Resize the window to see the effect of the flexbox parameters").
		AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})
	pages.AddPage("flex", flex, true, true).
		AddPage("modal", modal, false, false)
	return "Flex", pages
}
