# Mock Payment Service

A lightweight mock implementation of a payment gateway (similar to Stripe) using gRPC for local testing and development.
This service simulates the behavior of real payment APIs, including customer creation, payment intents, and webhooks — without external network calls or real transactions.

---

## Features

- Create and manage **mock customers** and **payment intents** via gRPC
- Simulate **card charges**, **refunds**, and **payment statuses**
- Support **webhook testing** for payment events
- Return realistic protobuf responses for integration testing
- Easy to run locally or in CI/CD environments

---

## gRPC Methods

### **1. CreateCustomer**
`rpc CreateCustomer(CreateCustomerRequest) returns (CreateCustomerResponse)`

Creates a mock customer.

**Request Fields**
- name: string
- email: string

**Response Fields**
- customer: Customer (id, object, name, email, created)

---

### **2. RetrieveCustomer**
`rpc RetrieveCustomer(RetrieveCustomerRequest) returns (RetrieveCustomerResponse)`

Returns customer details.

**Request Fields**
- id: string

**Response Fields**
- customer: Customer (id, object, name, email, created)

---

### **3. CreatePaymentIntent**
`rpc CreatePaymentIntent(CreatePaymentIntentRequest) returns (CreatePaymentIntentResponse)`

Creates a mock payment intent representing a charge attempt.

**Request Fields**
- amount: int64
- currency: string
- payment_method: string
- description: string

**Response Fields**
- payment_intent: PaymentIntent (id, object, amount, currency, status, client_secret, description, payment_method)

---

### **4. ConfirmPaymentIntent**
`rpc ConfirmPaymentIntent(ConfirmPaymentIntentRequest) returns (ConfirmPaymentIntentResponse)`

Simulates confirming a payment (e.g., charge a card).

**Request Fields**
- id: string

**Response Fields**
- payment_intent: PaymentIntent
- charges: Charges (data: []Charge)

---

### **5. CreateRefund**
`rpc CreateRefund(CreateRefundRequest) returns (CreateRefundResponse)`

Simulates refunding a completed charge.

**Request Fields**
- payment_intent: string
- amount: int64

**Response Fields**
- refund: Refund (id, object, amount, currency, status, payment_intent)

---

### **6. TestWebhook**
`rpc TestWebhook(TestWebhookRequest) returns (TestWebhookResponse)`

Simulates Stripe-style webhook delivery for testing event listeners.

**Request Fields**
- type: string
- data: string (JSON string)

**Response Fields**
- received: bool

---

## 🚀 Running the Server

To start the gRPC server:

```bash
go run main.go
```

The server will listen on port 50051 by default. You can set the `PORT` environment variable to use a different port:

```bash
PORT=8080 go run main.go
```

## 🧪 Running the Client Example

To run the example client (ensure the server is running first):

```bash
go run client/client.go
```

The client connects to `localhost:50051` by default. If the server is running on a different port, set the `PORT` environment variable:

```bash
PORT=8080 go run client/client.go
```

This will demonstrate various gRPC calls to the mock service.

## 📁 Project Structure

- `proto/`: Contains the protobuf definitions and generated Go code
- `data/`: Mock data definitions and helper functions
- `server/`: gRPC server implementation
- `client/`: Example client for testing the service
- `main.go`: Entry point to start the server
