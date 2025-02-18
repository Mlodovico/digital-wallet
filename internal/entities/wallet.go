package entities

import (
	"github.com/google/uuid"
)

type Wallet struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    UserID int    `json:"user_id"`
    Cards  []Card `json:"cards"`
}

func NewWallet(userID int, name string, cards []Card) *Wallet {
    return &Wallet{
        ID:     uuid.New().String(),
        Name:   name,
        UserID: userID,
        Cards:  cards,
    }
}