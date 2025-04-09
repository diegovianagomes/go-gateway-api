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

// CreateAccount creates a new account and validates duplicate API Key
// Returns ErrDuplicatedAPIKey if the key already exists
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

// UpdateBalance updates the balance of an account in a thread-safe way
// The amount can be positive (credit)
func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil
}

// FindByAPIKey searches for an account by API Key
func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return output, nil

}

// FindByID searches for an account by ID
func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return output, nil
}
