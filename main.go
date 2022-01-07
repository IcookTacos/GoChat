package main

import (
	"github.com/IcookTacos/GoChat/client"
	"github.com/IcookTacos/GoChat/server"
)

func main() {
	client.SetupClient()
	server.InitializeServer()
}
