package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Account represents a user account
type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

// generateAPIKey generates a random API key
func generateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// NewAccount creates a new account
func NewAccount(name, email string) *Account {

	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

// AddBalance adds the given amount to the account balance
func (a *Account) AddBalance(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil // TODO: remove this line after implementing the method
}
