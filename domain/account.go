package domain

import "github.com/tutysara/banking-go/errs"

type Account struct {
	CustomerId  string
	AccountId   string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// seconday port, ports are always interfaces
type AccountRepository interface {
	Save(a Account) (*Account, *errs.AppError) //its a pointer to return nil value
}
