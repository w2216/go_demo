package controller

import (
	"encoding/json"
	"fmt"
	"gin_jwt/model"
	"gin_jwt/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterController(ctx *gin.Context) {
	parms, _ := ctx.GetRawData() // 定义map或结构体
	var req map[string]interface{}
	err := json.Unmarshal(parms, &req) // 反序列化
	if err != nil {
		fmt.Println("err: ", err)
	}

	email := req["email"].(string)
	password := req["password"].(string)
	if email == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	admin := model.Admin{Email: email, Password: utils.Md5(password)}
	db := model.Db
	firstRes := model.Admin{}
	db.Where("email = ?", email).First(&firstRes)
	if firstRes.ID > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已注册: " + email,
		})
		return
	}
	db.Create(&admin)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": gin.H{"id": admin.ID},
	})
}

func LoginWebController(ctx *gin.Context) {
	parms, _ := ctx.GetRawData() // 定义map或结构体
	var req map[string]interface{}
	err := json.Unmarshal(parms, &req) // 反序列化
	if err != nil {
		fmt.Println("err: ", err)
	}

	email := req["email"].(string)
	password := req["password"].(string)

	db := model.Db
	firstRes := model.Admin{}
	db.Where("email = ? and password = ?", email, utils.Md5(password)).First(&firstRes).Debug()
	if firstRes.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	//生成token
	webUsers := utils.WebUsers{Email: firstRes.Email, Id: firstRes.ID}
	token, err := utils.WebToken(webUsers)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}
	db.Create(&model.LoginLog{Aid: firstRes.ID})

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": gin.H{"token": token},
	})
}
