package handlers

import "net/http"

func CardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get card by id
	case http.MethodPost:
		// create new card
	case http.MethodPut:
		// update card
	case http.MethodDelete:
		// delete card
	}

}