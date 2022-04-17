package service

import (
	"github.com/AvraamMavridis/randomcolor"
	"github.com/shopspring/decimal"
	"time"
	"ww/stocks/dto"
	"ww/stocks/models"
	"ww/stocks/orm"
	"ww/stocks/xerr"
)

func GetHistoryBySymbol(symbols []string)(result dto.KLineResult,err error){
	xerr.HandleErr(&err)

	var items []models.StockHistoryData
	qa := orm.NewQuery()
	dt := time.Now().AddDate(0,-3,0)
	qa.ExecuteTextQuery(&items,"select * from stock_history_data where symbol in (?) and trade_time_d>=? order by trade_time asc",symbols,dt)
	//qa.ExecuteTextQuery(&items,"select * from stock_history_data where symbol in (?)  order by trade_time asc limit 20",symbols)
	xerr.ThrowError(qa.Error)
	if len(items) == 0 {
		qa.ExecuteTextQuery(&items,"select * from stock_history_data where symbol in (?) and trade_time_d>=? order by trade_time asc","QQQ",dt)
	}

	// 按时间顺序补齐
	var orderDate []string
	for _, item := range items {
		if contains(orderDate,item.TradeTime) == false {
			orderDate = append(orderDate,item.TradeTime)
		}
	}

	mm := make(map[string][]models.StockHistoryData)
	for _, item := range items {
		mm[item.Symbol] = append(mm[item.Symbol],item)
	}
	mm2 := make(map[string]map[string]models.StockHistoryData)
	for symbol,items := range mm {
		m0 := make(map[string]models.StockHistoryData)
		for _,ii := range items {
			m0[ii.TradeTime] = ii
		}
		mm2[symbol] = m0
	}

	var lines []dto.KLineItem
	var colors []string
	var kLineItems []dto.KLineStyle

	for symbol,hisMap := range mm2 {
		var prePrice float64
		var per float64
		for _,dd := range orderDate {
			item,ok := hisMap[dd]
			if ok {
				// 如果可以找到时间上的历史数据
				var l dto.KLineItem
				l.Day = item.TradeTime
				l.Symbol = symbol
				if item.ClosePrice > 0 && per == 0 {
					per = tenPercent(item.ClosePrice)
				}

				l.Price = roundPrice(item.ClosePrice * per)
				prePrice = l.Price
				lines = append(lines,l)
			}else {
				var l dto.KLineItem
				l.Day = dd
				l.Symbol = symbol
				l.Price = prePrice
				lines = append(lines,l)
			}
		}

		color := randomcolor.GetRandomColorInHex()
		time.Sleep(time.Duration(1*time.Microsecond))
		colors = append(colors,color)

		var style dto.KLineStyle
		style.Name = symbol
		var marker dto.KLineMarker
		marker.Symbol = "square"
		var colorStyle dto.KLineColor
		colorStyle.Fill = color
		marker.Style = colorStyle
		style.Marker = marker
		kLineItems = append(kLineItems,style)
	}

	result.Colors = colors
	result.Data = lines
	result.Items = kLineItems
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func tenPercent(price float64)float64{
	if price < 1 {
		return price
	}
	v10 := 50/price

	v1 := decimal.NewFromFloat(v10)
	v2, _ := v1.Float64()

	return v2
}

func roundPrice(price float64)float64{
	if price < 1 {
		return price
	}
	v1 := decimal.NewFromFloat(price)
	v2, _ := v1.Round(3).Float64()

	return v2
}