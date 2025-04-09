package dto

import (
	"time"

	"github.com/diegovianagomes/go-gateway-api/internal/domain"
)

// Account represents a new user account
type CreateAccountInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

// When return the account data
type AccountOutput struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Balance float64 `json:"balance"`
	APIKey string `json:"api_key, omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// When user accout is created return a domain object with account input "Name", "Email"
func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

// When I have the domain object I want to transform it into DTO format
func FromAccount(account *domain.Account) *AccountOutput {
	return &AccountOutput{
		ID: 		account.ID,
		Name: 		account.Name,
		Email: 		account.Email,
		Balance: 	account.Balance,
		APIKey: 	account.APIKey,
		CreatedAt:	account.CreatedAt,
		UpdatedAt:	account.UpdatedAt,
	}
}

