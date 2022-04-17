package api

import (
	"github.com/gin-gonic/gin"
	"ww/stocks/dto"
	"ww/stocks/service/update"
	"ww/stocks/utils"
)

// +post /update/blacklist
func SetBlackList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param []string
		inputErr := ctx.BindJSON(&param)
		if inputErr != nil {
			utils.ApiErrorResponse(ctx,"参数转换错误")
			return
		}

		err := update.SetBlackList(param)

		utils.ApiResponse(ctx,nil,err)
	}
}

// +post /update/UpdateStockAttr
func UpdateStockAttr() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param dto.AttributeParam

		inputErr := ctx.BindJSON(&param)
		if inputErr != nil {
			utils.ApiErrorResponse(ctx,"参数转换错误")
			return
		}

		err := update.UpdateStockAttr(param)

		utils.ApiResponse(ctx,nil,err)
	}
}
