package domain

import "errors"
// ErrAccountNotFound is returned when an account is not found
// ErrDuplcatedAPIKey is returned when an attempt is made to create an account with a duplicate API key
// ErrInvoiceNotFound is returned when an invoice is not found
// ErrUnauthorizedAcces is returned when an attempt is made to access a resource that the user does not have permission to access
// ErrInvalidAmount is returned when an attempt is made to create an invoice with an invalid amount
// ErrInvalidStatus is returned when an attempt is made to create an invoice with an invalid status
var (
	
	ErrAccountNotFound = errors.New("account not found")
	ErrDuplicatedAPIKey = errors.New("api key already exists")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorizedAcces = errors.New("unauthorized acces")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)