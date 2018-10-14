package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/skanehira/mockapi/app/common"
)

// Endpoint endpoint info
type Endpoint struct {
	ID              string
	URL             string `gorm:"primary_key"`
	Method          string `gorm:"primary_key"`
	Description     string
	ResponseStatus  int
	ResponseHeaders string
	ResponseBody    string
	HistoryID       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

// SaveEndpoint creat new endpoint
func (d *DB) RegistEndpoint(endpoint *Endpoint) error {
	if err := d.Create(endpoint).Error; err != nil {
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

	if err := d.Model(&Endpoint{}).Scan(&list).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.NewErrNotFoundEndpoint(err)
		}
		return list, err
	}

	return list, nil
}
