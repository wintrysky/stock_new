package models

// StockUser table comment
type StockUser struct {
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// UserName column comments
	UserName string `gorm:"column:user_name" json:"user_name"`
	// UserPwd column comments
	UserPwd string `gorm:"column:user_pwd" json:"user_pwd"`
}

// TableName sets the insert table name for this struct type
func (s *StockUser) TableName() string {
	return "stock_user"
}
