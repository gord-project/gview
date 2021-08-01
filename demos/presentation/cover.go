package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/gord-project/gview"
)

const logo = `
   __        _
  / /__   __(_)__ _      __
 / __/ | / / / _ \ | /| / /
/ /_ | |/ / /  __/ |/ |/ /
\__/ |___/_/\___/|__/|__/

`

const (
	subtitle   = `tview - Rich Widgets for Terminal UIs`
	navigation = `Ctrl-N: Next slide    Ctrl-P: Previous slide    Ctrl-C: Exit`
)

// Cover returns the cover page.
func Cover(nextSlide func()) (title string, content gview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := gview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := gview.NewFrame(gview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, gview.AlignCenter, tcell.ColorWhite).
		AddText("", true, gview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, gview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := gview.NewFlex().
		SetDirection(gview.FlexRow).
		AddItem(gview.NewBox(), 0, 7, false).
		AddItem(gview.NewFlex().
			AddItem(gview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(gview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Start", flex
}
