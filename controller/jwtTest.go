package controller

import (
	"encoding/json"
	"fmt"
	"gin_jwt/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtController(ctx *gin.Context) {
	log.Println(ctx)
	user := utils.Users{
		Username: "hello",
		Password: "world",
	}
	token, err := utils.GenToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"token": token},
	})
}

func LoginController(ctx *gin.Context) {
	parms, _ := ctx.GetRawData() // 定义map或结构体
	var req map[string]interface{}
	err := json.Unmarshal(parms, &req) // 反序列化
	if err != nil {
		fmt.Println("err: ", err)
	}

	username := req["username"]
	password := req["password"]
	if !(username == "admin" && password == "123456") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	user := utils.Users{
		Username: "admin",
		Password: "123456",
	}
	//生成token
	token, err := utils.GenToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功",
		"data": gin.H{"token": token},
	})

}

func UserListController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "成功",
		"data": "[]",
	})
}
