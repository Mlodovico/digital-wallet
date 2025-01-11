package repository

import (
	"errors"
)


func DepositToCard(walletId string, cardNumber string, amount float64) error {
    wallet, err := GetWalletByID(walletId)
    if err != nil {
        return err
    }

    for i, card := range wallet.Card {
        if card.CardNumber == cardNumber {
            wallet.Card[i].Balance += amount
            return UpdateWallet(*wallet)
        }
    }

    return errors.New("card not found")
}

func WithdrawFromCard(walletId string, cardNumber string, amount float64) error {
    wallet, err := GetWalletByID(walletId)
    if err != nil {
        return err
    }

    for i, card := range wallet.Card {
        if card.CardNumber == cardNumber {
            if card.Balance < amount {
                return errors.New("insufficient funds")
            }
            wallet.Card[i].Balance -= amount
            return UpdateWallet(*wallet)
        }
    }

    return errors.New("card not found")
}