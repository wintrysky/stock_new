package service

import (
	"errors"
	"strings"
	"ww/stocks/dto"
	"ww/stocks/global"
	"ww/stocks/models"
	"ww/stocks/orm"
	"ww/stocks/xerr"
)

func Search(p dto.StockSearchCondition)(items []models.StockBasic,err error){
	defer xerr.HandleErr(&err)

	db := orm.NewQuery()
	db = db.Where(p.Symbol != "","symbol = ?",strings.ToUpper(p.Symbol))
	db = db.Where(p.IsOption == true && p.SearchCategory !="BK","is_option = ?",global.LogicY)
	db = db.Where(p.BuyTags != "","buy_tags = ?",p.BuyTags)
	db = db.Where(p.CompanyTags != "","company_tags = ?",p.CompanyTags)
	if p.ProfitTags == "GoodGain" {
		db = db.Where(true,"pe > ?",0)
	}else if p.ProfitTags == "Debt"{
		db = db.Where(true,"pe < ?",0)
	}
	switch p.MarketCapTags {
	case "TenToHundred": //10-100亿
		db = db.Where(true,"total_market_cap between ? and ?",10,100)
	case "LargerThanHundred": // 大于100亿
		db = db.Where(true,"total_market_cap > ?",100)
	default: // 大于10亿
		db = db.Where(true,"total_market_cap >= ?",10)
	}
	switch p.TradeAmount {
	case "LargerThanM": // 大于1000万
		db = db.Where(true,"trade_amount > ?",1000)
	case "LargerThanOne": // 大于1亿
		db = db.Where(true,"trade_amount > ?",10000)
	case "LargerThanTen": // 大于10亿
		db = db.Where(true,"trade_amount > ?",100000)
	case "LargerThanHundred": // 大于50亿
		db = db.Where(true,"trade_amount > ?",500000)
	default:
		db = db.Where(true,"trade_amount > ?",10)
	}
	switch p.SearchCategory {
	case "BK":
		db = db.Where(true,"is_block = ?",global.LogicY)
	case "Star":
		db = db.Where(true,"is_star = ?",global.LogicY)
	case "ETF":
		db = db.Where(true,"is_etf = ?",global.LogicY)
	case "China":
		db = db.Where(true,"is_china = ?",global.LogicY)
	case "Hot":
		db = db.Where(true,"is_hot = ?",global.LogicY)
	case "YesterdayHot":
		db = db.Where(true,"is_yestoday_hot = ?",global.LogicY)
	}

	db.GetItemWhere(&items,"blacklist_flag != ?",global.LogicY)
	return items,db.Error
}

func SearchBlock(blockList []string)(blocks []string,err error){
	defer xerr.HandleErr(&err)

	var items []models.StockBasic
	db := orm.NewQuery()
	db.GetItemWhere(&items, "symbol in (?) and is_block='Y'",blockList)
	xerr.ThrowError(db.Error)

	var industryList []string
	for _, item := range items {
		industryList = append(industryList,item.Name)
	}

	var items2 []models.StockBasic
	qc := orm.NewQuery()
	qc.GetItemWhere(&items2,"industry in (?)",industryList)

	for _, item := range items2 {
		blocks = append(blocks,item.Symbol)
	}
	return
}

func GetUser(userName,pwd string)(err error){
	var item models.StockUser
	qa := orm.NewQuery()
	qa.GetItemWhereFirst(&item,"user_name = ? and user_pwd = ?",userName,pwd)
	if qa.Error != nil {
		return qa.Error
	}

	if item.UserName != userName {
		return errors.New("用户登录失败")
	}

	return qa.Error
}

