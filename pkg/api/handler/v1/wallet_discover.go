package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phenix3443/dex-trade-bot/pkg/types"
)

// DiscoverWallets 获取活跃钱包列表（支持时间周期筛选和分页）
// @Summary Discover wallets
// @Description 获取活跃交易钱包列表，支持时间周期筛选（1d/7d/30d）与分页
// @Tags Wallets
// @Accept json
// @Produce json
// @Param period query string false "时间周期，可选值为 1d, 7d, 30d" Enums(1d, 7d, 30d) default(7d)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} PaginatedResponse
// @Router /api/v1/wallets/discover [get]
func DiscoverWallets(c *gin.Context) {
	period := c.DefaultQuery("period", "7d")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	allWallets := mockWalletsForPeriod(period)

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(allWallets) {
		start = len(allWallets)
	}
	if end > len(allWallets) {
		end = len(allWallets)
	}
	paged := allWallets[start:end]

	resp := PaginatedResponse{
		Page:     page,
		PageSize: pageSize,
		Total:    len(allWallets),
		Wallets:  paged,
	}

	c.JSON(http.StatusOK, resp)
}

type DiscoverWalletInfo struct {
	types.WalletStats
	LongPnL      float64  `json:"long_pnl"`       // 多头盈亏
	ShortPnL     float64  `json:"short_pnl"`      // 空头盈亏
	LongWinRate  string   `json:"long_win_rate"`  // 多头胜率（百分比）
	ShortWinRate string   `json:"short_win_rate"` // 空头胜率（百分比）
	TokenCount   int      `json:"token_count"`    // 曾交易的币种数量
	TradedTokens []string `json:"traded_tokens"`  // 曾交易币种列表（如 BTC, ETH, SOL）
	FirstTxTime  string   `json:"first_tx_time"`  // 第一次交易时间（ISO 8601 格式）
	LastTxTime   string   `json:"last_tx_time"`   // 最近一次交易时间
}

type PaginatedResponse struct {
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
	Total    int                  `json:"total"`
	Wallets  []DiscoverWalletInfo `json:"wallets"`
}

func mockWalletsForPeriod(_ string) []DiscoverWalletInfo {
	return []DiscoverWalletInfo{
		{
			WalletStats: types.WalletStats{
				Address:     "0x123...",
				PnL:         1234567.0,
				WinRate:     "85%",
				TradeCount:  1000,
				ROE:         0.15,
				AvgLeverage: 2.5,
				Funding:     0.02,
			},
			LongPnL:  1251905.0,
			ShortPnL: 796140.0,
		},
	}
}
