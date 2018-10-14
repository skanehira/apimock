package view

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/skanehira/mockapi/app/db"
)

type Keybind struct {
	endpointPanel *EndpointPanel
	responsePanel *ResponsePanel
	headersPanel  *HeadersPanel
	db            *db.DB
}

func NewKeybind(endpointPanel *EndpointPanel, responsePanel *ResponsePanel, headersPanel *HeadersPanel, db *db.DB) *Keybind {
	return &Keybind{
		endpointPanel: endpointPanel,
		responsePanel: responsePanel,
		headersPanel:  headersPanel,
		db:            db,
	}
}

// add endpoint
func (k *Keybind) AddEndpoint(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'c' {
		// TODO add
	}
	return event
}

// remove endpoint
func (k *Keybind) RemoveEndpoint(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'd' {
		// TODO remove
	}
	return event
}

// update endpoint
func (k *Keybind) UpdateEndpoint(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'i' {
		// TODO update
	}
	return event
}

// get endpoint
func (k *Keybind) GetEndpoint(event *tcell.EventKey) *tcell.EventKey {
	return event
}

func (k *Keybind) UpdatePanel(row, column int) {
	list, err := k.db.FindEndpointList()

	if err != nil {
		log.Println(err)
		return
	}

	if row < 1 {
		return
	}

	endpoint := list[row-1]

	k.responsePanel.SetResponseBody(endpoint.ResponseBody)
	k.headersPanel.SetHeaders(endpoint.ResponseHeaders)

	return
}
