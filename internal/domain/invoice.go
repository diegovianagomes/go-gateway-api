package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Status is the status of an invoice
type Status string

// StatusPending is the status of an invoice when it is pending
// StatusApproved is the status of an invoice when it is approved
// StatusRejected is the status of an invoice when it is rejected
const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "Failed"
)

// Invoice is a struct that represents an invoice
// ID is the ID of the invoice
// AccountID is the ID of the account that the invoice belongs to
// Amount is the amount of the invoice
// Status is the status of the invoice
// Description is the description of the invoice
// PaymentType is the payment type of the invoice
// CardLastDigits is the last 4 digits of the credit card used to pay the invoice
// CreatedAt is the time the invoice was created
// UpdatedAt is the time the invoice was last updated
type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// CreditCard is a struct that represents a credit card
// Number is the number of the credit card
// CVV is the CVV of the credit card
// ExpiryMonth is the expiry month of the credit card
// ExpiryYear is the expiry year of the credit card
// CardHolderName is the name of the card holder

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

// NewInvoice creates a new invoice
func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	// len(card.Number) - 4
	// 16 Number of the CC -> 16 - 4 = 12
	// card.Number[12:<Just read here>] -> "XXXX-XXXX-XXXX-1234"
	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// Process processes the invoice
// If the amount is greater than 10000, the invoice is approved
// If the amount is less than 10000, the invoice is rejected

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	return nil
}

// UpdateStatus updates the status of the invoice
func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
