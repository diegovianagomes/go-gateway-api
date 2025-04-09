package repository

import (
	"database/sql"
	"time"

	"github.com/diegovianagomes/go-gateway-api/internal/domain"
)

// implements interface to interact with db
type AccountRepository struct {
	db *sql.DB
}

// FindByID implements domain.AccountRepository.
func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	panic("unimplemented")
}

// constructor
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Method insert data into db
func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(` 
		INSERT INTO accounts(id, name, email, api_key, balance, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Method to find account by API Key
func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var created_at, updated_at time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at 
		FROM accounts 
		WHERE api_key = $1
	`, apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&created_at,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	account.CreatedAt = created_at
	account.UpdatedAt = updated_at
	return &account, nil
}

// Method to find account by ID
func (r *AccountRepository) FindById(id string) (*domain.Account, error) {
	var account domain.Account
	var created_at, updated_at time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&created_at,
		&updated_at,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	account.CreatedAt = created_at
	account.UpdatedAt = updated_at
	return &account, nil
}

// Method to update account balance
func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Block the account row to avoid concurrent updates
	// The FOR UPDATE clause locks the selected rows until the end of the transaction
	var currentBalance float64
	err = tx.QueryRow(`
		SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, account.ID).Scan(&currentBalance)
	// Error no rows means account not found or another error
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	// Update the account balance
	_, err = tx.Exec(`
		UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3
	`, currentBalance+account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}
	return tx.Commit()
	// As soon as you commit, this line is released into the database
}
