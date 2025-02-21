package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/repository"
)

func CardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")

		if id != "" {
			card, err := repository.GetCardByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(card)
		} else {
			// get all cards
		}
	case http.MethodPost:
		// create new card
	case http.MethodPut:
		// update card
	case http.MethodDelete:
		// delete card
	}

}