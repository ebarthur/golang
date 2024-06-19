package main

import (
	"fmt"
	"github.com/ebarthur/go-app-structure/internal/routes"
	"net/http"
)

func main() {

	router := routes.NewRouter()

	server := http.Server{
		Addr:    ":10000",
		Handler: router,
	}

	fmt.Println("Server is running at", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
