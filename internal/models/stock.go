package models

import (
	"time"
)

// StockHistoryEntry represents a single historical data point
type StockHistoryEntry struct {
	Date     time.Time `json:"date" db:"date"`
	Open     float64   `json:"open" db:"open"`
	High     float64   `json:"high" db:"high"`
	Low      float64   `json:"low" db:"low"`
	Close    float64   `json:"close" db:"close"`
	Volume   int64     `json:"volume" db:"volume"`
	AdjClose float64   `json:"adj_close" db:"adj_close"`
}

// StockWatchlist represents a user's watchlist entry
type StockWatchlist struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Symbol    string    `json:"symbol" db:"symbol"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// StockAlert represents price alerts for stocks
type StockAlert struct {
	ID          int64      `json:"id" db:"id"`
	UserID      int64      `json:"user_id" db:"user_id"`
	Symbol      string     `json:"symbol" db:"symbol"`
	AlertType   string     `json:"alert_type" db:"alert_type"` // "above", "below"
	TargetPrice float64    `json:"target_price" db:"target_price"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	TriggeredAt *time.Time `json:"triggered_at,omitempty" db:"triggered_at"`
}

// Stock request structure for fetching stock data
type StockRequest struct {
	Symbol string `json:"symbol" validate:"required,min=1,max=10"`
	Fields string `json:"fields,omitempty" validate:"omitempty,min=1,max=100"`
}

type StockListRequest struct {
	Symbols []string `json:"symbols" validate:"required,min=1,max=100"`
	Fields  string   `json:"fields,omitempty" validate:"omitempty,min=1,max=100"`
}

// Stock represents the main stock entity with current data
type Stock struct {
	Symbol           string    `json:"symbol" db:"symbol"`
	Name             string    `json:"name" db:"name"`
	CurrentPrice     float64   `json:"current_price" db:"last_price"`
	PreviousClose    float64   `json:"previous_close,omitempty"`
	Open             float64   `json:"open,omitempty"`
	DayHigh          float64   `json:"day_high,omitempty"`
	DayLow           float64   `json:"day_low,omitempty"`
	Volume           int64     `json:"volume,omitempty"`
	MarketCap        int64     `json:"market_cap,omitempty"`
	PeRatio          float64   `json:"pe_ratio,omitempty"`
	DividendYield    float64   `json:"dividend_yield,omitempty"`
	FiftyTwoWeekHigh float64   `json:"fifty_two_week_high,omitempty"`
	FiftyTwoWeekLow  float64   `json:"fifty_two_week_low,omitempty"`
	Sector           string    `json:"sector,omitempty"`
	Industry         string    `json:"industry,omitempty"`
	LastUpdated      time.Time `json:"last_updated" db:"created_at"`
}

// StockQuote represents real-time stock quote data
type StockQuote struct {
	Symbol        string    `json:"symbol"`
	Price         float64   `json:"price"`
	Change        float64   `json:"change"`
	ChangePercent float64   `json:"change_percent"`
	Volume        int64     `json:"volume"`
	Timestamp     time.Time `json:"timestamp"`
}

// TechnicalIndicators represents calculated technical analysis data
type TechnicalIndicators struct {
	Symbol         string                 `json:"symbol"`
	Timeframe      string                 `json:"timeframe"`
	SMA20          float64                `json:"sma_20,omitempty"`
	SMA50          float64                `json:"sma_50,omitempty"`
	SMA200         float64                `json:"sma_200,omitempty"`
	EMA12          float64                `json:"ema_12,omitempty"`
	EMA26          float64                `json:"ema_26,omitempty"`
	RSI            float64                `json:"rsi,omitempty"`
	MACD           *MACDData              `json:"macd,omitempty"`
	BollingerBands *BollingerBandsData    `json:"bollinger_bands,omitempty"`
	CalculatedAt   time.Time              `json:"calculated_at"`
}

// MACDData represents MACD indicator values
type MACDData struct {
	MACD      float64 `json:"macd"`
	Signal    float64 `json:"signal"`
	Histogram float64 `json:"histogram"`
}

// BollingerBandsData represents Bollinger Bands values
type BollingerBandsData struct {
	Upper  float64 `json:"upper"`
	Middle float64 `json:"middle"`
	Lower  float64 `json:"lower"`
}

// Portfolio represents a user's portfolio
type Portfolio struct {
	ID           int64                `json:"id" db:"id"`
	UserID       int64                `json:"user_id" db:"user_id"`
	Name         string               `json:"name" db:"name"`
	TotalValue   float64              `json:"total_value"`
	CashBalance  float64              `json:"cash_balance" db:"cash_balance"`
	DayChange    float64              `json:"day_change"`
	TotalReturn  float64              `json:"total_return"`
	Positions    []PortfolioPosition  `json:"positions,omitempty"`
	CreatedAt    time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at" db:"updated_at"`
}

// PortfolioPosition represents a stock position in a portfolio
type PortfolioPosition struct {
	ID              int64     `json:"id" db:"id"`
	PortfolioID     int64     `json:"portfolio_id" db:"portfolio_id"`
	Symbol          string    `json:"symbol" db:"symbol"`
	Quantity        int64     `json:"quantity" db:"quantity"`
	AverageCost     float64   `json:"average_cost" db:"average_cost"`
	CurrentPrice    float64   `json:"current_price,omitempty"`
	MarketValue     float64   `json:"market_value,omitempty"`
	UnrealizedGain  float64   `json:"unrealized_gain,omitempty"`
	UnrealizedGainPercent float64 `json:"unrealized_gain_percent,omitempty"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Transaction represents buy/sell transactions
type Transaction struct {
	ID          int64     `json:"id" db:"id"`
	PortfolioID int64     `json:"portfolio_id" db:"portfolio_id"`
	Symbol      string    `json:"symbol" db:"symbol"`
	Type        string    `json:"type" db:"type"` // "buy" or "sell"
	Quantity    int64     `json:"quantity" db:"quantity"`
	Price       float64   `json:"price" db:"price"`
	Fees        float64   `json:"fees" db:"fees"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	ExecutedAt  time.Time `json:"executed_at" db:"executed_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// StockScreenerRequest for filtering stocks
type StockScreenerRequest struct {
	MinPrice      *float64 `json:"min_price,omitempty"`
	MaxPrice      *float64 `json:"max_price,omitempty"`
	MinVolume     *int64   `json:"min_volume,omitempty"`
	MinMarketCap  *int64   `json:"min_market_cap,omitempty"`
	MaxMarketCap  *int64   `json:"max_market_cap,omitempty"`
	Sector        string   `json:"sector,omitempty"`
	Industry      string   `json:"industry,omitempty"`
	MinPeRatio    *float64 `json:"min_pe_ratio,omitempty"`
	MaxPeRatio    *float64 `json:"max_pe_ratio,omitempty"`
	SortBy        string   `json:"sort_by,omitempty"` // "price", "volume", "market_cap", etc.
	SortOrder     string   `json:"sort_order,omitempty"` // "asc" or "desc"
	Limit         int      `json:"limit,omitempty"`
	Offset        int      `json:"offset,omitempty"`
}
