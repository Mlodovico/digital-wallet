package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mlodovico/digital-wallet/internal/entities"
)

var jsonServerURL = "http://localhost:3000"

func GetAllWallets() ([]entities.Wallet, error) {
    resp, err := http.Get(jsonServerURL + "/wallets")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to fetch wallets")
    }

    var wallets []entities.Wallet
    if err := json.NewDecoder(resp.Body).Decode(&wallets); err != nil {
        return nil, err
    }

    return wallets, nil
}

func GetWalletByID(id string) (*entities.Wallet, error) {
    resp, err := http.Get(jsonServerURL + "/wallets/" + id)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("wallet not found")
    }

    var wallet entities.Wallet
    if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
        return nil, err
    }

    return &wallet, nil
}

func CreateWallet(wallet entities.Wallet) {
    body, err := json.Marshal(wallet)
    if err != nil {
        return
    }

    http.Post(jsonServerURL+"/wallets", "application/json", bytes.NewBuffer(body))
}

func UpdateWallet(wallet entities.Wallet) error {
    body, err := json.Marshal(wallet)
    if err != nil {
        return err
    }

    req, err := http.NewRequest(http.MethodPut, jsonServerURL+"/wallets", bytes.NewBuffer(body))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode != http.StatusOK {
        return errors.New("wallet not found")
    }

    return nil
}

func DeleteWallet(id string) error {
    req, err := http.NewRequest(http.MethodDelete, jsonServerURL+"/wallets/"+id, nil)
    if err != nil {
        return err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("wallet not found")
    }

    return nil
}
