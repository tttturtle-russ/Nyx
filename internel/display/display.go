package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tttturtle-russ/Nyx/internel/parser"
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
	funcs []*parser.Function,
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
	InitFuncList(funcs)
	//InitCodeView(funcs[0].Code)
	InitCodePage(funcs)
	memView := newPrimitive("Mem View")
	miscView := newPrimitive("Misc View")
	grid := tview.NewGrid().
		SetRows(10, 0, 10).
		SetColumns(30, 0, 30).
		SetBorders(true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(codePage, 0, 0, 6, 2, 0, 100, false).
		AddItem(funcList, 0, 2, 6, 1, 0, 100, true).
		AddItem(memView, 6, 2, 4, 1, 0, 100, false).
		AddItem(miscView, 6, 0, 4, 2, 0, 100, false)

	// Set the input to make it possible to change focus
	// TODO: find a better way to switch focus
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if app.GetFocus() == nil {
			app.SetFocus(codePage)
		}
		focus := app.GetFocus()
		switch event.Key() {
		case tcell.KeyCtrlL:
			if focus == codePage {
				app.SetFocus(funcList)
			} else if focus == memView {
				app.SetFocus(miscView)
			}
		case tcell.KeyCtrlH:
			if focus == funcList {
				app.SetFocus(codePage)
			} else if focus == miscView {
				app.SetFocus(memView)
			}
		case tcell.KeyCtrlJ:
			if focus == codePage {
				app.SetFocus(memView)
			} else if focus == funcList {
				app.SetFocus(miscView)
			}
		case tcell.KeyCtrlK:
			if focus == memView {
				app.SetFocus(codePage)
			} else if focus == miscView {
				app.SetFocus(funcList)
			}
		default:
			return event
		}
		return event
	})

	return &Screen{
		app:  app,
		grid: grid,
		//codeView: codep,
		funcView: funcList,
	}
}

func (s *Screen) Display() {
	err := s.app.SetRoot(s.grid, true).Run()
	if err != nil {
		return
	}
}
