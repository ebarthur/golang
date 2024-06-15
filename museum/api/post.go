package api

import (
	"encoding/json"
	"github.com/ebarthur/golang/museum/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {

	// specify request method
	if r.Method == http.MethodPost {

		// create a new exhibition item
		newExhibition := data.Exhibition{}

		// decode the request body into the new exhibition item
		err := json.NewDecoder(r.Body).Decode(&newExhibition)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// add the new exhibition item to the list of exhibitions
		data.Add(newExhibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Exhibition added successfully"))

	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
