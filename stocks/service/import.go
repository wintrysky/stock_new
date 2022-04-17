package service

import (
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"strings"
	"time"
	"ww/stocks/global"
	"ww/stocks/models"
	"ww/stocks/orm"
	"ww/stocks/utils"
	"ww/stocks/utils/csv"
	"ww/stocks/xerr"
)

type ImportDataSrv struct {

}

func (c *ImportDataSrv) ImportData(ctx *gin.Context,importType,dateString string) (err error) {
	defer xerr.HandleErr(&err)

	filePath := utils.LoadExcelToTempFolder(ctx)
	cls := &csv.FileService{}
	data := cls.ImportCSV(filePath) // data [][英文数据库字段名][值]

	sCols := make(map[string]string)
	isBlock := ""
	isHot := ""
	isChina := ""
	isETF := ""
	isHK := ""
	isStar := ""
	isOption :=""
	isYestodayHot := ""
	tradeDate := time.Now().AddDate(0,0,-1)
	if dateString != "" {
		tradeDate = utils.ParseShortTime(dateString)
	}
	day := tradeDate.Weekday()
	if day == 6 || day == 0 {
		xerr.ThrowErrorMessage("错误的导入日期")
	}

	var isAll bool
	switch importType {
	case "all":
		if len(data) < 8000 {
			xerr.ThrowErrorMessage("less then 8000 rows")
		}
		isAll = true
	case "bk": // block
		isBlock = global.LogicY
		sCols["is_block"] = "Y"
		c.checkSymbolExits("BK2015",data)
		c.updateFlag("update stock_basic set is_block = ''")
	case "hot":
		isHot = global.LogicY
		sCols["is_hot"] = "Y"
		c.updateFlag("update stock_basic set is_hot = ''")
	case "china":
		isChina = global.LogicY
		sCols["is_china"] = "Y"
		c.checkSymbolExits("FUTU",data)
		c.updateFlag("update stock_basic set is_china = ''")
	case "ETF":
		isETF = global.LogicY
		sCols["is_etf"] = "Y"
		c.checkSymbolExits("TQQQ",data)
		c.updateFlag("update stock_basic set is_etf = ''")
	case "hk":
		isHK = global.LogicY
		sCols["is_hk"] = "Y"
		c.updateFlag("update stock_basic set is_hk = ''")
	case "star":
		isStar = global.LogicY
		sCols["is_star"] = "Y"
		c.updateFlag("update stock_basic set is_star = ''")
	case "option":
		isOption = global.LogicY
		sCols["is_option"] = "Y"
		c.updateFlag("update stock_basic set is_option = ''")
	case "yestodayhot":
		isYestodayHot = global.LogicY
		sCols["is_yestoday_hot"] = "Y"
		sCols["is_yestoday_hot_date"] = "is_yestoday_hot_date"
	default:
		xerr.ThrowErrorMessage("错误的类型:%s",importType)
	}

	var tempList []models.StockBasic
	dt := time.Now()
	for _,value := range data {
		var tmp models.StockBasic
		mapper.MapperMap(value,&tmp)
		tmp.CreateTime = dt
		tmp.ID = 0
		tmp.Currency = "USD"
		tmp.TradeDate = null.TimeFrom(tradeDate)
		tmp.Symbol = strings.ToUpper(tmp.Symbol)
		if isBlock != "" {
			tmp.IsBlock = isBlock
		}
		if isHot != "" {
			tmp.IsHot = isHot
		}
		if isChina != "" {
			tmp.IsChina = isChina
		}
		if isETF != "" {
			tmp.IsEtf = isETF
		}
		if isHK != "" {
			tmp.IsHk = isHK
			tmp.Currency = "HKD"
		}
		if isStar != "" {
			tmp.IsStar = isStar
		}
		if isOption != "" {
			tmp.IsOption = isOption
		}
		if isYestodayHot != "" {
			tmp.IsYestodayHot = isYestodayHot
			tmp.IsYestodayHotDate = null.TimeFrom(tradeDate)
		}

		tmp.Key = tmp.Symbol
		tmp.UpdateTime = dt
		tmp.BlacklistFlag = "-"
		tempList = append(tempList,tmp)
	}

	c.refreshBasicData(tempList,sCols,isAll)
	c.refreshHistoryData(dateString,tempList)
	CalculateTradeRaceTask()
	return
}

func (c *ImportDataSrv)updateFlag(sql string){
	qc := orm.NewQuery()
	qc.Exec(sql)
	xerr.ThrowError(qc.Error)
}

func (c *ImportDataSrv)refreshBasicData(items []models.StockBasic,sCols map[string]string,isAll bool){
	// 待更新的symbols
	var symbols []string
	for _, item := range items {
		symbols = append(symbols,item.Symbol)
	}

	// 获取原来的数据
	var oldItems []models.StockBasic
	q1 :=orm.NewQuery()
	if isAll {
		q1.GetItemWhere(&oldItems,"1=1")
	}else{
		q1.GetItemWhere(&oldItems,"symbol in (?)",symbols)
	}

	xerr.ThrowError(q1.Error)
	oldItemMap := make(map[string]models.StockBasic)
	for _, item := range oldItems {
		oldItemMap[item.Symbol] = item
	}

	// 更新字段
	var needInsertItems []models.StockBasic
	for _,newItem := range items {
		oldItem,ok := oldItemMap[newItem.Symbol]
		if !ok {
			// 如果不存在
			if _,ok := sCols["is_block"];ok {
				newItem.IsChina = global.LogicN
				newItem.IsEtf = global.LogicN
				newItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_china"];ok {
				newItem.IsBlock = global.LogicN
				newItem.IsEtf = global.LogicN
				newItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_etf"];ok {
				newItem.IsBlock = global.LogicN
				newItem.IsChina = global.LogicN
				newItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_hk"];ok {
				newItem.IsEtf = global.LogicN
				newItem.IsBlock = global.LogicN
				newItem.IsChina = global.LogicN
			}
			needInsertItems = append(needInsertItems,newItem)
		}else{
			// 更新字段
			oldItem.ID = 0
			oldItem.Name = newItem.Name
			oldItem.CurrentPrice = newItem.CurrentPrice
			oldItem.IncreaseRateCurrDay = newItem.IncreaseRateCurrDay
			oldItem.IncreaseRate60day = newItem.IncreaseRate60day
			oldItem.IncreaseRateFormYear = newItem.IncreaseRateFormYear
			oldItem.OpenPrice = newItem.OpenPrice
			oldItem.YesterdayPrice = newItem.YesterdayPrice
			oldItem.HighPrice = newItem.HighPrice
			oldItem.LowPrice = newItem.LowPrice
			oldItem.Pe = newItem.Pe
			oldItem.TotalMarketCap = newItem.TotalMarketCap
			oldItem.Industry = newItem.Industry
			oldItem.IncreaseRate5day = newItem.IncreaseRate5day
			oldItem.IncreaseRate10day = newItem.IncreaseRate10day
			oldItem.IncreaseRate20day = newItem.IncreaseRate10day
			oldItem.IncreaseRate120day = newItem.IncreaseRate120day
			oldItem.IncreaseRate250day = newItem.IncreaseRate250day
			oldItem.TradeAmount = newItem.TradeAmount
			oldItem.UpdateTime = newItem.UpdateTime
			oldItem.TradeDate = newItem.TradeDate
			if _,ok := sCols["is_block"];ok {
				oldItem.IsBlock = global.LogicY
				oldItem.IsChina = global.LogicN
				oldItem.IsEtf = global.LogicN
				oldItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_hot"];ok {
				oldItem.IsHot = global.LogicY
			}
			if _,ok := sCols["is_china"];ok {
				oldItem.IsChina = global.LogicY
				oldItem.IsBlock = global.LogicN
				oldItem.IsEtf = global.LogicN
				oldItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_etf"];ok {
				oldItem.IsEtf = global.LogicY
				oldItem.IsBlock = global.LogicN
				oldItem.IsChina = global.LogicN
				oldItem.IsHk = global.LogicN
			}
			if _,ok := sCols["is_hk"];ok {
				oldItem.IsHk = global.LogicY
				oldItem.IsEtf = global.LogicN
				oldItem.IsBlock = global.LogicN
				oldItem.IsChina = global.LogicN
			}
			if _,ok := sCols["is_star"];ok {
				oldItem.IsStar = global.LogicY
			}
			if _,ok := sCols["is_option"];ok {
				oldItem.IsOption = global.LogicY
			}
			if _,ok := sCols["is_yestoday_hot"];ok {
				oldItem.IsYestodayHot = global.LogicY
				oldItem.IsYestodayHotDate = newItem.TradeDate
			}

			needInsertItems = append(needInsertItems,oldItem)
		}
	}


	tx := orm.BeginTransaction()
	defer tx.EndTransaction()

	// 删除
	q2 := tx.NewQuery()
	if isAll {
		q2.DeleteAll("delete from stock_basic")
	}else{
		if len(oldItems) >0 {
			q2.DeleteItem(oldItems)
		}
	}

	xerr.ThrowError(q2.Error)
	// 新增
	q3 := tx.NewQuery()
	q3.BatchInsert(needInsertItems)
	xerr.ThrowError(q3.Error)
}

func (c *ImportDataSrv)checkSymbolExits(symbol string,data []map[string]interface{}){
	hasSymbol := "N"
	for _, row :=range data {
		value := row["symbol"]
		if value == symbol {
			hasSymbol = "Y"
			break
		}
	}
	if hasSymbol != "Y" {
		xerr.ThrowErrorMessage("没有发现Symbol:%s",symbol)
	}
}

func (c *ImportDataSrv)refreshHistoryData(dateString string,data []models.StockBasic){
	var cols []string
	cols = append(cols,"close_price")
	cols = append(cols,"high_price")
	cols = append(cols,"low_price")
	cols = append(cols,"open_price")
	cols = append(cols,"trade_amount")

	tradeDate := time.Now().AddDate(0,0,-1)
	if dateString != "" {
		tradeDate = utils.ParseShortTime(dateString)
	}else{
		dateString = tradeDate.Format("2006-01-02")
	}

	var items []models.StockHistoryData
	for _, row := range data {
		var item models.StockHistoryData
		item.Symbol = row.Symbol
		item.ClosePrice = row.CurrentPrice
		item.HighPrice = row.HighPrice
		item.LowPrice = row.LowPrice
		item.OpenPrice = row.OpenPrice
		item.TradeTime = dateString
		item.TradeTimeD = tradeDate
		item.TradeAmount = row.TradeAmount
		items = append(items,item)
	}

	var subRows []models.StockHistoryData
	for _, row := range items {
		subRows = append(subRows, row)
		if len(subRows) == 500 {
			qc := orm.NewQuery()
			qc.UpdateOrInsertHistory(&subRows,[]string{})
			subRows = nil
			xerr.ThrowError(qc.Error)
		}
	}
	if subRows != nil && len(subRows) > 0 {
		qc := orm.NewQuery()
		qc.UpdateOrInsertHistory(&subRows,cols)
		xerr.ThrowError(qc.Error)
	}
}