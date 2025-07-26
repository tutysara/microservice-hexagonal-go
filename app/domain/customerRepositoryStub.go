package domain

// TODO: I think stub should be outside domain in driven side
// seondary adapter, implements secondary port
type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// helper to create the stub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Saran", "Bang", "x60100", "2020-01-01", "1"},
		{"1002", "Dhana", "Bang", "x60100", "2020-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
