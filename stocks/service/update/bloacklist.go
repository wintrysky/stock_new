package update

import (
	"ww/stocks/orm"
	"ww/stocks/xerr"
)

func SetBlackList(symbols []string) (err error) {
	defer xerr.HandleErr(&err)

	if len(symbols) == 0 {
		xerr.ThrowErrorMessage("symbol list cant by empty")
	}

	qa := orm.NewQuery()
	qa.Exec("update stock_basic set blacklist_flag = 'Y' where symbol in (?)",symbols)
	return
}
