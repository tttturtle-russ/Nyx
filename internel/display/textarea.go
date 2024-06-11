package display

import (
	"github.com/rivo/tview"
	"github.com/tttturtle-russ/Nyx/internel/parser"
)

//var codeView *tview.TextView
//
//// InitCodeView is the constructor of the CodeView struct.
//func InitCodeView(code string) {
//	codeView = tview.NewTextView()
//	codeView.SetTextAlign(tview.AlignLeft).
//		SetText(code).
//		SetScrollable(true).
//		SetRegions(true).
//		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
//			switch event.Key() {
//			case tcell.KeyRune:
//				switch event.Rune() {
//				case 'j':
//					row, column := codeView.GetScrollOffset()
//					codeView.ScrollTo(row+1, column)
//					codeView.Highlight(strconv.Itoa(row + 1))
//					return nil
//				case 'k':
//					row, column := codeView.GetScrollOffset()
//					codeView.ScrollTo(row-1, column)
//					codeView.Highlight(strconv.Itoa(row - 1))
//					return nil
//				}
//			default:
//				return event
//			}
//			return nil
//		}).SetTitle("Code View")
//}

var codePage *tview.Pages

// InitCodePage init the codePage and add corresponding function code
func InitCodePage(functions []*parser.Function) {
	codePage = tview.NewPages()
	for _, function := range functions {
		codePage.AddPage(function.Name, tview.NewTextView().SetText(function.Code).SetScrollable(true).ScrollToBeginning(), true, true)
	}
}
