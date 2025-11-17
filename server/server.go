package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/nerdgarten/mock-payment-service/data"
	"github.com/nerdgarten/mock-payment-service/types"
)

// PaymentServer exposes HTTP handlers for the mock payment API.
type PaymentServer struct{}

// NewPaymentServer creates a new PaymentServer instance.
func NewPaymentServer() *PaymentServer {
	return &PaymentServer{}
}

// RegisterRoutes registers HTTP endpoints on the provided mux.
func (s *PaymentServer) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/customers", s.handleCustomers)
	mux.HandleFunc("/customers/", s.handleCustomerByID)
	mux.HandleFunc("/payment-intents", s.handlePaymentIntents)
	mux.HandleFunc("/payment-intents/confirm", s.handleConfirmPaymentIntent)
	mux.HandleFunc("/refunds", s.handleCreateRefund)
	mux.HandleFunc("/webhooks/test", s.handleTestWebhook)

	// New payment gateway endpoints
	mux.HandleFunc("/accounts/", s.handleGetAccount)
	mux.HandleFunc("/deposit", s.handleDeposit)
	mux.HandleFunc("/withdraw", s.handleWithdraw)
	mux.HandleFunc("/refund", s.handleRefund)
	mux.HandleFunc("/process-payment", s.handleProcessPayment)
}

func (s *PaymentServer) handleCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST CreateCustomer called name=%s email=%s", req.Name, req.Email)
	customer := data.CreateMockCustomer(req.Name, req.Email)
	writeJSON(w, http.StatusCreated, types.CreateCustomerResponse{Customer: *customer})
}

func (s *PaymentServer) handleCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "missing customer id")
		return
	}
	log.Printf("REST RetrieveCustomer called id=%s", id)
	customer := data.GetMockCustomer(id)
	if customer == nil {
		writeError(w, http.StatusNotFound, "customer not found")
		return
	}
	writeJSON(w, http.StatusOK, types.RetrieveCustomerResponse{Customer: *customer})
}

func (s *PaymentServer) handlePaymentIntents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.CreatePaymentIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST CreatePaymentIntent called amount=%d currency=%s", req.Amount, req.Currency)
	intent := data.CreateMockPaymentIntent(req.Amount, req.Currency, req.PaymentMethod, req.Description)
	writeJSON(w, http.StatusCreated, types.CreatePaymentIntentResponse{PaymentIntent: *intent})
}

func (s *PaymentServer) handleConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.ConfirmPaymentIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	if strings.TrimSpace(req.ID) == "" {
		writeError(w, http.StatusBadRequest, "payment intent id is required")
		return
	}
	log.Printf("REST ConfirmPaymentIntent called id=%s", req.ID)
	intent, charges := data.ConfirmMockPaymentIntent(req.ID)
	if intent == nil {
		writeError(w, http.StatusNotFound, "payment intent not found")
		return
	}
	resp := types.ConfirmPaymentIntentResponse{PaymentIntent: *intent}
	if charges != nil {
		resp.Charges = *charges
	}
	writeJSON(w, http.StatusOK, resp)
}

func (s *PaymentServer) handleCreateRefund(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.CreateRefundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	if strings.TrimSpace(req.PaymentIntent) == "" {
		writeError(w, http.StatusBadRequest, "payment_intent is required")
		return
	}
	log.Printf("REST CreateRefund called payment_intent=%s amount=%d", req.PaymentIntent, req.Amount)
	refund := data.CreateMockRefund(req.PaymentIntent, req.Amount)
	writeJSON(w, http.StatusCreated, types.CreateRefundResponse{Refund: *refund})
}

func (s *PaymentServer) handleTestWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.TestWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST TestWebhook called type=%s", req.Type)
	writeJSON(w, http.StatusOK, types.TestWebhookResponse{Received: true})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, types.ErrorResponse{Error: message})
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func (s *PaymentServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	paymentTypeStr := strings.TrimPrefix(r.URL.Path, "/accounts/")
	if paymentTypeStr == "" {
		writeError(w, http.StatusBadRequest, "missing payment type")
		return
	}

	var paymentType types.PaymentType
	switch paymentTypeStr {
	case "cash":
		paymentType = types.PaymentTypeCash
	case "mobilebanking":
		paymentType = types.PaymentTypeMobileBanking
	case "creditcard":
		paymentType = types.PaymentTypeCreditCard
	case "meowth-wallet":
		paymentType = types.PaymentTypeMeowthWallet
	default:
		writeError(w, http.StatusBadRequest, "invalid payment type")
		return
	}

	log.Printf("REST GetAccount called type=%s", paymentType)
	account := data.GetAccount(paymentType)
	if account == nil {
		writeError(w, http.StatusNotFound, "account not found")
		return
	}
	writeJSON(w, http.StatusOK, account)
}

func (s *PaymentServer) handleDeposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST Deposit called type=%s amount=%d", req.Type, req.Amount)
	result := data.Deposit(req.Type, req.Amount)
	writeJSON(w, http.StatusOK, result)
}

func (s *PaymentServer) handleWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST Withdraw called type=%s amount=%d", req.Type, req.Amount)
	result := data.Withdraw(req.Type, req.Amount)
	writeJSON(w, http.StatusOK, result)
}

func (s *PaymentServer) handleRefund(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.RefundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST Refund called type=%s amount=%d reference=%s", req.Type, req.Amount, req.ReferenceID)
	result := data.Refund(req.Type, req.Amount, req.ReferenceID)
	writeJSON(w, http.StatusOK, result)
}

func (s *PaymentServer) handleProcessPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	var req types.ProcessPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	log.Printf("REST ProcessPayment called type=%s amount=%d orderID=%s", req.Type, req.Amount, req.OrderID)
	result := data.ProcessPayment(req.Type, req.Amount, req.OrderID)
	writeJSON(w, http.StatusOK, result)
}
