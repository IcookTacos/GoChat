package client

type Client interface {
	SendMessage(message string) error
}
