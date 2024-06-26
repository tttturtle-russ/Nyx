package display

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/tttturtle-russ/Nyx/internel/parser"
)

var (
	funcList *tview.List
)

func InitFuncList(funcs []*parser.Function) {
	funcList = tview.NewList()
	for _, _func := range funcs[:len(funcs)-1] {
		funcList.AddItem(_func.Name, fmt.Sprintf("0x%08X", _func.StartOffset), 0, func() {
			codePage.SwitchToPage(_func.Name)
		})
	}
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
