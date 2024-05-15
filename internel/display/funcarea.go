package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

type FuncArea struct {
	funcs []string
}

func NewFuncArea(funcs []string) *FuncArea {
	return &FuncArea{funcs: funcs}
}

func (f *FuncArea) Build() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter).
		SetText(strings.Join(f.funcs, "\n")).
		SetScrollable(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'j':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row+1, column)
				return nil
			case 'k':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row-1, column)
				return nil
			default:
				return event
			}
		}).SetBackgroundColor(tcell.ColorBlue)
	return textView
}
