package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/entities"
	"github.com/mlodovico/digital-wallet/internal/repository"
)

func WalletHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        id := r.URL.Query().Get("id")
        if id != "" {
            wallet, err := repository.GetWalletByID(id)
            if err != nil {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
            }
            json.NewEncoder(w).Encode(wallet)
        } else {
            wallets, err := repository.GetAllWallets()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(wallets)
        }
    case http.MethodPost:
        var data entities.Card
        json.NewDecoder(r.Body).Decode(&data)

        card := entities.NewCard(data.CompletedName, data.CardNumber, data.PaymentCardType, data.Balance, data.ExpMonth, data.ExpYear)
        wallet := entities.NewWallet(1, "John Doe", []entities.Card{*card})

        if !card.IsCardValid() {
            http.Error(w, "invalid card", http.StatusBadRequest)
            return
        }

        repository.CreateWallet(*wallet)
        w.WriteHeader(http.StatusCreated)
    case http.MethodPut:
        var data entities.Wallet
        json.NewDecoder(r.Body).Decode(&data)

        err := repository.UpdateWallet(data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusOK)
    case http.MethodDelete:
        idStr := r.URL.Query().Get("id")
  
        err := repository.DeleteWallet(idStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusOK)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}