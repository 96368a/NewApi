package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H) {
	res := gin.H{
		"code": code,
	}
	if data != nil {
		for k, v := range data {
			res[k] = v
		}
	}
	ctx.JSON(httpStatus, res)
}

func Success(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusOK, 200, data)
}

func Fail(ctx *gin.Context, code int, msg string) {
	Response(ctx, http.StatusOK, code, gin.H{
		"msg": msg,
	})
}
