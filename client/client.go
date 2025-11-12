package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nerdgarten/mock-payment-service/types"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	client := &http.Client{Timeout: 5 * time.Second}

	// Example 1: Create Customer
	createCustomerReq := types.CreateCustomerRequest{Name: "Ruff", Email: "ruff@example.com"}
	var createCustomerResp types.CreateCustomerResponse
	if err := postJSON(client, baseURL+"/customers", createCustomerReq, &createCustomerResp); err != nil {
		log.Printf("CreateCustomer failed: %v", err)
	} else {
		log.Printf("Created customer: %s (%s)", createCustomerResp.Customer.Name, createCustomerResp.Customer.ID)
	}

	// Example 2: Retrieve Customer
	var retrieveCustomerResp types.RetrieveCustomerResponse
	if err := getJSON(client, baseURL+"/customers/cus_mock_12345", &retrieveCustomerResp); err != nil {
		log.Printf("RetrieveCustomer failed: %v", err)
	} else {
		log.Printf("Retrieved customer: %s (%s)", retrieveCustomerResp.Customer.Name, retrieveCustomerResp.Customer.Email)
	}

	// Example 3: Create Payment Intent
	createIntentReq := types.CreatePaymentIntentRequest{
		Amount:        1200,
		Currency:      "thb",
		PaymentMethod: "pm_mock_visa",
		Description:   "Food delivery payment",
	}
	var createIntentResp types.CreatePaymentIntentResponse
	if err := postJSON(client, baseURL+"/payment-intents", createIntentReq, &createIntentResp); err != nil {
		log.Printf("CreatePaymentIntent failed: %v", err)
	} else {
		log.Printf("Created payment intent: %s, amount: %d %s", createIntentResp.PaymentIntent.ID, createIntentResp.PaymentIntent.Amount, createIntentResp.PaymentIntent.Currency)
	}

	// Example 4: Confirm Payment Intent
	confirmIntentReq := types.ConfirmPaymentIntentRequest{ID: "pi_mock_98765"}
	var confirmIntentResp types.ConfirmPaymentIntentResponse
	if err := postJSON(client, baseURL+"/payment-intents/confirm", confirmIntentReq, &confirmIntentResp); err != nil {
		log.Printf("ConfirmPaymentIntent failed: %v", err)
	} else {
		log.Printf("Confirmed payment intent: %s, status: %s", confirmIntentResp.PaymentIntent.ID, confirmIntentResp.PaymentIntent.Status)
		if len(confirmIntentResp.Charges.Data) > 0 {
			log.Printf("Charge: %s, amount: %d", confirmIntentResp.Charges.Data[0].ID, confirmIntentResp.Charges.Data[0].Amount)
		}
	}

	// Example 5: Create Refund
	createRefundReq := types.CreateRefundRequest{PaymentIntent: "pi_mock_98765", Amount: 600}
	var createRefundResp types.CreateRefundResponse
	if err := postJSON(client, baseURL+"/refunds", createRefundReq, &createRefundResp); err != nil {
		log.Printf("CreateRefund failed: %v", err)
	} else {
		log.Printf("Created refund: %s, amount: %d %s", createRefundResp.Refund.ID, createRefundResp.Refund.Amount, createRefundResp.Refund.Currency)
	}

	// Example 6: Test Webhook
	testWebhookReq := types.TestWebhookRequest{
		Type: "payment_intent.succeeded",
		Data: `{"object": {"id": "pi_mock_98765", "amount": 1200, "currency": "thb"}}`,
	}
	var testWebhookResp types.TestWebhookResponse
	if err := postJSON(client, baseURL+"/webhooks/test", testWebhookReq, &testWebhookResp); err != nil {
		log.Printf("TestWebhook failed: %v", err)
	} else {
		log.Printf("Webhook test received: %t", testWebhookResp.Received)
	}
}

func postJSON(client *http.Client, url string, payload any, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, out)
}

func getJSON(client *http.Client, url string, out any) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, out)
}

func decodeResponse(resp *http.Response, out any) error {
	if resp.StatusCode >= http.StatusBadRequest {
		data, _ := io.ReadAll(resp.Body)
		var apiErr types.ErrorResponse
		if err := json.Unmarshal(data, &apiErr); err == nil && apiErr.Error != "" {
			return fmt.Errorf("request failed (%d): %s", resp.StatusCode, apiErr.Error)
		}
		return fmt.Errorf("request failed (%d): %s", resp.StatusCode, strings.TrimSpace(string(data)))
	}
	if out == nil {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}
