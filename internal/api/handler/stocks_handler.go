package handler

import (
	_ "encoding/json"
	_ "fmt"
	_ "net/http"
	_ "time"

	_ "go-flow/internal/repository"

	_ "github.com/gin-gonic/gin"
)

type StocksHandler struct {
	// stockRepo repository.StockRepository
}
