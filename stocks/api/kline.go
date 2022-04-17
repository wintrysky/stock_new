package api
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ww/stocks/dto"
	"ww/stocks/global"
	"ww/stocks/service"
	"ww/stocks/utils"
)


func GetHistoryBySymbol(ctx *gin.Context) {

	var param dto.KLineParam
	inputErr := ctx.BindJSON(&param)
	if inputErr != nil {
		utils.ApiErrorResponse(ctx,"参数转换错误")
		return
	}

	if param.IsBlock == "Y" {
		items, err := service.SearchBlock(param.CompareList)
		if err != nil {
			ctx.JSON(http.StatusOK, &global.Response{Success: false, Message: "0"})
			return
		}
		param.CompareList = items
	}
	if param.OptionName == "ALL" {
		param.CompareList = append(param.CompareList,"QQQ")
		param.CompareList = append(param.CompareList,"DIA")
		param.CompareList = append(param.CompareList,"SPY")
	}else{
		param.CompareList = append(param.CompareList,param.OptionName)
	}

	result,err := service.GetHistoryBySymbol(param.CompareList)

	if err != nil {
		ctx.JSON(http.StatusOK, &global.Response{Success: false, Message: "0"})
	}else{
		ctx.JSON(http.StatusOK, &global.Response{Success: true, Data: result})
	}
}
