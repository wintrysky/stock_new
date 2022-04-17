package api

import (
	"github.com/gin-gonic/gin"
	"ww/stocks/service"
	"ww/stocks/utils"
)

// +get /tasks/CalculateTradeRaceTask
func CalculateTradeRaceTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		service.CalculateTradeRaceTask()

		utils.ApiResponse(ctx,nil,nil)
	}
}
