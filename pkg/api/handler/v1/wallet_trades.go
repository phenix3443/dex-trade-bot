package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetWalletTrades 获取钱包历史交易记录
// @Summary 获取钱包交易历史
// @Description 返回某个钱包地址的所有交易历史（开仓/平仓）
// @Tags Wallets
// @Produce json
// @Param wallet_address path string true "钱包地址"
// @Success 200 {array} types.TradeRecord
// @Router /api/v1/wallets/{wallet_address}/trades [get]
func GetWalletTrades(c *gin.Context) {
	address := c.Param("wallet_address")
	trades := mockWalletsTrades(common.HexToAddress(address))
	c.JSON(http.StatusOK, trades)
}

type TradeRecord struct {
	Time       string   `json:"time"`        // 时间（ISO 字符串）
	Coin       string   `json:"coin"`        // 币种（如 ETH）
	Direction  string   `json:"direction"`   // Open Short / Close Short 等
	Price      float64  `json:"price"`       // 成交价格
	Size       float64  `json:"size"`        // 成交数量
	TradeValue float64  `json:"trade_value"` // 成交金额
	Fee        string   `json:"fee"`         // 手续费（USDC）
	ClosedPnL  *float64 `json:"closed_pnl"`  // 已实现盈亏（可为 null）
	OrderID    string   `json:"order_id"`    // 订单 ID
}

func mockWalletsTrades(_ common.Address) []TradeRecord {
	return []TradeRecord{
		{
			Time:       "2025-04-06T02:40:00Z",
			Coin:       "ETH",
			Direction:  "Close Short",
			Price:      1790.10,
			Size:       1.12,
			TradeValue: 2000.44,
			Fee:        "0.4601 USDC",
			ClosedPnL:  ptrFloat(556.86),
			OrderID:    "84150511875",
		},
		{
			Time:       "2025-04-06T02:40:00Z",
			Coin:       "ETH",
			Direction:  "Open Short",
			Price:      1790.00,
			Size:       1.12,
			TradeValue: 2000.32,
			Fee:        "0.4601 USDC",
			ClosedPnL:  nil,
			OrderID:    "84150509895",
		},
	}
}
