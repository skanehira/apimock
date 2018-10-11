package db

import (
	"log"
	"time"
)

// Endpoint endpoint info
type Endpoint struct {
	ID              string
	URL             string
	Method          string
	ResponseBody    string
	ResponseHeaders map[string]string
	ResponseStatus  int
	HistoryID       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

// SaveEndpoint save endpoint to database
func (d *DB) RegistEndpoint(endpoint *Endpoint) error {
	if err := d.Create(endpoint).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetEndpoint get endpoint info
func (d *DB) FindEndpoint(url, method string) (*Endpoint, error) {
	endpoint := &Endpoint{
		URL:    url,
		Method: method,
	}

	if err := d.Find(endpoint).Error; err != nil {
		return endpoint, err
	}

	return endpoint, nil
}

// GetEndpointList get endpoints info list
func (d *DB) FindEndpointList() ([]*Endpoint, error) {
	var list []*Endpoint
	list = append(list, &Endpoint{})
	return list, nil
}
