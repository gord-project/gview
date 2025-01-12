package gview

import "github.com/gdamore/tcell/v2"

// Button is labeled box that triggers an action when selected.
//
// See https://github.com/Bios-Marcel/cordless/tview/wiki/Button for an example.
type Button struct {
	*Box

	// The text to be displayed before the input area.
	label string

	// The label color.
	labelColor tcell.Color

	// The label color when the button is in focus.
	labelColorActivated tcell.Color

	// The background color when the button is in focus.
	backgroundColorActivated tcell.Color

	// An optional function which is called when the button was selected.
	selected func()

	// An optional function which is called when the user leaves the button. A
	// key is provided indicating which key was pressed to leave (tab or backtab).
	blur func(tcell.Key)
}

// NewButton returns a new input field.
func NewButton(label string) *Button {
	box := NewBox().SetBackgroundColor(Styles.ContrastBackgroundColor)
	box.SetRect(0, 0, TaggedStringWidth(label)+4, 1)
	button := &Button{
		Box:                      box,
		label:                    label,
		labelColor:               Styles.PrimaryTextColor,
		labelColorActivated:      Styles.InverseTextColor,
		backgroundColorActivated: Styles.PrimaryTextColor,
	}
	button.SetMouseHandler(func(event *tcell.EventMouse) bool {
		if event.Buttons() == tcell.Button1 {
			button.selected()
			return true
		}

		return false
	})
	return button
}

// SetLabel sets the button text.
func (b *Button) SetLabel(label string) *Button {
	b.label = label
	return b
}

// GetLabel returns the button text.
func (b *Button) GetLabel() string {
	return b.label
}

// SetLabelColor sets the color of the button text.
func (b *Button) SetLabelColor(color tcell.Color) *Button {
	b.labelColor = color
	return b
}

// SetLabelColorActivated sets the color of the button text when the button is
// in focus.
func (b *Button) SetLabelColorActivated(color tcell.Color) *Button {
	b.labelColorActivated = color
	return b
}

// SetBackgroundColorActivated sets the background color of the button text when
// the button is in focus.
func (b *Button) SetBackgroundColorActivated(color tcell.Color) *Button {
	b.backgroundColorActivated = color
	return b
}

// SetSelectedFunc sets a handler which is called when the button was selected.
func (b *Button) SetSelectedFunc(handler func()) *Button {
	b.selected = handler
	return b
}

// SetBlurFunc sets a handler which is called when the user leaves the button.
// The callback function is provided with the key that was pressed, which is one
// of the following:
//
//   - KeyEscape: Leaving the button with no specific direction.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (b *Button) SetBlurFunc(handler func(key tcell.Key)) *Button {
	b.blur = handler
	return b
}

// Draw draws this primitive onto the screen.
func (b *Button) Draw(screen tcell.Screen) bool {
	// Draw the box.
	borderColor := b.borderColor
	backgroundColor := b.backgroundColor
	if IsVtxxx {
		b.reverse = false
	}
	if b.focus.HasFocus() {
		b.backgroundColor = b.backgroundColorActivated
		b.borderColor = b.labelColorActivated
		if IsVtxxx {
			b.reverse = true
		}
		defer func() {
			b.borderColor = borderColor
		}()
	}

	res := b.Box.Draw(screen)
	if !res {
		return false
	}

	b.backgroundColor = backgroundColor

	// Draw label.
	x, y, width, height := b.GetInnerRect()
	if width > 0 && height > 0 {
		y = y + height/2
		labelColor := b.labelColor
		if b.focus.HasFocus() {
			labelColor = b.labelColorActivated
		}
		Print(screen, b.label, x, y, width, AlignCenter, labelColor)
	}

	return true
}

// InputHandler returns the handler for this primitive.
func (b *Button) InputHandler() InputHandlerFunc {
	return b.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) *tcell.EventKey {
		// Process key event.
		switch key := event.Key(); key {
		case tcell.KeyEnter: // Selected.
			if b.selected != nil {
				b.selected()
			}
			return nil
		case tcell.KeyBacktab, tcell.KeyTab, tcell.KeyEscape: // Leave. No action.
			if b.blur != nil {
				b.blur(key)
			}
			return nil
		}

		return event
	})
}
