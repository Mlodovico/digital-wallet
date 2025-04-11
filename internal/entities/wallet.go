package entities

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    UserID int    `json:"user_id"`
    DocumentID string `json:"document_id"`
    BirthDate string `json:"birth_date"`
    Cards  []Card `json:"cards"`
}

func NewWallet(userID int, name string, documentId string, birthDate time.Time, cards []Card) *Wallet {
    return &Wallet{
        ID:     uuid.New().String(),
        Name:   name,
        UserID: userID,
        DocumentID: documentId,
        BirthDate: birthDate.Format("2006-01-02"),
        Cards:  cards,
    }
}

func (w *Wallet) IsDocumentIDValid() bool {
    if (len(w.DocumentID ) < 1) {
        return false
    }
    for _, char := range w.DocumentID {
        if char < '0' || char > '9' {
            return false
        }
    }
    return true
}
