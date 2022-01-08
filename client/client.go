package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var port string = "8080"
var address string = "localhost:8080"
var userName string

func init() {
	fmt.Println("=======================================")
	fmt.Println("	Welcome to GoChat!					")
	fmt.Println("=======================================\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	userName, _ = reader.ReadString('\n')
	// Removing newline from userName makes formatting easier
	userName = strings.TrimSuffix(userName, "\n")
}

func main() {

	connection, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Could not connect to %v", address)
		log.Fatal("Error: %v", err)
		return
	}

	log.Printf("Sucesfully connected to %v, welcome %v ", address, userName)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		sendMessageToServer(text, connection)
	}
}

func sendMessageToServer(text string, connection net.Conn) {
	buffer := []byte(text)
	connection.Write(buffer)
}

func recieveMessageFromServer() {

}
