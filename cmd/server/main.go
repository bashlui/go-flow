package server

import (
	"context"
	"log"
	"time"

	"go-flow/internal/repository"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := repository.NewDBConnection(ctx)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer conn.Close(ctx)

	// Test a simple query
	var version string
	err = conn.QueryRow(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Printf("Database version: %s", version)
	log.Println("Database connection test successful!")
}
