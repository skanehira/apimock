package db

// Request user request info
type Request struct {
	ID      string
	Headers []string
	Queries []string
	Body    string
}
