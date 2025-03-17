package repository

import (
	"bytes"
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

func updateWallet(wallet entities.Wallet) error {
    body, err := json.Marshal(wallet)

    if err != nil {
        return err
    }

    req, err := http.NewRequest(http.MethodPut, jsonServerURL + "/wallets/" + wallet.ID, bytes.NewBuffer(body))

    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to update wallet")
    }

    return nil
}

func CreateNewCard(walletId string, card entities.Card) error {
    wallet, err := GetWalletByID(walletId)

    if err != nil {
        return err
    }

    if wallet.Cards == nil {
        wallet.Cards = []entities.Card{}
    }

    card.ID = uuid.New().String()
    wallet.Cards = append(wallet.Cards, card)
    return updateWallet(*wallet)
}

func GetCardByID(id string) ([]entities.Card, error) {
    resp, err := GetWalletByID(id)
    if err != nil {
        return nil, err
    }

    return resp.Cards, nil
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