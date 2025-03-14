package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/entities"
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
			http.Error(w, "invalid card number", http.StatusBadRequest)
			return
		}

	case http.MethodPost:
		id := r.URL.Query().Get("id")

		if id != "" {
			var card entities.Card

			if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
				http.Error(w, "invalid card data", http.StatusBadRequest)
				return
			}

			repository.CreateNewCard(id, card)
			json.NewEncoder(w).Encode(card)
			return
		} else {
			http.Error(w, "invalid wallet id", http.StatusBadRequest)
			return
		}

	case http.MethodPut:
		// update card
	case http.MethodDelete:
		// delete card
	}

}