package server

import (
	"fmt"
)

var port int = 8080

func InitializeServer() {
	fmt.Printf("Initializing server on %d \n", port)
}

func Send(userName string, message string) {
	fmt.Printf("%s:%s", userName, message)
}
