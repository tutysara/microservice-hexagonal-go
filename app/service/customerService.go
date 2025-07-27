package service

import (
	"github.com/tutysara/banking-go/app/domain"
	"github.com/tutysara/banking-go/errs"
)

// primary port
// TODO: should this be inside the domain?
// yes this is indise the domain/business logic side, I guess for service a separate package is used
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer() (*domain.Customer, *errs.AppError)
}

// business logic implementing primary port
// depends on repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id) // primary port is connected to seconday port
}

// helper function to create new customer service
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
