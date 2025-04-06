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
// GetWalletFills 获取钱包成交明细
// @Summary 获取钱包成交记录
// @Description 返回某个钱包地址的所有 Fills（成交记录）
// @Tags Wallets
// @Produce json
// @Param wallet_address path string true "钱包地址"
// @Success 200 {array} types.FillRecord
// @Router /api/v1/wallets/{wallet_address}/fills [get]
func GetWalletFills(c *gin.Context) {
	records := mockWalletsFills(common.HexToAddress(c.Param("address")))
	c.JSON(http.StatusOK, records)
}

type FillRecord struct {
	Coin       string   `json:"coin"`          // ETH
	Direction  string   `json:"direction"`     // Short / Long
	Type       string   `json:"type"`          // Increase / Decrease
	Size       float64  `json:"size"`          // 成交数量
	AvgPrice   float64  `json:"avg_price"`     // 成交均价
	PnL        *float64 `json:"pnl,omitempty"` // 盈亏（可能为 null）
	FillCount  string   `json:"fill_count"`    // 填充次数（如 "1 Fills"）
	LatestFill string   `json:"latest_fill"`   // 最近成交时间（ISO 字符串）
	OrderID    string   `json:"order_id"`      // 订单号
}

func mockWalletsFills(_ common.Address) []FillRecord {
	return []FillRecord{
		{
			Coin:       "ETH",
			Direction:  "Short",
			Type:       "Decrease",
			Size:       1.1175,
			AvgPrice:   1790.1,
			PnL:        ptrFloat(556.86),
			FillCount:  "1 Fills",
			LatestFill: "2025-04-06T02:40:00Z",
			OrderID:    "84150511875",
		},
	}
}

func ptrFloat(v float64) *float64 {
	return &v
}
