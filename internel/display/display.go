package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Screen struct {
	grid *tview.Grid
}

func NewScreen(
	text, funcs, basicInfo []string,
) *Screen {
	//grid := tview.NewGrid().
	//	SetSize(2, 2, 0, 0).
	//	SetBorders(true).
	//	AddItem(tview.NewTextView(), 0, 0, 0, 0, 0, 0, true)
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
	menu := newPrimitive("Menu")
	//main := newPrimitive("Main content\n" + strings.Join(text, "\n"))
	main := NewTextArea(text).Build()
	sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(10, 0, 10).
		SetColumns(30, 0, 30).
		SetBorders(true)

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			modal := tview.NewModal().SetText("Help")

		}
	})

	// Layout for screens wider than 100 cells.
	grid.AddItem(main, 0, 0, 6, 2, 0, 100, true).
		AddItem(menu, 0, 2, 6, 1, 0, 100, false).
		AddItem(sideBar, 6, 2, 4, 1, 0, 100, false).
		AddItem(newPrimitive("Bottom"), 6, 0, 4, 2, 0, 100, false)

	return &Screen{grid: grid}
}

func (s *Screen) Display() {
	err := tview.NewApplication().SetRoot(s.grid, true).Run()
	if err != nil {
		return
	}
	tview.NewApplication().
}
