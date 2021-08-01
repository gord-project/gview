package main

import (
	"github.com/gord-project/gview"
	"strings"
)

func main() {
	app := gview.NewApplication()
	textView := gview.NewTextView().
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		}).
		SetScrollable(true).
		SetText(strings.Repeat("OwO\n", 100)).
		SetIndicateOverflow(true).
		SetBorderSides(true, true, true, true).
		SetBorder(true).
		SetBorderPadding(0, 0, 0, 10)
	flex := gview.NewFlex()
	flex.SetDirection(gview.FlexRow)
	flex.AddItem(gview.NewTextView(), 1, -1, false)
	flex.AddItem(textView, 10, -1, false)

	root := gview.NewTreeNode("Root")
	root.AddChild(gview.NewTreeNode("A"))
	root.AddChild(gview.NewTreeNode("B"))
	root.AddChild(gview.NewTreeNode("C"))
	root.AddChild(gview.NewTreeNode("D"))
	root.AddChild(gview.NewTreeNode("E"))
	root.AddChild(gview.NewTreeNode("F"))
	tree := gview.NewTreeView().SetRoot(root)
	tree.SetBorder(true).SetIndicateOverflow(true)
	flex.AddItem(tree, 6, -1, false)

	list := gview.NewList()
	list.AddItem("A", "", 0, nil)
	list.AddItem("B", "", 0, nil)
	list.AddItem("C", "", 0, nil)
	list.AddItem("D", "", 0, nil)
	list.AddItem("E", "", 0, nil)
	list.AddItem("F", "", 0, nil)
	list.SetBorder(true).SetIndicateOverflow(true)
	list.ShowSecondaryText(false)
	flex.AddItem(list, 6, -1, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
