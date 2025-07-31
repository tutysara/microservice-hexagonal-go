package service

import (
	"time"

	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/dto"
	"github.com/tutysara/banking-go/errs"
)

// primary port, service interface -- used by handlers and other drivers
type AccountService interface {
	NewAccount(a dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// primary port adapter
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{ // TODO: LRN: why is DTO to Domain not in domain package? Guess DTO->Domain in service, Domain->DTO in Domain package
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
