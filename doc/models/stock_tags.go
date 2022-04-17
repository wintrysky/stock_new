package models

import (
	"github.com/guregu/null"
)

// StockTags table comment
type StockTags struct {
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// Symbol symbol
	Symbol string `gorm:"column:symbol" json:"symbol"`
	// TagName tag name
	TagName string `gorm:"column:tag_name" json:"tag_name"`
}

// TableName sets the insert table name for this struct type
func (s *StockTags) TableName() string {
	return "stock_tags"
}
