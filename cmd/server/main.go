package main

import (
	"context"
	"log"
	"os"

	"go-flow/internal/api/handler"
	"go-flow/internal/repository"
	"go-flow/internal/service"

	"github.com/gin-gonic/gin"
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

	// Initialize services
	avService := service.NewAlphaVantageService(os.Getenv("ALPHA_VANTAGE_API_KEY"))

	// Initialize repository with real database connection
	stockRepo := repository.NewStockRepository(conn)

	// Initialize handlers
	stocksHandler := handler.NewStocksHandler(stockRepo, avService)

	// Set up Gin router
	r := gin.Default()

	// Set up routes
	api := r.Group("/api")
	{
		stocks := api.Group("/stocks")
		{
			stocks.GET("", stocksHandler.GetStocks)
			stocks.GET("/:id", stocksHandler.GetStockByID)
			stocks.POST("/fetch/:symbol", stocksHandler.FetchStockData)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
