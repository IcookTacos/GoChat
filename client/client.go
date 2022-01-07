package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/IcookTacos/GoChat/server"
)

type User struct {
	name string
}

func SetupClient() {
	reader := bufio.NewReader(os.Stdin)
	server.InitializeServer()
	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	user := User{name}
	fmt.Printf("Welcome %s ", user.name)
}
