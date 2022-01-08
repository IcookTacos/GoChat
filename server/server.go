package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var PORT string

const BUFFER_LENGTH = 1024

var connectedClients []net.Conn

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
		log.Fatal("Could not listen on port")
		return
	}
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
	connectedClients = append(connectedClients, connection)
	log.Printf("Connection accepted: %v", connection.RemoteAddr())
	_, err := io.Copy(connection, connection)
	if err == io.EOF {
		log.Printf("Recieved EOF. Client dissconnected")
		return
	} else if err != nil {
		log.Fatal("copy error")
	}
}

//func listenToConnections() {
//buffer := make([]byte, BUFFER_LENGTH)
//for _, connection := range connectedClients {
//n, err := connection.Read(buffer)
//if err != nil {
//log.Printf("Could not recieve from %v", connection.LocalAddr())
//return
//}
//message := string(buffer[:n])
//fmt.Println(message)
//}
//}

func printConnections(s []net.Conn) {
	log.Printf("connections: %d", len(s))
	log.Print(s)
}
