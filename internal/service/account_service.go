package service

import (
	"github.com/diegovianagomes/go-gateway-api/internal/domain"
	"github.com/diegovianagomes/go-gateway-api/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

// constructor
func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	// Check if API key already exists
	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey

	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil

}
