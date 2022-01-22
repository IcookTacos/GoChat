package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var address string = "172.31.81.78:2309"

func main() {
	fmt.Printf("Connecting to %v ... \n",address)
	connection, err := net.Dial("tcp", address)
	clientReader := bufio.NewReader(os.Stdin)
	//serverReader := bufio.NewReader(connection)

	if err != nil {
		log.Fatal("Could not connect to server")
		return
	}

	log.Printf("Sucesfully connected to %v", address)

	for {
		fmt.Print(">> ")
		text, _ := clientReader.ReadString('\n')
		go sendMessageToServer(text, connection)
	}
}

func sendMessageToServer(text string, connection net.Conn) {
	fmt.Println("Sending to server...")
	buffer := []byte(text)
	connection.Write(buffer)
}
