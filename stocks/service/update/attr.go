package update

import (
	"strings"
	"time"
	"ww/stocks/dto"
	"ww/stocks/orm"
	"ww/stocks/xerr"
)

func UpdateStockAttr(p dto.AttributeParam)(err error){
	xerr.HandleErr(&err)

	if len(p.IDs) == 0 {
		xerr.ThrowErrorMessage("Please select at least one row")
	}

	var conns []interface{}
	sql := "update stock_basic set update_time = ?,"
	conns = append(conns,time.Now())
	if p.BuyTags != "" {
		if p.BuyTags == "--" {
			sql += "buy_tags = '',"
		}else{
			sql += "buy_tags = ?,"
			conns = append(conns,p.BuyTags)
		}
	}
	if p.CompanyTags != "" {
		if p.CompanyTags == "--" {
			sql += "company_tags = '',"
		}else{
			sql += "company_tags = ?,"
			conns = append(conns,p.CompanyTags)
		}
	}
	if p.ProfitTags != "" {
		if p.ProfitTags == "--" {
			sql += "profit_tags = '',"
		}else{
			sql += "profit_tags = ?,"
			conns = append(conns,p.ProfitTags)
		}
	}
	if p.Description != "" {
		sql += "description = ?"
		conns = append(conns,p.Description)
	}
	sql = strings.TrimRight(sql,",")
	sql += " where symbol in (?)"

	conns = append(conns,p.IDs)

	qa := orm.NewQuery()
	qa.Exec(sql,conns...)
	return qa.Error
}
