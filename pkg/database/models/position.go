package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Address     string  `gorm:"index;type:text;comment:钱包地址" json:"address"`
	Coin        string  `gorm:"comment:币种" json:"coin"`
	Type        string  `gorm:"comment:持仓类型 (如 SHORT - 25X)" json:"type"`
	Size        float64 `gorm:"comment:仓位数量" json:"size"`
	PositionUSD float64 `gorm:"comment:仓位价值 (USD)" json:"position_usd"`
	EntryPrice  float64 `gorm:"comment:开仓价格" json:"entry_price"`
	MarkPrice   float64 `gorm:"comment:标记价格" json:"mark_price"`
	PnL         float64 `gorm:"comment:当前盈亏" json:"pnl"`
	ROE         float64 `gorm:"comment:当前 ROE (%)" json:"roe"`
	LiqPrice    float64 `gorm:"comment:清算价格" json:"liq_price"`
	Margin      string  `gorm:"comment:保证金描述" json:"margin"`
	Funding     float64 `gorm:"comment:当前资金费" json:"funding"`
}
