package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const ADDRESS = "localhost"
const PORT = "8080"

var port string = "8080"
var address string = "localhost:8080"

func main() {

	connection, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Could not connect to %v", address)
		log.Fatal("Error: %v", err)
		return
	}

	log.Printf("Connection: ", connection)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(connection, text+"\n")

		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("> " + message)
	}

}
