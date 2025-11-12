# Mock Payment Service

A mock Stripe-like REST API for local integration testing. The service exposes HTTP endpoints for customer management, payment intent workflows, refunds, and webhook simulations without talking to external services.

## Features

- Pure HTTP+JSON contract (no gRPC dependencies)
- Deterministic responses ideal for automated tests
- Built-in mock datasets for customers, payment intents, charges, and refunds
- Example Go client demonstrating endpoint usage

## REST Endpoints

| Method | Path                       | Description                                                    |
| ------ | -------------------------- | -------------------------------------------------------------- |
| `POST` | `/customers`               | Create a mock customer.                                        |
| `GET`  | `/customers/{id}`          | Retrieve a customer by ID.                                     |
| `POST` | `/payment-intents`         | Create a mock payment intent.                                  |
| `POST` | `/payment-intents/confirm` | Confirm an existing payment intent and generate a mock charge. |
| `POST` | `/refunds`                 | Create a refund linked to a payment intent.                    |
| `POST` | `/webhooks/test`           | Simulate webhook delivery.                                     |

Requests and responses use the JSON models defined in `types/types.go`. Errors follow the structure:

```json
{
  "error": "description"
}
```

## Running the Server

```bash
go run main.go
```

The service listens on `:50051` by default. Override by setting `PORT`:

```bash
PORT=8080 go run main.go
```

## Sample Requests

### Create Customer

```bash
curl -X POST http://localhost:50051/customers \
	-H "Content-Type: application/json" \
	-d '{"name":"Ruff","email":"ruff@example.com"}'
```

### Confirm Payment Intent

```bash
curl -X POST http://localhost:50051/payment-intents/confirm \
	-H "Content-Type: application/json" \
	-d '{"id":"pi_mock_98765"}'
```

## Project Structure

- `types/` – shared request/response models
- `data/` – mock datasets and helper functions
- `server/` – HTTP handlers and route registration
- `client/` – REST demo client
- `main.go` – server entrypoint

## Example Client

Run the included client to exercise all endpoints:

```bash
go run client/client.go
```

The client targets `http://localhost:50051` (configurable with `PORT`).
