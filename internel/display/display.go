package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

type Screen struct {
	app      *tview.Application
	grid     *tview.Grid
	codeView *tview.TextView
	funcView *tview.TextView
	memView  *tview.TextView
	miscView *tview.TextView
}

func InitScreen(
	text, funcs, basicInfo []string,
) *Screen {
	app := tview.NewApplication()
	newPrimitive := func(text string) tview.Primitive {
		t := tview.NewTextView()
		t.SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetScrollable(true).
			SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Rune() {
				case 'j':
					row, column := t.GetScrollOffset()
					t.ScrollTo(row+1, column)
					return nil
				case 'k':
					row, column := t.GetScrollOffset()
					t.ScrollTo(row-1, column)
					return nil
				default:
					return event
				}
			})
		return t
	}
	funcView := InitFuncView(funcs).Build()
	codeView := NewCodeView(text).Build()
	codeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				row, column := codeView.GetScrollOffset()
				codeView.ScrollTo(row+1, column)
				codeView.Highlight(strconv.Itoa(row + 1))
				return nil
			case 'k':
				row, column := codeView.GetScrollOffset()
				codeView.ScrollTo(row-1, column)
				codeView.Highlight(strconv.Itoa(row - 1))
				return nil
			}
		case tcell.KeyCtrlL:
			app.SetFocus(funcView)
		default:
			return event
		}
		return nil
	})
	memView := newPrimitive("Side Bar")
	miscView := newPrimitive("Bottom")
	grid := tview.NewGrid().
		SetRows(10, 0, 10).
		SetColumns(30, 0, 30).
		SetBorders(true)

	//grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Rune() {
	//	case 'h':
	//		modal := tview.NewModal().SetText("Help")
	//
	//	}
	//})

	// Layout for screens wider than 100 cells.
	grid.AddItem(codeView, 0, 0, 6, 2, 0, 100, true).
		AddItem(funcView, 0, 2, 6, 1, 0, 100, false).
		AddItem(memView, 6, 2, 4, 1, 0, 100, false).
		AddItem(miscView, 6, 0, 4, 2, 0, 100, false)

	return &Screen{
		app:      app,
		grid:     grid,
		codeView: codeView,
		funcView: funcView,
	}
}

func (s *Screen) Display() {
	err := s.app.SetRoot(s.grid, true).Run()
	if err != nil {
		return
	}
}
