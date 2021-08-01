package main

import "github.com/gord-project/gview" // Center returns a new primitive which shows the provided primitive in its

// Center, given the provided primitive's size.
func Center(width, height int, p gview.Primitive) gview.Primitive {
	return gview.NewFlex().
		AddItem(gview.NewBox(), 0, 1, false).
		AddItem(gview.NewFlex().
			SetDirection(gview.FlexRow).
			AddItem(gview.NewBox(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(gview.NewBox(), 0, 1, false), width, 1, true).
		AddItem(gview.NewBox(), 0, 1, false)
}
