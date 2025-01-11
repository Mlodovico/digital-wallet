package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"unicode"

	"github.com/mlodovico/digital-wallet/internal/models"
	"github.com/mlodovico/digital-wallet/internal/repository"
)

func CreditCardNumberValid(cardNumber string) bool {
    var sum int
    alt := false

    for i := len(cardNumber) - 1; i > -1; i-- {
        n := int(cardNumber[i] - '0')

        if alt {
            n *= 2

            if n > 9 {
                n -= 9
            }
        }

        sum += n
        alt = !alt
    }

    return sum%10 == 0
}

func IsValidCardType(cardNumber string) bool {
    if len(cardNumber) < 1 {
        return false
    }

    switch cardNumber[0] {
        case '4':
            return true
        case '5':
            return true
        default:
            return false
    }
}

func IsExpired(expMonth, expYear int) bool {
    currentYear, currentMounth, _ := time.Now().Date()

    if expYear < currentYear || (expYear == currentYear && expMonth < int(currentMounth)) {
        return true
    }

    return false
}

func IsNumeric(card string) bool {
    for _, character := range card {
        if !unicode.IsDigit(character) {
            return false
        }
    }

    return true
}

func WalletHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        wallets := repository.GetAllWallets()
        json.NewEncoder(w).Encode(wallets)
    case http.MethodPost:
        var wallet models.Wallet
        json.NewDecoder(r.Body).Decode(&wallet)

        if CreditCardNumberValid(wallet.CardNumber) || IsValidCardType(wallet.CardNumber) || IsExpired(wallet.ExpMonth, wallet.ExpYear) {
            http.Error(w, "invalid card", http.StatusBadRequest)
            return
        }

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