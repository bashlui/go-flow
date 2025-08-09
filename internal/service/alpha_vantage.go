package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AlphaVantageService struct {
	apiKey     string
	httpClient *http.Client
}

func NewAlphaVantageService(apiKey string) *AlphaVantageService {
	godotenv.Load() // Load .env file if it exists

	return &AlphaVantageService{
		apiKey: os.Getenv("ALPHA_VANTAGE_API_KEY"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Struct to parse Alpha Vantage response
type StockData struct {
	Symbol        string  `json:"symbol"`
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	LastRefreshed string  `json:"last_refreshed"`
}

type DailyResponse struct {
	MetaData struct {
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
	} `json:"Meta Data"`
	TimeSeries map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

func (s *AlphaVantageService) GetDailyStockData(symbol string) ([]StockData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, s.apiKey)

	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Alpha Vantage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response from Alpha Vantage: %d", resp.StatusCode)
	}

	var dailyResponse DailyResponse
	if err := json.NewDecoder(resp.Body).Decode(&dailyResponse); err != nil {
		return nil, fmt.Errorf("failed to decode Alpha Vantage response: %w", err)
	}

	var stockData []StockData
	for date, data := range dailyResponse.TimeSeries {
		// Convert string values to appropriate types
		open, _ := strconv.ParseFloat(data.Open, 64)
		high, _ := strconv.ParseFloat(data.High, 64)
		low, _ := strconv.ParseFloat(data.Low, 64)
		close, _ := strconv.ParseFloat(data.Close, 64)
		volume, _ := strconv.ParseInt(data.Volume, 10, 64)

		stockData = append(stockData, StockData{
			Symbol:        dailyResponse.MetaData.Symbol,
			Date:          date,
			Open:          open,
			High:          high,
			Low:           low,
			Close:         close,
			Volume:        volume,
			LastRefreshed: dailyResponse.MetaData.LastRefreshed,
		})
	}

	return stockData, nil
}
