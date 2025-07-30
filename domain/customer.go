package domain

import "github.com/tutysara/banking-go/errs"

type Customer struct {
	Id          string `db:"customer_id"` // LRN: TODO: Go compiler doesn't protect from this, runtime error on mismatch between db and dto
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// secondary port
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError) // to return nil when customer is not found
}
