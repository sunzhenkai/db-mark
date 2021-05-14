package client

// Client interface
type Client interface {
	Get(string) string
	Put(string, string)
}
