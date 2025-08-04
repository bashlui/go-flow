package server

import (
	"context"
	"log"
	"time"

	"goflow-project/internal/repository"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := repository.NewDBConnection(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())
}
