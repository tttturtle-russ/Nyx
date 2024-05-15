package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

type TextArea struct {
	// text is the text that will be displayed in the text view.
	text []string
	// focus is the index of the text that is currently focused.
	focus uint64
}

// NewTextArea is the constructor of the TextArea struct.
func NewTextArea(text []string) *TextArea {
	return &TextArea{
		text:  text,
		focus: 0,
	}
}

// Build is the method that builds the text view
// and configure it with specific settings.
func (t *TextArea) Build() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter).
		SetText(strings.Join(t.text, "\n")).
		SetScrollable(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			// j and k are used to scroll the text view up and down just like vim keymap.
			case 'j':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row+1, column)
				if t.focus < uint64(len(t.text)-1) {
					t.focus++
				}

				return nil
			case 'k':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row-1, column)
				if t.focus > 0 {
					t.focus--
				}
				return nil
			default:
				return event
			}
		}).
		SetBackgroundColor(tcell.ColorBlue)
	return textView
}
