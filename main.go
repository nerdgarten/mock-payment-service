package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nerdgarten/mock-payment-service/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
	}

	mux := http.NewServeMux()
	server.NewPaymentServer().RegisterRoutes(mux)

	addr := ":" + port
	log.Printf("Mock Payment REST server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
