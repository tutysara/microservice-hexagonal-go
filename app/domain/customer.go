package domain

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
}
