package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Address        string  `gorm:"primaryKey;type:text;comment:钱包地址" json:"address"`
	WinRate        float64 `gorm:"comment:总胜率（%）" json:"win_rate"`
	TradeCount     int     `gorm:"comment:总成交笔数" json:"trade_count"`
	PnL            float64 `gorm:"comment:累计盈亏 (USD)" json:"pnl"`
	ROE            float64 `gorm:"comment:投资回报率 (%)" json:"roe"`
	AvgLeverage    float64 `gorm:"comment:平均杠杆倍数" json:"avg_leverage"`
	Funding        float64 `gorm:"comment:累计资金费" json:"funding"`
	PerpEquity     float64 `gorm:"comment:永续合约账户权益 (margin + uPnL)" json:"perp_equity"`
	MarginUsage    float64 `gorm:"comment:保证金使用率 (%)" json:"margin_usage"`
	CurrentWinRate string  `gorm:"comment:当前胜率 (n/n %)" json:"current_win_rate"`
	DirectionBias  string  `gorm:"comment:方向偏好 (LONG/SHORT)" json:"direction_bias"`
	LongExposure   float64 `gorm:"comment:多头仓位占比 (%)" json:"long_exposure"`
	ShortExposure  float64 `gorm:"comment:空头仓位占比 (%)" json:"short_exposure"`
}
