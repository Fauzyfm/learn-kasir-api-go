package handler

import (
	"encoding/json"
	"fmt"
	"kasir-api/internal/models"
	"kasir-api/internal/service"
	"net/http"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("DEBUG: JSON Decode Error: %v\n", err)
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("DEBUG: Raw request: %+v\n", req)
	fmt.Printf("DEBUG: Request items count: %d\n", len(req.Items))
	for i, item := range req.Items {
		fmt.Printf("DEBUG: Item %d - ProductID: %d, Quantity: %d\n", i, item.ProductID, item.Quantity)
	}

	if len(req.Items) == 0 {
		fmt.Println("DEBUG: Items array is empty!")
		http.Error(w, "Items cannot be empty", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
