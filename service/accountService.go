package service

import (
	"time"

	"github.com/tutysara/banking-go/domain"
	"github.com/tutysara/banking-go/dto"
	"github.com/tutysara/banking-go/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

// primary port, service interface -- used by handlers and other drivers
type AccountService interface {
	NewAccount(a dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(t dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	// validate the input data dto
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// validate business condition (account should have balance > amount)
	if req.IsWithDrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Amount should be less than account balance")
		}
	}
	// make the transaction
	// convert from dto to domain and call save function in domain

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToDto()
	return &response, nil

}
