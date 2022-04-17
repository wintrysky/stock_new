package service

import (
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
	"time"
	"ww/stocks/dto"
	"ww/stocks/models"
	"ww/stocks/orm"
)

// 计算5/10/20/60日涨跌幅
func CalculateTradeRaceTask(){
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	log.Info("CALL CalculateTradeRaceTask at",time.Now())

	// 1、取所有的stock_basic,遍历
	var p dto.StockSearchCondition
	p.IsOption = true
	oldItems, err := Search(p)
	if err != nil {
		log.Error(err)
		return
	}

	for _, item := range oldItems {
		runOne(item.Symbol)
	}
}

func runOne(symbol string){
	// 2、对于每一只symbol，取最新的60条数据
	historyList,err := getHistory(symbol)
	if err != nil {
		log.Error(err)
		return
	}
	if len(historyList) < 5 {
		return
	}
	// 计算涨幅
	var r5 float64 = 0
	var r10 float64 = 0
	var r20 float64 = 0
	var r60 float64 = 0
	colMap := make(map[string]interface{})
	if len(historyList)>=5 {
		r5 = calRate(historyList,5)
		colMap["trade_rate_5day"] = r5
	}
	if len(historyList)>=10 {
		r10 = calRate(historyList,10)
		colMap["trade_rate_10day"] = r10
	}
	if len(historyList)>=20 {
		r20 = calRate(historyList,20)
		colMap["trade_rate_20day"] = r20
	}
	if len(historyList)>=60 {
		r60 = calRate(historyList,60)
		colMap["trade_rate_60day"] = r60
	}
	if len(colMap) == 0 {
		return
	}
	colMap["trade_rate_date"] = null.TimeFrom(time.Now())

	// 更新stock_basic
	qa := orm.NewQuery()
	qa.UpdateTradeRate(symbol,colMap)
	if qa.Error != nil {
		log.Error(err)
	}
}

func getHistory(symbol string)(items []models.StockHistoryData,err error){
	qa := orm.NewQuery()
	sql := "SELECT * FROM stock_history_data WHERE symbol=? ORDER BY trade_time_d desc LIMIT 60"
	qa.ExecuteTextQuery(&items,sql,symbol)
	if qa.Error != nil {
		return items,qa.Error
	}
	return
}

func calRate(items []models.StockHistoryData,day int)float64{
	last := items[0]
	var first models.StockHistoryData
	var idx int
	for _, item := range items {
		first = item
		if idx == day -1 {
			break
		}
		idx++
	}

	if first.TradeAmount == 0 {
		return 0
	}
	var rate float64
	rate = 100*(last.TradeAmount - first.TradeAmount) / first.TradeAmount

	return rate
}

