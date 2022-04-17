package template

var FutuColumnMap map[string]string

// string,float,int,datetime,percent,money
func InitFutuColumnMap(){
	FutuColumnMap = make(map[string]string)
	FutuColumnMap["代码"] = "symbol,string"
	FutuColumnMap["名称"] = "name,string"
	FutuColumnMap["今开"] = "open_price,float"
	FutuColumnMap["昨收"] = "yesterday_price,float"
	FutuColumnMap["最低"] = "low_price,float"
	FutuColumnMap["最高"] = "high_price,float"
	FutuColumnMap["最新价"] = "current_price,float"
	FutuColumnMap["市盈率(静)"] = "pe,stringToFloat"
	FutuColumnMap["总市值"] = "total_market_cap,money"
	FutuColumnMap["资产规模"] = "total_market_cap,money" // ETF
	FutuColumnMap["涨跌幅"] = "increase_rate_curr_day,percent"
	FutuColumnMap["5日涨跌幅"] = "increase_rate_5day,percent"
	FutuColumnMap["10日涨跌幅"] = "increase_rate_10day,percent"
	FutuColumnMap["20日涨跌幅"] = "increase_rate_20day,percent"
	FutuColumnMap["60日涨跌幅"] = "increase_rate_60day,percent"
	FutuColumnMap["120日涨跌幅"] = "increase_rate_120day,percent"
	FutuColumnMap["250日涨跌幅"] = "increase_rate_250day,percent"
	FutuColumnMap["年初至今涨跌幅"] = "increase_rate_form_year,percent"
	FutuColumnMap["所属行业"] = "industry,string"
	FutuColumnMap["成交额"] = "trade_amount,money2"
}