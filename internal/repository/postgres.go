package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func NewDBConnection(ctx context.Context) (*pgx.Conn, error) {
	// 1. Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using environment variables.")
	}

	// 2. Get the database URL from an environment variable.
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// 3. Connect to the database using the pgx driver.
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// 4. Ping the database to ensure the connection is active.
	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return conn, nil
}
