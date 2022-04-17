package dto


type StockSearchCondition struct {
	// 可交易期权股票
	IsOption bool `json:"is_option"`
	// StrongHold
	BuyTags string `json:"buy_tags"`
	CompanyTags string `json:"company_tags"`
	ProfitTags string `json:"profit_tags"`
	// Symbol 代码
	Symbol string `json:"symbol"`
	// 市值 OneToTen、TenToHundred、LargerOne、LargerTen、LargerThanHundred
	MarketCapTags string `json:"market_cap_tags"`
	// 成交额
	TradeAmount string `json:"trade_amount"`

	// 板块集合BK、明星股Star、ETF、中概股China、热门Hot、昨日强势股YesterdayHot、创新高Top
	SearchCategory string `json:"search_category"`
	// 当前页 默认1
	Current int `json:"current"`
	// 每页记录数
	PageSize int `json:"pageSize"`
}

