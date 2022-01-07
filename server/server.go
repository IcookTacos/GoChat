package server

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func InitializeServer() {
	http.HandleFunc("/", Send)
	http.ListenAndServe(":8080", nil)
}

func Send(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "text %s", request.URL.Path[1:])
}
