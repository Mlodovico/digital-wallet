package entities

import (
	"time"
	"unicode"

	"github.com/google/uuid"
)

type Card struct {
    ID    string    `json:"id"`
    CompletedName string `json:"completed_name"`
    CardNumber string `json:"card_number"`
    PaymentCardType string `json:"payment_card_type"`
    Balance  float64 `json:"balance"`
    ExpMonth int `json:"exp_month"`
    ExpYear int `json:"exp_year"`
}

func NewCard(completedName string, cardNumber string, PaymentCardType string, balance float64, expMonth, expYear int) *Card {
	return &Card{
		ID: uuid.New().String(),
		CompletedName: completedName,
		CardNumber: cardNumber,
        PaymentCardType: PaymentCardType,
		Balance: balance,
		ExpMonth: expMonth,
		ExpYear: expYear,
	}
}

func (c *Card) IsCardValid() bool {
    return c.CreditCardNumberValid() || c.IsValidCardType() || !c.IsExpired() || c.IsNumeric()
}

func (c *Card) CreditCardNumberValid() bool {
    var sum int
    alt := false

    for i := len(c.CardNumber) - 1; i > -1; i-- {
        n := int(c.CardNumber[i] - '0')

        if alt {
            n *= 2

            if n > 9 {
                n -= 9
            }
        }

        sum += n
        alt = !alt
    }

    return sum%10 == 0
}

func (c *Card) IsValidCardType() bool {
    if len(c.CardNumber) < 1 {
        return false
    }

    switch c.CardNumber[0] {
        case '4':
            return true
        case '5':
            return true
        default:
            return false
    }
}

func (c *Card) IsExpired() bool {
    currentYear, currentMounth, _ := time.Now().Date()

    if c.ExpYear < currentYear || (c.ExpYear == currentYear && c.ExpMonth < int(currentMounth)) {
        return true
    }

    return false
}

func (c *Card) IsNumeric() bool {
    for _, character := range c.CardNumber {
        if !unicode.IsDigit(character) {
            return false
        }
    }

    return true
}