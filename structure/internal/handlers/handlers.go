package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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
// eg: curl -H "Authorization: Bearer token" localhost:10000/api/v1/data
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

const URL = "https://catfact.ninja/fact"

// Response represents the structure of the response from the cat fact API
type Response struct {
	Fact string `json:"fact"`
}

// HandleCatFact returns a random fact about cats.
//
// It accesses an external API and returns a random fact.
// This route requires authentication.
//
// @GET /api/v1/fact
//
// eg: curl -H "Authorization: Bearer token" http://localhost:10000/api/v1/fact
func HandleCatFact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := GetCatFact(URL)
		if err != nil {
			http.Error(w, "error while loading fact", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// GetCatFact fetches a random cat fact from the provided URL
func GetCatFact(url string) (*Response, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while fetching fact: %w", err)
	}
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %w", err)
	}

	var resp Response
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling response: %w", err)
	}

	return &resp, nil
}
