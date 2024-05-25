package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

type CodeView struct {
	// text is the text that will be displayed in the text view.
	text []string
	// focus is the index of the text that is currently focused.
	focus uint64
}

// NewCodeView is the constructor of the CodeView struct.
func NewCodeView(text []string) *CodeView {
	return &CodeView{
		text:  text,
		focus: 0,
	}
}

// Build is the method that builds the text view
// and configure it with specific settings.
func (t *CodeView) Build() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter).
		SetText(strings.Join(t.text, "\n")).
		SetScrollable(true).
		SetRegions(true).
		SetBackgroundColor(tcell.ColorBlue)
	return textView
}
