package server

import (
	"fmt"
)

var port int = 8080

func InitializeServer() {
	fmt.Printf("Initializing server on %d \n", port)
}

func send(user string, message string) {
	fmt.Println(message)
}
