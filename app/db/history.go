package db

// History request and response history
type History struct {
	ID       string
	Endpoint *Endpoint
	Request  *Request
}

// SaveHistory save history info to database
func (d *DB) RegistHistory(endpoint *Endpoint, history *History) error {
	return nil
}

// GetHistory get history info
func (d *DB) GetHistory(id string) (*History, error) {
	history := &History{}
	return history, nil
}

// GetHistoryList get history info list
func (d *DB) GetHistoryList() ([]*History, error) {
	var list []*History

	list = append(list, &History{})
	return list, nil
}
