package domain

import "github.com/tutysara/banking-go/dto"

const WITHDRAWAL string = "withdrawal"

type Transaction struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t Transaction) IsWithDrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	} else {
		return false
	}
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId: t.TransactionId,
		Amount:        t.Amount,
	}
}
