package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phenix3443/dex-trade-bot/pkg/types"
)

func GetWalletDetail(c *gin.Context) {
	address := c.Param("wallet_address")
	info := mockWalletDetail(address)
	c.JSON(http.StatusOK, gin.H{"wallet": info})
}

type WalletDetail struct {
	types.WalletStats
	PerpEquity     float64 `json:"perp_equity"`
	MarginUsage    float64 `json:"margin_usage"`
	CurrentWinRate string  `json:"current_win_rate"`
	DirectionBias  string  `json:"direction_bias"`
	LongExposure   float64 `json:"long_exposure"`
	ShortExposure  float64 `json:"short_exposure"`
	Hypurrscan     string  `json:"hypurrscan"`
}

func mockWalletDetail(address string) WalletDetail {
	return WalletDetail{
		WalletStats: types.WalletStats{
			Address:     address,
			PnL:         12345.67,
			WinRate:     "75%",
			TradeCount:  100,
			ROE:         10.5,
			AvgLeverage: 2.0,
			Funding:     0.01,
		},
		PerpEquity:     10000.0,
		MarginUsage:    0.25,
		CurrentWinRate: "55%",
		DirectionBias:  "LONG",
		LongExposure:   5000.0,
		ShortExposure:  3000.0,
		Hypurrscan:     getWalletHypurrscan(address),
	}
}

func getWalletHypurrscan(address string) string {
	return fmt.Sprintf("https://hypurrscan.io/address/ data for %s", address)
}
