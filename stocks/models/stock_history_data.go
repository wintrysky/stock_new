package models

import (
	"time"
)

// StockHistoryData table comment
type StockHistoryData struct {
	// ClosePrice column comments
	ClosePrice float64 `gorm:"column:close_price" json:"close_price"`
	// HighPrice column comments
	HighPrice float64 `gorm:"column:high_price" json:"high_price"`
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// LowPrice column comments
	LowPrice float64 `gorm:"column:low_price" json:"low_price"`
	// OpenPrice column comments
	OpenPrice float64 `gorm:"column:open_price" json:"open_price"`
	// Symbol column comments
	Symbol string `gorm:"column:symbol" json:"symbol"`
	// TradeTime yyyyMMdd
	TradeTime string `gorm:"column:trade_time" json:"trade_time"`
	TradeTimeD time.Time `gorm:"column:trade_time_d" json:"trade_time_d"`
	// 交易总金额
	TradeAmount float64 `gorm:"column:trade_amount" json:"trade_amount"`
}

// TableName sets the insert table name for this struct type
func (s *StockHistoryData) TableName() string {
	return "stock_history_data"
}
