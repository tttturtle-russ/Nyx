package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	_ "strconv"
)

type Screen struct {
	app      *tview.Application
	grid     *tview.Grid
	codeView *tview.TextView
	funcView *tview.List
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
			SetTitle(text)
		return t
	}
	InitFuncView(app, funcs)
	codeView := InitCodeView(text).Build()

	memView := newPrimitive("Mem View")
	miscView := newPrimitive("Misc View")
	grid := tview.NewGrid().
		SetRows(10, 0, 10).
		SetColumns(30, 0, 30).
		SetBorders(true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(codeView, 0, 0, 6, 2, 0, 100, true).
		AddItem(funcList, 0, 2, 6, 1, 0, 100, false).
		AddItem(memView, 6, 2, 4, 1, 0, 100, false).
		AddItem(miscView, 6, 0, 4, 2, 0, 100, false)

	// Set the input to make it possible to change focus
	// TODO: find a better way to switch focus
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if app.GetFocus() == nil {
			app.SetFocus(codeView)
		}
		v, _ := app.GetFocus().(*tview.TextView)
		title := v.GetTitle()
		switch event.Key() {
		case tcell.KeyCtrlL:
			if title == "Code View" {
				app.SetFocus(funcList)
			} else if title == "Mem View" {
				app.SetFocus(miscView)
			}
		case tcell.KeyCtrlH:
			if title == "Func View" {
				app.SetFocus(codeView)
			} else if title == "Misc View" {
				app.SetFocus(memView)
			}
		case tcell.KeyCtrlJ:
			if title == "Code View" {
				app.SetFocus(memView)
			} else if title == "Func View" {
				app.SetFocus(miscView)
			}
		case tcell.KeyCtrlK:
			if title == "Mem View" {
				app.SetFocus(codeView)
			} else if title == "Misc View" {
				app.SetFocus(funcList)
			}
		default:
			return event
		}
		return event
	})

	return &Screen{
		app:      app,
		grid:     grid,
		codeView: codeView,
		funcView: funcList,
	}
}

func (s *Screen) Display() {
	err := s.app.SetRoot(s.grid, true).Run()
	if err != nil {
		return
	}
}
