package entities

import "github.com/google/uuid"

type Wallet struct {
	ID    string    `json:"id"`
    Name string `json:"name"`
	UserId int `json:"user_id"`
    Card [] Card `json:"card"`
}

func NewWallet(id int, Name string, card [] Card) *Wallet {
	return &Wallet{
		ID: uuid.New().String(),
		Name: Name,
		Card: card,
	}
}