package main

import (
	"context"
	"log"
	"os"

	pb "github.com/nerdgarten/mock-payment-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	// Example 1: Create Customer
	createCustomerReq := &pb.CreateCustomerRequest{
		Name:  "Ruff",
		Email: "ruff@example.com",
	}
	createCustomerResp, err := client.CreateCustomer(context.Background(), createCustomerReq)
	if err != nil {
		log.Printf("CreateCustomer failed: %v", err)
	} else {
		log.Printf("Created customer: %s (%s)", createCustomerResp.Customer.Name, createCustomerResp.Customer.Id)
	}

	// Example 2: Retrieve Customer
	retrieveCustomerReq := &pb.RetrieveCustomerRequest{
		Id: "cus_mock_12345",
	}
	retrieveCustomerResp, err := client.RetrieveCustomer(context.Background(), retrieveCustomerReq)
	if err != nil {
		log.Printf("RetrieveCustomer failed: %v", err)
	} else {
		log.Printf("Retrieved customer: %s (%s)", retrieveCustomerResp.Customer.Name, retrieveCustomerResp.Customer.Email)
	}

	// Example 3: Create Payment Intent
	createIntentReq := &pb.CreatePaymentIntentRequest{
		Amount:        1200,
		Currency:      "thb",
		PaymentMethod: "pm_mock_visa",
		Description:   "Food delivery payment",
	}
	createIntentResp, err := client.CreatePaymentIntent(context.Background(), createIntentReq)
	if err != nil {
		log.Printf("CreatePaymentIntent failed: %v", err)
	} else {
		log.Printf("Created payment intent: %s, amount: %d %s", createIntentResp.PaymentIntent.Id, createIntentResp.PaymentIntent.Amount, createIntentResp.PaymentIntent.Currency)
	}

	// Example 4: Confirm Payment Intent
	confirmIntentReq := &pb.ConfirmPaymentIntentRequest{
		Id: "pi_mock_98765",
	}
	confirmIntentResp, err := client.ConfirmPaymentIntent(context.Background(), confirmIntentReq)
	if err != nil {
		log.Printf("ConfirmPaymentIntent failed: %v", err)
	} else {
		log.Printf("Confirmed payment intent: %s, status: %s", confirmIntentResp.PaymentIntent.Id, confirmIntentResp.PaymentIntent.Status)
		if len(confirmIntentResp.Charges.Data) > 0 {
			log.Printf("Charge: %s, amount: %d", confirmIntentResp.Charges.Data[0].Id, confirmIntentResp.Charges.Data[0].Amount)
		}
	}

	// Example 5: Create Refund
	createRefundReq := &pb.CreateRefundRequest{
		PaymentIntent: "pi_mock_98765",
		Amount:        600,
	}
	createRefundResp, err := client.CreateRefund(context.Background(), createRefundReq)
	if err != nil {
		log.Printf("CreateRefund failed: %v", err)
	} else {
		log.Printf("Created refund: %s, amount: %d %s", createRefundResp.Refund.Id, createRefundResp.Refund.Amount, createRefundResp.Refund.Currency)
	}

	// Example 6: Test Webhook
	testWebhookReq := &pb.TestWebhookRequest{
		Type: "payment_intent.succeeded",
		Data: `{"object": {"id": "pi_mock_98765", "amount": 1200, "currency": "thb"}}`,
	}
	testWebhookResp, err := client.TestWebhook(context.Background(), testWebhookReq)
	if err != nil {
		log.Printf("TestWebhook failed: %v", err)
	} else {
		log.Printf("Webhook test received: %t", testWebhookResp.Received)
	}
}
