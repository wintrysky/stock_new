package models

import (
	"github.com/guregu/null"
	"time"
)

// StockBasic table comment
type StockBasic struct {
	// BlacklistFlag 是否黑名单
	BlacklistFlag string `gorm:"column:blacklist_flag" json:"blacklist_flag"`
	// BuyTags 强烈持有:StrongHold,谨慎持有:Careful,做多观察中:WatchCall,做空观察中:WatchPut,做空:Put,黑名单:BlackList
	BuyTags string `gorm:"column:buy_tags" json:"buy_tags"`
	// CalculateDate 计算周涨跌幅完成时间,如果无效时间,表示不可交易
	CalculateDate null.Time `gorm:"column:calculate_date" json:"calculate_date"`
	// CompanyTags 高使用率:HighUsageRate,护城河:StrongCityMoat,高成长:HighGrowth,新兴行业:EmergingIndustry
	CompanyTags string `gorm:"column:company_tags" json:"company_tags"`
	// CompareYesLastWeek 昨日收盘对比上周最高点涨跌幅0.1表示涨了10%,-0.1表示跌了10%
	CompareYesLastWeek null.Float `gorm:"column:compare_yes_last_week" json:"compare_yes_last_week"`
	// ConID IB股票标记值
	ConID null.Int `gorm:"column:con_id" json:"con_id"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Currency USD
	Currency string `gorm:"column:currency" json:"currency"`
	// CurrentPrice 最新价
	CurrentPrice float64 `gorm:"column:current_price" json:"current_price"`
	// Description 描述
	Description null.String `gorm:"column:description" json:"description"`
	// Exchange ISLAND/SMART
	Exchange null.String `gorm:"column:exchange" json:"exchange"`
	// HighPrice 最高
	HighPrice float64 `gorm:"column:high_price" json:"high_price"`
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// Ikey column comments
	Ikey null.String `gorm:"column:ikey" json:"ikey"`
	// IncreaseRate10day 10日涨跌幅
	IncreaseRate10day null.Float `gorm:"column:increase_rate_10day" json:"increase_rate_10day"`
	// IncreaseRate120day 120日涨跌幅
	IncreaseRate120day null.Float `gorm:"column:increase_rate_120day" json:"increase_rate_120day"`
	// IncreaseRate20day 20日涨跌幅
	IncreaseRate20day null.Float `gorm:"column:increase_rate_20day" json:"increase_rate_20day"`
	// IncreaseRate250day 250日涨跌幅
	IncreaseRate250day null.Float `gorm:"column:increase_rate_250day" json:"increase_rate_250day"`
	// IncreaseRate5day 5日涨跌幅
	IncreaseRate5day null.Float `gorm:"column:increase_rate_5day" json:"increase_rate_5day"`
	// IncreaseRate60day 60日涨跌幅
	IncreaseRate60day null.Float `gorm:"column:increase_rate_60day" json:"increase_rate_60day"`
	// IncreaseRateCurrDay 涨跌幅
	IncreaseRateCurrDay null.Float `gorm:"column:increase_rate_curr_day" json:"increase_rate_curr_day"`
	// IncreaseRateFormYear 年初至今涨跌幅
	IncreaseRateFormYear null.Float `gorm:"column:increase_rate_form_year" json:"increase_rate_form_year"`
	// IncreaseRateYesterday 涨跌幅
	IncreaseRateYesterday null.Float `gorm:"column:increase_rate_yesterday" json:"increase_rate_yesterday"`
	// Industry 所属行业
	Industry string `gorm:"column:industry" json:"industry"`
	// IsBlock 板块
	IsBlock string `gorm:"column:is_block" json:"is_block"`
	// IsChina 中概股
	IsChina string `gorm:"column:is_china" json:"is_china"`
	// IsEtf ETF
	IsEtf string `gorm:"column:is_etf" json:"is_etf"`
	// IsHk 香港
	IsHk string `gorm:"column:is_hk" json:"is_hk"`
	// IsHot 热门
	IsHot string `gorm:"column:is_hot" json:"is_hot"`
	// IsOption 是否期权
	IsOption string `gorm:"column:is_option" json:"is_option"`
	// IsRisk 政策风险
	IsRisk string `gorm:"column:is_risk" json:"is_risk"`
	// IsStar STAR
	IsStar string `gorm:"column:is_star" json:"is_star"`
	// IsTooHigh 高位可空
	IsTooHigh string `gorm:"column:is_too_high" json:"is_too_high"`
	// IsYestodayHot 昨日强势股
	IsYestodayHot string `gorm:"column:is_yestoday_hot" json:"is_yestoday_hot"`
	// IsYestodayHotDate 昨日强势股
	IsYestodayHotDate null.Time `gorm:"column:is_yestoday_hot_date" json:"is_yestoday_hot_date"`
	// LastWeekRate 上周涨跌率,0.1表示涨了10%,-0.1表示跌了10%
	LastWeekRate null.Float `gorm:"column:last_week_rate" json:"last_week_rate"`
	// LowPrice 最低
	LowPrice float64 `gorm:"column:low_price" json:"low_price"`
	// Name 名称
	Name string `gorm:"column:name" json:"name"`
	// OpenPrice 开盘
	OpenPrice float64 `gorm:"column:open_price" json:"open_price"`
	// Pe 市盈率(静)
	Pe float64 `gorm:"column:pe" json:"pe"`
	// PrimExchange NASDAQ.NMS/NYSE
	PrimExchange null.String `gorm:"column:prim_exchange" json:"prim_exchange"`
	// ProfitTags 盈利良好:GoodGain,亏损:Debt
	ProfitTags string `gorm:"column:profit_tags" json:"profit_tags"`
	// SecType STK/OPT
	SecType null.String `gorm:"column:sec_type" json:"sec_type"`
	// Symbol 代码
	Symbol string `gorm:"column:symbol" json:"symbol"`
	// TotalMarketCap 总市值
	TotalMarketCap float64 `gorm:"column:total_market_cap" json:"total_market_cap"`
	// TradeDate 股票交易日
	TradeDate null.Time `gorm:"column:trade_date" json:"trade_date"`
	// UpdateTime column comments
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// YesterdayPrice 昨收
	YesterdayPrice float64 `gorm:"column:yesterday_price" json:"yesterday_price"`
}

// TableName sets the insert table name for this struct type
func (s *StockBasic) TableName() string {
	return "stock_basic"
}
