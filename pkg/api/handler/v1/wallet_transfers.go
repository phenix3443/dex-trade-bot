package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetWalletTransfers 获取钱包充值和提现记录
// @Summary 获取钱包转账记录（充值/提现）
// @Description 查询某个钱包地址的充值与提现明细
// @Tags Wallets
// @Produce json
// @Param wallet_address path string true "钱包地址"
// @Success 200 {array} types.TransferRecord
// @Router /api/v1/wallets/{wallet_address}/transfers [get]
func GetWalletTransfers(c *gin.Context) {
	address := c.Param("wallet_address")
	records := mockWalletsTransfers(common.HexToAddress(address))
	c.JSON(http.StatusOK, records)
}

type TransferRecord struct {
	Time        string  `json:"time"`        // 时间（ISO 8601）
	Type        string  `json:"type"`        // 类型（Deposit / Withdraw）
	Asset       string  `json:"asset"`       // 资产类型（如 USDC、ETH）
	Amount      float64 `json:"amount"`      // 金额
	Transaction string  `json:"transaction"` // 交易哈希（支持点击查看链上）
}

func mockWalletsTransfers(_ common.Address) []TransferRecord {
	return []TransferRecord{
		{
			Time:        "2025-04-05T11:20:00Z",
			Type:        "Withdraw",
			Asset:       "-",
			Amount:      4999999,
			Transaction: "0x1e...e205",
		},
		{
			Time:        "2025-03-29T01:03:00Z",
			Type:        "Withdraw",
			Asset:       "-",
			Amount:      9499999,
			Transaction: "0xf8...4474",
		},
	}
}
