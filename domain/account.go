package domain

import (
	"github.com/tutysara/banking-go/dto"
	"github.com/tutysara/banking-go/errs"
)

type Account struct {
	CustomerId  string  `db:"customer_id"`
	AccountId   string  `db:"account_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

// seconday port, ports are always interfaces
type AccountRepository interface {
	Save(a Account) (*Account, *errs.AppError) //its a pointer to return nil value
	FindBy(id string) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	naResponse := dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
	return naResponse
}

func (a Account) CanWithdraw(reqAmount float64) bool {
	return a.Amount > reqAmount
}
