package repository

import (
	"context"
	"go-flow/internal/models"
)

// StockRepository defines the interface for stock data operations
type StockRepository interface {
	// Basic stock operations
	GetStock(ctx context.Context, symbol string) (*models.Stock, error)
	GetStocks(ctx context.Context, symbols []string) ([]*models.Stock, error)
	CreateStock(ctx context.Context, stock *models.Stock) error
	UpdateStock(ctx context.Context, stock *models.Stock) error
	DeleteStock(ctx context.Context, symbol string) error
	ListStocks(ctx context.Context, limit, offset int) ([]*models.Stock, error)

	// Historical data operations
	GetStockHistory(ctx context.Context, symbol string, limit int) ([]*models.StockHistoryEntry, error)
	GetStockHistoryByDateRange(ctx context.Context, symbol string, startDate, endDate string) ([]*models.StockHistoryEntry, error)
	CreateStockHistoryEntry(ctx context.Context, entry *models.StockHistoryEntry) error
	CreateStockHistoryBatch(ctx context.Context, symbol string, entries []*models.StockHistoryEntry) error

	// Watchlist operations
	GetUserWatchlist(ctx context.Context, userID int64) ([]*models.StockWatchlist, error)
	AddToWatchlist(ctx context.Context, watchlist *models.StockWatchlist) error
	RemoveFromWatchlist(ctx context.Context, userID int64, symbol string) error
	IsInWatchlist(ctx context.Context, userID int64, symbol string) (bool, error)

	// Alert operations
	GetUserAlerts(ctx context.Context, userID int64) ([]*models.StockAlert, error)
	GetActiveAlerts(ctx context.Context) ([]*models.StockAlert, error)
	CreateAlert(ctx context.Context, alert *models.StockAlert) error
	UpdateAlert(ctx context.Context, alert *models.StockAlert) error
	DeleteAlert(ctx context.Context, alertID int64) error
	TriggerAlert(ctx context.Context, alertID int64) error

	// Portfolio operations
	GetUserPortfolios(ctx context.Context, userID int64) ([]*models.Portfolio, error)
	GetPortfolio(ctx context.Context, portfolioID int64) (*models.Portfolio, error)
	CreatePortfolio(ctx context.Context, portfolio *models.Portfolio) error
	UpdatePortfolio(ctx context.Context, portfolio *models.Portfolio) error
	DeletePortfolio(ctx context.Context, portfolioID int64) error

	// Position operations
	GetPortfolioPositions(ctx context.Context, portfolioID int64) ([]*models.PortfolioPosition, error)
	GetPosition(ctx context.Context, portfolioID int64, symbol string) (*models.PortfolioPosition, error)
	CreatePosition(ctx context.Context, position *models.PortfolioPosition) error
	UpdatePosition(ctx context.Context, position *models.PortfolioPosition) error
	DeletePosition(ctx context.Context, positionID int64) error

	// Transaction operations
	GetPortfolioTransactions(ctx context.Context, portfolioID int64, limit, offset int) ([]*models.Transaction, error)
	GetUserTransactions(ctx context.Context, userID int64, limit, offset int) ([]*models.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	GetTransaction(ctx context.Context, transactionID int64) (*models.Transaction, error)

	// Stock screening/filtering
	ScreenStocks(ctx context.Context, request *models.StockScreenerRequest) ([]*models.Stock, error)
	GetTopGainers(ctx context.Context, limit int) ([]*models.Stock, error)
	GetTopLosers(ctx context.Context, limit int) ([]*models.Stock, error)
	GetMostActive(ctx context.Context, limit int) ([]*models.Stock, error)

	// Market data aggregations
	GetSectorPerformance(ctx context.Context) (map[string]float64, error)
	GetMarketSummary(ctx context.Context) (map[string]interface{}, error)
}
