package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ww/stocks/dto"
	"ww/stocks/service"
	"ww/stocks/utils"
)

// +post /login/account
func Account() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p dto.LoginParam
		inputErr := ctx.BindJSON(&p)
		if inputErr != nil {
			utils.ApiErrorResponse(ctx,"参数转换错误")
			return
		}

		var result dto.LoginResponse
		result.Status = "error"
		result.Type = "account"
		result.CurrentAuthority = "user"
		err := service.GetUser(p.UserName,p.Password)
		if err == nil {
			result.Status = "ok"
		}

		ctx.JSON(http.StatusOK, result)


	}
}
