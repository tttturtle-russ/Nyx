package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

type CodeView struct {
	// text is the text that will be displayed in the text view.
	text []string
	// focus is the index of the text that is currently focused.
	focus uint64
}

// InitCodeView is the constructor of the CodeView struct.
func InitCodeView(text []string) *CodeView {
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
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'j':
					row, column := textView.GetScrollOffset()
					textView.ScrollTo(row+1, column)
					textView.Highlight(strconv.Itoa(row + 1))
					return nil
				case 'k':
					row, column := textView.GetScrollOffset()
					textView.ScrollTo(row-1, column)
					textView.Highlight(strconv.Itoa(row - 1))
					return nil
				}
			default:
				return event
			}
			return nil
		}).SetTitle("Code View")
	return textView
}
