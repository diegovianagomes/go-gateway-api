package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/diegovianagomes/go-gateway-api/internal/service"
	"github.com/diegovianagomes/go-gateway-api/internal/dto"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandle(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler {accountService: accountService}
}

// Create POST/accounts, returns 201 created and 400-500 erros
func (h *AccountHandler) Create (w http.ResponseWriter, r* http.Request) {
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nill {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.accountService.CreateAccount(input)
	if err != nill {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Get processes GET /accounts, requires X-API-Key in header
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API Key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}