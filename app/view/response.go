package view

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ResponsePanel struct {
	View       *tview.TextView
	name       string
	FixedSize  int
	Proportion int
}

func NewResponsePanel(name string) *ResponsePanel {
	panel := &ResponsePanel{
		name:       name,
		FixedSize:  0,
		Proportion: 2,
	}

	view := tview.NewTextView().SetTextColor(tcell.ColorWhite)
	view.SetBorder(true).SetTitle(panel.name).SetTitleAlign(tview.AlignLeft)

	panel.View = view

	return panel
}

func (r *ResponsePanel) Name() string {
	return r.name
}

func (r *ResponsePanel) SetResponseBody(body string) {
	r.View.SetText(body)
}
