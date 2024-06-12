package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Write([]byte("Welcome to API 1.0.1"))
	}
}

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", handleHello)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Error while opening the server")
	}

}
