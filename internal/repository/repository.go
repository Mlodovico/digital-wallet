package repository

import (
	"errors"

	"github.com/mlodovico/digital-wallet/internal/models"
)

var wallets = []models.Wallet{
    {ID: 1, Balance: 100.0},
    {ID: 2, Balance: 200.0},
}

func GetAllWallets() []models.Wallet {
    return wallets
}

func GetWalletByID(id int) (*models.Wallet, error) {
    for _, wallet := range wallets {
        if wallet.ID == id {
            return &wallet, nil
        }
    }
    return nil, errors.New("wallet not found")
}

func CreateWallet(wallet models.Wallet) {
    wallets = append(wallets, wallet)
}

func UpdateWallet(wallet models.Wallet) error {
    for i, w := range wallets {
        if w.ID == wallet.ID {
            wallets[i] = wallet
            return nil
        }
    }
    return errors.New("wallet not found")
}

func DeleteWallet(id int) error {
    for i, w := range wallets {
        if w.ID == id {
            wallets = append(wallets[:i], wallets[i+1:]...)
            return nil
        }
    }
    return errors.New("wallet not found")
}

func DepositToWallet(id int, amount float64) error {
    wallet, err := GetWalletByID(id)
    if err != nil {
        return err
    }
    wallet.Balance += amount
    return UpdateWallet(*wallet)
}

func WithdrawFromWallet(id int, amount float64) error {
    wallet, err := GetWalletByID(id)
    if err != nil {
        return err
    }
    if wallet.Balance < amount {
        return errors.New("insufficient funds")
    }
    wallet.Balance -= amount
    return UpdateWallet(*wallet)
}