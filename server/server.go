package server

import (
	"context"
	"fmt"
	"log"

	"github.com/nerdgarten/mock-payment-service/data"
	pb "github.com/nerdgarten/mock-payment-service/proto"
)

// PaymentServer implements the PaymentServiceServer interface
type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
}

// NewPaymentServer creates a new PaymentServer instance
func NewPaymentServer() *PaymentServer {
	return &PaymentServer{}
}

// CreateCustomer creates a new mock customer
func (s *PaymentServer) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	log.Printf("CreateCustomer called with name: %s, email: %s", req.Name, req.Email)
	customer := data.CreateMockCustomer(req.Name, req.Email)
	return &pb.CreateCustomerResponse{Customer: customer}, nil
}

// RetrieveCustomer retrieves a customer by ID
func (s *PaymentServer) RetrieveCustomer(ctx context.Context, req *pb.RetrieveCustomerRequest) (*pb.RetrieveCustomerResponse, error) {
	log.Printf("RetrieveCustomer called with id: %s", req.Id)
	customer := data.GetMockCustomer(req.Id)
	if customer == nil {
		return nil, fmt.Errorf("customer not found: %s", req.Id)
	}
	return &pb.RetrieveCustomerResponse{Customer: customer}, nil
}

// CreatePaymentIntent creates a new mock payment intent
func (s *PaymentServer) CreatePaymentIntent(ctx context.Context, req *pb.CreatePaymentIntentRequest) (*pb.CreatePaymentIntentResponse, error) {
	log.Printf("CreatePaymentIntent called with amount: %d, currency: %s, payment_method: %s", req.Amount, req.Currency, req.PaymentMethod)
	intent := data.CreateMockPaymentIntent(req.Amount, req.Currency, req.PaymentMethod, req.Description)
	return &pb.CreatePaymentIntentResponse{PaymentIntent: intent}, nil
}

// ConfirmPaymentIntent confirms a payment intent
func (s *PaymentServer) ConfirmPaymentIntent(ctx context.Context, req *pb.ConfirmPaymentIntentRequest) (*pb.ConfirmPaymentIntentResponse, error) {
	log.Printf("ConfirmPaymentIntent called with id: %s", req.Id)
	intent, charges := data.ConfirmMockPaymentIntent(req.Id)
	if intent == nil {
		return nil, fmt.Errorf("payment intent not found: %s", req.Id)
	}
	return &pb.ConfirmPaymentIntentResponse{
		PaymentIntent: intent,
		Charges:       charges,
	}, nil
}

// CreateRefund creates a new mock refund
func (s *PaymentServer) CreateRefund(ctx context.Context, req *pb.CreateRefundRequest) (*pb.CreateRefundResponse, error) {
	log.Printf("CreateRefund called with payment_intent: %s, amount: %d", req.PaymentIntent, req.Amount)
	refund := data.CreateMockRefund(req.PaymentIntent, req.Amount)
	return &pb.CreateRefundResponse{Refund: refund}, nil
}

// TestWebhook simulates webhook delivery
func (s *PaymentServer) TestWebhook(ctx context.Context, req *pb.TestWebhookRequest) (*pb.TestWebhookResponse, error) {
	log.Printf("TestWebhook called with type: %s", req.Type)
	// In a real implementation, this would trigger webhook delivery
	// For mock purposes, we just acknowledge receipt
	return &pb.TestWebhookResponse{Received: true}, nil
}
