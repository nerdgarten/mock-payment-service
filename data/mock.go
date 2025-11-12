package data

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nerdgarten/mock-payment-service/types"
)

// MockCustomers stores mock customer data
var MockCustomers = map[string]*types.Customer{
	"cus_mock_12345": {
		ID:      "cus_mock_12345",
		Object:  "customer",
		Name:    "Ruff",
		Email:   "ruff@example.com",
		Created: 1734567890,
	},
	"cus_mock_67890": {
		ID:      "cus_mock_67890",
		Object:  "customer",
		Name:    "John Doe",
		Email:   "john@example.com",
		Created: 1734567800,
	},
}

// MockPaymentIntents stores mock payment intents
var MockPaymentIntents = map[string]*types.PaymentIntent{
	"pi_mock_98765": {
		ID:            "pi_mock_98765",
		Object:        "payment_intent",
		Amount:        1200,
		Currency:      "thb",
		Status:        "requires_confirmation",
		ClientSecret:  "pi_mock_98765_secret_abc123",
		Description:   "Food delivery payment",
		PaymentMethod: "pm_mock_visa",
	},
}

// MockCharges stores mock charges
var MockCharges = map[string]*types.Charge{
	"ch_mock_555": {
		ID:            "ch_mock_555",
		Status:        "succeeded",
		Amount:        1200,
		Currency:      "thb",
		PaymentMethod: "pm_mock_visa",
	},
}

// MockRefunds stores mock refunds
var MockRefunds = map[string]*types.Refund{
	"re_mock_444": {
		ID:            "re_mock_444",
		Object:        "refund",
		Amount:        600,
		Currency:      "thb",
		Status:        "succeeded",
		PaymentIntent: "pi_mock_98765",
	},
}

// GenerateCustomerID generates a mock customer ID
func GenerateCustomerID() string {
	return fmt.Sprintf("cus_mock_%d", rand.Intn(100000))
}

// GeneratePaymentIntentID generates a mock payment intent ID
func GeneratePaymentIntentID() string {
	return fmt.Sprintf("pi_mock_%d", rand.Intn(100000))
}

// GenerateChargeID generates a mock charge ID
func GenerateChargeID() string {
	return fmt.Sprintf("ch_mock_%d", rand.Intn(100000))
}

// GenerateRefundID generates a mock refund ID
func GenerateRefundID() string {
	return fmt.Sprintf("re_mock_%d", rand.Intn(100000))
}

// CreateMockCustomer creates a new mock customer
func CreateMockCustomer(name, email string) *types.Customer {
	id := GenerateCustomerID()
	customer := &types.Customer{
		ID:      id,
		Object:  "customer",
		Name:    name,
		Email:   email,
		Created: time.Now().Unix(),
	}
	MockCustomers[id] = customer
	return customer
}

// GetMockCustomer retrieves a customer by ID
func GetMockCustomer(id string) *types.Customer {
	return MockCustomers[id]
}

// CreateMockPaymentIntent creates a new mock payment intent
func CreateMockPaymentIntent(amount int64, currency, paymentMethod, description string) *types.PaymentIntent {
	id := GeneratePaymentIntentID()
	intent := &types.PaymentIntent{
		ID:            id,
		Object:        "payment_intent",
		Amount:        amount,
		Currency:      currency,
		Status:        "requires_confirmation",
		ClientSecret:  fmt.Sprintf("%s_secret_%s", id, generateRandomString(6)),
		Description:   description,
		PaymentMethod: paymentMethod,
	}
	MockPaymentIntents[id] = intent
	return intent
}

// ConfirmMockPaymentIntent confirms a payment intent and creates a charge
func ConfirmMockPaymentIntent(id string) (*types.PaymentIntent, *types.Charges) {
	intent := MockPaymentIntents[id]
	if intent == nil {
		return nil, nil
	}
	intent.Status = "succeeded"
	chargeID := GenerateChargeID()
	charge := &types.Charge{
		ID:            chargeID,
		Status:        "succeeded",
		Amount:        intent.Amount,
		Currency:      intent.Currency,
		PaymentMethod: intent.PaymentMethod,
	}
	MockCharges[chargeID] = charge
	charges := &types.Charges{
		Data: []types.Charge{*charge},
	}
	return intent, charges
}

// CreateMockRefund creates a new mock refund
func CreateMockRefund(paymentIntent string, amount int64) *types.Refund {
	id := GenerateRefundID()
	refund := &types.Refund{
		ID:            id,
		Object:        "refund",
		Amount:        amount,
		Currency:      "thb", // Assuming THB for simplicity
		Status:        "succeeded",
		PaymentIntent: paymentIntent,
	}
	MockRefunds[id] = refund
	return refund
}

// generateRandomString generates a random string of given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
