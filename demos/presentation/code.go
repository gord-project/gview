package main

import (
	"fmt"
	"github.com/gord-project/gview"
)

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
func Code(p gview.Primitive, width, height int, code string) gview.Primitive {
	// Set up code view.
	codeView := gview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	codeView.SetBorderPadding(1, 1, 2, 0)
	fmt.Fprint(codeView, code)

	return gview.NewFlex().
		AddItem(Center(width, height, p), 0, 1, true).
		AddItem(codeView, codeWidth, 1, false)
}
