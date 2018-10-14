package view

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/skanehira/mockapi/app/common"
	"github.com/skanehira/mockapi/app/db"
)

type EndpointPanel struct {
	Table        *tview.Table
	name         string
	TableHeaders []string
	FixedSize    int
	Proportion   int
}

func NewEndpointPanel(name string) *EndpointPanel {
	panel := &EndpointPanel{
		name: name,
		TableHeaders: []string{
			"METHOD",
			"URL",
			"DESCRIPTION",
			"CREATED",
		},
		FixedSize:  0,
		Proportion: 1,
	}

	table := tview.NewTable().SetBorders(false).SetSeparator('|').SetSelectable(true, false)

	for i, h := range panel.TableHeaders {
		cell := &tview.TableCell{
			Text:            h,
			Align:           tview.AlignCenter,
			Color:           tcell.ColorYellow,
			BackgroundColor: tcell.ColorGray,
			Attributes:      tcell.AttrBold,
			NotSelectable:   true,
		}

		cell.SetExpansion(1)
		table.SetCell(0, i, cell)
	}

	table.SetTitle(panel.name).SetBorder(true).SetTitleAlign(tview.AlignCenter)

	panel.Table = table
	return panel
}

func (e *EndpointPanel) SetEndpointList(endpoints []*db.Endpoint) error {
	for i, endpoint := range endpoints {
		i++

		data := []string{
			endpoint.Method,
			endpoint.URL,
			endpoint.Description,
			common.ParseDateToString(endpoint.CreatedAt),
		}

		for j, d := range data {
			e.Table.SetCell(i, j, tview.NewTableCell(d).SetTextColor(tcell.ColorWhite))
		}
	}

	return nil
}

func (e *EndpointPanel) SetKeybinding(keybind *Keybind) {
	e.Table.SetSelectionChangedFunc(keybind.UpdatePanel)
}

func (e *EndpointPanel) Name() string {
	return e.name
}
