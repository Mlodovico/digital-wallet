package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

            wallet, err := repository.GetWalletByID(id)
            if err != nil {
                http.Error(w, "wallet not found", http.StatusNotFound)
                return
            }

            for _, c := range wallet.Cards {
                if c.CardNumber == card.CardNumber {
                    http.Error(w, "card already exists", http.StatusBadRequest)
                    return
                }
            }

            card.ID = uuid.New().String()

            err = repository.CreateNewCard(id, card)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(card)
            return
		} else {
			http.Error(w, "invalid wallet id", http.StatusBadRequest)
			return
		}

	case http.MethodPut:
		walletId := r.URL.Query().Get("wallet-id")
		cardId := r.URL.Query().Get("card-id")

		if walletId != "" && cardId != "" {
			var udpatedCard entities.Card
			
			if err := json.NewDecoder(r.Body).Decode(&udpatedCard); err != nil {
				http.Error(w, "invalid card data", http.StatusBadRequest)
				return
			}

			updatedCard.ID = cardId

			err := repository.UpdateCard(walletId, updatedCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("card updated with success"))
		} else {
			http.Error(w, "invalid wallet or card id", http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		cardID := r.URL.Query().Get("card-id")
		walletID := r.URL.Query().Get("wallet-id")

		if cardID != "" && walletID != "" {
			err := repository.DeleteCard(walletID, cardID)

			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("card deleted with success"))
		} else {
			http.Error(w, "invalid card number", http.StatusBadRequest)
			return
		}
	}

}