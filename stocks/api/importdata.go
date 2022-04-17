package api

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"ww/stocks/service"
	"ww/stocks/utils"
)

// +post /import/csv
func ImportCSV() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// nginx下不能用下划线
		importType := ctx.Request.Header.Get("importtype")
		dateString := ctx.Request.Header.Get("dateString")
		tagName := ctx.Request.Header.Get("tag")
		if tagName != "" {
			tagName,_ = url.QueryUnescape(tagName)
		}

		var err error
		cls := &service.ImportDataSrv{}
		err = cls.ImportData(ctx,importType,dateString)

		utils.ApiResponse(ctx,nil,err)
	}
}
