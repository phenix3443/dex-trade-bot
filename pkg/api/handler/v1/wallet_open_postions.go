package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetWalletPositions 获取钱包当前持仓列表
// @Summary 获取钱包当前持仓
// @Description 查询某个钱包当前的持仓列表（Open Positions）
// @Tags Wallets
// @Produce json
// @Param wallet_address path string true "钱包地址"
// @Success 200 {array} types.OpenPosition
// @Router /api/v1/wallets/{wallet_address}/positions [get]
func GetWalletPositions(c *gin.Context) {
	positions := mockWalletsOpenPositions(common.HexToAddress(c.Param("address")))
	c.JSON(http.StatusOK, positions)
}

type OpenPosition struct {
	Coin        string  `json:"coin"`         // ETH
	Type        string  `json:"type"`         // SHORT - 25X
	Size        float64 `json:"size"`         // 仓位数量
	PositionUSD float64 `json:"position_usd"` // 仓位价值（美元）
	EntryPrice  float64 `json:"entry_price"`  // 开仓价
	MarkPrice   float64 `json:"mark_price"`   // 当前标记价
	PnL         float64 `json:"pnl"`          // 盈亏金额
	ROE         float64 `json:"roe"`          // ROE%
	LiqPrice    float64 `json:"liq_price"`    // 清算价格
	Margin      string  `json:"margin"`       // 保证金类型/金额
	Funding     float64 `json:"funding"`      // 资金费用
}

func mockWalletsOpenPositions(_ common.Address) []OpenPosition {
	return []OpenPosition{
		{
			Coin:        "ETH",
			Type:        "SHORT - 25X",
			Size:        77695.63,
			PositionUSD: 139813289.24,
			EntryPrice:  2288.41,
			MarkPrice:   1799.50,
			PnL:         37986618.59,
			ROE:         534.12,
			LiqPrice:    2894.52,
			Margin:      "$5,592,531.57 (Cross)",
			Funding:     3381728.89,
		},
	}
}
