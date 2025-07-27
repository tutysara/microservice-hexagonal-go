package domain

import "github.com/tutysara/banking-go/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

// secondary port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errs.AppError) // to return nil when customer is not found
}
