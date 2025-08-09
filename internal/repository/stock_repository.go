package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go-flow/internal/models"
	"go-flow/internal/service"

	"github.com/jackc/pgx/v5"
)

type StockRepository interface {
	GetAll() ([]models.Stock, error)
	GetByID(id string) (*models.Stock, error)
	GetBySymbol(symbol string) (*models.Stock, error)
	SaveStock(stock *models.Stock) error
	SaveStockData(data []service.StockData) error
	SaveStockHistory(entries []models.StockHistoryEntry) error
}

type PostgresStockRepository struct {
	conn *pgx.Conn
}

func NewStockRepository(conn *pgx.Conn) StockRepository {
	return &PostgresStockRepository{
		conn: conn,
	}
}

func (r *PostgresStockRepository) GetAll() ([]models.Stock, error) {
	ctx := context.Background()

	query := `
        SELECT symbol, name, last_price, created_at 
        FROM stocks 
        ORDER BY created_at DESC
    `

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query stocks: %w", err)
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		var lastPrice sql.NullFloat64

		err := rows.Scan(
			&stock.Symbol,
			&stock.Name,
			&lastPrice,
			&stock.LastUpdated,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stock: %w", err)
		}

		if lastPrice.Valid {
			stock.CurrentPrice = lastPrice.Float64
		}

		stocks = append(stocks, stock)
	}

	return stocks, rows.Err()
}

func (r *PostgresStockRepository) GetByID(id string) (*models.Stock, error) {
	// For now, treat ID as symbol since your primary key is symbol
	return r.GetBySymbol(id)
}

func (r *PostgresStockRepository) GetBySymbol(symbol string) (*models.Stock, error) {
	ctx := context.Background()

	query := `
        SELECT symbol, name, last_price, created_at 
        FROM stocks 
        WHERE symbol = $1
    `

	var stock models.Stock
	var lastPrice sql.NullFloat64

	err := r.conn.QueryRow(ctx, query, symbol).Scan(
		&stock.Symbol,
		&stock.Name,
		&lastPrice,
		&stock.LastUpdated,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("stock with symbol %s not found", symbol)
		}
		return nil, fmt.Errorf("failed to get stock: %w", err)
	}

	if lastPrice.Valid {
		stock.CurrentPrice = lastPrice.Float64
	}

	return &stock, nil
}

func (r *PostgresStockRepository) SaveStock(stock *models.Stock) error {
	ctx := context.Background()

	query := `
        INSERT INTO stocks (symbol, name, last_price, created_at) 
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (symbol) 
        DO UPDATE SET 
            name = EXCLUDED.name,
            last_price = EXCLUDED.last_price,
            created_at = EXCLUDED.created_at
    `

	_, err := r.conn.Exec(ctx, query,
		stock.Symbol,
		stock.Name,
		stock.CurrentPrice,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to save stock: %w", err)
	}

	return nil
}

func (r *PostgresStockRepository) SaveStockData(data []service.StockData) error {
	ctx := context.Background()

	// Start a transaction
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// First, ensure the stock exists in the stocks table
	if len(data) > 0 {
		firstEntry := data[0]
		stockQuery := `
            INSERT INTO stocks (symbol, name, last_price, created_at) 
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (symbol) 
            DO UPDATE SET last_price = EXCLUDED.last_price
        `

		_, err = tx.Exec(ctx, stockQuery,
			firstEntry.Symbol,
			firstEntry.Symbol, // Using symbol as name for now
			firstEntry.Close,
			time.Now(),
		)
		if err != nil {
			return fmt.Errorf("failed to insert/update stock: %w", err)
		}
	}

	// Then save the historical data
	historyQuery := `
        INSERT INTO stock_history (symbol, date, open, high, low, close, volume, adj_close) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (symbol, date) 
        DO UPDATE SET 
            open = EXCLUDED.open,
            high = EXCLUDED.high,
            low = EXCLUDED.low,
            close = EXCLUDED.close,
            volume = EXCLUDED.volume,
            adj_close = EXCLUDED.adj_close
    `

	for _, entry := range data {
		// Parse the date string to time.Time
		date, err := time.Parse("2006-01-02", entry.Date)
		if err != nil {
			continue // Skip invalid dates
		}

		_, err = tx.Exec(ctx, historyQuery,
			entry.Symbol,
			date,
			entry.Open,
			entry.High,
			entry.Low,
			entry.Close,
			entry.Volume,
			entry.Close, // Using close as adj_close for now
		)
		if err != nil {
			return fmt.Errorf("failed to save stock history entry: %w", err)
		}
	}

	// Commit the transaction
	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *PostgresStockRepository) SaveStockHistory(entries []models.StockHistoryEntry) error {
	ctx := context.Background()

	query := `
        INSERT INTO stock_history (symbol, date, open, high, low, close, volume, adj_close) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (symbol, date) 
        DO UPDATE SET 
            open = EXCLUDED.open,
            high = EXCLUDED.high,
            low = EXCLUDED.low,
            close = EXCLUDED.close,
            volume = EXCLUDED.volume,
            adj_close = EXCLUDED.adj_close
    `

	for _, entry := range entries {
		_, err := r.conn.Exec(ctx, query,
			entry.Date, // Assuming you add Symbol field to StockHistoryEntry
			entry.Date,
			entry.Open,
			entry.High,
			entry.Low,
			entry.Close,
			entry.Volume,
			entry.AdjClose,
		)
		if err != nil {
			return fmt.Errorf("failed to save stock history entry: %w", err)
		}
	}

	return nil
}
