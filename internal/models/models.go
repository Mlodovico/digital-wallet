package models

type Wallet struct {
    ID    int    `json:"id"`
    CardNumber string `json:"card_number"`
    Balance  float64 `json:"balance"`
    ExpMonth int `json:"exp_month"`
    ExpYear int `json:"exp_year"`
}