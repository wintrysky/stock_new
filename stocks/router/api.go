package router

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"ww/stocks/api"
	"ww/stocks/middleware"
)

// UseRouters 定义路由
func UseRouters(eng *gin.Engine) {
	rg := eng.Group("/")
	var errWriter io.Writer = os.Stderr
	rg.Use(middleware.ErrorHandle(errWriter))
	useBizRouters(rg)
}

func useBizRouters(rg *gin.RouterGroup) {
	rg.POST("/login/account", api.Account())

	rg.POST("/search/stocks", api.SearchStocks())

	rg.POST("/import/csv", api.ImportCSV())

	rg.POST("/update/blackList", api.SetBlackList())
	rg.POST("/update/updateStockAttr", api.UpdateStockAttr())
	rg.POST("/kline/getHistoryBySymbol", api.GetHistoryBySymbol)

	rg.GET("/tasks/CalculateTradeRaceTask", api.CalculateTradeRaceTask())

}

