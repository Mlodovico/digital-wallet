package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mlodovico/digital-wallet/internal/models"
	"github.com/mlodovico/digital-wallet/internal/repository"
)

func WalletHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        wallets := repository.GetAllWallets()
        json.NewEncoder(w).Encode(wallets)
    case http.MethodPost:
        var wallet models.Wallet
        json.NewDecoder(r.Body).Decode(&wallet)
        repository.CreateWallet(wallet)
        w.WriteHeader(http.StatusCreated)
    case http.MethodPut:
        var wallet models.Wallet
        json.NewDecoder(r.Body).Decode(&wallet)
        err := repository.UpdateWallet(wallet)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusOK)
    case http.MethodDelete:
        idStr := r.URL.Query().Get("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "invalid id", http.StatusBadRequest)
            return
        }
        err = repository.DeleteWallet(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusOK)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}