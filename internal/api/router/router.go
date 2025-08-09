package router

import (
	"go-flow/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, stocksHandler *handler.StocksHandler) {
	api := router.Group("/api")
	{
		stocks := api.Group("/stocks")
		{
			stocks.GET("", stocksHandler.GetStocks)
			stocks.GET("/:id", stocksHandler.GetStockByID)
			stocks.POST("/fetch/:symbol", stocksHandler.FetchStockData)
		}
	}
}
