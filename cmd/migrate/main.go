package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"go-flow/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Create database connection
	ctx := context.Background()
	conn, err := repository.NewDBConnection(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close(ctx)

	// Read and execute the migration file
	migrationFile := "db/migrations/000001_create_initial_tables.up.sql"

	// Read the SQL file
	file, err := os.Open(migrationFile)
	if err != nil {
		log.Fatal("Failed to open migration file:", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	// Execute the SQL
	_, err = conn.Exec(ctx, string(content))
	if err != nil {
		log.Fatal("Failed to execute migration:", err)
	}

	fmt.Println("Migration executed successfully!")
}
