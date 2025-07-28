package service

import (
	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/errs"
)

// primary port
// TODO: should this be inside the domain?
// yes this is indise the domain/business logic side, I guess for service a separate package is used
type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer() (*domain.Customer, *errs.AppError)
}

// business logic implementing primary port
// depends on repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id) // primary port is connected to seconday port
}

// helper function to create new customer service
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
