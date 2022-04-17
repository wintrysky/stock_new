package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ww/stocks/global"
)

func ApiResponse(ctx *gin.Context,rst interface{},err error){
	if err != nil {
		ctx.JSON(http.StatusOK, &global.Response{Success: false, Message: err.Error()})
	}else{
		ctx.JSON(http.StatusOK, &global.Response{Success: true,Data: rst})
	}
}

func ApiErrorResponse(ctx *gin.Context,errMsg string){
	ctx.JSON(http.StatusOK, &global.Response{Success: false, Message: errMsg})
}