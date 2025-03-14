package repository

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/mlodovico/digital-wallet/internal/entities"
)

func DepositToCard(walletId string, cardNumber string, amount float64) error {
    wallet, err := GetWalletByID(walletId)
    if err != nil {
        return err
    }

    for i, card := range wallet.Cards {
        if card.CardNumber == cardNumber {
            wallet.Cards[i].Balance += amount
            return UpdateWallet(*wallet)
        }
    }

    return errors.New("card not found")
}

func CreateNewCard(walletId string, card entities.Card) {
    wallet, err := GetWalletByID(walletId)

    if err != nil {
        return
    }

    if wallet.Cards == nil {
        wallet.Cards = []entities.Card{}
    }

    card.ID = uuid.New().String()
    wallet.Cards = append(wallet.Cards, card)
    return
}

func GetCardByID(id string) (*entities.Card, error) {
    resp, err := http.Get(jsonServerURL + id)
    if err != nil {
        return nil, err
    }
    
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("card not found")
    }

    var card entities.Card
    if err := json.NewDecoder(resp.Body).Decode(&card); err != nil {
        return nil, err
    }

    return nil, errors.New("card not found")
}

func WithdrawFromCard(walletId string, cardNumber string, amount float64) error {
    wallet, err := GetWalletByID(walletId)
    if err != nil {
        return err
    }

    for i, card := range wallet.Cards {
        if card.CardNumber == cardNumber {
            if card.Balance < amount {
                return errors.New("insufficient funds")
            }
            wallet.Cards[i].Balance -= amount
            return UpdateWallet(*wallet)
        }
    }

    return errors.New("card not found")
}