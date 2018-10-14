package view

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/skanehira/mockapi/app/config"
	"github.com/skanehira/mockapi/app/db"
)

type View struct {
	active int
	*tview.Application
	panels    map[string]Panel
	primitive tview.Primitive
	config    *config.Config
	db        *db.DB
	keybind   *Keybind
}

type Panel interface {
	Name() string
}

var (
	EndpointPanelName = "Endpoint List"
	HeadersPanelName  = "Headers"
	ResponsePanelName = "Response Body"
)

func New(db *db.DB, config *config.Config) *View {
	view := &View{
		active:      0,
		Application: tview.NewApplication(),
		panels:      make(map[string]Panel),
		config:      config,
		db:          db,
	}

	endpointPanel := NewEndpointPanel(EndpointPanelName)
	responsePanel := NewResponsePanel(ResponsePanelName)
	headersPanel := NewHeadersPanel(HeadersPanelName)

	view.panels[EndpointPanelName] = endpointPanel
	view.panels[HeadersPanelName] = headersPanel
	view.panels[ResponsePanelName] = responsePanel

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(endpointPanel.Table, endpointPanel.FixedSize, endpointPanel.Proportion, true).
			AddItem(headersPanel.View, headersPanel.FixedSize, headersPanel.Proportion, true).
			AddItem(responsePanel.View, responsePanel.FixedSize, responsePanel.Proportion, true),
			0, 2, true)

	view.primitive = flex

	view.SetFocus(endpointPanel.Table)

	view.keybind = NewKeybind(endpointPanel, responsePanel, headersPanel, view.db)

	return view
}

func (v *View) GetEndpointList() ([]*db.Endpoint, error) {
	endpoints, err := v.db.FindEndpointList()
	if err != nil {
		return endpoints, err
	}

	return endpoints, nil
}

func (v *View) Setup() error {
	// set data
	endpointList, err := v.GetEndpointList()

	if err != nil {
		return err
	}

	endpoint := v.GetEndpointPanel()
	response := v.GetResponsePanel()
	headers := v.GetHeadersPanel()

	endpoint.SetEndpointList(endpointList)
	response.SetResponseBody(endpointList[0].ResponseBody)
	headers.SetHeaders(endpointList[0].ResponseHeaders)

	// set keybinding
	v.GetEndpointPanel().SetKeybinding(v.keybind)
	v.SetKeybinding()

	return nil
}

func (v *View) SetKeybinding() {
	panels := map[int]tview.Primitive{
		0: v.GetEndpointPanel().Table,
		1: v.GetResponsePanel().View,
		2: v.GetHeadersPanel().View,
	}

	v.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'l':
			v.active = (v.active + 1) % len(v.panels)
			v.SetFocus(panels[v.active])
		case 'h':
			v.active = (v.active - 1) % len(v.panels)
			v.SetFocus(panels[v.active])
		}

		return event
	})
}

func (v *View) GetEndpointPanel() *EndpointPanel {
	return v.panels[EndpointPanelName].(*EndpointPanel)
}

func (v *View) GetResponsePanel() *ResponsePanel {
	return v.panels[ResponsePanelName].(*ResponsePanel)
}

func (v *View) GetHeadersPanel() *HeadersPanel {
	return v.panels[HeadersPanelName].(*HeadersPanel)
}

func (v *View) Run() error {
	if err := v.SetRoot(v.primitive, true).Run(); err != nil {
		return err
	}
	return nil
}
