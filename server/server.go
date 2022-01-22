package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
)

var PORT string

const BUFFER_LENGTH = 1024

func init() {
	flag.StringVar(&PORT, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if PORT == "" {
		PORT = "8080"
	}

}

func main() {
	log.Printf("Establishing connection...")
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal("Could not listen on port" + PORT)
		return
	}
	log.Printf("listening on port %v", PORT)
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Could not accept connection")
			return
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	log.Printf("Connection accepted: %v", connection.RemoteAddr())
	clientReader := bufio.NewReader(connection)
	for {
		clientMessage, err := clientReader.ReadString('\n')
		if err != nil {
			log.Print("Client disconnected")
			return
		}
		log.Printf("Client message: %v", clientMessage)
	}
}
