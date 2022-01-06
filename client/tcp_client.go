package client

import (
	"fmt"
	"net"
)

type TcpChatClient struct {
	conn net.Conn
}

func (c *TcpChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	fmt.Println("Conn", conn)
	fmt.Println("Error", err)
	if err == nil {
		c.conn = conn
	}
	return err
}
