package dto

import (
	"strings"

	"github.com/tutysara/banking-go/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (req NewAccountRequest) Validate() *errs.AppError { // LRN:TODO: kept as pointer to return nil
	if req.Amount < 5000 {
		return errs.NewValidationError("Amount must be atleast 5000")
	}

	if strings.ToLower(req.AccountType) != "savings" && strings.ToLower(req.AccountType) != "checking" {
		return errs.NewValidationError("AccountType should be savings or checking")
	}
	return nil
}
