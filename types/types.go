package types

// Customer represents a mock customer record.
type Customer struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created int64  `json:"created"`
}

// CreateCustomerRequest is the payload used to create a customer.
type CreateCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateCustomerResponse wraps the created customer.
type CreateCustomerResponse struct {
	Customer Customer `json:"customer"`
}

// RetrieveCustomerResponse wraps a retrieved customer record.
type RetrieveCustomerResponse struct {
	Customer Customer `json:"customer"`
}

// PaymentIntent models an intent to collect a payment.
type PaymentIntent struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	ClientSecret  string `json:"client_secret"`
	Description   string `json:"description"`
	PaymentMethod string `json:"payment_method"`
}

// CreatePaymentIntentRequest defines the required parameters to create an intent.
type CreatePaymentIntentRequest struct {
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Description   string  `json:"description"`
}

// CreatePaymentIntentResponse wraps the created payment intent.
type CreatePaymentIntentResponse struct {
	PaymentIntent PaymentIntent `json:"payment_intent"`
}

// Charge represents a processed charge linked to a payment intent.
type Charge struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}

// Charges is a collection wrapper used for responses.
type Charges struct {
	Data []Charge `json:"data"`
}

// ConfirmPaymentIntentRequest identifies which intent to confirm.
type ConfirmPaymentIntentRequest struct {
	ID string `json:"id"`
}

// ConfirmPaymentIntentResponse returns the updated intent and associated charges.
type ConfirmPaymentIntentResponse struct {
	PaymentIntent PaymentIntent `json:"payment_intent"`
	Charges       Charges       `json:"charges"`
}

// Refund represents a returned payment.
type Refund struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	PaymentIntent string `json:"payment_intent"`
}

// CreateRefundRequest specifies the refund payload.
type CreateRefundRequest struct {
	PaymentIntent string  `json:"payment_intent"`
	Amount        float64 `json:"amount"`
}

// CreateRefundResponse wraps the mock refund result.
type CreateRefundResponse struct {
	Refund Refund `json:"refund"`
}

// TestWebhookRequest imitates a webhook trigger payload.
type TestWebhookRequest struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// TestWebhookResponse acknowledges webhook delivery.
type TestWebhookResponse struct {
	Received bool `json:"received"`
}

// ErrorResponse standardises error messages returned by the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// PaymentType represents different payment methods
type PaymentType string

const (
	PaymentTypeCash          PaymentType = "cash"
	PaymentTypeMobileBanking PaymentType = "mobilebanking"
	PaymentTypeCreditCard    PaymentType = "creditcard"
	PaymentTypeMeowthWallet  PaymentType = "meowth-wallet"
)

// Account represents a payment account with balance
type Account struct {
	Type    PaymentType `json:"type"`
	Balance float64     `json:"balance"`
}

// DepositRequest represents a deposit request
type DepositRequest struct {
	Type   PaymentType `json:"type"`
	Amount float64     `json:"amount"`
}

// DepositResponse represents a deposit response
type DepositResponse struct {
	Success       bool    `json:"success"`
	TransactionID string  `json:"transaction_id"`
	Message       string  `json:"message"`
	Account       Account `json:"account"`
}

// WithdrawRequest represents a withdrawal request
type WithdrawRequest struct {
	Type   PaymentType `json:"type"`
	Amount float64     `json:"amount"`
}

// WithdrawResponse represents a withdrawal response
type WithdrawResponse struct {
	Success       bool    `json:"success"`
	TransactionID string  `json:"transaction_id"`
	Message       string  `json:"message"`
	Account       Account `json:"account"`
}

// RefundRequest represents a refund request
type RefundRequest struct {
	Type        PaymentType `json:"type"`
	Amount      float64     `json:"amount"`
	ReferenceID string      `json:"reference_id"`
}

// RefundResponse represents a refund response
type RefundResponse struct {
	Success       bool    `json:"success"`
	TransactionID string  `json:"transaction_id"`
	Message       string  `json:"message"`
	Account       Account `json:"account"`
}

// ProcessPaymentRequest represents a payment processing request
type ProcessPaymentRequest struct {
	Type    PaymentType `json:"type"`
	Amount  float64     `json:"amount"`
	OrderID string      `json:"order_id"`
}

// ProcessPaymentResponse represents a payment processing response
type ProcessPaymentResponse struct {
	Success       bool    `json:"success"`
	TransactionID string  `json:"transaction_id"`
	Message       string  `json:"message"`
	OrderID       string  `json:"order_id"`
	Account       Account `json:"account"`
}
