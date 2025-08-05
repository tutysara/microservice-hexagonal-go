package dto

import (
	"strings"

	"github.com/tutysara/banking-go/errs"
	"github.com/tutysara/banking-go/logger"
)

type TransactionRequest struct {
	CustomerId      string
	AccountId       string
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

func (req TransactionRequest) Validate() *errs.AppError {
	if req.Amount < 0 {
		return errs.NewValidationError("Amount should be > 0")
	}

	if strings.ToLower(req.TransactionType) != WITHDRAWAL && strings.ToLower(req.TransactionType) != DEPOSIT {
		logger.Info("req.TransactionType==" + req.TransactionType)
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}

	return nil
}

func (req TransactionRequest) IsWithDrawal() bool {
	return req.TransactionType == WITHDRAWAL
}
