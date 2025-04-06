package models

import (
	"time"

	"gorm.io/gorm"
)

type Fill struct {
	gorm.Model
	Address    string    `gorm:"index;type:text;comment:钱包地址" json:"address"`
	Coin       string    `gorm:"comment:币种" json:"coin"`
	Direction  string    `gorm:"comment:方向 (Short / Long)" json:"direction"`
	Type       string    `gorm:"comment:开减仓类型 (Increase / Decrease)" json:"type"`
	Size       float64   `gorm:"comment:成交数量" json:"size"`
	AvgPrice   float64   `gorm:"comment:成交均价" json:"avg_price"`
	PnL        *float64  `gorm:"comment:成交盈亏，可为空" json:"pnl,omitempty"`
	FillCount  int       `gorm:"comment:成交次数" json:"fill_count"`
	LatestFill time.Time `gorm:"comment:最近成交时间" json:"latest_fill"`
	OrderID    string    `gorm:"index;comment:订单 ID" json:"order_id"`
}
