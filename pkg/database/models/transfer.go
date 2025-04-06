package models

import (
	"time"

	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Address     string    `gorm:"index;type:text;comment:钱包地址" json:"address"`
	Time        time.Time `gorm:"comment:操作时间" json:"time"`
	Type        string    `gorm:"comment:操作类型 (Deposit / Withdraw)" json:"type"`
	Asset       string    `gorm:"comment:资产类型，如 USDC" json:"asset"`
	Amount      float64   `gorm:"comment:操作金额" json:"amount"`
	Transaction string    `gorm:"comment:交易哈希" json:"transaction"`
}
