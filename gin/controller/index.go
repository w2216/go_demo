package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "支付宝测试！")
}
