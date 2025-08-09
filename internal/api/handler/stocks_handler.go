package handler

import (
	"go-flow/internal/repository"
	"go-flow/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StocksHandler struct {
	stockRepo repository.StockRepository
	avService *service.AlphaVantageService
}

func NewStocksHandler(repo repository.StockRepository, avService *service.AlphaVantageService) *StocksHandler {
	return &StocksHandler{
		stockRepo: repo,
		avService: avService,
	}
}

// GetStocks returns all stocks from the database
func (h *StocksHandler) GetStocks(c *gin.Context) {
	stocks, err := h.stockRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stocks"})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// GetStockByID returns a specific stock by ID
func (h *StocksHandler) GetStockByID(c *gin.Context) {
	id := c.Param("id")
	stock, err := h.stockRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

// FetchStockData gets data from Alpha Vantage and stores it
func (h *StocksHandler) FetchStockData(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symbol is required"})
		return
	}

	// Get data from Alpha Vantage
	stockData, err := h.avService.GetDailyStockData(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock data: " + err.Error()})
		return
	}

	// Store in database
	if err := h.stockRepo.SaveStockData(stockData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save stock data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetched and stored stock data",
		"count":   len(stockData),
	})
}
