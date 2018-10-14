package view

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type HeadersPanel struct {
	View       *tview.TextView
	name       string
	FixedSize  int
	Proportion int
}

func NewHeadersPanel(name string) *HeadersPanel {
	panel := &HeadersPanel{
		name:       name,
		FixedSize:  0,
		Proportion: 1,
	}

	view := tview.NewTextView().SetTextColor(tcell.ColorWhite)
	view.SetBorder(true).SetTitle(panel.name).SetTitleAlign(tview.AlignLeft)

	panel.View = view

	return panel
}

func (h *HeadersPanel) Name() string {
	return h.name
}

func (h *HeadersPanel) SetHeaders(headers string) {
	h.View.SetText(headers)
}
