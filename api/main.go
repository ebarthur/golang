package main

import (
	"log"
)

func main() {

	server := NewAPIServer(":4032")
	err := server.Run()

	if err != nil {
		log.Fatal(err)
	}
}
