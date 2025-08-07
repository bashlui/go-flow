package models

import "time"

// StockResponse represents the response structure for stock API calls
type StockResponse struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Change    float64   `json:"change"`
	Timestamp time.Time `json:"timestamp"`
	Volume    int64     `json:"volume,omitempty"`
}

// StockHistoryResponse represents historical stock data response
type StockHistoryResponse struct {
	Symbol  string              `json:"symbol"`
	History []StockHistoryEntry `json:"history"`
	Period  string              `json:"period,omitempty"`
}

// StockWatchlistResponse represents user watchlist response
type StockWatchlistResponse struct {
	Watchlist []StockWatchlist `json:"watchlist"`
	Count     int              `json:"count"`
}

// StockAlertResponse represents stock alert response
type StockAlertResponse struct {
	Alerts []StockAlert `json:"alerts"`
	Count  int          `json:"count"`
}

// NewsResponse represents the response structure for news API calls
type NewsResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Source      string    `json:"source"`
	PublishedAt time.Time `json:"published_at"`
	Category    string    `json:"category,omitempty"`
}

// APIResponse is a generic wrapper for API responses
type APIResponse[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
