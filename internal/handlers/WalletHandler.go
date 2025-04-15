package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
        var wallet entities.Wallet
        if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        if wallet.UserID == 0 || wallet.Name == "" || wallet.DocumentID == "" || len(wallet.Cards) == 0 {
            http.Error(w, "missing required fields", http.StatusBadRequest)
            return
        }

        for _, card := range wallet.Cards {
            if card.CompletedName == "" || card.CardNumber == "" || card.PaymentCardType == "" || card.Balance == 0 || card.ExpMonth == 0 || card.ExpYear == 0 {
                http.Error(w, "missing required fields", http.StatusBadRequest)
                return
            }

            if !card.IsCardValid() {
                http.Error(w, "invalid card", http.StatusBadRequest)
                return
            }
        }

        wallet.ID = uuid.New().String()
        repository.CreateWallet(wallet)
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