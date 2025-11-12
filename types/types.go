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
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
	Description   string `json:"description"`
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
	PaymentIntent string `json:"payment_intent"`
	Amount        int64  `json:"amount"`
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
