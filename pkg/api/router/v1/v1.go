package v1

import (
	"github.com/gin-gonic/gin"

	handler "github.com/phenix3443/dex-trade-bot/pkg/api/handler/v1"
)

func RegisterWalletRoutes(rg *gin.RouterGroup) {
	wallet := rg.Group("/wallets")
	{
		wallet.GET("/discover", handler.DiscoverWallets)
		wallet.GET("/:address", handler.GetWalletDetail)
		wallet.GET("/:address/positions", handler.GetWalletPositions)
		wallet.GET("/:address/fills", handler.GetWalletFills)
		wallet.GET("/:address/trades", handler.GetWalletTrades)
		wallet.GET("/:address/transfers", handler.GetWalletPositions)
	}
}
