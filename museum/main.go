package main

import (
	"fmt"
	"github.com/ebarthur/golang/museum/api"
	"github.com/ebarthur/golang/museum/data"
	"html/template"
	"net/http"
)

// handleHello is a function that handles the /hello route
func handleHello(w http.ResponseWriter, r *http.Request) {

	// if the request METHOD is GET, we will return a welcome message
	if r.Method == "GET" {
		w.Write([]byte("Welcome to API 1.0.1"))
	}
}

// This handler serves a .tmpl file
func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")

	// handle errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while parsing the template"))
	}

	// render the template with the second exhibition
	err = html.Execute(w, data.GetAll())

	// handle errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while executing the template"))
	}
}

func main() {

	// create a new server mux
	server := http.NewServeMux()

	// create routes
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)

	// render exhibitions
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/create", api.Post)

	// serve static files
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":3333", server)

	if err != nil {
		fmt.Println("Error while opening the server")
	}

}
