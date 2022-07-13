package router

import (
	"gin_jwt/controller"
	"gin_jwt/middelware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	//用户登录
	r.GET("/jwt/login", controller.LoginController)
	// 获取token
	r.GET("/jwt/index", controller.JwtController)
	//使用中间件
	r.Use(middelware.JWTAuth())

	//获取列表数据
	r.GET("/jwt/list", controller.UserListController)
}
