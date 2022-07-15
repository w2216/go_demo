package controller

import (
	"fmt"
	"gin_jwt/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RdbSetController(ctx *gin.Context) {
	rdb := model.Rdb
	rctx := model.RdbCtx

	time := time.Now().Format("2006-01-02 15:04:05")

	result, err := rdb.Set(rctx, "time", time, 0).Result()
	if err != nil {
		fmt.Println("redis.ERROR: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"data": result},
	})
}

func RdbGetController(ctx *gin.Context) {
	rdb := model.Rdb
	rctx := model.RdbCtx

	result, err := rdb.Get(rctx, "time").Result()
	if err != nil {
		fmt.Println("redis.ERROR: ", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"data": result},
	})
}

func RdbDelController(ctx *gin.Context) {
	rdb := model.Rdb
	rctx := model.RdbCtx

	result, err := rdb.Del(rctx, "time").Result()
	if err != nil {
		fmt.Println("redis.ERROR: ", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"data": result},
	})
}
