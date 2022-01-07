package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var PORT string

func init() {
	flag.StringVar(&PORT, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if PORT == "" {
		PORT = "8080"
	}

}

func main() {
	log.Printf("listening on port %v", PORT)
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal("Cannot listen on port")
		log.Fatal("Error: %s", err)
		return
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Could not accept connection")
			log.Fatal("Error: %s", err)
			return
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	log.Printf("Connection accepted %v", connection.RemoteAddr())
	_, err := io.Copy(connection, connection)
	if err == io.EOF {
		log.Printf("Recieved EOF. Client dissconnected")
		return
	} else if err != nil {
		log.Fatal("copy error")
		log.Fatal("Error: %s", err)
	}

}
