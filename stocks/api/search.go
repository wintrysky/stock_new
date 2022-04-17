package api

import (
	"github.com/gin-gonic/gin"
	"ww/stocks/dto"
	"ww/stocks/service"
	"ww/stocks/utils"
)

// +post /search/stocks
func SearchStocks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param dto.StockSearchCondition
		inputErr := ctx.BindJSON(&param)
		if inputErr != nil {
			utils.ApiErrorResponse(ctx,"参数转换错误")
			return
		}

		items,err := service.Search(param)

		utils.ApiResponse(ctx,items,err)
	}
}
