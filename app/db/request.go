package db

// Request user request info
type Request struct {
	ID      string
	Headers string
	Query   string
	Body    string
}
