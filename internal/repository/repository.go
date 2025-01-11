package repository

import (
	"errors"

	"github.com/mlodovico/digital-wallet/internal/entities"
)

var cards = []entities.Card{
    *entities.NewCard("John Doe", "4111111111111111", 100.0, 1, 2024),
    *entities.NewCard("Jane Smith", "5500000000000004", 200.0, 6, 2023),
    *entities.NewCard("Alice Johnson", "340000000000009", 300.0, 12, 2025),
    *entities.NewCard("Bob Brown", "30000000000004", 400.0, 9, 2022),
}

var wallets = []entities.Wallet{
    *entities.NewWallet(1, "John Doe", cards),
}

func GetAllWallets() []entities.Wallet {
    return wallets
}

func GetWalletByID(id string) (*entities.Wallet, error) {
    for _, wallet := range wallets {
        if wallet.ID == id {
            return &wallet, nil
        }
    }
    return nil, errors.New("wallet not found")
}

func CreateWallet(wallet entities.Wallet) {
    wallets = append(wallets, wallet)
}

func UpdateWallet(wallet entities.Wallet) error {
    for i, w := range wallets {
        if w.ID == wallet.ID {
            wallets[i] = wallet
            return nil
        }
    }
    return errors.New("wallet not found")
}

func DeleteWallet(id string) error {
    for i, w := range wallets {
        if w.ID == id {
            wallets = append(wallets[:i], wallets[i+1:]...)
            return nil
        }
    }
    return errors.New("wallet not found")
}

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