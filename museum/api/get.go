package api

import (
	"encoding/json"
	d "github.com/ebarthur/golang/museum/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// api/exhibitions?id=1
	id := r.URL.Query()["id"]

	// specify request method
	if r.Method == "GET" {

		// if url string contains id
		if id != nil {

			// convert id to int
			idx, err := strconv.Atoi(id[0])

			// if there is no err and the requested id is in the range of the exhibitions
			if err == nil && idx <= len(d.GetAll()) {

				// return exhibition with the requested id
				json.NewEncoder(w).Encode(d.GetAll()[idx])

				// handle errors
			} else {
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
			}
			// if url string does not contain id
		} else {
			json.NewEncoder(w).Encode(d.GetAll())
		}
	} else {
		// handle invalid request method to this endpoint
		http.Error(w, "Invalid request for this endpoint", http.StatusBadRequest)
	}
}
