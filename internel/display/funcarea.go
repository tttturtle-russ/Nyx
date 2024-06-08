package display

import (
	"github.com/rivo/tview"
)

var (
	funcList *tview.List
)

func InitFuncView(app *tview.Application, funcs []string) {
	funcList = tview.NewList()
	for _, _func := range funcs[:len(funcs)-2] {
		funcList.AddItem(_func, "", 0, nil)
	}
	_func := funcs[len(funcs)-1]
	funcList.AddItem(_func, "", 'q', func() {
		app.Stop()
	})
	funcList.SetTitle("Func View")
}

//func (f *FuncView) Build() *tview.TextView {
//	textView := tview.NewTextView()
//	textView.SetTextAlign(tview.AlignCenter).
//		SetText(strings.Join(f.funcs, "\n")).
//		SetRegions(true).
//		SetScrollable(true).
//		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
//			if !textView.HasFocus() {
//				return event
//			}
//			switch event.Rune() {
//			case 'j':
//				row, column := textView.GetScrollOffset()
//				textView.ScrollTo(row+1, column)
//				if f.currentFunc < uint64(len(f.funcs)-1) {
//					f.currentFunc++
//					textView.Highlight(strconv.Itoa(int(f.currentFunc)))
//				}
//				return nil
//			case 'k':
//				row, column := textView.GetScrollOffset()
//				textView.ScrollTo(row-1, column)
//				if f.currentFunc > 0 {
//					f.currentFunc--
//					textView.Highlight(strconv.Itoa(int(f.currentFunc)))
//				}
//				return nil
//			default:
//				return event
//			}
//		}).
//		SetBackgroundColor(tcell.ColorBlue).
//		SetTitle("Func View")
//	return textView
//}
