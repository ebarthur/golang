package handlers

import (
	"fmt"
	"net/http"
)

// HandlePing returns "pong" to validate that the server works alright
//
// @GET /api/ping eg: curl localhost:10000/api/ping
//
// Response: pong
func HandlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "pong")
		if err != nil {
			panic(err)
		}
	}
}

// HandleData returns "Some data in this API"
//
// # Route requires authentication
//
// @GET /api/v1/data
//
// eg: curl -H "Authorization: Bearer token" localhost:10000/api/ping
//
// Response: "Some data in this API"
func HandleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := "Some data in this API"
		_, err := fmt.Fprintf(w, data)
		if err != nil {
			panic(err)
		}
	}
}
