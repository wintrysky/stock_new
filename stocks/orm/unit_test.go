package orm

import (
	"github.com/spf13/cast"
	"os"
	"testing"
	"ww/stocks/models"
	"ww/stocks/xlog"
)

func TestMain(m *testing.M) {
	initLog()
	dbSettings := &DbSettings{
		DbName:          os.Getenv("stock.db.name"),
		Host:            os.Getenv("stock.db.host"),
		Port:            cast.ToInt(os.Getenv("stock.db.port")),
		User:            os.Getenv("stock.db.user"),
		Key:             os.Getenv("stock.db.key"),
		Password:        os.Getenv("stock.db.password"),
	}
	err := InitDB(dbSettings)
	if err != nil {
		panic(err)
	}

	m.Run()
}

func initLog() error {
	defer xlog.Sync()

	err := xlog.Init(&xlog.LogSettings{
		Level:    xlog.DefaultLevel,
		Path:     xlog.DefaultPath,
		FileName: xlog.DefaultFileName,
		CataLog:  xlog.DefaultCataLog,
		Caller:   xlog.DefaultCaller,
	})
	return err
}

func TestUpdateItem(t *testing.T) {
	var item models.StockBasic
	item.ID = 2
	item.Industry = "xxxxx"
	qa := NewQuery()
	qa.UpdateItem(&item,[]string{"industry"})
	if qa.Error != nil {
		t.Error(qa.Error)
	}
}

func TestGetItemWhereFirst(t *testing.T) {
	qa := NewQuery()
	var item models.StockBasic
	qa.GetItemWhereFirst(&item,"symbol = ?","UMAV")
	var item2 models.StockBasic
	qb := NewQuery()
	qb.GetItemWhereFirst(&item2,"name = ?","VERITEQ CORPORAT")
	if qb.Error != nil {
		t.Error(qa.Error)
	}
}

// 测试事务--批量更新
func TestBatchUpdate(t *testing.T){
	var items []models.StockBasic
	qa := NewQuery()
	qa.GetItemWhere(&items,"id < ?",7)

	if qa.RowsAffected == 0 {
		t.Error("没有找到数据")
	}

	for idx := range items {
		items[idx].Industry = "aaaa" + cast.ToString(idx)
		items[idx].IsChina = "N"
	}

	items[4].Symbol = "NNNN"
	items[5].Symbol = "NNNN"

	tx := BeginTransaction()
	defer tx.EndTransaction()
	tx.BatchUpdate(items,[]string{"industry","is_china","symbol"})
	if tx.Error != nil {
		t.Error(tx.Error.Error())
	}
}
