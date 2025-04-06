package models

import (
	"time"

	"gorm.io/gorm"
)

type Trade struct {
	gorm.Model
	Address    string    `gorm:"index;type:text;comment:钱包地址" json:"address"`
	Time       time.Time `gorm:"comment:交易时间" json:"time"`
	Coin       string    `gorm:"comment:币种" json:"coin"`
	Direction  string    `gorm:"comment:方向 (Open Short / Close Short)" json:"direction"`
	Price      float64   `gorm:"comment:成交价格" json:"price"`
	Size       float64   `gorm:"comment:成交数量" json:"size"`
	TradeValue float64   `gorm:"comment:成交总价值 (USD)" json:"trade_value"`
	Fee        string    `gorm:"comment:手续费" json:"fee"`
	ClosedPnL  *float64  `gorm:"comment:已实现盈亏，可为空" json:"closed_pnl"`
	OrderID    string    `gorm:"index;comment:订单 ID" json:"order_id"`
}
