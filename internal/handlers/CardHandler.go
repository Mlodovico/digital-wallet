package handlers

import "net/http"

func CardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Get all cards
	case http.MethodPost:
		// Create a new card
	case http.MethodPut:
		// Update a card
	case http.MethodDelete:
		// Delete a card
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}