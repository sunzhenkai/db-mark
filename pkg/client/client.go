package client

// Client interface
type Client interface {
	Get(string) string
	Put(string) bool
	GetList(string) []string
	PutList([]string) bool
}
