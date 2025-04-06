package types

type WalletStats struct {
	Address     string  `json:"address"`
	PnL         float64 `json:"pnl"`
	WinRate     string  `json:"win_rate"`
	TradeCount  int     `json:"trade_count"`
	ROE         float64 `json:"roe"`
	AvgLeverage float64 `json:"avg_leverage"`
	Funding     float64 `json:"funding"`
}
