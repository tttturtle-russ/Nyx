package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

type FuncView struct {
	funcs       []string
	currentFunc uint64
}

func InitFuncView(funcs []string) *FuncView {
	return &FuncView{
		funcs:       funcs,
		currentFunc: 0,
	}
}

func (f *FuncView) Build() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter).
		SetText(strings.Join(f.funcs, "\n")).
		SetRegions(true).
		SetScrollable(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'j':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row+1, column)
				if f.currentFunc < uint64(len(f.funcs)-1) {
					f.currentFunc++
					textView.Highlight(strconv.Itoa(int(f.currentFunc)))
				}
				return nil
			case 'k':
				row, column := textView.GetScrollOffset()
				textView.ScrollTo(row-1, column)
				if f.currentFunc > 0 {
					f.currentFunc--
					textView.Highlight(strconv.Itoa(int(f.currentFunc)))
				}
				return nil
			default:
				return event
			}
		}).SetBackgroundColor(tcell.ColorBlue)
	return textView
}
